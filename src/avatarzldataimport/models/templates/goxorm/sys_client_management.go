package goxorm

type SysClientManagement struct {
	Id     int    `json:"id" xorm:"not null pk autoincr comment('ID') INT(11)"`
	ImgUrl string `json:"img_url" xorm:"comment('医院logo') VARCHAR(500)"`
	OrgId  int    `json:"org_id" xorm:"comment('机构ID') INT(11)"`
	Remark string `json:"remark" xorm:"comment('医院备注') VARCHAR(500)"`
}
