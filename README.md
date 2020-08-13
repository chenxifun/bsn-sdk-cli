# bsn-sdk-cli

1. 生成sm2、secp25r1、secp256k1 格式的密钥，并且生成测试的签名数据
2. 调用网关的 cli 程序

```
bsnCli secret --algo sm2 --data abc123
秘钥类型： sm2
签名数据： abc123
开始生成秘钥！
私钥：
-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgBrli71yGjPjSDG9C
Sh7lwNlGoSzVnTkWwrCc15CkBWugCgYIKoEcz1UBgi2hRANCAAQr34+G4tmB1fxx
/zhdTz8D++/TcHghExKrqSLOPtEAyupD729SXYlZb39HXRk7tZvgWiIhTmKpXOPR
Afg6oS8x
-----END PRIVATE KEY-----

公钥：
-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEK9+PhuLZgdX8cf84XU8/A/vv03B4
IRMSq6kizj7RAMrqQ+9vUl2JWW9/R10ZO7Wb4FoiIU5iqVzj0QH4OqEvMQ==
-----END PUBLIC KEY-----

签名数据： abc123
签名结果： MEUCIQDGVOJ6LKJCkcdhnbpn1mNF+cLXZgb6AM2EH9e8uA+XMwIgLhZq08uovKv9JfnVF5ZR4mx729Zj6hd4gUn0Lf7hNSU=

```