package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
)

func InsertPayDao(payEntity entity.OrderEntity) {
	DB := mysqlconn.GetDBCon()
	insertSQL := "INSERT INTO OrderData(orderNum, orderType, LinkId) values (?, ?, ?)"
	_, err := DB.Exec(insertSQL, payEntity.OrderNum, payEntity.OrderType, payEntity.LinkId)
	if err != nil {
		log.Warning.Println("插入失败：", err)
	}
}

func GetOrderById(OrderType int, LinkId int) (order entity.OrderEntity) {
	DB := mysqlconn.GetDBCon()
	selectSQL := "SELECT orderNum, createDate, isPay FROM OrderData WHERE orderType = ? and LinkId = ?"
	query, err := DB.Query(selectSQL, OrderType, LinkId)
	if err != nil {
		log.Warning.Println(err)
	}
	err = query.Scan(&order.OrderNum, &order.CreateDate, &order.IsPay)
	if err != nil {
		log.Warning.Println(err)
	}
	return
}

func UpdatePayDao(isPay bool, orderNum string) {
	DB := mysqlconn.GetDBCon()
	updateSQL := "UPDATE OrderData set isPay = ? where orderNum = ?"
	_, err := DB.Exec(updateSQL, isPay, orderNum)
	if err != nil {
		log.Warning.Println(err)
	}
}
