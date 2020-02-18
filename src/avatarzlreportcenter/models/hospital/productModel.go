package hospital

type Product struct {
	Id          int    `gorm:"column:id; type:int(11); primary_key;AUTO_INCREMENT; not null" json:"id"`
	Img         string `gorm:"column:img; type:varchar(20); not null;default '' " json:"img"`
	ProductCode string `gorm:"column:product_code; type:varchar(255); not null " json:"product_code"`
	ProductName string `gorm:"column:product_name; type:varchar(255); not null " json:"product_name"`
	ProductType string `gorm:"column:product_name; type:varchar(10); not null " json:"product_type"`
	CompanyId   int    `gorm:"column:company_id; type:int(11); not null " json:"company_id"`
	Category    string `gorm:"column:category; type:int(11); not null " json:"category"`
	Ingredient  string `gorm:"column:category; type:int(11); not null " json:"category"`
}
