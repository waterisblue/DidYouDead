package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
)

func GetAllSupply() (supplyList []entity.SupplyEntity) {
	DB := mysqlconn.GetDBCon()
	selectSQL := `SELECT id, sourceName, sourceDetail, imgUrl, type, price, subTitle, createTime FROM SupplySource where isActive = 1`
	query, err := DB.Query(selectSQL)
	if err != nil {
		log.Warning.Println("查询供应商资源失败")
		return
	}
	for query.Next() {
		var supply entity.SupplyEntity
		query.Scan(&supply.Id, &supply.SourceName, &supply.SourceDetail, &supply.ImgUrl, &supply.Type, &supply.Price, &supply.SubTitle, &supply.CreateTime)
		supplyList = append(supplyList, supply)
	}
	return
}

func AddSupply(supply entity.SupplyEntity) int {
	DB := mysqlconn.GetDBCon()
	insertSQL := `INSERT INTO SupplySource(sourceName, sourceDetail, imgUrl, type, price, subTitle) values (?, ?, ?, ?, ?, ?)`
	_, err := DB.Exec(insertSQL, supply.SourceName, supply.SourceDetail, supply.ImgUrl, supply.Type, supply.Price, supply.SubTitle)
	if err != nil {
		log.Warning.Println(err)
		return -1
	}
	return 0
}

func DeleteSupply(supply entity.SupplyEntity) int {
	DB := mysqlconn.GetDBCon()
	deleteSQl := `UPDATE SupplySource set isActive = 0 where id = ?`
	_, err := DB.Exec(deleteSQl, supply.Id)
	if err != nil {
		log.Warning.Println(supply, err)
		return -1
	}
	return 0
}

func SelectSupplyByType(types int) (supplyList []entity.SupplyEntity) {
	DB := mysqlconn.GetDBCon()
	selectSQL := `SELECT id, sourceName, sourceDetail, imgUrl, type, price, subTitle, createTime FROM SupplySource where isActive = 1 and type = ?`
	query, err := DB.Query(selectSQL, types)
	if err != nil {
		log.Warning.Println("查询供应商资源失败")
		return
	}
	for query.Next() {
		var supply entity.SupplyEntity
		query.Scan(&supply.Id, &supply.SourceName, &supply.SourceDetail, &supply.ImgUrl, &supply.Type, &supply.Price, &supply.SubTitle, &supply.CreateTime)
		supplyList = append(supplyList, supply)
	}
	return
}
