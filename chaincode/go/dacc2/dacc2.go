package main

import (
	"crypto/sha256"
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
	Auth() peer.Response
	//	add two functions
	UpdateRequestRecord(shim.ChaincodeStubInterface, []string) peer.Response
	UpdateResponseRecord(shim.ChaincodeStubInterface, []string) peer.Response
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
	} else if function == "Synchro" {
		return cc.Synchro()
	} else if function == "GetRequest" {
		return cc.GetRequest(APIstub, args)
	} else if function == "GetResponse" {
		return cc.GetResponse(APIstub, args)
	} else if function == "UpdateRequestRecord" {
		return cc.UpdateRequestRecord(APIstub, args)
	} else if function == "UpdateResponseRecord" {
		return cc.UpdateResponseRecord(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name...")
}

// update request record
func (cc *Chaincode) UpdateRequestRecord(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	// arg0: requestKey
	// arg1: requestRecord
	if len(args) != 2 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	abacRequest, err := cc.parseRequest(args[1])
	if err != nil {
		return shim.Error("403")
	}
	requestKey := fmt.Sprintf("%x", abacRequest.AS.UserId + "_" + abacRequest.AO.Signer + "_request");
	if(requestKey != args[0]){
		return shim.Error("Incorrect of argument. Expecting requestKey or requestRecord.....")
	}
	key_hash := fmt.Sprintf("%x", sha256.Sum256([]byte(requestKey)))
	err = APIstub.PutState(key_hash, abacRequest.ToByte())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(key_hash))

}

// update response record
func (cc *Chaincode) UpdateResponseRecord(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	// arg0: responseKey
	// arg1: responseRecord
	if len(args) != 2 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	abacResponse, err := cc.parseResponse(args[1])
	if err != nil {
		return shim.Error("403")
	}
	responseKey := fmt.Sprintf("%x", abacResponse.Owner + "_" + abacResponse.RequestId + "_request");
	if(responseKey != args[0]){
		return shim.Error("Incorrect of argument. Expecting requestKey or requestRecord.....")
	}
	key_hash := fmt.Sprintf("%x", sha256.Sum256([]byte(responseKey)))
	err = APIstub.PutState(key_hash, abacResponse.ToByte())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(key_hash))
}

func (cc *Chaincode) Auth() peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) Synchro() peer.Response {
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
	b.Timestamp = time.Now().Unix()
	return b, err
}

func (cc *Chaincode) GetRequest(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	requestAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(requestAsBytes)
}
func (cc *Chaincode) GetResponse(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
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
	if len(args) != 1 {
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
