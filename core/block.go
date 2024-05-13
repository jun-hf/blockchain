package core

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
)

type Hash [32]uint8

func NewHash() Hash {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return Hash(b)
}

type Header struct {
	Version uint32
	PrevHash Hash
	TimeStamp uint64
	Height uint32
	Nonce uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.PrevHash); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.TimeStamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.Height); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, h.Nonce)
}

func (h *Header) DecodeValue(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return fmt.Errorf("fail to read version: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevHash); err != nil {
		return fmt.Errorf("fail to read prevhash: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &h.TimeStamp); err != nil {
		return fmt.Errorf("fail to read timestamp: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return fmt.Errorf("fail to read height: %w", err)
	}
	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

type Block struct {
	Header Header
	Transactions []Transaction
}