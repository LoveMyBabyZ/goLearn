package goxorm

import (
	"time"
)

type MdAssay struct {
	AddTime    time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	AssayName  string    `json:"assay_name" xorm:"not null default '' comment('化验名称') VARCHAR(200)"`
	CatMax     string    `json:"cat_max" xorm:"not null default '' comment('猫最大值') VARCHAR(30)"`
	CatMin     string    `json:"cat_min" xorm:"not null default '' comment('猫最小值') VARCHAR(30)"`
	CategoryId string    `json:"category_id" xorm:"not null default '' comment('目录ID') VARCHAR(30)"`
	DogMax     string    `json:"dog_max" xorm:"not null default '' comment('犬最大值') VARCHAR(30)"`
	DogMin     string    `json:"dog_min" xorm:"not null default '' comment('犬最小值') VARCHAR(30)"`
	From       int       `json:"from" xorm:"not null default 1 comment('数据来源（1：医院自建；2：数据导入）') TINYINT(1)"`
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IsDeleted  int       `json:"is_deleted" xorm:"not null default 0 comment('是否删除（0：未删除；1：已删除）') TINYINT(1)"`
	IsHistory  int       `json:"is_history" xorm:"not null default 0 comment('是否历史数据（0：非历史数据；1：历史数据）') TINYINT(1)"`
	ItemCode   string    `json:"item_code" xorm:"not null default '' comment('总部编码') VARCHAR(30)"`
	LastModify time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	NuanId     int       `json:"nuan_id" xorm:"default 0 comment('化验指标ID') INT(11)"`
	OrgId      int       `json:"org_id" xorm:"not null default 0 comment('机构ID') INT(11)"`
	OtherMax   string    `json:"other_max" xorm:"not null default '' comment('其他最大值') VARCHAR(30)"`
	OtherMin   string    `json:"other_min" xorm:"not null default '' comment('其他最小值') VARCHAR(30)"`
	Remark     string    `json:"remark" xorm:"not null default '' comment('注意事项/备注') VARCHAR(2000)"`
	Type       int       `json:"type" xorm:"not null default 0 comment('化验类型(1->自动项目;2->手动项目)') TINYINT(2)"`
	UnitId     int       `json:"unit_id" xorm:"not null default 0 comment('单位ID') INT(11)"`
	UnitName   string    `json:"unit_name" xorm:"not null default '' comment('单位名称') VARCHAR(30)"`
}
