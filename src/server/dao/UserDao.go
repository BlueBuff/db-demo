package dao

import (
	"github.com/jinzhu/gorm"
	"hdg.com/db-demo/src/server/model"
	"errors"
)

type UserDao interface {
	GetUserById(id int) (*model.User, error)
}

type UserDaoImpl struct {
	UserDao
	db *gorm.DB
}

func NewUserDaoImpl(db *gorm.DB) UserDao {
	userDao := new(UserDaoImpl)
	userDao.db = db
	return userDao
}

func (dao *UserDaoImpl) GetUserById(id int) (*model.User, error) {
	user := model.User{}
	db := dao.db.First(&user, id)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("not found ... ")
	}
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}
