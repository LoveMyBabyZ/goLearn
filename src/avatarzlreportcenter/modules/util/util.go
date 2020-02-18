package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"reflect"
	"strings"
)

//encode struct to  json
func StructToJson(structModel interface{}) ([]byte, error) {
	json, err := json.Marshal(structModel)
	return json, err
}

//encode struct to json string
func StructToJsonStr(structModel interface{}) (string, error) {
	json, err := json.Marshal(structModel)
	return string(json), err
}

//decode json to map
func JsonToMap(jsonData []byte) (mapData map[string]interface{}, err error) {
	err = json.Unmarshal(jsonData, &mapData)
	return mapData, err
}

//string md5
func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

//get new uuid
func GetNewUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}

func EncodeB64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return string(base64Text)
}

// 获取结构体key
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fieldNum := t.NumField()
	fmt.Println(t)
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

func GetTagName(structName interface{}) []string  {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string,0,fieldNum)
	for i:= 0;i<fieldNum;i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result,tagName)
	}
	return result
}
