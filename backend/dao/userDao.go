package dao

import (
	"dyd/entity"
	"dyd/mysqlconn"
)

func GetUserByAccount(username string) (exist bool, administrator string) {
	exist = false
	db := mysqlconn.GetDBCon()
	selectSQL := "SELECT account, administrator FROM SysUser WHERE account = ?"
	rows, _ := db.Query(selectSQL, username)
	defer rows.Close()

	if rows.Next() {
		exist = true
		rows.Scan(&administrator)
	}
	return
}

func AddUser(user *entity.UserInfo) (err error) {
	db := mysqlconn.GetDBCon()
	insertSQL := "INSERT INTO SysUser(account, password) VALUES(?, ?)"
	_, err = db.Exec(insertSQL, user.Username, user.Password)
	return
}
