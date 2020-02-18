package goxorm

import (
	"time"
)

type THosDepartment struct {
	AddTime            time.Time `json:"add_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	CompanyCode        string    `json:"company_code" xorm:"comment('所属公司的编码') VARCHAR(64)"`
	CompanyName        string    `json:"company_name" xorm:"comment('所属公司名称') VARCHAR(64)"`
	DepartmentType     int       `json:"department_type" xorm:"comment('组织区域的类型') TINYINT(3)"`
	DepartmentTypeCn   string    `json:"department_type_cn" xorm:"comment('组织区域的类型名称') VARCHAR(16)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ParentCode         string    `json:"parent_code" xorm:"comment('上级组织的编码') VARCHAR(64)"`
	ParentName         string    `json:"parent_name" xorm:"comment('上级组织的名称') VARCHAR(64)"`
	ParentType         int       `json:"parent_type" xorm:"comment('上级组织的类型') TINYINT(3)"`
	StructCode         string    `json:"struct_code" xorm:"comment('组织的编码，权限系统分配的组织标识，ID一样的唯一标识') VARCHAR(64)"`
	StructCompanyId    int       `json:"struct_company_id" xorm:"comment('所属公司的ID') INT(11)"`
	StructDepartSort   int       `json:"struct_depart_sort" xorm:"comment('组织所属区域的编号') INT(3)"`
	StructDepartSortCn string    `json:"struct_depart_sort_cn" xorm:"comment('组织所属区域的类型名称') VARCHAR(16)"`
	StructId           int       `json:"struct_id" xorm:"not null comment('组织ID') unique INT(11)"`
	StructName         string    `json:"struct_name" xorm:"comment('组织名称') VARCHAR(64)"`
	StructParentId     int       `json:"struct_parent_id" xorm:"comment('上级组织的ID') index INT(11)"`
	StructShortName    string    `json:"struct_short_name" xorm:"comment('组织简称') VARCHAR(64)"`
	StructStatus       int       `json:"struct_status" xorm:"comment('组织机构状态，0运营，1停用，99删除') INT(11)"`
	StructType         int       `json:"struct_type" xorm:"comment('组织类型，10代表公司') TINYINT(3)"`
}
