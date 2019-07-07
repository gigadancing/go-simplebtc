package crypto

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const KeySize = 24

//
func genKey(key []byte) []byte {
	kkey := make([]byte, KeySize) // 最终秘钥
	l := len(key)                 // 传入秘钥长度
	if l > KeySize {
		kkey = append(kkey, key[:KeySize]...) // 截取
	} else {
		div := KeySize / l // 求商
		mod := KeySize % l // 求余
		// 填充
		for i := 0; i < div; i++ {
			kkey = append(kkey, key...)
		}
		kkey = append(kkey, key[:mod]...)
	}
	return kkey
}

func Test3DesFile(t *testing.T) {
	var (
		command  string
		filename string
	)
	fmt.Println("input command, E -- [encrypt] | D -- [decrypt]")
	_, _ = fmt.Scanln(&command)

	if command == "E" {
		var (
			password string
			confirm  string
		)
		fmt.Println("input filename:")
		_, _ = fmt.Scanln(&filename)
		fmt.Println("input key:")
		_, _ = fmt.Scanln(&password)
		fmt.Println("confirm key:")
		_, _ = fmt.Scanln(&confirm)

		if !bytes.Equal([]byte(password), []byte(confirm)) {
			fmt.Println("key not equal")
		} else {
			key := genKey([]byte(password))
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				panic(err)
			}
			encrypted := Encrypt3Des(data, key)
			index := strings.LastIndex(filename, ".")
			newFilename := filename[:index] + "_encrypted" + filename[index:]
			if err := ioutil.WriteFile(newFilename, encrypted, 0777); err != nil {
				panic(err)
			}
			fmt.Println("your file is encrypted.")
		}
	} else if command == "D" {
		var password string
		fmt.Println("input encrypted file:")
		_, _ = fmt.Scanln(&filename)
		fmt.Println("input key:")
		_, _ = fmt.Scanln(password)
		key := genKey([]byte(password))
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		src := Decrypt3Des(data, key)
		if len(src) == 0 {
			fmt.Println("key error")
		} else {
			index := strings.LastIndex(filename, ".")
			newFilename := filename[:index] + "_decrypted" + filename[index:]
			if err := ioutil.WriteFile(newFilename, src, 0777); err != nil {
				panic(err)
			}
			fmt.Println("your file is decrypted.")
		}
	}
}
