package goxorm

type THosDictCountry struct {
	Abbr   string `json:"abbr" xorm:"not null default '' comment('缩写') VARCHAR(45)"`
	Cninfo string `json:"cninfo" xorm:"not null default '' VARCHAR(100)"`
	Id     int    `json:"id" xorm:"not null pk autoincr INT(10)"`
}
