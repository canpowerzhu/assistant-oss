package main

import (
	"flag"
	"fmt"
	"kane.zhu/osstools/utils"
	"os"
	"path/filepath"
)

var (
	uploadpath   string
	enpoint      string
	accesskey    string
	accesssecret string
	osspath      string
	bucketname   string
)

func main() {

	//flag参数
	uploadpath := flag.String("uploadpath", "E:\\迅雷下载\\22.iso", "请输入需要上传的文件或者文件夹的绝对路径:")
	flag.Parse()
	//uploadpath := utils.Localuploadpath()

	if !utils.Exists(*uploadpath) {
		fmt.Println("需要上传的文件或者文件夹" + *uploadpath + "不存在")
		os.Exit(-1)

	}

	if utils.IsDir(*uploadpath) {
		fmt.Println("需要操作的是一个文件夹")
		utils.GetALlfile(*uploadpath, *uploadpath)
	}

	if utils.IsFile(*uploadpath) {
		//这里的uploadpath是上传文件的本地路径  E:\deploy\layui\layui.all.js
		//fileName 是远端oss的相对路径 layui\layui.all.js
		paths, fileName := filepath.Split(*uploadpath)
		fmt.Println("需要操作的是一个文件" + fileName + "; 路径是：" + paths)
		utils.Uploadfile(*uploadpath, fileName)
	}

}
