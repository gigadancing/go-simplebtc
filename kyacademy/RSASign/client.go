package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"os"
)

/*
 * 发送消息，需要对数据进行签名
 */

func SignData(plainTxt []byte) []byte {
	//
	h := md5.New()
	h.Write(plainTxt)
	hashed := h.Sum(nil)

	// 读取私钥
	f, err := os.Open(".\\kyacademy\\RSASign\\private.pem")
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

	// 将字节数组转换成私钥类型
	block, _ := pem.Decode(buff)
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	// 签名
	opts := &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.MD5,
	}
	sig, err := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, opts)
	if err != nil {
		panic(err)
	}

	return sig
}

// 通过TCP将数据和签名结果发送给接收端
func Send(data []byte) {
	addr, _ := net.ResolveTCPAddr("tcp4", ":10086")
	conn, _ := net.DialTCP("tcp", nil, addr)
	// 将数据通过tcp协议发送给接收方
	_, _ = conn.Write(data)
	fmt.Println("发送结束...")
}

func main() {
	// 源数据
	src := "cdd block chain"
	length := len(src)
	// 获得签名结果
	sig := SignData([]byte(src))
	// 发送的数据
	data := make([]byte, length+len(sig)+1)
	data[0] = byte(length)
	copy(data[1:length+1], src)
	copy(data[length+1:], sig)
	// data数据由两部分组成，源数据长度+源数据+签名结果
	Send(data)
	//fmt.Println(sig)
}
