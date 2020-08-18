# bsn-sdk-cli


## 1. 生成sm2、secp25r1、secp256k1 格式的密钥，并且生成测试的签名数据

> go版本需要1.13及以上
> go get github.com/chenxifun/bsn-sdk-cli
> 代理`go env -w GOPROXY=https://goproxy.cn,direct`
```
bsn-sdk-cli --algo sm --data 123456
参数
    --algo  算法类型可选为【sm2、secp25r1、secp256k1】
       sm2              适用于国密签名算法的密钥
       secp256r1        适用于r1签名算法的密钥
       secp256k1        适用于k1签名算法的密钥
    --data  待签名的数据

```
示例：
```
bsn-sdk-cli secret --algo sm2 --data abc123

```

> openssl 生成密钥参考[这里](http://kb.bsnbase.com/webdoc/real/Pub4028813e711a7c39017176a7dc154026.html?STATE=0&OPERATE=3)