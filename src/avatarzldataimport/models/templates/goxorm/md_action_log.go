package goxorm

import (
	"time"
)

type MdActionLog struct {
	Action             string    `json:"action" xorm:"not null default '' comment('请求操作名称') VARCHAR(70)"`
	ClientEdtion       string    `json:"client_edtion" xorm:"not null default '' comment('小暖客户端发行版标识') VARCHAR(20)"`
	Controller         string    `json:"controller" xorm:"not null default '' comment('请求控制器名称') VARCHAR(50)"`
	Cookies            string    `json:"cookies" xorm:"not null default '' comment('cookie数据 JSON格式存储') VARCHAR(500)"`
	CreateTime         time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('日志记录创建时间') DATETIME"`
	CurlDuration       float32   `json:"curl_duration" xorm:"not null default 0 comment('cURL远程请求总时长') FLOAT"`
	DbDuration         float32   `json:"db_duration" xorm:"not null default 0 comment('数据库查询总时长') FLOAT"`
	ExceptionMessage   string    `json:"exception_message" xorm:"not null default '' comment('异常信息') VARCHAR(500)"`
	ExceptionName      string    `json:"exception_name" xorm:"not null default '' comment('服务端异常名称') VARCHAR(50)"`
	HostName           string    `json:"host_name" xorm:"not null default '' comment('服务器主机域名') VARCHAR(50)"`
	HostPort           int       `json:"host_port" xorm:"not null default 0 comment('服务器主机端口号') SMALLINT(5)"`
	HttpReferer        string    `json:"http_referer" xorm:"not null default '' comment('http来源地址') VARCHAR(255)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Method             string    `json:"method" xorm:"not null default '' comment('HTTP请求方法') VARCHAR(10)"`
	Orgid              int       `json:"orgid" xorm:"not null default 0 comment('机构(医院)id') INT(11)"`
	Params             string    `json:"params" xorm:"not null default '' comment('请求参数 JSON格式存储') VARCHAR(500)"`
	RequestContentType string    `json:"request_content_type" xorm:"not null default '' comment('HTTP请求的数据格式') VARCHAR(255)"`
	RequestDuration    float32   `json:"request_duration" xorm:"not null default 0 comment('服务端处理时长') FLOAT"`
	Sessions           string    `json:"sessions" xorm:"not null default '' comment('服务端SESSION数据 JSON格式存储') VARCHAR(500)"`
	UserHost           string    `json:"user_host" xorm:"not null default '' VARCHAR(50)"`
	UserId             int       `json:"user_id" xorm:"not null default 0 comment('用户id') INT(11)"`
	UserIp             string    `json:"user_ip" xorm:"not null default '' comment('客户端用户ip地址 未来可以考虑使用int保存') VARCHAR(15)"`
	UserName           string    `json:"user_name" xorm:"not null default '' comment('用户名称') VARCHAR(30)"`
}
