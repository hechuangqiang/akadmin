package dao

import (
	"akadmin/entity"
	"strconv"
	"testing"
)

func Test_AddUser(t *testing.T) {
	user := &entity.User{Name: "张三", LoginName: "zhangsan", Pwd: "123"}
	AddUser(user)
}

func Test_GetUsers(t *testing.T) {
	t.Log("test_getUsers")
	users := GetUsers()
	t.Log("users = ", users)
}

func Test_DelUser(t *testing.T) {
	ids := ""
	for i := 0; i < 1000; i++ {
		ids = ids + "," + strconv.Itoa(i)
	}
	DelUser(ids)
}
