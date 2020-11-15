package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/yjzhu1106/fabric-single/chaincode/go/m"
	"time"
)

type DACCContract interface {
	Init(shim.ChaincodeStubInterface) peer.Response
	Invoke(shim.ChaincodeStubInterface) peer.Response
	Synchro() peer.Response
	RequestAccess(shim.ChaincodeStubInterface, []string) peer.Response
	ResponseAccess(shim.ChaincodeStubInterface, []string) peer.Response
	CheckAccess(shim.ChaincodeStubInterface, []string) peer.Response
	GetRequest(shim.ChaincodeStubInterface, []string) peer.Response
	GetResponse(shim.ChaincodeStubInterface, []string) peer.Response
	//ActionRecord(shim.ChaincodeStubInterface, []string) peer.Response
	Auth() peer.Response
}

type Chaincode struct {

}

func NewDACCContract() DACCContract {
	return new(Chaincode)
}

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "RequestAccess" {
		return cc.RequestAccess(APIstub, args)
	} else if function == "ResponseAccess" {
		return cc.ResponseAccess(APIstub, args)
	} else if function == "CheckAccess" {
		return cc.CheckAccess(APIstub, args)
	} else if function == "Synchro"{
		return cc.Synchro()
	}else if function == "GetRequest"{
		return cc.GetRequest(APIstub, args)
	}else if function == "GetResponse"{
		return cc.GetResponse(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name...")
}
func (cc *Chaincode) Auth() peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode)Synchro()peer.Response{
	return shim.Success(m.OK)
}

func (cc *Chaincode) parseRequest(str string) (m.ABACRequest, error) {
	b := m.ABACRequest{}
	err := json.Unmarshal([]byte(str), &b)
	return b, err
}
func (cc *Chaincode) parseResponse(str string) (m.ABACResponse, error) {
	b := m.ABACResponse{}
	err := json.Unmarshal([]byte(str), &b)
	b.Timestamp =time.Now().Unix()
	return b, err
}

func (cc *Chaincode) GetRequest(APIstub shim.ChaincodeStubInterface, args []string) peer.Response{
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	requestAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(requestAsBytes)
}
func (cc *Chaincode) GetResponse(APIstub shim.ChaincodeStubInterface, args []string) peer.Response{
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	responseAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(responseAsBytes)
}

func (cc *Chaincode) RequestAccess(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	abacRequest, err := cc.parseRequest(args[0])
	if err != nil {
		return shim.Error("403")
	}

	//attrs := abacRequest.GetAttrs()
	// gain the key of request
	//policyId := attrs.GetId()
	// put the request to blockchain, wait the response from the owner
	err = APIstub.PutState(abacRequest.GetId(), abacRequest.ToByte())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(abacRequest.GetId()))
}

func (cc *Chaincode) ResponseAccess(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	abacResponse, err := cc.parseResponse(args[0])
	if err != nil {
		return shim.Error("403")
	}
	responseId := abacResponse.GetId()
	err = APIstub.PutState(responseId, abacResponse.ToByte())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(responseId))

}
func (cc *Chaincode) CheckAccess(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1{
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	// {"Args":["policyId"]}
	resp := APIstub.InvokeChaincode("apmc", [][]byte{[]byte("QueryPolicy"), []byte(args[0])}, "iot-channel")
	//policyAsByte, err := APIstub.GetState(args[0])
	if resp.GetStatus() != 200 {
		return shim.Error("403")
	}
	policy := m.Policy{}
	err := json.Unmarshal(resp.GetPayload(), &policy)
	if err != nil {
		return shim.Error("check policy false")
	}
	return shim.Success([]byte("check policy true"))
}

//func (cc *Chaincode) ActionRecord(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
//
//}

func main() {
	err := shim.Start(NewDACCContract())
	if err != nil {
		fmt.Printf("Error creating the smart contract: %s", err)
	}
}
