package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
	"testing"
)

func TestBase58And64(t *testing.T) {
	hasher := sha256.New()
	hasher.Write([]byte("themoonstone"))
	bytes := hasher.Sum(nil)
	fmt.Printf("%x\n", bytes)

	hash160 := ripemd160.New()
	hash160.Write([]byte("themoonstone"))
	bytes160 := hash160.Sum(nil)
	fmt.Printf("%x\n", bytes160)

	// base64
	// 编码
	msg := "this is the eg of base64 encode"
	fmt.Println("源数据:", msg)
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("base64编码后:", encoded)
	// 解码
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("base64解码后:", string(decoded))

	encoded = base58.Encode([]byte(msg))
	fmt.Println("base58编码:", encoded)
	decoded, err = base58.Decode(encoded)
	fmt.Println("base58解码:", string(decoded))

}
