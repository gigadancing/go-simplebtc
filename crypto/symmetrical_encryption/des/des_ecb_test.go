package des

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
	"simplebtc/crypto/symmetrical_encryption"
	"testing"
)

// ECB模式加密
func DesEcbEncrypt(src, key []byte) []byte {
	cipherBlock, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	l := cipherBlock.BlockSize()                       // 8字节
	data := symmetrical_encryption.ZeroPadding(src, l) // 尾部填充0
	out := make([]byte, len(data))                     // 加密后的密文
	dst := out
	for len(data) > 0 {
		cipherBlock.Encrypt(dst, data[:l]) // 每次加密8字节
		data = data[l:]                    // 向后移动8字节
		dst = dst[l:]
	}

	return out
}

// ECB模式解密
func DesEcbDecrypt(src, key []byte) []byte {
	cipherBlock, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	l := cipherBlock.BlockSize()
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		cipherBlock.Decrypt(dst, src[:l])
		src = src[l:]
		dst = dst[l:]
	}

	return symmetrical_encryption.ZeroUnpadding(out)
}

func TestDesEcbEncrypt(t *testing.T) {
	key := []byte("12345678")
	data := []byte("hello world")
	encrypted := DesEcbEncrypt(data, key)
	fmt.Println("加密后密文：", hex.EncodeToString(encrypted))
	src := DesEcbDecrypt(encrypted, key)
	fmt.Println("解密后：", string(src))
}
