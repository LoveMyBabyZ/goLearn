package goxorm

import (
	"time"
)

type TMycatOrderDetail struct {
	CreateTime  time.Time `json:"create_time" xorm:"not null DATETIME"`
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ItemId      string    `json:"item_id" xorm:"not null default '' VARCHAR(10)"`
	OrderNumber string    `json:"order_number" xorm:"not null default '' VARCHAR(50)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null DATETIME"`
}
