package main

import (
	"fmt"
	"testing"
)
var s = Sign_result{[]byte("daadsa"),[]byte("afrqdsad")}

func TestSign_ToBytes(t *testing.T) {
	fmt.Sprintf("%x",s.ToBytes())
}

var s1 = Sign{
	"100/100",
	1,
	"-----BEGIN ecc private key-----\nMIHcAgEBBEIB5Rsn6WarZZ0FGo4/2prJYNYdxMvyoBLqaf/GinelALg6WFutM7si\nbPzu3zeT6FLqUcp95NoUJ0B4Is+ArT6scqWgBwYFK4EEACOhgYkDgYYABAA/1/Yb\nq1raIn+fk3qUZXyzH8oNDybYs8qPr+5BtWK3/ljC6lKv2HGmz4jLcmBATRHo8QbL\nwxFx2nvX3j+TT0LnZgDuCWZmkLlCocIJsPgUFjUf/qp+npG8NSJS2cuYS2owg3lz\ntrV7zCKtvR07e1w5mb1njItA1dDbKeIygJHMrsnTxg==\n-----END ecc private key-----\n",
}
var v = Verify{
	"100/100",
	s,
	"-----BEGIN ecc private key-----\nMIHcAgEBBEIB5Rsn6WarZZ0FGo4/2prJYNYdxMvyoBLqaf/GinelALg6WFutM7si\nbPzu3zeT6FLqUcp95NoUJ0B4Is+ArT6scqWgBwYFK4EEACOhgYkDgYYABAA/1/Yb\nq1raIn+fk3qUZXyzH8oNDybYs8qPr+5BtWK3/ljC6lKv2HGmz4jLcmBATRHo8QbL\nwxFx2nvX3j+TT0LnZgDuCWZmkLlCocIJsPgUFjUf/qp+npG8NSJS2cuYS2owg3lz\ntrV7zCKtvR07e1w5mb1njItA1dDbKeIygJHMrsnTxg==\n-----END ecc private key-----\n",
}

func TestSign_ToByte(t *testing.T) {
	fmt.Sprintf("%x",s1.ToByte())
}

func TestVerify_ToByte(t *testing.T) {
	fmt.Sprintf("%x",v.ToByte())
}