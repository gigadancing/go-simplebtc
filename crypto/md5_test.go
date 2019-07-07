package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	// 第一种方式
	// 优点：可以对hash对象多次复用
	// 缺点：代码繁琐
	h := md5.New()
	h.Write([]byte("hello world"))
	res := h.Sum(nil) // 128位的散列值
	fmt.Println(hex.EncodeToString(res))
	fmt.Printf("%x\n", res)
	h.Reset()
	h.Write([]byte("go go go!"))
	res = h.Sum(nil)
	fmt.Println(hex.EncodeToString(res))

	// 第二种
	// 优点：代码简洁
	// 缺点：对象不能复用
	data := md5.Sum([]byte("hello world"))
	fmt.Println(hex.EncodeToString(data[:]))
}
