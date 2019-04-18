package util

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"log"
)

// 整数转哈希
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panicf("int to []byte failed: %v\n", err)
	}
	return buff.Bytes()
}

// 哈希转字符转
func HexToString(data []byte) string {
	return hex.EncodeToString(data)
}

// json转数组
func JsonToSlice(jstr string) []string {
	var strSli []string
	if err := json.Unmarshal([]byte(jstr), &strSli); err != nil {
		log.Panicf("json to string err: %v\n", err)
	}
	return strSli
}
