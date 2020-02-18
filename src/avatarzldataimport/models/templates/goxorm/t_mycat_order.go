package goxorm

import (
	"time"
)

type TMycatOrder struct {
	Amount      string    `json:"amount" xorm:"not null default 0.00 DECIMAL(11,2)"`
	Count       string    `json:"count" xorm:"not null default 0 DECIMAL(11)"`
	CreateDate  time.Time `json:"create_date" xorm:"not null DATETIME"`
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OrderNumber string    `json:"order_number" xorm:"not null default '' VARCHAR(50)"`
	UpdateDate  time.Time `json:"update_date" xorm:"not null DATETIME"`
}
