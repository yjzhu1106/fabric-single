package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/yjzhu1106/fabric-single/chaincode/go/m"
)

type APMContract interface {
	Init(shim.ChaincodeStubInterface) peer.Response
	Invoke(shim.ChaincodeStubInterface) peer.Response
	AddPolicy(shim.ChaincodeStubInterface, []string) peer.Response
	UpdatePolicy(shim.ChaincodeStubInterface, []string) peer.Response
	DeletePolicy(shim.ChaincodeStubInterface, []string) peer.Response
	QueryPolicy(shim.ChaincodeStubInterface, []string) peer.Response
	Auth(shim.ChaincodeStubInterface) bool
	Synchro() peer.Response
}
type Chaincode struct {

}

func NewAPMContract() APMContract {
	return new(Chaincode)
}

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "AddPolicy" {
		return cc.AddPolicy(APIstub, args)
	} else if function == "UpdatePolicy" {
		return cc.UpdatePolicy(APIstub, args)
	} else if function == "DeletePolicy" {
		return cc.DeletePolicy(APIstub, args)
	} else if function == "QueryPolicy" {
		return cc.QueryPolicy(APIstub, args)
	}else if function == "Synchro" {
		return cc.Synchro()
	}
	return shim.Error("Invalid Smart Contract function name...")
}

func (cc *Chaincode)Synchro() peer.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) parsePolicy(arg string) (m.Policy, error) {
	policyAsBytes := []byte(arg)
	policy := m.Policy{}
	err := json.Unmarshal(policyAsBytes, &policy)
	return policy, err
}

func (cc *Chaincode) AddPolicy(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	policy, err := cc.parsePolicy(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	//queryconsole, _ := APIstub.GetState(policy.GetID())
	queryconsole := cc.QueryPolicy(APIstub, []string{policy.GetID()})
	if queryconsole.GetStatus() == 200 {
		// return shim.Error("policy exited...")
		err = APIstub.PutState(policy.GetID(), policy.ToBytes())
		if err != nil {
		return shim.Error(err.Error())
		}
		return shim.Success([]byte(policy.GetID()))
	}

	err = APIstub.PutState(policy.GetID(), policy.ToBytes())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(policy.GetID()))
}

func (cc *Chaincode) UpdatePolicy(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	policy, err := cc.parsePolicy(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	r := cc.QueryPolicy(APIstub, []string{args[0]})
	if r.GetStatus() != 200 {
		// return shim.Error("policy not exit......")
		err = APIstub.PutState(args[0], policy.ToBytes())
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(policy.ToBytes())
	}
	err = APIstub.PutState(args[0], policy.ToBytes())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(policy.ToBytes())
	//return cc.AddPolicy(APIstub, args)
}

func (cc *Chaincode) DeletePolicy(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	err := APIstub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(m.OK)
}

func (cc *Chaincode) QueryPolicy(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1.....")
	}
	policyAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if policyAsBytes == nil{
		return shim.Error("the policy not exit")
	}
	return shim.Success(policyAsBytes)
}

func (cc *Chaincode) Auth(APIStub shim.ChaincodeStubInterface) bool{
	creatorByte, err:= APIStub.GetCreator()
	si := &msp.SerializedIdentity{}
	err = proto.Unmarshal(creatorByte,si)
	mspId := si.GetMspid()
	cert := string(si.GetIdBytes())
	if err != nil{
		return false
	}
	if string(mspId) == ""{
		return false
	}
	if cert == ""{
		return false
	}
	return true
}

func main() {
	err := shim.Start(NewAPMContract())
	if err != nil {
		fmt.Printf("Error creating the smart contract: %s", err)
	}
}
