package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

/*
 * 利用秘钥通过DES算法实现明文加密
 * 利用秘钥通过DES算法实现密文解密
 *
 * 在加解密之前，首先需要补码和去码
 *
 *
 */

// 补码
func PKCS5Padding(originData []byte, blockSize int) []byte {
	padding := blockSize - len(originData)%8
	paddingTxt := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(originData, paddingTxt...)
}

// 去码
func PKCS5Unpadding(cipherTxt []byte) []byte {
	length := len(cipherTxt)
	unpadding := int(cipherTxt[length-1])

	return cipherTxt[:length-unpadding]
}

// DES加密
func DesEncrypt(src []byte, key []byte) []byte {
	// 校验秘钥是否合法
	if len(key) != 8 {
		return nil
	}
	// DES加密算法，秘钥的长度必须为8位
	block, _ := des.NewCipher(key)
	// 补码
	paddingTxt := PKCS5Padding(src, block.BlockSize())
	// 设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, key)
	// 加密
	crypted := make([]byte, len(paddingTxt)) // 存放加密后的密文
	blockMode.CryptBlocks(crypted, paddingTxt)

	return crypted
}

//
func DesDecrypt(cipherTxt []byte, key []byte) []byte {
	if len(key) != 8 {
		return nil
	}

	block, _ := des.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key)
	src := make([]byte, len(cipherTxt))
	blockMode.CryptBlocks(src, cipherTxt)
	return PKCS5Unpadding(src)
}

func main() {
	key := "hellowor"
	src := "chendada"
	cipherTxt := DesEncrypt([]byte(src), []byte(key))
	fmt.Println("src:", []byte(src))

	fmt.Println("===============DES Encrypt===============")
	fmt.Println("binary:", cipherTxt)
	fmt.Println("hex:", hex.EncodeToString(cipherTxt))
	fmt.Println("base64:", base64.StdEncoding.EncodeToString(cipherTxt))

	fmt.Println("===============DES Decrypt===============")
	data := DesDecrypt(cipherTxt, []byte(key))
	fmt.Println("plain txs:", data)

	// 对称加密中，加密与解密是互逆的
	// DES加密中，秘钥长度必须为8字节；3DES秘钥长度必须为24字节；AES秘钥长度必须为16、24、32字节
}
