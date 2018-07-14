package service

import (
	"hdg.com/db-demo/src/server/model"
	"hdg.com/db-demo/src/server/dao"
	"hdg.com/db-demo/src/server/common"
	"errors"
	"fmt"
	"hdg.com/db-demo/src/server/cache"
	"time"
)

type UserService interface {
	GetUser(id int) *model.User
	GetCacheUser(id int) *model.User
}

type UserServiceImpl struct {
	UserService
	dao   dao.UserDao
	cache cache.Cache
}

func NewUserServiceImpl() UserService {
	userService := new(UserServiceImpl)
	db, ok := common.DBPool.GetDB(common.DB_RESOURCE_DEFAULT)
	if !ok {
		panic(errors.New("db resource is failed"))
	}
	cache := cache.NewUserRedisCache(common.RedisClient)
	userService.cache = cache
	userService.dao = dao.NewUserDaoImpl(db)
	return userService
}

func (service *UserServiceImpl) GetUser(id int) *model.User {
	user, err := service.dao.GetUserById(id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = service.cache.Set("", user, time.Second*60)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func (service *UserServiceImpl) GetCacheUser(id int) *model.User{
	user:=new(model.User)
	user.Id=id
	err:=service.cache.Get("",user)
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return user
}
