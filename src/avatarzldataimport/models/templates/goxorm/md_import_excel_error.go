package goxorm

import (
	"time"
)

type MdImportExcelError struct {
	AddTime          time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	ErrorInfo        string    `json:"error_info" xorm:"not null comment('错误信息（json）') LONGTEXT"`
	Filepath         string    `json:"filePath" xorm:"not null comment('文件路径') TEXT"`
	Id               int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	ImportFields     string    `json:"import_fields" xorm:"not null comment('字段名（json）') TEXT"`
	ImportTitle      string    `json:"import_title" xorm:"not null default '' comment('标题') VARCHAR(200)"`
	LastModify       time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	ModuleId         int       `json:"module_id" xorm:"not null default 0 comment('模块id') INT(11)"`
	ModuleName       string    `json:"module_name" xorm:"not null default '' comment('模块名称') VARCHAR(200)"`
	OperatorId       int       `json:"operator_id" xorm:"not null default 0 comment('操作人ID') INT(11)"`
	OperatorName     string    `json:"operator_name" xorm:"not null default '' comment('操作人') VARCHAR(60)"`
	OrgId            int       `json:"org_id" xorm:"not null default 0 comment('机构ID') INT(11)"`
	OriginImportData string    `json:"origin_import_data" xorm:"not null comment('初始数据（json）') LONGTEXT"`
	Type             int       `json:"type" xorm:"not null default 1 comment('数据类型（1：错误数据；2：忽略数据；3：入库数据）') TINYINT(1)"`
	Work             int       `json:"work" xorm:"not null default 0 comment('模板类型id(1:目录；2:品类)') INT(11)"`
}
