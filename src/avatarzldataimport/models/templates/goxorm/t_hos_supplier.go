package goxorm

import (
	"time"
)

type THosSupplier struct {
	AddTime                    time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	AddType                    int       `json:"add_type" xorm:"not null default 1 comment('数据来源类型(1->总部自建;2->医院申请建立)') TINYINT(2)"`
	Address                    string    `json:"address" xorm:"not null default '' comment('办公地址') VARCHAR(200)"`
	AnimalLicenseCode          string    `json:"animal_license_code" xorm:"not null default '' comment('兽医经营许可证编号') VARCHAR(100)"`
	AnimalLicenseEndTime       string    `json:"animal_license_end_time" xorm:"not null default '' comment('兽医经营许可证结束日期') VARCHAR(30)"`
	AnimalLicenseImg           string    `json:"animal_license_img" xorm:"comment('兽医经营许可证照片(json)') TEXT"`
	AnimalLicenseRegisterTime  string    `json:"animal_license_register_time" xorm:"not null default '' comment('兽医经营许可证注册日期') VARCHAR(30)"`
	AnimalLicenseStartTime     string    `json:"animal_license_start_time" xorm:"not null default '' comment('兽医经营许可证开始日期') VARCHAR(30)"`
	CompanyBusinessEndTime     string    `json:"company_business_end_time" xorm:"not null default '' comment('企业营业结束日期') VARCHAR(30)"`
	CompanyBusinessStartTime   string    `json:"company_business_start_time" xorm:"not null default '' comment('企业营业开始日期') VARCHAR(30)"`
	CompanyEstablishTime       string    `json:"company_establish_time" xorm:"not null default '' comment('企业成立日期') VARCHAR(30)"`
	CompanyId                  int       `json:"company_id" xorm:"not null default 0 comment('公司ID') unique(UNIQUE_COMPANY_SUPPLIER_CODE) INT(11)"`
	CompanyType                string    `json:"company_type" xorm:"not null default '' comment('企业类型') VARCHAR(30)"`
	EnName                     string    `json:"en_name" xorm:"default '' comment('英文名称') VARCHAR(100)"`
	FinanceAccount             string    `json:"finance_account" xorm:"not null default '' comment('收款账号') VARCHAR(30)"`
	FinanceAccountName         string    `json:"finance_account_name" xorm:"not null default '' comment('收款户名') VARCHAR(30)"`
	FinanceAccountType         string    `json:"finance_account_type" xorm:"not null default '' comment('收款账号性质') VARCHAR(30)"`
	FinanceBank                string    `json:"finance_bank" xorm:"not null default '' comment('收款开户行') VARCHAR(30)"`
	FinanceCurrencyType        string    `json:"finance_currency_type" xorm:"not null default '' comment('收款币种') VARCHAR(30)"`
	FinanceInvoiceTaxRate      string    `json:"finance_invoice_tax_rate" xorm:"not null default '' comment('发票税点') VARCHAR(30)"`
	FinanceInvoiceType         string    `json:"finance_invoice_type" xorm:"not null default '' comment('发票性质') VARCHAR(30)"`
	FinanceMail                string    `json:"finance_mail" xorm:"not null default '' comment('财务邮箱') VARCHAR(100)"`
	FinanceName                string    `json:"finance_name" xorm:"not null default '' comment('财务联系人') VARCHAR(30)"`
	FinancePhone               string    `json:"finance_phone" xorm:"not null default '' comment('财务电话') VARCHAR(30)"`
	Id                         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IsDealer                   int       `json:"is_dealer" xorm:"not null default 0 comment('是否经销商(1->是;0->否;)') TINYINT(2)"`
	IsDeleted                  int       `json:"is_deleted" xorm:"not null default 0 comment('是否删除(0->未删除;1->已删除;)') TINYINT(2)"`
	IsProducer                 int       `json:"is_producer" xorm:"not null default 0 comment('是否生产商(1->是;0->否;)') TINYINT(2)"`
	LastModify                 time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	LegalName                  string    `json:"legal_name" xorm:"not null default '' comment('法人代表') VARCHAR(30)"`
	LicenseImg                 string    `json:"license_img" xorm:"comment('营业执照照片(json)') TEXT"`
	OperatorId                 int       `json:"operator_id" xorm:"not null default 0 comment('操作人ID') INT(11)"`
	OperatorName               string    `json:"operator_name" xorm:"not null default '' comment('操作人') VARCHAR(60)"`
	PayPeriodType              string    `json:"pay_period_type" xorm:"not null default '' comment('结算周期') VARCHAR(30)"`
	RegisterCapital            string    `json:"register_capital" xorm:"not null default '' comment('注册资本') VARCHAR(30)"`
	RegisterCapitalCurrency    string    `json:"register_capital_currency" xorm:"not null default '' comment('注册资本币种') VARCHAR(30)"`
	SalesmanMail               string    `json:"salesman_mail" xorm:"not null default '' comment('销售邮箱') VARCHAR(100)"`
	SalesmanMobile             string    `json:"salesman_mobile" xorm:"not null default '' comment('销售手机') VARCHAR(30)"`
	SalesmanName               string    `json:"salesman_name" xorm:"not null default '' comment('销售联系人') VARCHAR(30)"`
	SalesmanPhone              string    `json:"salesman_phone" xorm:"not null default '' comment('销售座机') VARCHAR(30)"`
	ServiceProvince            string    `json:"service_province" xorm:"not null comment('供应商服务省份(json)') TEXT"`
	SimpleName                 string    `json:"simple_name" xorm:"not null default '' comment('拼音简写') VARCHAR(30)"`
	Status                     int       `json:"status" xorm:"not null default 1 comment('收款账号是否停用(1->正常;2->停用)') TINYINT(2)"`
	SupplierAgentAuthorization string    `json:"supplier_agent_authorization" xorm:"not null comment('供应商代理授权') TEXT"`
	SupplierCode               string    `json:"supplier_code" xorm:"not null default '' comment('供应商编码') unique(UNIQUE_COMPANY_SUPPLIER_CODE) VARCHAR(30)"`
	SupplierName               string    `json:"supplier_name" xorm:"not null default '' comment('供应商名称') VARCHAR(30)"`
	SupplierTaxCode            string    `json:"supplier_tax_code" xorm:"not null default '' comment('供应商税号') VARCHAR(30)"`
	TaxCode                    string    `json:"tax_code" xorm:"not null default '' comment('统一社会信用代码(税号)') VARCHAR(200)"`
	Tips                       string    `json:"tips" xorm:"default '' comment('备注') VARCHAR(2000)"`
}
