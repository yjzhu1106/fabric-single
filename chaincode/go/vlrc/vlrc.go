package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	//"github.com/yjzhu1106/fabric-single/chaincode/go/m"
)

type VLRContract interface {
	Init(shim.ChaincodeStubInterface) peer.Response
	Invoke(shim.ChaincodeStubInterface) peer.Response
	UpdateRecord(shim.ChaincodeStubInterface, []string) peer.Response
	GetRecord(shim.ChaincodeStubInterface, []string) peer.Response
	Synchro() peer.Response
}

type Chaincode struct {
}

func NewVLRContract() VLRContract {
	return new(Chaincode)
}

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("OK"))
}
func (cc *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "UpdateRecord" {
		return cc.UpdateRecord(APIstub, args)
	} else if function == "GetRecord" {
		return cc.GetRecord(APIstub, args)
	}else if function == "Synchro"{
		return cc.Synchro()
	}
	return shim.Error("Invalid Smart Contract function name...")
}

func (cc *Chaincode)Synchro() peer.Response {
	return shim.Success([]byte("OK"))
}

func (cc *Chaincode) UpdateRecord(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of argumentcc. Expecting 2.....")
	}
	err := APIStub.PutState(args[0],[]byte(args[1]))
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("OK"))
}
func (cc *Chaincode) GetRecord(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	recordHistory, _ := APIStub.GetHistoryForKey(args[0])
	recordStrings := "Visit Record as follow : init"
	for recordHistory.HasNext(){
		record, err := recordHistory.Next()
		if err != nil{
			return shim.Error(err.Error())
		}
		recordString :=string(record.Value)
		recordStrings = recordStrings + ", "+recordString
	}
	return shim.Success([]byte(recordStrings))
}

func main() {
	err := shim.Start(NewVLRContract())
	if err != nil {
		fmt.Printf("Error creating the smart contract: %s", err)
	}
}
