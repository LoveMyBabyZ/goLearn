package goxorm

import (
	"time"
)

type MdOperateLog struct {
	AddTime        time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastModify     time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	Module         int       `json:"module" xorm:"not null default 0 comment('模块编码') INT(11)"`
	ObjCode        string    `json:"obj_code" xorm:"comment('操作对象编码（json格式）') TEXT"`
	OperateContent string    `json:"operate_content" xorm:"not null comment('操作内容') TEXT"`
	OperateData    string    `json:"operate_data" xorm:"comment('操作数据内容（json格式）') TEXT"`
	OperateObj     string    `json:"operate_obj" xorm:"not null default '' comment('操作对象') VARCHAR(30)"`
	OperateType    int       `json:"operate_type" xorm:"not null default 0 comment('操作类型') INT(11)"`
	OperatorId     int       `json:"operator_id" xorm:"not null default 0 comment('操作人ID') INT(11)"`
	OperatorName   string    `json:"operator_name" xorm:"not null default '' comment('操作人') VARCHAR(60)"`
	OrgId          int       `json:"org_id" xorm:"not null default 0 comment('机构ID') INT(11)"`
}
