package m

import (
	"fmt"
	"testing"
	"time"
)

var r = ABACRequest{
	AS{"01","request_all","MAFFIWQHIFHSAO"},
	AO{"01","01","NFJAFJAFA","abc","http://127.0.0.1"},
}

var re = ABACResponse{
	"001",
	"01",
	"011",
	200,
	123321,
	1236565,
}

var a = Attrs{"01","01",time.Now().Unix()}

func TestABACRequest_ToByte(t *testing.T) {
	fmt.Sprintf("%x",r.ToByte())
}

func TestABACRequest_GetAttrs(t *testing.T) {
	fmt.Sprintf("%x",r.GetAttrs())
}

func TestAttrs_GetId(t *testing.T) {
	fmt.Sprintf("%x",a.GetId())
}

func TestABACResponse_ToByte(t *testing.T) {
	fmt.Sprintf("%x", re.ToByte())
}

func TestNewABACResponse(t *testing.T) {
	b := []byte("")
	re, err:=NewABACResponse(b)
	fmt.Sprintf("%x\n",re)
	fmt.Sprintf("%x",err)
}

func TestABACRequest_GetId(t *testing.T) {
	fmt.Sprintf("%x",r.GetId())
}

func TestABACResponse_GetId(t *testing.T) {
	fmt.Sprintf("%x",re.GetId())
}