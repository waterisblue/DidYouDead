package entity

import "time"

type OrderEntity struct {
	Id         int        `json:"id,omitempty"`
	OrderNum   string     `json:"orderNum,omitempty"`
	OrderType  int        `json:"orderType,omitempty"`
	IsPay      bool       `json:"isPay,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LinkId     int        `json:"linkId,omitempty"`
}
