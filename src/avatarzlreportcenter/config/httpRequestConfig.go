package config

type HttpHost struct {
	RpCourseList   string `json:"Host"`
	RpCourseDetail string `json:"Host"`
}

var HttpRequestConfig HttpHost