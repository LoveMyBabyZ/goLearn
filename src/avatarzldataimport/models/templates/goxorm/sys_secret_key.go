package goxorm

import (
	"time"
)

type SysSecretKey struct {
	CreateDate time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OrgId      int       `json:"org_id" xorm:"not null default 0 comment('机构Id') INT(11)"`
	SecretKey  string    `json:"secret_key" xorm:"comment('密钥') VARCHAR(50)"`
}
