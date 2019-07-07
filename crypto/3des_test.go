package crypto

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"testing"
)

// 3DES
// 明文经过3次DES处理才能变成最后的密文，由于DES秘钥的长度实质上是56比特，因此3DES的秘钥长度是56*3=168比特。
// 3DES并不是经过3次加密（加密->加密->加密）而是加密->解密->加密的过程，这中设计是为了让3DES能够兼容普通的DES。
// 当3DES所有秘钥都相同时，3DES也就是普通的DES。因此DES加密的密文也可以使用3DES来进行解密。
// 密码算法不能依靠算法的不公开性来保证密码算法的安全性。反而应该公开算法思想，如果大家都不能破解，才是安全的密码算法。
//
//

func Encrypt3Des(src, key []byte) []byte {
	cipherBlock, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}

	text := PaddingText(src, cipherBlock.BlockSize())
	mode := cipher.NewCBCEncrypter(cipherBlock, key[:cipherBlock.BlockSize()])
	dst := make([]byte, len(text))
	mode.CryptBlocks(dst, text)

	return dst
}

//
func Decrypt3Des(src, key []byte) []byte {
	cipherBlock, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(cipherBlock, key[:cipherBlock.BlockSize()])
	dst := make([]byte, len(src))
	mode.CryptBlocks(dst, src)
	dst = UnpaddingText(dst)
	return dst
}

func Test3Des(t *testing.T) {
	key := []byte("12345678abcdefgh87654321")
	data := []byte("hello world")

	fmt.Println("明文：", string(data))
	encrypted := Encrypt3Des(data, key)
	fmt.Println("3DES加密后：", hex.EncodeToString(encrypted))
	decrypted := Decrypt3Des(encrypted, key)
	fmt.Println("3DES解密后：", string(decrypted))
}
