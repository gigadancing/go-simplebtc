package symmetrical_encryption

import "bytes"

// 填充最后一个分组
// src：待填充的数据 blockSize：块大小
func PaddingText(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize                   // 最后一个分组需要填充的字节数
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding) // 填充的数据
	text := append(src, paddingText...)
	// 将填充的数据和源数据进行拼接
	return text
}

// 删除填充数据
func UnpaddingText(src []byte) []byte {
	l := len(src)
	num := int(src[l-1])
	text := src[:l-num]

	return text
}

// src：待填充数据 blockSize：块大小
func ZeroPadding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{0}, padding)

	return append(src, paddingText...)
}

// 去除尾部填充的0
func ZeroUnpadding(src []byte) []byte {
	return bytes.TrimRightFunc(src, func(r rune) bool {
		return r == rune(0)
	})
}
