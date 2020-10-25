package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/newham/hamtask"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"os"
)

var BASE_URL = "http://localhost:8001/fabric-single/test"

func main() {
	ccName, fName, n := GetArgs()
	if n == -1 {
		println("please input the number of client......")
		return
	}
	println("Chaincode Name : ", ccName)
	println("Function Name : ", fName)
	println("Client number : ", n)
	start := time.Now().UnixNano()
	if ccName == "apmc" {
		if fName == "AddPolicy" {
			println("Test the AddPolicy function start.....")
			Loop(1, n, TestAPMCAddPolicy)
		} else if fName == "UpdatePolicy" {
			Loop(1, n, TestAPMCUpdatePolicy)
		} else if fName == "DeletePolicy" {
			Loop(1, n, TestAPMCDeletePolicy)
		} else if fName == "QueryPolicy" {
			println("Test the QueryPolicy function start.....")
			Loop(1, n, TestAPMCQueryPolicy)
		}
	} else if ccName == "dacc" {
		if fName == "RequestAccess" {
			Loop(1, n, TestDACCRequestAccess)
		} else if fName == "ResponseAccess" {
			Loop(1, n, TestDACCResponseAccess)
		} else if fName == "CheckAccess" {
			Loop(1, n, TestDACCCheckAccess)
		}
	} else if ccName == "pdcc" {
		if fName == "AddData" {
			Loop(1, n, TestPDCCAddData)
		} else if fName == "GetData" {
			Loop(1, n, TestPDCCGetData)
		} else if fName == "GetURL" {
			Loop(1, n, TestPDCCGetURL)
		}

	}
	end := time.Now().UnixNano()
	println("cost time : ", end-start, "second")
	println("tcp : ", (end-start)/int64(n),"second")
}
func GetID(i int) int {
	return i +410000
}
func GetArgs() (string, string, int) {
	args := os.Args
	// args:=[]string{"asd","apmc","AddPolicy","1000"}
	if len(args) > 3 {
		ccName := args[1]
		fName := args[2]
		n, err := strconv.Atoi(args[3])
		if err != nil {
			println("bad count.....")
			return "", "", -1
		}
		return ccName, fName, n
	}
	return "", "", -1
}

func Loop(loop int, n int, f func(v chan int, n int)) {
	c := make(chan int, 100000)
	for i := 0; i < loop; i++ {
		go func() {
			f(c, n)
		}()
	}
	// wait channel end
	for i := 0; i < loop; i++ {
		<-c
	}
}

func TestCC(c chan int, n int, f func(int) string) {
	i := 0
	hamtask.NewSimpleWorker(n, func(i int, d hamtask.Data) {
		url := d.String()
		_, err := http.Get(url)
		if err != nil {
			log.Println("Error", err.Error())
		}

	}, func() hamtask.Data {
		i++
		return hamtask.String(f(i))
	}, n).Start()
	c <- 1
}

func TestAPMCAddPolicy(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("apmc", "AddPolicy", GetPolicyRequest(i, 1, 1))
	})
}

func GetPolicyRequest(i, p_url, p_data int) string {
	t := time.Now().Unix()
	return fmt.Sprintf(`{"AS":{"userId":"%s","role":"requester","PKuser":"MAIEKSAKJD="},"AO":{"dataId":"%s","signer":"user1","sign_data":"MAJSKDS=","resourceKey":"be9ecaa87baefef74b9c574092466ce332513cddf981fcba08087a1be9adc5ba","url":"https://localhost"},"AP":{"auth_sign":"MASHFKLJFS=","p_data":%d,"p_url":%d},"AE":{"createTime":%d,"endTime":%d,"address":"10.10.39.70","sign_PKuser":"MAHFHISOAHFS="}}`, GetUserID(i), GetDataId(i), p_data, p_url, t, t+1000000)
}

func TestAPMCUpdatePolicy(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("apmc", "UpdatePolicy", GetPolicyRequest(i, 1, 0))
	})
}

func TestAPMCDeletePolicy(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("apmc", "DeletePolicy", GetPolicyId(i))
	})
}
func TestAPMCQueryPolicy(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("apmc", "QueryPolicy", GetPolicyId(i))
	})
}

func TestDACCRequestAccess(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("dacc", "RequestAccess", GetAccessRequest(i))
	})
}
func GetAccessRequest(i int) string {
	return fmt.Sprintf(`{"AS":{"userId":"%s","role":"requester","PKuser":"MAIEKSAKJD="},"AO":{"dataId":"%s","signer":"user1","sign_data":"MAJSKDS=","resourceKey":"be9ecaa87baefef74b9c574092466ce332513cddf981fcba08087a1be9adc5ba","url":"https://localhost"}}`, GetUserID(i), GetDataId(i))
}

func TestDACCResponseAccess(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("dacc", "ResponseAccess", GetAccessResponse("dhashdias", "djkslajdkja", 1))
	})
}
func GetAccessResponse(policyId string, requestId string, status int) string {
	t := time.Now().Unix()
	return fmt.Sprintf(`{"policyId":"%s","owner":"user1","requestId":"%s","status":%d,"endTime":%d,"timestamp":%d}`, policyId, requestId, status, t+100000, t)
}
func TestDACCCheckAccess(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("dacc", "CheckAccess", GetPolicyId(i))
	})
}

func TestPDCCAddData(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("pdcc", "AddData", GetDataRequest("jsadsa", "180", "http://localhost"))
	})
}

func TestPDCCGetData(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("pdcc", "GetData", GetPolicyId(i))
	})
}

func TestPDCCGetURL(v chan int, n int) {
	TestCC(v, n, func(i int) string {
		return GetURL("pdcc", "GetURL", GetPolicyId(i))
	})
}

func GetDataRequest(resourceId, data, url string) string {
	t := time.Now().Unix()
	return fmt.Sprintf(`{"resourceId":"%s","data":"%s","url":"%s"}`, resourceId, data, url, t)
}

func GetPolicyId(i int) string {
	return GetSHA256(GetUserID(i), GetDataId(i))
}

func GetSHA256(args ...string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(args, ""))))
}

func GetRequest(i int) string {
	return fmt.Sprintf(`{"AS":{"userId":"%s","role":"requester","PKuser":"ECC"},"AO":{"dataId":"%s","signer":"user1","sign_data":"","resourceKey":"be9ecaa87baefef74b9c574092466ce332513cddf981fcba08087a1be9adc5ba","url":"https://localhost"}}`, GetUserID(i), GetDataId(i))
}

func GetUserID(i int) string {
	//return strconv.Itoa(GetID(i))
	return fmt.Sprintf("user%d", GetID(i))
}

func GetDataId(i int) string {
	return fmt.Sprintf("data%d", GetID(i))
}
func GetURL(ccName string, fName string, args ...string) string {
	//println("args",args)
	return fmt.Sprintf("%s?cc_name=%s&f_name=%s&args=%s", BASE_URL, ccName, fName, args[0])
}

