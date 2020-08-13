package sm

import (
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
)

var (
	default_uid = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
)

func NewKey(data string) {
	fmt.Println("开始生成秘钥！")
	pk, err := sm2.GenerateKey()
	if err != nil {
		fmt.Println("生成私钥时出现错误：", err.Error())
		return
	}

	pb, err := sm2.WritePrivateKeytoMem(pk, nil)

	if err != nil {
		fmt.Println("转换私钥时出现错误：", err.Error())
		return
	}

	fmt.Println("私钥：")
	fmt.Println(string(pb))

	pukb, err := sm2.WritePublicKeytoMem(&pk.PublicKey, nil)
	if err != nil {
		fmt.Println("转换公钥时出现错误：", err.Error())
		return
	}
	fmt.Println("公钥：")
	fmt.Println(string(pukb))

	fmt.Println("签名数据：", data)

	digest := []byte(data)

	r, s, err := sm2.Sm2Sign(pk, digest, default_uid)
	sign, err := sm2.SignDigitToSignData(r, s)

	if err != nil {
		fmt.Println("签名时出现错误：", err.Error())
		return
	}

	fmt.Println("签名结果：", base64.StdEncoding.EncodeToString(sign))
}
