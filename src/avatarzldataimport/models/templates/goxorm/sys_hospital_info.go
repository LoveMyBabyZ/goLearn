package goxorm

import (
	"time"
)

type SysHospitalInfo struct {
	AddTime             time.Time `json:"add_time" xorm:"comment('添加时间') DATETIME"`
	Address             string    `json:"address" xorm:"not null default '' comment('地址：地址') VARCHAR(128)"`
	Brand               string    `json:"brand" xorm:"not null default '' comment('品牌') VARCHAR(16)"`
	BrandCode           string    `json:"brand_code" xorm:"not null default '' comment('品牌内部编码') VARCHAR(16)"`
	BrandId             int       `json:"brand_id" xorm:"not null comment('品牌ID') INT(11)"`
	BusinessLicenseName string    `json:"business_license_name" xorm:"not null default '' comment('营业执照名称') VARCHAR(255)"`
	City                string    `json:"city" xorm:"not null default '' comment('地址：城市') VARCHAR(32)"`
	ClinicId            string    `json:"clinic_id" xorm:"not null default '' comment('公司内部编码') VARCHAR(16)"`
	ClinicName          string    `json:"clinic_name" xorm:"not null default '' comment('医院名称') VARCHAR(128)"`
	ClinicShortname     string    `json:"clinic_shortname" xorm:"not null default '' comment('医院简称') VARCHAR(64)"`
	ClinicStatus        int       `json:"clinic_status" xorm:"not null comment('基础信息_医院状态：0未知，1营业，2暂停，3闭店，4交接，5筹备，99删除') INT(11)"`
	ClinicType          string    `json:"clinic_type" xorm:"not null default '' comment('基础信息_医院类型：单选：未知/中心医院/中型全科医院/社区医院/专科医院/宠物店') VARCHAR(32)"`
	CreateTime          time.Time `json:"create_time" xorm:"comment('创院时间') DATE"`
	DataTime            time.Time `json:"data_time" xorm:"not null comment('数据时间') DATE"`
	Dean                string    `json:"dean" xorm:"not null default '' comment('院长') VARCHAR(128)"`
	DeanNumber          string    `json:"dean_number" xorm:"not null default '' comment('院长电话') VARCHAR(16)"`
	HasLingshou         int       `json:"has_lingshou" xorm:"not null comment('是否有零售业务：1是') INT(11)"`
	HasMedcine          int       `json:"has_medcine" xorm:"not null comment('是否有医疗业务：1是') INT(11)"`
	HasMeirong          int       `json:"has_meirong" xorm:"not null comment('是否有美容业务：1是') INT(11)"`
	HospitalPhone       string    `json:"hospital_phone" xorm:"not null default '' comment('前台电话') VARCHAR(20)"`
	HospitalProfile     string    `json:"hospital_profile" xorm:"not null default '' comment('医院简介') VARCHAR(400)"`
	Id                  int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	InnerContactor      string    `json:"inner_contactor" xorm:"not null default '' comment('第二联系人|运营经理') VARCHAR(16)"`
	InnerContactorPhone string    `json:"inner_contactor_phone" xorm:"not null default '' comment('第二联系人|运营经理电话') VARCHAR(16)"`
	LastModify          time.Time `json:"last_modify" xorm:"comment('最后修改时间') DATETIME"`
	Latitude            string    `json:"latitude" xorm:"not null default 0.000000 comment('纬度') DECIMAL(10,6)"`
	Longitude           string    `json:"longitude" xorm:"not null default 0.000000 comment('经度') DECIMAL(10,6)"`
	MeituanId           string    `json:"meituan_id" xorm:"not null default '0' comment('美团id') VARCHAR(63)"`
	ParentId            int       `json:"parent_id" xorm:"not null comment('所属组织ID') INT(11)"`
	Province            string    `json:"province" xorm:"not null default '' comment('地址：省份') VARCHAR(16)"`
	SubBrand            string    `json:"sub_brand" xorm:"not null default '' comment('子品牌') VARCHAR(32)"`
	SubBrandId          int       `json:"sub_brand_id" xorm:"not null default 0 comment('子品牌Id') INT(11)"`
	System              string    `json:"system" xorm:"comment('关联ID_HIS系统    单选：小暖/林特/迅德/汉思/其他') VARCHAR(16)"`
	SystemId            int       `json:"system_id" xorm:"comment('小暖系统内部ID') INT(11)"`
}
