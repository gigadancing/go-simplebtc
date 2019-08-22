package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net"
)

/*
 * 对称加密中，DES，3DES，AES
 */
// AES对称加密，需要首先对明文补码
// PKCS5的分组是以8为单位
// PKCS7的分组长度为1-255
// 补码
func PKCS7Padding(src []byte, blockSize int) []byte {
	pad := blockSize - len(src)%blockSize
	padArr := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(src, padArr...)
}

// AES加密
func AESEncrypt(src []byte, key []byte) []byte {
	// 校验秘钥是否合法
	block, _ := aes.NewCipher(key)
	// 对明文进行补码
	padtxt := PKCS7Padding(src, block.BlockSize())
	// 设置加密模式
	mode := cipher.NewCBCEncrypter(block, key)
	// 创建密文缓冲区
	cypted := make([]byte, len(padtxt))
	// 加密
	mode.CryptBlocks(cypted, padtxt)
	return cypted
}

/*
 * 客户端，向服务端发送数据
 * 广域网 UDP+NAT
 */

func main() {
	crypted := AESEncrypt([]byte("helloworld"), []byte("2113648236482368"))
	// 构建服务器连接
	conn, _ := net.ResolveTCPAddr("tcp", ":10086")
	// 连接拨号
	n, _ := net.DialTCP("tcp", nil, conn)
	// 发送数据
	_, _ = n.Write(crypted)
	fmt.Println("send over")
}
