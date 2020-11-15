package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Policy struct {
	AS AS `json:"AS"`
	AO AO `json:"AO"`
	AP AP `json:"AP"`
	AE AE `json:"AE"`
}



type AS struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	//Group  string `json:"group"`
	PKUser string `json:"PKuser"`
}

type AO struct {
	DataId    string `json:"dataId"`
	Signer    string `json:"signer"`
	Sign_data string `json:"sign_data"`
	ResourceKey string `json:"resourceKey"`
	URL       string `json:"url"`
}
type sign struct{
	Rtext string `json:"rtext"`
	Stext string `json:"stext"`
}

type AP struct {
	Auth_sign sign `json:"auth_sign"`
	P_url  int `json:"p_url"`
	P_data int `json:"p_data"`
}

type AE struct {
	CreateTime  int64 `json:"createTime"`
	EndTime     int64 `json:"endTime"`
	Address     string `json:"address"`
	Sign_PKuser sign `json:"sign_PKuser"`
}

func (p *Policy) ToBytes() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return b
}

func (p *Policy) GetID() string {
	t := fmt.Sprintf("%x", time.Now().Unix())
	return fmt.Sprintf("%x", sha256.Sum256([]byte(p.AS.UserId+p.AO.DataId+t)))
}



