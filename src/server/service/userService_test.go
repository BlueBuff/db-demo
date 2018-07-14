package service

import (
	"testing"
	"fmt"
)

func TestUserServiceImpl_GetUser(t *testing.T) {
	userDao:=NewUserServiceImpl()
	user:=userDao.GetUser(1)
	fmt.Println(user)
}
