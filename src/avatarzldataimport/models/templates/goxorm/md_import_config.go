package goxorm

import (
	"time"
)

type MdImportConfig struct {
	AddTime      time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	DefaultValue string    `json:"default_value" xorm:"not null default '' comment('非必填项的默认值') VARCHAR(30)"`
	DictType     string    `json:"dict_type" xorm:"not null default '' comment('字典型字段的字典类型ID') VARCHAR(200)"`
	EnumValue    string    `json:"enum_value" xorm:"not null default '' comment('枚举型字段的合法枚举值（json）') VARCHAR(200)"`
	Field        string    `json:"field" xorm:"not null default '' comment('数据表字段名称') VARCHAR(200)"`
	FieldType    int       `json:"field_type" xorm:"not null default 0 comment('字段类型（1：需验重字段；2：必填项字典；3：非必填项字典：4：枚举值字段；5：数值型字段；6：普通必填项；7：生产商编码；8：子类型；）') TINYINT(2)"`
	Id           int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IgnoreValue  string    `json:"ignore_value" xorm:"not null default '' comment('验重字段的忽略验重值（json）') VARCHAR(200)"`
	IsMust       string    `json:"is_must" xorm:"not null default '' comment('是否必填（0：非必填；1：必填）') VARCHAR(30)"`
	LastModify   time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	ModuleId     int       `json:"module_id" xorm:"not null default 0 comment('模块id') INT(11)"`
	ModuleName   string    `json:"module_name" xorm:"not null default '' comment('模块名称') VARCHAR(200)"`
	Name         string    `json:"name" xorm:"not null default '' comment('名称(标题或字段)') VARCHAR(200)"`
	NumberLength string    `json:"number_length" xorm:"not null default '' comment('数值型字段的小数点保留最大长度（0：必须整数；2：大于0且最多小数点后2位；5：大于0且最多小数点后5为）') VARCHAR(200)"`
	Order        int       `json:"order" xorm:"not null default 0 comment('字段排序值') TINYINT(3)"`
	OrgId        int       `json:"org_id" xorm:"not null comment('机构ID') INT(11)"`
	Remark       string    `json:"remark" xorm:"not null comment('模板字段填写注意事项') TEXT"`
	Type         string    `json:"type" xorm:"not null default '' comment('类型：1-标题  2-字段') VARCHAR(30)"`
	Width        int       `json:"width" xorm:"not null default 0 comment('字段宽度') TINYINT(3)"`
}
