package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Resource struct {
	ResourceId string `json:"resourceId"`
	DATA       string `json:"data"`
	URL        string `json:"url"`
	Timestamp  int64  `json:"timestamp"`
}

type ResourceSign struct {
	Data string `json:"data"`
	Rtext  []byte `json:"rtext"`
	Stext  []byte `json:"stext"`
	Timestamp int64 `json:"timestamp"`
}

func (s *ResourceSign) ToByte() []byte {
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	return b
}
func (s *ResourceSign) GetId() string {
	t := fmt.Sprintf("%x", s.Timestamp)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s.Data+t)))
}

func (r *Resource) ToByte() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return b
}

func (r *Resource) GetId() string {
	// t := fmt.Sprintf("%x", r.Timestamp)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(r.ResourceId)))
}

func NewResource(b []byte) (Resource, error) {
	r := Resource{}
	err := json.Unmarshal(b, &r)
	return r, err
}
