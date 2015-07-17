package dao

import (
	"akadmin/entity"
	"log"
)

//添加用户
func AddUser(user *entity.User) {
	row := mysqlDB.QueryRow("select id from user where loginName=?", user.LoginName)
	err := row.Scan()
	log.Println("err = ", err.Error())
	if err != nil {
		panic("loginName = " + user.LoginName + " is exists ")
	}
	result, err := mysqlDB.Exec("insert into user(name,LoginName,pwd) values(?,?,?)", user.Name, user.LoginName, user.Pwd)
	checknil(err)
	id, err := result.LastInsertId()
	checknil(err)
	rowNum, err := result.RowsAffected()
	checknil(err)
	log.Println("id = ", id, " , rowNum = ", rowNum)
}

//修改用户
func UpdateUser(user *entity.User) {
	stmt, err := mysqlDB.Prepare("update user set name=?,loginName=?,pwd=? where id=?")
	checknil(err)
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.LoginName, user.Pwd, user.Id)
	checknil(err)
}

//删除用户
func DelUser(ids string) {
	result, err := mysqlDB.Exec("delete from user where id in (" + ids + ")")
	checknil(err)
	rowNum, err := result.RowsAffected()
	checknil(err)
	log.Println("delete rowsAffected num : ", rowNum)
}

//获取用户列表
func GetUsers() (users []entity.User) {
	row := mysqlDB.QueryRow("select count(id) from user")
	var total int
	err := row.Scan(&total)
	checknil(err)
	users = make([]entity.User, total)
	rows, err := mysqlDB.Query("select * from user")
	checknil(err)
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&users[i].Id, &users[i].Name, &users[i].LoginName, &users[i].Pwd)
		checknil(err)
	}
	log.Println("users = ", users)
	return
}

func checknil(err error) {
	if err != nil {
		panic(err.Error())
	}
}
