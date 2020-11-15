package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/yjzhu1106/fabric-single/chaincode/go/m"
	"math/big"
	"time"
)

type PDCCContract interface {
	Init(shim.ChaincodeStubInterface) peer.Response
	Invoke(shim.ChaincodeStubInterface) peer.Response
	Synchro() peer.Response
	AddData(shim.ChaincodeStubInterface, []string) peer.Response
	//AddURL(shim.ChaincodeStubInterface, []string) peer.Response
	GetData(shim.ChaincodeStubInterface, []string) peer.Response
	GetURL(shim.ChaincodeStubInterface, []string) peer.Response
}
type Chaincode struct {
}

func NewPDCCContract() PDCCContract {
	return new(Chaincode)
}
func (cc *Chaincode) Synchro() peer.Response {
	return shim.Success(m.OK)
}
func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(m.OK)
}
func (cc *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "AddData" {
		return cc.AddData(APIstub, args)
	} else if function == "GetData" {
		return cc.GetData(APIstub, args)
	} else if function == "GetURL" {
		return cc.GetURL(APIstub, args)
	} else if function == "Synchro" {
		return cc.Synchro()
	}
	return shim.Error("Invalid Smart Contract function name...")
}
func (cc *Chaincode) parseResource(arg string) (m.Resource, error) {
	resourcesByte := []byte(arg)
	resource := m.Resource{}
	err := json.Unmarshal(resourcesByte, &resource)
	resource.Timestamp = time.Now().Unix()
	return resource, err
}

func (cc *Chaincode) AddData(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	// {"Args":["resourceId","data","url"],["owner"],["PKowner"]}+timestamp
	resource, err := cc.parseResource(args[0])
	if err != nil {
		return shim.Error("bad resouece......")
	}
	err = APIstub.PutState(resource.GetId(), resource.ToByte())
	if err != nil {
		return shim.Error(err.Error())
	}
	user, _ := APIstub.GetState(args[1])
	if user == nil {
		err := APIstub.PutState(args[1], []byte(args[2]))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(resource.GetId()))
	} else {
		return shim.Success([]byte(resource.GetId()))
	}
}

func (cc *Chaincode) GetData(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	//{"Args":["policyId"]}
	resp := APIstub.InvokeChaincode("apmc", [][]byte{[]byte("QueryPolicy"), []byte(args[0])}, "iot-channel")
	//policyAsByte, err := APIstub.GetState(args[0])
	if resp.GetStatus() != 200 {
		return shim.Error("403")
	}
	policy := m.Policy{}
	err := json.Unmarshal(resp.GetPayload(), &policy)
	if err != nil {
		return shim.Error("403.")
	}

	keyAsByte, _ := APIstub.GetState(policy.AO.Signer)
	info, p_data, _, err := cc.Auth(policy, string(keyAsByte))
	if err != nil {
		return shim.Error(err.Error())
	}
	if info != "" {
		return shim.Error(info)
	}
	resource := m.Resource{}
	if p_data == 1 {
		resourceAsByte, err := APIstub.GetState(policy.AO.ResourceKey)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = json.Unmarshal(resourceAsByte, &resource)
		if err != nil {
			return shim.Error("403")
		}
	}
	data := resource.DATA
	//resourceAsByte, err := APIstub.GetState(policy.AO.ResourceKey)
	//return shim.Success(resourceAsByte)
	mspid := policy.AS.UserId
	resp2 := APIstub.InvokeChaincode("vlrc", [][]byte{[]byte("UpdateRecord"), []byte(resource.ResourceId), []byte(mspid)}, "iot-channel")
	if resp2.GetStatus() != 200 {
		return shim.Error("403")
	}
	return shim.Success([]byte(data))
}
func (cc *Chaincode) GetURL(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	//{"Args":["policyId"]}
	resp := APIstub.InvokeChaincode("apmc", [][]byte{[]byte("QueryPolicy"), []byte(args[0])}, "iot-channel")
	//policyAsByte, err := APIstub.GetState(args[0])
	if resp.GetStatus() != 200 {
		return shim.Error("403")
	}
	policy := m.Policy{}
	err := json.Unmarshal(resp.GetPayload(), &policy)
	if err != nil {
		return shim.Error("403.")
	}
	//keyAsByte, _ := APIstub.GetState(policy.AO.Signer)
	//info, _, p_url, err := cc.Auth(policy, string(keyAsByte))
	info := ""
	p_url := policy.AP.P_url

	if err != nil {
		return shim.Error(err.Error())
	}
	if info != "" {
		return shim.Error(info)
	}
	resource := m.Resource{}
	if p_url == 1 {
		resourceAsByte, err := APIstub.GetState(policy.AO.ResourceKey)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = json.Unmarshal(resourceAsByte, &resource)
		if err != nil {
			return shim.Error("403")
		}
	}
	url := resource.URL
	return shim.Success([]byte(url))
}

func (cc *Chaincode) Auth(policy m.Policy, key string) (string, int, int, error) {
	p_data := 0
	p_url := 0
	if time.Now().Unix() > policy.AE.EndTime {
		return "the policy time out", 0, 0, nil
	}
	if policy.AP.P_data == 1 {
		p_data = 1
	}
	if policy.AP.P_url == 1 {
		p_url = 1
	}
	rtext := []byte(policy.AP.Auth_sign.Rtext)
	stext := []byte(policy.AP.Auth_sign.Stext)
	data := string(policy.AP.P_data) + string(policy.AP.P_url)
	//time.Sleep(2000000000)
	return "", p_data, p_url, nil
	verify_result, err := verify(data, rtext, stext, key)

	if !verify_result {
		return "the verify signature failed...", 0, 0, err
	}
	return "", p_data, p_url, err
}

func verify(data string, rtext, stext []byte, key string) (bool, error) {
	// {"Args":["data","sign_result","puStr"]}

	publicKey, err := GetPuKey(key)
	//计算哈希值
	hash := sha256.New()
	hash.Write([]byte(data))
	bytes := hash.Sum(nil)
	//验证数字签名
	var r, s big.Int
	r.UnmarshalText(rtext)
	s.UnmarshalText(stext)
	verify := ecdsa.Verify(publicKey, bytes, &r, &s)
	if verify {
		return true, err
	} else {
		return false, err
	}
}

func GetPuKey(pu string) (*ecdsa.PublicKey, error) {
	//读取私钥
	buf := []byte(pu)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := publicInterface.(*ecdsa.PublicKey)
	return publicKey, err

}

func main() {
	err := shim.Start(NewPDCCContract())
	if err != nil {
		fmt.Printf("Error creating the smart contract: %s", err)
	}
}
