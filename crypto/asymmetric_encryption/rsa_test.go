package asymmetric_encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

//
func RsaGenKey(bits int) error {
	// rand.Reader是一个全局共享的密码随机生成器
	privKey, err := rsa.GenerateKey(rand.Reader, bits) // 私钥里面包含公钥
	if err != nil {
		panic(err)
	}

	// x509是通用的证书格式：序列号 签名算法 颁发者 有效时间 持有者 公钥
	// PKCS：RSA实验室与其他安全系统开发商制定的一系列标准，有1-15个标准
	priStream := x509.MarshalPKCS1PrivateKey(privKey)

	// 将私钥设置到pem格式的块中
	// pem是一种证书或私钥的格式：
	// ---------BEGIN RSA Private Key ------------------
	// ......
	// ---------END RSA Private Key---------------------
	block := pem.Block{
		Type:  "RSA Private Key",
		Bytes: priStream,
	}
	privFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privFile.Close()

	// 将块编码到文件
	if err := pem.Encode(privFile, &block); err != nil {
		panic(err)
	}

	pubKey := privKey.PublicKey
	pubStream := x509.MarshalPKCS1PublicKey(&pubKey)
	block = pem.Block{
		Type:  "RSA Public Key",
		Bytes: pubStream,
	}
	pubFile, err := os.Create("public.pem")
	defer pubFile.Close()
	if err != nil {
		panic(err)
	}
	if err := pem.Encode(pubFile, &block); err != nil {
		panic(err)
	}

	return nil
}

// 用公钥加密
func RsaPubKeyEncrypt(src []byte, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat() // 获取文件信息
	if err != nil {
		panic(err)
	}

	data := make([]byte, info.Size())
	if _, err := file.Read(data); err != nil { // 读取公钥
		panic(err)
	}

	block, _ := pem.Decode(data)                         // 解码
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes) // 解析出公钥
	if err != nil {
		panic(err)
	}

	msg, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src) // 使用公钥加密
	if err != nil {
		panic(err)
	}

	return msg, nil
}

func RsaPrivateKeyDecrypt(src []byte, filename string) ([]byte, error) {
	msg := []byte("")
	file, err := os.Open(filename)
	if err != nil {
		return msg, err
	}
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}

	buff := make([]byte, info.Size())
	_, err = file.Read(buff)
	if err != nil {
		return msg, err
	}

	block, _ := pem.Decode(buff)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	msg, err = rsa.DecryptPKCS1v15(rand.Reader, privKey, src)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

func TestRsaEncrypt(t *testing.T) {
	src := []byte("hello world")
	fmt.Println("明文：", string(src))
	encrypted, err := RsaPubKeyEncrypt(src, "public.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后密文：", hex.EncodeToString(encrypted))
	decrypted, err := RsaPrivateKeyDecrypt(encrypted, "private.pem")
	fmt.Println("解密后明文：", string(decrypted))
}

func TestRsaGenKey(t *testing.T) {
	if err := RsaGenKey(256); err != nil {
		panic(err)
	}
}
