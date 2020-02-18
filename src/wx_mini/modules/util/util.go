package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
)

//encode struct to  json
func StructToJson(structModel interface{}) ([]byte, error) {
	json, err := json.Marshal(structModel)
	return json, err
}

//encode struct to json string
func StructTojsonStr(structModel interface{}) (string, error) {
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


