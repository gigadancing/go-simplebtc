package crypto

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
	"testing"
)

// DSA(Digital Signature Algorithm)，数字签名算法
// 作用：
//     保证数据完整性
//     确保数据的来源
func TestDSA(t *testing.T) {
	var (
		param dsa.Parameters // 秘钥的预参数
		priv  dsa.PrivateKey
	)
	// 随机设置合法的参数到params
	// 根据第三个参数决定L和N的长度，长度越长，加密强度越高
	if err := dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160); err != nil {
		panic(err)
	}

	// 生成私钥
	priv.Parameters = param
	if err := dsa.GenerateKey(&priv, rand.Reader); err != nil {
		panic(err)
	}

	// 签名
	msg := []byte("hello world")
	r, s, err := dsa.Sign(rand.Reader, &priv, msg)
	if err != nil {
		panic(err)
	}

	pub := priv.PublicKey // 通过私钥获取公钥
	msg1 := []byte("hello world")
	if !dsa.Verify(&pub, msg1, r, s) {
		fmt.Println("验证失败")
	}

	fmt.Println("验证通过")
}
