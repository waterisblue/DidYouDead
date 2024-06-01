package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
)

func GetUserByAccount(username string) (exist bool, userdetail entity.UserDetail) {
	var DB = mysqlconn.GetDBCon()
	exist = false
	selectSQL := "SELECT account, password, administrator FROM SysUser WHERE account = ?"
	rows, _ := DB.Query(selectSQL, username)
	if rows.Next() {
		defer rows.Close()
		exist = true
		rows.Scan(&userdetail.Username, &userdetail.Password, &userdetail.Administrator)
	}
	return
}

func AddUser(user *entity.UserInfo) (err error) {
	db := mysqlconn.GetDBCon()
	insertSQL := "INSERT INTO SysUser(account, password) VALUES(?, ?)"
	_, err = db.Exec(insertSQL, user.Username, user.Password)
	return
}

func GetAllUser() (userList []entity.UserDetail) {
	var DB = mysqlconn.GetDBCon()

	selectSQL := "SELECT account, administrator FROM SysUser"
	rows, err := DB.Query(selectSQL)
	if err != nil {
		log.Warning.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user entity.UserDetail
		rows.Scan(&user.Username, &user.Administrator)
		userList = append(userList, user)
	}
	return
}

func RemoveUserByAccount(account string) {
	var DB = mysqlconn.GetDBCon()

	deleteSQL := "DELETE FROM SysUser WHERE account = ?"
	_, err := DB.Exec(deleteSQL, account)
	if err != nil {
		log.Warning.Println(err)
	}
	return
}

func GrantUser(user entity.UserDetail) {
	var DB = mysqlconn.GetDBCon()
	log.Info.Println("更改权限：", user)

	updateSQL := `UPDATE SysUser set administrator = ? where account = ?`
	_, err := DB.Exec(updateSQL, !user.Administrator, user.Username)
	if err != nil {
		log.Warning.Println(err)
	}
	return
}
