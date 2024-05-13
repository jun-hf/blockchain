package main

import (
	"bytes"
	"fmt"

	"github.com/jun-hf/blockchain/core"
)

func main() {
	by := make([]byte, 32)
	for i := range by {
		by[i] = byte(i)
	}
	hash := core.Hash(by)
	h := &core.Header{Version: 3, PrevHash: hash, TimeStamp: 3, Height: 3, Nonce: 3}
	network := new(bytes.Buffer)
	fmt.Println(h.EncodeBinary(network))
	newH := &core.Header{}
	fmt.Println(newH.DecodeValue(network))
	fmt.Printf("%+v \n", newH)
}