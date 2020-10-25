package m

import (
	"fmt"
	"testing"
)

var p =Policy{
	AS{"01","owner","9de81ac790473b233fb0857fbdabd4944bb92106e0a28dd25f48fa0b902c1757_pk"},
	AO{"01","01","MEQCIEWAfdLXMBX1NfAxalBL/Iv","abc","https://127.0.0.1"},
	AP{"YIBD6V8gAiABqOMjHmSQU",1,1},
	AE{31321313,313184389,"10.10.39.70","FF1XIDf+ww8u9"},
}

func TestPolicy_ToBytes(t *testing.T) {
	fmt.Sprintf("%x",p.ToBytes())
}

func TestPolicy_GetID(t *testing.T) {
	fmt.Sprintf("%x",p.GetID())
}