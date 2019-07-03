package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"testing"
)

// DES算法把64位的明文输入块变为64位的密文输出块,它所使用的密钥也是64位（实际用到了56位，
// 第8、16、24、32、40、48、56、64位是校验位， 使得每个密钥都有奇数个1）
//

// 填充最后一个分组
// src：待填充的数据 blockSize：分组大小
func PaddingText(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize                   // 最后一个分组需要填充的字节数
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding) // 填充的数据
	text := append(src, paddingText...)                         // 将填充的数据和源数据进行拼接
	return text
}

// 删除填充数据
func UnpaddingText(src []byte) []byte {
	l := len(src)
	num := int(src[l-1])
	text := src[:l-num]
	return text
}

// 使用des算法进行家吗
// src：待加密铭文 key：秘钥
func EncryptDes(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	l := block.BlockSize()                    // 块的大小
	text := PaddingText(src, l)               // 对源数据填充
	iv := []byte("12345678")                  // 初始化向量
	mode := cipher.NewCBCEncrypter(block, iv) // 创建CBC加密模式
	encrypted := make([]byte, len(text))
	mode.CryptBlocks(encrypted, text)
	return encrypted
}

// 使用des解密
// src：密文 key：秘钥
func DecryptDes(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	mode := cipher.NewCBCDecrypter(block, iv) // 创建CBC解密模式
	decrypted := make([]byte, len(src))
	mode.CryptBlocks(decrypted, src)

	return UnpaddingText(decrypted)
}

func TestEncryptDes(t *testing.T) {
	src := []byte("hello world")
	key := []byte("87654321")
	encrypted := EncryptDes(src, key)
	fmt.Println("encrypted:", encrypted)
	fmt.Println("encrypted:", hex.EncodeToString(encrypted))
	fmt.Println("---------------------------------------------")
	decrypted := DecryptDes(encrypted, key)
	fmt.Println("decrypted:", string(decrypted))
}
