package goxorm

import (
	"time"
)

type MdSupplier struct {
	AddTime                    time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Address                    string    `json:"address" xorm:"not null default '' comment('办公地址') VARCHAR(200)"`
	AnimalLicenseCode          string    `json:"animal_license_code" xorm:"not null default '' comment('兽医经营许可证编号') VARCHAR(100)"`
	AnimalLicenseEndTime       string    `json:"animal_license_end_time" xorm:"not null default '' comment('兽医经营许可证结束日期') VARCHAR(30)"`
	AnimalLicenseImg           string    `json:"animal_license_img" xorm:"comment('兽医经营许可证照片(json)') TEXT"`
	AnimalLicenseRegisterTime  string    `json:"animal_license_register_time" xorm:"not null default '' comment('兽医经营许可证注册日期') VARCHAR(30)"`
	AnimalLicenseStartTime     string    `json:"animal_license_start_time" xorm:"not null default '' comment('兽医经营许可证开始日期') VARCHAR(30)"`
	CompanyBusinessEndTime     string    `json:"company_business_end_time" xorm:"not null default '' comment('企业营业结束日期') VARCHAR(30)"`
	CompanyBusinessStartTime   string    `json:"company_business_start_time" xorm:"not null default '' comment('企业营业开始日期') VARCHAR(30)"`
	CompanyEstablishTime       string    `json:"company_establish_time" xorm:"not null default '' comment('企业成立日期') VARCHAR(30)"`
	CompanyType                int       `json:"company_type" xorm:"not null default 0 comment('企业类型') INT(11)"`
	EnName                     string    `json:"en_name" xorm:"default '' comment('英文名称') VARCHAR(100)"`
	FinanceAccount             string    `json:"finance_account" xorm:"not null default '' comment('收款账号') VARCHAR(30)"`
	FinanceAccountName         string    `json:"finance_account_name" xorm:"not null default '' comment('收款户名') VARCHAR(30)"`
	FinanceAccountType         int       `json:"finance_account_type" xorm:"not null default 0 comment('收款账号性质') INT(11)"`
	FinanceBank                string    `json:"finance_bank" xorm:"not null default '' comment('收款开户行') VARCHAR(30)"`
	FinanceCurrencyType        int       `json:"finance_currency_type" xorm:"not null default 0 comment('收款币种') INT(11)"`
	FinanceInvoiceTaxRate      string    `json:"finance_invoice_tax_rate" xorm:"not null default '' comment('发票税率') VARCHAR(30)"`
	FinanceInvoiceType         int       `json:"finance_invoice_type" xorm:"not null default 0 comment('发票性质') INT(11)"`
	FinanceMail                string    `json:"finance_mail" xorm:"not null default '' comment('财务邮箱') VARCHAR(100)"`
	FinanceName                string    `json:"finance_name" xorm:"not null default '' comment('财务联系人') VARCHAR(30)"`
	FinancePhone               string    `json:"finance_phone" xorm:"not null default '' comment('财务电话') VARCHAR(30)"`
	From                       int       `json:"from" xorm:"not null default 1 comment('数据来源（1：医院自建；2：数据导入）') TINYINT(1)"`
	Id                         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IsDealer                   int       `json:"is_dealer" xorm:"not null default 0 comment('是经销商(1->是;0->否;)') TINYINT(2)"`
	IsDeleted                  int       `json:"is_deleted" xorm:"not null default 0 comment('是否删除(0->未删除;1->已删除;)') TINYINT(2)"`
	IsHistory                  int       `json:"is_history" xorm:"not null default 0 comment('是否历史数据（0：非历史数据；1：历史数据）') TINYINT(1)"`
	IsProducer                 int       `json:"is_producer" xorm:"not null default 0 comment('是生产商(1->是;0->否;)') TINYINT(2)"`
	ItemCode                   string    `json:"item_code" xorm:"not null default '' comment('总部编码') VARCHAR(30)"`
	LastModify                 time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	LegalName                  string    `json:"legal_name" xorm:"not null default '' comment('法人代表') VARCHAR(30)"`
	LicenseImg                 string    `json:"license_img" xorm:"comment('营业执照照片(json)') TEXT"`
	NuanId                     int       `json:"nuan_id" xorm:"not null default 0 comment('小暖ID') index(idx_org_nuan_id) INT(11)"`
	OldFaxNumber               string    `json:"old_fax_number" xorm:"not null default '' comment('传真号码') VARCHAR(30)"`
	OldLeader                  string    `json:"old_leader" xorm:"not null default '' comment('负责人') VARCHAR(30)"`
	OldPostalCode              string    `json:"old_postal_code" xorm:"not null default '' comment('邮政编码') VARCHAR(30)"`
	OldScopeBusiness           string    `json:"old_scope_business" xorm:"not null default '' comment('经营范围') VARCHAR(30)"`
	OldTelephone               string    `json:"old_telephone" xorm:"not null default '' comment('联系电话1') VARCHAR(30)"`
	OperatorId                 int       `json:"operator_id" xorm:"not null default 0 comment('操作人ID') INT(11)"`
	OperatorName               string    `json:"operator_name" xorm:"not null default '' comment('操作人') VARCHAR(60)"`
	OrgId                      int       `json:"org_id" xorm:"not null default 0 comment('机构ID') index(idx_org_nuan_id) INT(11)"`
	PayPeriodType              int       `json:"pay_period_type" xorm:"not null default 0 comment('结算周期') INT(11)"`
	RegisterCapital            string    `json:"register_capital" xorm:"not null default '' comment('注册资本') VARCHAR(30)"`
	RegisterCapitalCurrency    int       `json:"register_capital_currency" xorm:"not null default 0 comment('注册资本币种') INT(11)"`
	Remark                     string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(2000)"`
	SalesmanMail               string    `json:"salesman_mail" xorm:"not null default '' comment('销售邮箱') VARCHAR(100)"`
	SalesmanMobile             string    `json:"salesman_mobile" xorm:"not null default '' comment('销售手机') VARCHAR(30)"`
	SalesmanName               string    `json:"salesman_name" xorm:"not null default '' comment('销售联系人') VARCHAR(30)"`
	SalesmanPhone              string    `json:"salesman_phone" xorm:"not null default '' comment('销售座机') VARCHAR(30)"`
	ServiceProvince            string    `json:"service_province" xorm:"comment('供应商服务区域(json)') TEXT"`
	SimpleName                 string    `json:"simple_name" xorm:"not null default '' comment('拼音简写') VARCHAR(30)"`
	Status                     int       `json:"status" xorm:"not null default 1 comment('供应商状态(1->启用;2->冻结;3->停用)') TINYINT(2)"`
	SupplierAgentAuthorization string    `json:"supplier_agent_authorization" xorm:"comment('供应商代理授权') TEXT"`
	SupplierName               string    `json:"supplier_name" xorm:"not null default '' comment('供应商名称') VARCHAR(200)"`
	SupplierTaxCode            string    `json:"supplier_tax_code" xorm:"not null default '' comment('供应商税号') VARCHAR(30)"`
	TaxCode                    string    `json:"tax_code" xorm:"not null default '' comment('统一社会信用代码(营业执照编号)') VARCHAR(200)"`
	Type                       int       `json:"type" xorm:"not null default 1 comment('供应商类型标识(1->医院自建;2->总部下发)') TINYINT(2)"`
}
