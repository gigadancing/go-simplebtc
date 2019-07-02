package bc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

// 钱包本质上就是存储了一些公私钥对
// 钱包
type Wallet struct {
	PrivateKey ecdsa.PrivateKey // 私钥
	PublicKey  []byte           // 公钥
}

// 创建钱包
func NewWallet() *Wallet {
	privKey, pubKey := newKeyPair()
	return &Wallet{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
}

// 生成公私钥对
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	// 椭圆曲线加密
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicf("ecdsa generate key failed: %v\n", err)
	}
	pubKey := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
	return *priv, pubKey
}
