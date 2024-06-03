package entity

import "time"

type SupplyEntity struct {
	Id           int        `json:"id,omitempty"`
	Type         int        `json:"type,omitempty"`
	SourceName   string     `json:"sourcename,omitempty"`
	SubTitle     string     `json:"subtitle,omitempty"`
	SourceDetail string     `json:"sourcedetail,omitempty"`
	ImgUrl       string     `json:"imgurl,omitempty"`
	CreateTime   *time.Time `json:"createtime,omitempty"`
	Price        int        `json:"price"`
	IsActive     bool       `json:"isactive,omitempty"`
}
