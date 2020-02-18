package goxorm

import (
	"time"
)

type SysSmsTemplate struct {
	Brand          string    `json:"brand" xorm:"comment('主品牌名称') VARCHAR(50)"`
	BrandId        int       `json:"brand_id" xorm:"comment('主品牌Id') INT(11)"`
	BusinessType   int       `json:"business_type" xorm:"comment('业务类型(1:押金 2:储值卡 3:次卡 4:退货 10:订单商品打折)') INT(11)"`
	CreateDate     time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	CreateUserid   int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	SubBrand       string    `json:"sub_brand" xorm:"not null comment('子品牌名称') VARCHAR(50)"`
	SubBrandId     int       `json:"sub_brand_id" xorm:"not null default 0 comment('子品牌Id') INT(11)"`
	TemplateId     int       `json:"template_id" xorm:"comment('模板Id') INT(11)"`
	TemplateType   int       `json:"template_type" xorm:"comment('模板类型(1:退款 2:折扣3:CRM)') INT(11)"`
}
