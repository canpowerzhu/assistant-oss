package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"kane.zhu/osstools/conf"
	"log"
	"os"
	"time"
)

type Ossconfig struct {
	Bucketname   string `json:"Bucketname"`
	Endpoint     string `json:"Endpoint"`
	Accesskey    string `json:"Accesskey"`
	Accesssecret string `json:"Accesssecret"`
	Osspath      string `json:"Osspath"`
}

var timeObj = time.Now()
var str = timeObj.Format("20060102030405")

func Uploadfile(filepath string, subpath string) {

	res := conf.OssJson()
	ossconfig := Ossconfig{}
	err := json.Unmarshal([]byte(res), &ossconfig)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := oss.New(ossconfig.Endpoint, ossconfig.Accesskey, ossconfig.Accesssecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucketName := ossconfig.Bucketname
	objectName := ossconfig.Osspath + "/" + str + "/" + subpath
	locaFilename := filepath
	content, err := os.Stat(locaFilename)

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	//判断是否进行分片上传 1,074,316,776 大于500MB 进行分片上传
	if content.Size() > 537158388 {
		fmt.Println("文件" + locaFilename + "分片上传")
		fileparts := content.Size()/537158388 + 1
		fmt.Println("分片数量", fileparts)
		chunks, err := oss.SplitFileByPartNum(locaFilename, int(fileparts))
		fd, err := os.Open(locaFilename)
		defer fd.Close()

		// 指定存储类型为标准存储。
		storageType := oss.ObjectStorageClass(oss.StorageStandard)

		// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。
		imur, err := bucket.InitiateMultipartUpload(objectName, storageType)
		// 步骤2：上传分片。
		var parts []oss.UploadPart
		for _, chunk := range chunks {
			fd.Seek(chunk.Offset, os.SEEK_SET)
			// 调用UploadPart方法上传每个分片。
			part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number, oss.Progress(&OssProgressListener{}))
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(-1)
			}
			parts = append(parts, part)
		}

		// 指定Object的读写权限为公共读，默认为继承Bucket的读写权限。
		objectAcl := oss.ObjectACL(oss.ACLPublicRead)

		// 步骤3：完成分片上传，指定文件读写权限为公共读。
		cmur, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		fmt.Println("cmur:", cmur)
	} else {
		fmt.Println("文件" + locaFilename + "采用简单上传")
		err = bucket.PutObjectFromFile(objectName, locaFilename, oss.Progress(&OssProgressListener{}))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

	}

}
