package r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto/secp256r1"
	utils "github.com/chenxifun/bsn-sdk-cli/common"
)

func NewKey(data string) {
	fmt.Println("开始生成秘钥！")

	curve := elliptic.P256()
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		fmt.Println("生成私钥时出现错误：", err.Error())
		return
	}

	rawKey, err := utils.PrivateKeyToPEM(privKey, nil)

	if err != nil {
		fmt.Println("转换私钥时出现错误：", err.Error())
		return
	}
	fmt.Println("私钥：")
	fmt.Println(string(rawKey))

	puk, err := utils.PublicKeyToPEM(privKey.Public(), nil)
	if err != nil {
		fmt.Println("转换公钥时出现错误：", err.Error())
		return
	}
	fmt.Println("公钥：")
	fmt.Println(string(puk))

	fmt.Println("签名数据：", data)

	h := sha256.New()
	h.Write([]byte(data))
	digest := h.Sum(nil)

	prk, _ := secp256r1.LoadPrivateKey(string(rawKey))
	sign, err := secp256r1.SignECDSA(prk, digest)
	if err != nil {
		fmt.Println("签名时出现错误：", err.Error())
		return
	}

	fmt.Println("签名结果：", base64.StdEncoding.EncodeToString(sign))
}
