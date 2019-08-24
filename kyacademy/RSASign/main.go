package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

/*
 * 用公钥加密，私钥解密
 * 用私钥签名，公钥验证
 * 公钥是公开的，任何人都可以使用公钥，私钥非公开（保存好）
 */

// 公钥加密，私钥解密过程
func crypt() {
	// 创建私钥
	priv, _ := rsa.GenerateKey(rand.Reader, 1024)
	fmt.Println("系统产生的私钥：", priv)
	//  创建公钥
	pub := priv.PublicKey
	fmt.Println("系统产生的公钥：", pub)
	// 准备加密的明文
	src := []byte("hello cdd")
	// 公钥加密
	cipherTxt, _ := rsa.EncryptOAEP(md5.New(), rand.Reader, &pub, src, nil)
	fmt.Println("加密后的密文：", base64.StdEncoding.EncodeToString(cipherTxt))
	// 私钥解密
	plainTxt, _ := rsa.DecryptOAEP(md5.New(), rand.Reader, priv, cipherTxt, nil)
	fmt.Println("解密后的结果：", string(plainTxt))
}

// 私钥签名，公钥验证
func sign() {
	// 创建私钥
	priv, _ := rsa.GenerateKey(rand.Reader, 1024)
	// 创建公钥
	pub := priv.PublicKey
	// 明文
	src := []byte("block chain is very good.")
	// 对明文散列
	h := md5.New()
	h.Write(src)
	hashed := h.Sum(nil)
	// 私钥签名
	// 数字签名的作用是验证是否被篡改，A->B，当B收到数据时，验证是否由A传递的消息
	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sig, _ := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, &opts)
	fmt.Println("签名结果：", sig)
	// 公钥验证签名
	err := rsa.VerifyPSS(&pub, crypto.MD5, hashed, sig, &opts)
	if err != nil {
		fmt.Println("验证失败...")
	}
	fmt.Println("验证成功...")
}

func main() {
	crypt()
	sign()
}
