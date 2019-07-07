package des

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"simplebtc/crypto/symmetrical_encryption"
	"testing"
)

// DES算法把64位的明文输入块变为64位的密文输出块,它所使用的密钥也是64位（实际用到了56位，
// 第8、16、24、32、40、48、56、64位是校验位， 使得每个密钥都有奇数个1）
// 模式
//     CBC(Cipher Block Chaining)，密文分组链接模式
//     ECB(Electronic CodeBook)，电子密码本模式
//     CFB(Cipher FeedBack)，密文反馈模式：前一个密文分组会被送回到密码算法的输入
//     OFB(Output FeedBack)，输出反馈模式：密码算法的出处会反馈到密码算法的输入

// 使用des算法进行家吗
// src：待加密铭文 key：秘钥
func EncryptDes(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	l := block.BlockSize()                             // 块的大小
	text := symmetrical_encryption.PaddingText(src, l) // 对源数据填充
	iv := []byte("12345678")                           // 初始化向量
	mode := cipher.NewCBCEncrypter(block, iv)          // 创建CBC加密模式
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

	return symmetrical_encryption.UnpaddingText(decrypted)
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
