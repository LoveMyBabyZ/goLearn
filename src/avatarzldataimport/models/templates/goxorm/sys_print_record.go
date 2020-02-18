package goxorm

import (
	"time"
)

type SysPrintRecord struct {
	Author     int       `json:"author" xorm:"comment('打印者ID') INT(11)"`
	AuthorName string    `json:"author_name" xorm:"comment('打印者') VARCHAR(100)"`
	BrandId    int       `json:"brand_id" xorm:"default 0 comment('品牌ID') INT(11)"`
	CustomerId int       `json:"customer_id" xorm:"not null comment('客户ID') INT(11)"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	NoteCode   string    `json:"note_code" xorm:"not null comment('文书编号') VARCHAR(100)"`
	NoteId     int       `json:"note_id" xorm:"not null comment('文书ID') INT(11)"`
	NoteName   string    `json:"note_name" xorm:"not null comment('文书名') VARCHAR(100)"`
	OrgId      int       `json:"org_id" xorm:"not null comment('机构ID') index INT(11)"`
	PetId      int       `json:"pet_id" xorm:"not null comment('宠物ID') INT(11)"`
	PrintDate  time.Time `json:"print_date" xorm:"comment('打印时间') DATETIME"`
	Url        string    `json:"Url" xorm:"comment('打印文件路径') VARCHAR(256)"`
}
