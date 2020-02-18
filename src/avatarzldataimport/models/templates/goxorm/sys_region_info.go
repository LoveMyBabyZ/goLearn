package goxorm

type SysRegionInfo struct {
	Areaname  string `json:"areaname" xorm:"VARCHAR(255)"`
	Code      int    `json:"code" xorm:"not null INT(11)"`
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Lat       string `json:"lat" xorm:"VARCHAR(127)"`
	Level     int    `json:"level" xorm:"INT(11)"`
	Lng       string `json:"lng" xorm:"VARCHAR(127)"`
	Parent    int    `json:"parent" xorm:"INT(11)"`
	Position  string `json:"position" xorm:"VARCHAR(255)"`
	Shortname string `json:"shortname" xorm:"VARCHAR(255)"`
	Sort      int    `json:"sort" xorm:"INT(10)"`
}
