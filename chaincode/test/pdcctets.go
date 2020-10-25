package main

import (
	"encoding/json"
	"fmt"
	"github.com/yjzhu1106/fabric-single/chaincode/go/m"
)

var policyArgs = []string{"{\"AS\":{\"userId\":\"user22\",\"role\":\"owner\",\"PKuser\":\"MAKNFKSLFHSAHFLSFA=\"},\"AO\":{\"dataId\":\"data1\",\"signer\":\"user1\",\"sign_data\":\"dksalhfklhfkajlfkql\",\"resourceKey\":\"resource1\",\"url\":\"https://localhost\"},\"AP\":{\"auth_sign\":\"user1Sign\",\"p_data\":1,\"p_url\":1},\"AE\":{\"createdTime\":1575468182,\"endTime\":1976468182,\"address\":\"10.10.39.70\",\"sign_PKuser\":\"ahdjksahjkfhsjkah=\"}}"}
var resourceArgs = []string{"{\"resourceId\":\"R0001\",\"data\":\"183/80\",\"url\":\"https://localhost\",\"timestamp\":4214121}"}

func addData() string{
	resource, err := parseResource(resourceArgs[0])
	if err != nil {
		return err.Error()
	}
	println("the resource id is :")
	println(resource.GetId())
	println("the resource is :")
	println(resource.ToByte())
	return "OK"
}

func GetData() string{
	policy,err := parsePolicy(policyArgs[0])
	if err != nil {
		return err.Error()
	}
	println(policy.GetID())
	policyAsByte:=policy.ToBytes()

	println("the policy byte is:")
	println(policyAsByte)


	policyString := fmt.Sprintf("%x",policyAsByte)
	fmt.Sprintf("%x",policyAsByte)

	println(policyString)

	policy2 := m.Policy{}
	//err = json.Unmarshal([]byte(policyString), &policy2)
	err = json.Unmarshal(policyAsByte, &policy2)
	if err != nil {
		return err.Error()
	}
	println("the policy is :")
	println(policy2.GetID())
	return "success"
	//info, p_data, _ := cc.Auth(policy)
	//if info != "" {
	//	return shim.Error(info)
	//}
	//resource := m.Resource{}
	//if p_data == 1 {
	//	resourceAsByte, err := APIstub.GetState(policy.AO.ResourceKey)
	//	if err != nil {
	//		return shim.Error(err.Error())
	//	}
	//	err = json.Unmarshal(resourceAsByte, &resource)
	//	if err != nil {
	//		return shim.Error("403")
	//	}
	//}
	//data:=resource.DATA
	////resourceAsByte, err := APIstub.GetState(policy.AO.ResourceKey)
	////return shim.Success(resourceAsByte)
	//return shim.Success([]byte(data))
}

func main() {
	result1 := addData()
	result2:=GetData()
	println(result1)
	println(result2)
}

func parseResource(arg string) (m.Resource, error) {
	resourcesByte := []byte(arg)
	resource := m.Resource{}
	err := json.Unmarshal(resourcesByte, &resource)
	//resource.Timestamp = time.Now().Unix()
	return resource, err
}

func parsePolicy(arg string) (m.Policy, error) {
	policyAsBytes := []byte(arg)
	policy := m.Policy{}
	err := json.Unmarshal(policyAsBytes, &policy)
	return policy, err
}



