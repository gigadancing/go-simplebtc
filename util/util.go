package util

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"log"
)

//
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panicf("int to []byte failed: %v\n", err)
	}
	return buff.Bytes()
}

//
func HexToString(data []byte) string {
	return hex.EncodeToString(data)
}
