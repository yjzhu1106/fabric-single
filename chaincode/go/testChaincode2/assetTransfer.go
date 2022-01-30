package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/yjzhu1106/fabric-single/chaincode/go/m"
)

type ASSERTContract interface {
	Init(shim.ChaincodeStubInterface) peer.Response
	Invoke(shim.ChaincodeStubInterface) peer.Response
	CreateAsset(shim.ChaincodeStubInterface, []string) peer.Response
	UpdateAsset(shim.ChaincodeStubInterface, []string) peer.Response
	DeleteAsset(shim.ChaincodeStubInterface, []string) peer.Response
	QueryAsset(shim.ChaincodeStubInterface, []string) peer.Response
	Synchro() peer.Response
}
type Chaincode struct {
}

func NewASSETContract() ASSERTContract {
	return new(Chaincode)
}

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "CreateAsset" {
		return cc.CreateAsset(APIstub, args)
	} else if function == "UpdateAsset" {
		return cc.UpdateAsset(APIstub, args)
	} else if function == "DeleteAsset" {
		return cc.DeleteAsset(APIstub, args)
	} else if function == "QueryAsset" {
		return cc.QueryAsset(APIstub, args)
	} else if function == "Synchro" {
		return cc.Synchro()
	}
	return shim.Error("Invalid Smart Contract function name...")
}

func (cc *Chaincode) Synchro() peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) parsePolicy(arg string) (m.Policy, error) {
	policyAsBytes := []byte(arg)
	policy := m.Policy{}
	err := json.Unmarshal(policyAsBytes, &policy)
	return policy, err
}

func (cc *Chaincode) CreateAsset(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := APIstub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(args[0]))
}

func (cc *Chaincode) UpdateAsset(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	r := cc.QueryAsset(APIstub, []string{args[0]})
	if r.GetStatus() != 200 {
		// return shim.Error("policy not exit......")
		err := APIstub.PutState(args[0], []byte(args[1]))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(args[0]))
	}
	err := APIstub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(args[0]))

}

func (cc *Chaincode) DeleteAsset(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := APIstub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(m.OK)
}

func (cc *Chaincode) QueryAsset(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	resultByte, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if resultByte == nil {
		return shim.Error("the policy not exit")
	}
	return shim.Success(resultByte)
}

func main() {
	err := shim.Start(NewASSETContract())
	if err != nil {
		fmt.Printf("Error creating the smart contract: %s", err)
	}
}
