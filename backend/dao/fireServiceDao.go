package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
)

func InsertFireService(fireService entity.FireService) {
	DB := mysqlconn.GetDBCon()
	insertSQL := `INSERT INTO FireService 
    	(account, sex, idNum, locate, phoneNum, orderTime, funeralParlor, fireService, urnStyle, cemetery, name, age)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := DB.Exec(insertSQL, fireService.Username, fireService.Sex, fireService.IdNum, fireService.Locate, fireService.PhoneNum, fireService.OrderTime,
		fireService.FuneralParlor, fireService.FireService, fireService.UrnStyle, fireService.Cemetery, fireService.Name, fireService.Age)
	if err != nil {
		log.Warning.Println("插入FireService时发生错误，插入结构：", fireService, err)
		return
	}
	return
}

func GetFireServiceByUserName(username string) (exists bool, fireService entity.FireService) {
	exists = false
	DB := mysqlconn.GetDBCon()
	selectSQL := `SELECT sex, idnum, locate, phonenum, ordertime, funeralparlor, fireservice, urnstyle, cemetery, name, age FROM FireService WHERE account = ?`
	rows, _ := DB.Query(selectSQL, username)
	if rows.Next() {
		exists = true
		rows.Scan(&fireService.Sex, &fireService.IdNum, &fireService.Locate, &fireService.PhoneNum, &fireService.OrderTime,
			&fireService.FuneralParlor, &fireService.FireService, &fireService.UrnStyle, &fireService.Cemetery, &fireService.Name, &fireService.Age)
	}
	return
}
