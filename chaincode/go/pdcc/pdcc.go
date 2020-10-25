package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/yjzhu1106/fabric-single/chaincode/go/m"
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
	Auth(policy m.Policy) (string, int, int)
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
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	// {"Args":["resourceId","data","url"]}+timestamp
	resource, err := cc.parseResource(args[0])
	if err != nil {
		return shim.Error("bad resouece......")
	}
	err = APIstub.PutState(resource.GetId(), resource.ToByte())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(resource.GetId()))
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

	info, p_data, _ := cc.Auth(policy)
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
	data:=resource.DATA
	//resourceAsByte, err := APIstub.GetState(policy.AO.ResourceKey)
	//return shim.Success(resourceAsByte)
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
	info, _, p_url := cc.Auth(policy)
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

func (cc *Chaincode) Auth(policy m.Policy) (string, int, int) {
	p_data := 0
	p_url := 0
	// check the AE
	if time.Now().Unix() > policy.AE.EndTime {
		return "the policy time out", 0, 0
	}
	// check the identity of request
	// if policy.AS.PKUser != policy.AE.Sign_PKuser {
	// 	return "the policy have the incorrect requester...", 0, 0
	// }
	if policy.AP.P_data == 1 {
		p_data = 1
	}
	if policy.AP.P_url == 1 {
		p_url = 1
	}
	return "", p_data, p_url
}

func main() {
	err := shim.Start(NewPDCCContract())
	if err != nil {
		fmt.Printf("Error creating the smart contract: %s", err)
	}
}
