package goxorm

type THosDictUnit struct {
	CompanyId int    `json:"company_id" xorm:"not null default 0 comment('公司ID') unique(UNIQUE_NAME) INT(11)"`
	Id        int    `json:"id" xorm:"not null pk autoincr comment('小暖系统的id') INT(11)"`
	Name      string `json:"name" xorm:"not null default '' comment('单位名称') unique(UNIQUE_NAME) VARCHAR(200)"`
}
