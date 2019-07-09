package crypto

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	// 生成私钥
	priv, _ := rsa.GenerateKey(rand.Reader, 1024)
	// 消息
	msg := []byte("hello world")
	// 对消息进行散列处理
	hash := md5.New()
	hash.Write(msg)
	hashed := hash.Sum(nil)
	// 签名
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sig, err := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(sig))

	// 获取公钥
	pub := &priv.PublicKey
	err = rsa.VerifyPSS(pub, crypto.MD5, hashed, sig, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("验证通过")
}
