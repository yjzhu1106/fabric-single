package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
)

type Sign_result struct {
	Rtext []byte `json:"rtext"`
	Stext []byte `json:"stext"`
}

type Sign struct {
	Data string `json:"data"`
	Random int	`json:"random""`
	Key  string `json:"key"`
}

type Verify struct{
	Data string `json:"data"`
	SignR Sign_result `json:"sign_result"`
	Key  string `json:"key"`
}

func (s *Sign) ToByte() []byte {
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	return b
}
func (v *Verify) ToByte() []byte {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return b
}


func (s *Sign_result) ToBytes() []byte {
	b, err:=json.Marshal(s)
	if err != nil{
		return nil
	}
	return b
}

func GetPrKey(pr string) *ecdsa.PrivateKey {
	//读取私钥
	buf := []byte(pr)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func GetPuKey(pu string) *ecdsa.PublicKey {
	//读取私钥
	buf := []byte(pu)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicInterface.(*ecdsa.PublicKey)
	return publicKey
}
