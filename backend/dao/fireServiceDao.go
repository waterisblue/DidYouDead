package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
)

func InsertFireService(fireService entity.FireService, UserName string) {
	DB := mysqlconn.GetDBCon()
	// 将该用户已拥有的火化方式置0
	updateSQL := `UPDATE FireService SET isActive = 0 WHERE account = ?`
	_, err := DB.Exec(updateSQL, UserName)
	if err != nil {
		log.Warning.Println("将用户拥有的FireService取消激活失败！", fireService, err)
		return
	}
	insertSQL := `INSERT INTO FireService 
    	(account, sex, idNum, locate, phoneNum, orderTime, funeralParlor, fireService, urnStyle, cemetery, name, age, amount)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = DB.Exec(insertSQL, UserName, fireService.Sex, fireService.IdNum, fireService.Locate, fireService.PhoneNum, fireService.OrderTime,
		fireService.FuneralParlor, fireService.FireService, fireService.UrnStyle, fireService.Cemetery, fireService.Name, fireService.Age, fireService.Amount)
	if err != nil {
		log.Warning.Println("插入FireService时发生错误，插入结构：", fireService, err)
		return
	}
	return
}

func GetFireServiceByUserName(username string) (exists bool, fireService entity.FireServiceJoinOrderData) {
	exists = false
	DB := mysqlconn.GetDBCon()
	selectSQL := `SELECT id, sex, idnum, locate, phonenum, ordertime, funeralparlor, fireservice, urnstyle, cemetery, name, age, amount FROM FireService WHERE account = ? and isActive = 1`
	selectPaySQL := `SELECT isPay from OrderData where LinkId = ? and orderType = 0`
	rows, _ := DB.Query(selectSQL, username)
	if rows.Next() {
		exists = true
		err := rows.Scan(&fireService.Id, &fireService.Sex, &fireService.IdNum, &fireService.Locate, &fireService.PhoneNum, &fireService.OrderTime,
			&fireService.FuneralParlor, &fireService.FireService, &fireService.UrnStyle, &fireService.Cemetery, &fireService.Name, &fireService.Age, &fireService.Amount)
		if err != nil {
			log.Warning.Println("从查询中获取数据失败，", fireService, err)
			return
		}
		query, err := DB.Query(selectPaySQL, fireService.Id)
		if err != nil {
			log.Warning.Println(err)
		}
		if query.Next() {
			err := query.Scan(&fireService.IsPay)
			if err != nil {
				log.Warning.Println(err)
			}
		} else {
			fireService.IsPay = false
		}
	}
	return
}

func CalAmount(amountIds entity.AmountCalEntity) int {
	DB := mysqlconn.GetDBCon()
	selectSQL := `select sum(price) amount from SupplySource where id in (?, ?, ?, ?)`
	row := DB.QueryRow(selectSQL, amountIds.CemeteryId, amountIds.FireServiceId, amountIds.FuneralParlorId, amountIds.UrnStyleId)
	var res int
	err := row.Scan(&res)
	if err != nil {
		log.Warning.Println(amountIds, err)
	}
	return res
}
