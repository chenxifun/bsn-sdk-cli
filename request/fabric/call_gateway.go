package fabric

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	path2 "path"
)

var (
	privKey = ""
)

func goRequest() {
	if api == "" {
		fmt.Println("api 网关地址不能为空")
		return
	} else {
		fmt.Println("网关地址：", api)
	}
	if appCode == "" {
		fmt.Println("appCode 不能为空")
		return
	} else {
		fmt.Println("应用编号：", appCode)
	}

	if userCode == "" {
		fmt.Println("userCode 不能为空")
		return
	} else {
		fmt.Println("用户编号：", userCode)
	}

	keyByte, err := getKey(key)
	if err != nil {
		fmt.Println("读取私钥时出现错误：", err.Error())
		return
	}
	privKey = string(keyByte)

	if !isTrust {

		fmt.Println("创建MSP文件夹：", mspDir)
		err = os.MkdirAll(path2.Dir(mspDir), 0700)
		if err != nil {
			fmt.Println("创建MSP文件夹异常：", err.Error())
			return
		}
	}

}

func getKey(key string) ([]byte, error) {
	if key == "" {
		return nil, errors.New("私钥地址不能为空")
	} else {
		fmt.Println("私钥路径：", key)
	}

	return ioutil.ReadFile(key)

}
