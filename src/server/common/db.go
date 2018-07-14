package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
	"log"
)

type DB_RESOURCE string

const (
	DB_RESOURCE_DEFAULT DB_RESOURCE = "test"
)

var DBPool Resource

type Resource interface {
	GetDB(name DB_RESOURCE) (*gorm.DB, bool)
	PutDB(name DB_RESOURCE, db *gorm.DB) bool
}

type DBResource struct {
	pool map[DB_RESOURCE]*gorm.DB
}

func NewDBResource() Resource {
	resource := new(DBResource)
	resource.pool = make(map[DB_RESOURCE]*gorm.DB)
	return resource
}

func (resource *DBResource) GetDB(name DB_RESOURCE) (*gorm.DB, bool) {
	if db, ok := resource.pool[name]; !ok {
		return nil, false
	} else {
		return db, true
	}
}

func (resource *DBResource) PutDB(name DB_RESOURCE, db *gorm.DB) bool {
	if db == nil {
		return false
	}
	if _, ok := resource.pool[name]; ok {
		return false
	} else {
		resource.pool[name] = db
		return true
	}
}

func init() {
	resource := NewDBResource()
	for _, dbConfig := range ConfigurationContext.DBConfigs {
		fmt.Println(dbConfig)
		db, err := gorm.Open(dbConfig.Config.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbConfig.Config.UserName, dbConfig.Config.Password, dbConfig.Config.Host, dbConfig.Config.Port, dbConfig.Config.DataBaseName))
		if err != nil {
			panic(err)
		}
		db.LogMode(dbConfig.Config.Mode)
		db.DB().SetConnMaxLifetime(time.Minute * dbConfig.Config.ConnMaxLifetime)
		db.DB().SetMaxOpenConns(dbConfig.Config.MaxOpenNum)
		db.DB().SetMaxIdleConns(dbConfig.Config.MaxIdleNum)
		ok := resource.PutDB(DB_RESOURCE(dbConfig.Name), db)
		if !ok {
			log.Fatal("the db add failed ...")
		}
	}
	fmt.Println(resource)
	DBPool = resource
}
