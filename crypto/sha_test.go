package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestSHA(t *testing.T) {
	hash := sha256.New()
	hash.Write([]byte("hello world"))
	res := hash.Sum(nil)
	fmt.Println("SHA256:", hex.EncodeToString(res))

	hash = sha3.New224()
	hash.Write([]byte("hello world"))
	res = hash.Sum(nil)
	fmt.Println("SHA224:", hex.EncodeToString(res))

}
