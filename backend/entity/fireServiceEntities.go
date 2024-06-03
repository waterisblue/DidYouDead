package entity

import (
	"time"
)

type FireService struct {
	Name          string     `json:"Name,omitempty"`
	Sex           string     `json:"Sex,omitempty"`
	IdNum         string     `json:"IdNum,omitempty"`
	Age           int        `json:"Age,omitempty"`
	Locate        string     `json:"Locate,omitempty"`
	PhoneNum      string     `json:"PhoneNum,omitempty"`
	OrderTime     *time.Time `json:"OrderTime,omitempty"`
	FuneralParlor string     `json:"FuneralParlor,omitempty"`
	FireService   string     `json:"FireService,omitempty"`
	UrnStyle      string     `json:"UrnStyle,omitempty"`
	Cemetery      string     `json:"Cemetery,omitempty"`
	Amount        int        `json:"Amount,omitempty"`
}

type FireServiceJoinOrderData struct {
	Id            int        `json:"id,omitempty"`
	Name          string     `json:"Name,omitempty"`
	Sex           string     `json:"Sex,omitempty"`
	IdNum         string     `json:"IdNum,omitempty"`
	Age           int        `json:"Age,omitempty"`
	Locate        string     `json:"Locate,omitempty"`
	PhoneNum      string     `json:"PhoneNum,omitempty"`
	OrderTime     *time.Time `json:"OrderTime,omitempty"`
	FuneralParlor string     `json:"FuneralParlor,omitempty"`
	FireService   string     `json:"FireService,omitempty"`
	UrnStyle      string     `json:"UrnStyle,omitempty"`
	Cemetery      string     `json:"Cemetery,omitempty"`
	IsPay         bool       `json:"IsPay,omitempty"`
	Amount        int        `json:"Amount,omitempty"`
}

type AmountCalEntity struct {
	FuneralParlorId int `json:"funeralParlorId,omitempty"`
	FireServiceId   int `json:"fireServiceId,omitempty"`
	UrnStyleId      int `json:"urnStyleId,omitempty"`
	CemeteryId      int `json:"cemeteryId,omitempty"`
}
