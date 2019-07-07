package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"simplebtc/crypto/symmetrical_encryption"
	"testing"
)

// AES加密
func EncryptAes(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	text := symmetrical_encryption.PaddingText(src, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(text))
	mode.CryptBlocks(dst, text)

	return dst
}

// AES解密
func DecryptAes(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	mode.CryptBlocks(dst, src)

	return symmetrical_encryption.UnpaddingText(dst)
}

func TestAes(t *testing.T) {
	src := []byte("hello world")
	key := []byte("1234567887654321abcdefghabcdefgh") // 16 24 32
	encrypted := EncryptAes(src, key)
	fmt.Println("加密后密文：", hex.EncodeToString(encrypted))
	decrypted := DecryptAes(encrypted, key)
	fmt.Println("解密后明文：", string(decrypted))
}
