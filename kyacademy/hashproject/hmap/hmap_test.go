package hmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	InitBuckets()
	AddKV("a", "hello world")
	AddKV("b", "ni hao")
	AddKV("c", "xie xie")
	fmt.Println(GetValue("a"))
	fmt.Println(GetValue("b"))
	fmt.Println(GetValue("c"))
	fmt.Println(GetValue("abc"))
	AddKV("d", "i love china")
	AddKV("abc", "xie xie")
	fmt.Println(GetValue("d"))
	fmt.Println(GetValue("abc"))
}
