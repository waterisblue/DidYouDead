package entity

import "time"

type FireService struct {
	ID            int
	Username      string
	Name          string
	Sex           string
	IdNum         string
	Age           int
	Locate        string
	PhoneNum      string
	OrderTime     *time.Time
	FuneralParlor string
	FireService   string
	UrnStyle      string
	Cemetery      string
	CreateDate    *time.Time
}
