package utils

import (
	"io/ioutil"
	"strings"
)

func GetALlfile(basepath string, folderpath string) error {

	rd, err := ioutil.ReadDir(folderpath)
	for _, fi := range rd {

		//文件夹判断是否继续递归
		if fi.IsDir() {
			GetALlfile(basepath, folderpath+"\\"+fi.Name())
		} else {
			//本地上传文件的绝对路径
			localfilepath := folderpath + "\\" + fi.Name()

			//定义获取osss上传的相对路径 减去basepath的长度 获取后面字符
			subpath := localfilepath[len(basepath):]
			subpath = strings.Replace(subpath, "\\", "/", -1)

			// 需要用到的是subpath 和本地路径 localfilepath

			Uploadfile(localfilepath, subpath)

		}

	}
	return err
}
