package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type ABACResponse struct {
	PolicyId  string `json:"policyId"`
	Owner     string `json:"owner"`
	RequestId string `json:"requestId"`
	Status    int `json:"status"`
	EndTime   int64 `json:"endTime"`
	Timestamp int64 `json:"timestamp"`
}


func NewABACResponse(b []byte) (ABACResponse, error) {
	re := ABACResponse{}
	err := json.Unmarshal(b, &re)
	return re, err
}

func (re *ABACResponse) ToByte() []byte {
	b, err := json.Marshal(re)
	if err != nil {
		return nil
	}
	return b
}

func (re *ABACResponse) GetId() string{
	return fmt.Sprintf("%x", sha256.Sum256([]byte(re.PolicyId)))
}