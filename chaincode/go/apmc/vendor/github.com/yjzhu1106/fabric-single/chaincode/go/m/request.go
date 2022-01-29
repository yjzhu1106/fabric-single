package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type ABACRequest struct {
	AS AS
	AO AO
}

type Attrs struct {
	UserId    string
	DataId    string
	Timestamp int64
}

func (r *ABACRequest) ToByte() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return b
}

func (a *Attrs) GetId() string {
	createTime := fmt.Sprintf("%x", time.Now().Unix())

	return fmt.Sprintf("%x", sha256.Sum256([]byte(a.UserId+a.DataId+createTime)))
}

func (r *ABACRequest) GetAttrs() Attrs {
	return Attrs{UserId: r.AS.UserId, DataId: r.AO.DataId, Timestamp: time.Now().Unix()}
}

func (r *ABACRequest) GetId() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(r.AS.UserId+r.AO.Signer)))
}