package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	//"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/yjzhu1106/fabric-single/chaincode/go/m"
	"math/big"
)

type PDPCContract interface {
	Init(shim.ChaincodeStubInterface) peer.Response
	Invoke(shim.ChaincodeStubInterface) peer.Response
	SignData(shim.ChaincodeStubInterface, []string) peer.Response
	VerifyData(shim.ChaincodeStubInterface, []string) peer.Response
	//EnData(shim.ChaincodeStubInterface, []string) peer.Response
	//DeData(shim.ChaincodeStubInterface, []string) peer.Response
	Synchro() peer.Response
}

type Chaincode struct {
}

func NewPDPCContract() PDPCContract {
	return new(Chaincode)
}

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "SignData" {
		return cc.SignData(APIstub, args)
	} else if function == "VerifyData" {
		return cc.VerifyData(APIstub, args)
	//} else if function == "EnData" {
	//	return cc.EnData(APIstub, args)
	//} else if function == "DeData" {
	//	return cc.DeData(APIstub, args)
	} else if function == "Synchro" {
		return cc.Synchro()
	}
	return shim.Error("Invalid Smart Contract function name...")
}

func (cc *Chaincode) Synchro() peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) parseSign(str string) (Sign, error) {
	b := Sign{}
	signAsByte := []byte(str)
	err := json.Unmarshal(signAsByte, &b)
	return b, err
}

func (cc *Chaincode) parseVerify(str string) (Verify, error) {
	b := Verify{}
	err := json.Unmarshal([]byte(str), &b)
	return b, err
}

func (cc *Chaincode) SignData(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	// {"Args":["data","prStr"]}
	signData,err := cc.parseSign(args[0])
	if err != nil{
		return shim.Error("403")
	}
	privateKey := GetPrKey(signData.Key)
	//计算哈希值
	hash := sha256.New()
	//填入数据
	hash.Write([]byte(signData.Data))
	bytes1 := hash.Sum(nil)
	//对哈希值生成数字签名

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, bytes1)
	if err != nil {
		return shim.Error(err.Error())
	}
	rtext, _ := r.MarshalText()
	stext, _ := s.MarshalText()
	//sign_result := Sign_result{rtext, stext}
	t := time.Now().Unix()
	resourceSign := m.ResourceSign{
		signData.Data,
		rtext,
		stext,
		t,
	}
	err = APIstub.PutState(resourceSign.GetId(),resourceSign.ToByte())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(resourceSign.GetId()))
}

func (cc *Chaincode) VerifyData(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 3.....")
	}
	// {"Args":["data","sign_result","puStr"]}
	verifyData,err := cc.parseVerify(args[0])
	if err != nil{
		return shim.Error("403")
	}
	publicKey := GetPuKey(verifyData.Key)
	//计算哈希值
	hash := sha256.New()
	hash.Write([]byte(verifyData.Data))
	bytes := hash.Sum(nil)
	//验证数字签名
	var r, s big.Int
	signR :=verifyData.SignR

	r.UnmarshalText(signR.Rtext)
	s.UnmarshalText(signR.Stext)
	verify := ecdsa.Verify(publicKey, bytes, &r, &s)
	if verify {
		return shim.Success(m.OK)
	} else {
		return shim.Success(m.False)
	}

}

//func (cc *Chaincode) EnData(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
//	// adpot the ecc to encrypt the data
//	if len(args) != 2 {
//		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
//	}
//	// public
//	// {"Args":["data","prStr"]}
//	privateKey := GetPrKey(args[1])
//	prk2 := ecies.ImportECDSA(privateKey)
//	puk2 := prk2.PublicKey
//	ct, err := ecies.Encrypt(rand.Reader, &puk2, []byte(args[0]), nil, nil)
//	if err != nil {
//		return shim.Error(err.Error())
//	}
//	return shim.Success(ct)
//}
//
//func (cc *Chaincode) DeData(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
//	if len(args) != 2 {
//		return shim.Error("Incorrect number of argumentcc. Expecting 3.....")
//	}
//	// private
//	// {"Args":["ct","prStr"]}
//	privateKey := GetPrKey(args[1])
//	prk2 := ecies.ImportECDSA(privateKey)
//	pt, err := prk2.Decrypt([]byte(args[0]), nil, nil)
//	if err != nil {
//		return shim.Error(err.Error())
//	}
//	return shim.Success(pt)
//}

func main() {
	err := shim.Start(NewPDPCContract())
	if err != nil {
		fmt.Printf("Error creating the smart contract: %s", err)
	}
}
