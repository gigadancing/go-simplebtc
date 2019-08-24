package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"os"
)

/*
 * 接收网络传送过来的数据
 */

func Recv() []byte {
	listener, _ := net.Listen("tcp", ":10086")
	defer listener.Close()
	// 监听端口，并接收
	for {
		conn, _ := listener.Accept()
		buff := make([]byte, 2048)
		for {
			size, _ := conn.Read(buff)
			return buff[:size]
		}
	}
}

func main() {
	// 接收数据
	data := Recv()

	// 拆分数据
	length := int(data[0])
	fmt.Println("数据长度：", length)
	plainTxt := data[1 : length+1]
	fmt.Println("接收的明文：", string(plainTxt))
	sig := data[length+1:]

	// 读取公钥
	f, err := os.Open(".\\kyacademy\\RSASign\\public.pem")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	buff := make([]byte, info.Size())
	_, err = f.Read(buff)
	if err != nil {
		panic(err)
	}

	// 公钥验签
	block, _ := pem.Decode(buff)
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 将接收的数据散列
	h := md5.New()
	h.Write(plainTxt)
	hashed := h.Sum(nil)

	// 验签
	err = rsa.VerifyPSS(pubKey, crypto.MD5, hashed, sig, nil)
	if err != nil {
		fmt.Println("验证失败, err：", err)
		return
	}

	fmt.Println("验证成功...")
}
