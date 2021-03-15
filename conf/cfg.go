package conf

import (
	"encoding/json"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	//"os"
)

//
//// 返回json结构体
//type Repay struct {
//	Code  uint64 `json:"code"`
//	Message string `json:"message"`
//}
//
//var     JsonRepay Repay
//
////结构体 首字母必须大写
//type Cfg struct {
//	Bucketname string
//	Endpoint string
//	Accesskey string
//	Accesssecret string
//	Osspath string
//}
//
//var     Conf Cfg
//
//func init(){
//	//open config yaml
//	file, err := os.Open("conf.yaml")
//	if err != nil{
//		panic(err)
//	}
//	bytes, err := ioutil.ReadAll(file)
//	Conf = Cfg{}
//	err = yaml.Unmarshal(bytes, &Conf)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(Conf)
//
//}

// 返回oss json结构体
type ossjson struct {
	Bucketname   string `json:"Bucketname"`
	Endpoint     string `json:"Endpoint"`
	Accesskey    string `json:"Accesskey"`
	Accesssecret string `json:"Accesssecret"`
	Osspath      string `json:"Osspath"`
}

var ENDPOINT = "oss-accelerate.aliyuncs.com"
var Accesskey = ""
var Accesssecret = ""
var Osspath = "osstool"
var Bucketname = "moppowar"
var Ossjson ossjson

func OssJson() string {
	p := ossjson{}
	p.Endpoint = ENDPOINT
	p.Accesskey = Accesskey
	p.Accesssecret = Accesssecret
	p.Osspath = Osspath
	p.Bucketname = Bucketname

	data, _ := json.Marshal(p)
	return string(data)

}
