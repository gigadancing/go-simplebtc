package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net"
)

// 去码
func PKCS7Unpadding(src []byte) []byte {
	length := len(src)
	tail := src[length-1]
	return src[:length-int(tail)]
}

// 解密
func AESDecrypt(cipherTxt []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	mode := cipher.NewCBCDecrypter(block, key)
	// 创建明文缓存
	buff := make([]byte, len(cipherTxt))
	// 解密
	mode.CryptBlocks(buff, cipherTxt)
	// 去码
	return PKCS7Unpadding(buff)
}

/*
 * TCP是通过服务端监听端口，客户端发送数据，实现网络中数据传输
 */

func main() {
	// 监听电脑中的端口，1024-65535
	listener, _ := net.Listen("tcp", ":10086")

	// 延迟关闭
	defer listener.Close()

	// 通过循环等待客户端连接
	for {
		// 只有客户端连接成功才会往下执行
		conn, _ := listener.Accept()
		// 创建缓存，存放客户端发送的数据
		data := make([]byte, 1024)

		for {
			//接收客户端发送的数据，n是数据的大小
			n, _ := conn.Read(data)
			// data[:n]接收到的密文
			// 解密
			fmt.Println("密文为:", data[:n])
			fmt.Println("明文为:", string(AESDecrypt(data[:n], []byte("2113648236482368"))))
			break
		}

	}

}
