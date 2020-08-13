package k1

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto/secp256k1"
)

func NewKey(data string) {
	fmt.Println("开始生成秘钥！")
	pk, err := secp256k1.NewSecp256k1Key()
	if err != nil {
		fmt.Println("生成私钥时出现错误：", err.Error())
		return
	}

	pb, err := secp256k1.PrivateKeyToPEM(pk)

	if err != nil {
		fmt.Println("转换私钥时出现错误：", err.Error())
		return
	}

	fmt.Println("私钥：")
	fmt.Println(string(pb))

	pukb, err := secp256k1.PublicKeyToPEM(&pk.PublicKey)
	if err != nil {
		fmt.Println("转换公钥时出现错误：", err.Error())
		return
	}
	fmt.Println("公钥：")
	fmt.Println(string(pukb))

	fmt.Println("签名数据：", data)

	//哈希
	h := sha256.New()
	h.Write([]byte(data))
	digest := h.Sum(nil)

	sign, err := secp256k1.SignECDSA(pk, digest)
	if err != nil {
		fmt.Println("签名时出现错误：", err.Error())
		return
	}

	fmt.Println("签名结果：", base64.StdEncoding.EncodeToString(sign))
}
