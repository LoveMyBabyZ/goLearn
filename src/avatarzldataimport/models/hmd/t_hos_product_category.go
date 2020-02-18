package hmd

import (
	"avatarzldataimport/databases/db"
	"fmt"
	"time"
)

type THosProductCategory struct {
	CategoryId   string    `json:"category_id" xorm:"not null comment('目录ID') unique(category_id_UNIQUE) VARCHAR(50)"`
	CategoryName string    `json:"category_name" xorm:"not null comment('目录名称') VARCHAR(20)"`
	CompanyId    string    `json:"company_id" xorm:"not null default '' comment('公司编码') unique(category_id_UNIQUE) VARCHAR(50)"`
	CompanyName  string    `json:"company_name" xorm:"default '' comment('公司名称') VARCHAR(50)"`
	CreateTime   time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	IdPath       string    `json:"id_path" xorm:"VARCHAR(100)"`
	IsDeleted    int       `json:"is_deleted" xorm:"not null default 0 comment('是否已删除，1是，0否') TINYINT(3)"`
	NamePath     string    `json:"name_path" xorm:"VARCHAR(200)"`
	ParentId     string    `json:"parent_id" xorm:"not null default '0' comment('上级目录ID') VARCHAR(50)"`
	ProductType  string    `json:"product_type" xorm:"default '' comment('产品类型') unique(category_id_UNIQUE) VARCHAR(10)"`
	UpdateTime   time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// 获取目录
func GetCategoryPath() (arr map[string]string, err error) {
	Category := make([]THosProductCategory, 0)
	err = db.DB.Cols("name_path","category_id").Find(&Category)
	if err != nil {
		fmt.Println(err)
	}
	arr = make(map[string]string)
	for _, v := range Category {
		arr[v.NamePath] = v.CategoryId
	}
	return
}
