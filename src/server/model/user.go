package model

import (
	"time"
	"encoding/json"
	"log"
	"fmt"
	"errors"
)

type User struct {
	Id         int       `json:"id"         gorm:"id;AUTO_INCREMENT"`
	UserName   string    `json:"userName"   gorm:"user_name"`
	Password   string    `json:"-"          gorm:"password"`
	CreateTime time.Time `json:"createTime" gorm:"create_time"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Serialization() ([]byte,error){
	data,err:=json.Marshal(u)
	if err!=nil {
		return nil,err
	}
	return data,nil
}

func (u *User) UnSerialization(str []byte)error{
	err := json.Unmarshal(str,u)
	return err
}

func (u *User) ToString() string{
	data,err:=u.Serialization()
	if err!=nil {
		log.Fatal(err)
		return ""
	}
	return string(data)
}

/**
调用此方法的时候，一定要保证ID不等于0
 */
func (u *User) GetStringKey() string{
	if u.Id == 0{
		panic(errors.New("the id not == 0"))
	}
	return fmt.Sprintf("db:user:str:%d",u.Id)
}
