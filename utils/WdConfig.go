package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DBConfig struct {
	user   string `json:"user"`
	passwd string `json:"passwd"`
	ip     string `json:"ip"`
	port   int16  `json:"port"`
	DB     string `json:"DB"`
}

// func (r *DBConfig) MarshalJSON() ([]byte, error) {
// 	arr := []interface{}{r.user, r.passwd, r.ip, r.port, r.DB}
// 	return json.Marshal(arr)
// }

// func (r *DBConfig) UnmarshalJSON(bs []byte) error {
// 	arr := []interface{}{}
// 	json.Unmarshal(bs, &arr)
// 	// TODO: add error handling here.
// 	r.user = arr[0].(string)
// 	r.passwd = arr[1].(string)
// 	r.ip = arr[2].(string)
// 	r.port = arr[3].(int16)
// 	r.DB = arr[4].(string)
// 	return nil
// }

var cMysqlConfig *WdConfig
var cServerConfig *WdConfig

// for k, v := range conf.JsonData {
// 	fmt.Printf("key[%s] value[%s]\n", k, v)
// }

func ServerConfig() *WdConfig {
	if nil == cServerConfig {
		cServerConfig = WdConfigNew()
		return cServerConfig
	}
	return cServerConfig
}

func MysqlConfig() *WdConfig {
	if nil == cMysqlConfig {
		cMysqlConfig = WdConfigNew()
		return cMysqlConfig
	}
	return cMysqlConfig
}

type WdConfig struct {
	// file     *os.File
	JsonData map[string]interface{}
	// cConfig DBConfig
}

func WdConfigNew() *WdConfig {
	newConfig := new(WdConfig)
	// newConfig.cConfig = DBConfig{}
	return newConfig
}

func (r *WdConfig) Init(configPath string) error {
	fp, err := os.Open(configPath)
	if err != nil {
		return err
	}

	defer fp.Close()
	byteValue, _ := ioutil.ReadAll(fp)
	json.Unmarshal([]byte(byteValue), &r.JsonData)
	// fmt.Println(JsonData)

	// for k, v := range JsonData {
	// 	fmt.Printf("key[%s] value[%s]\n", k, v)
	// 	myMap := v.(map[string]interface{})
	// 	fmt.Println(myMap["user"].(string))
	// 	fmt.Printf("[%s]\n", reflect.TypeOf(v))
	// }
	return nil
}
