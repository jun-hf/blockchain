package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr NetAddr
	consumeCh chan RPC
	mu sync.RWMutex
	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr: addr,
		consumeCh: make(chan RPC, 1024),
		peers: make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.peers[tr.Addr()] = tr
	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%v does not exist in peer", to)
	}
	peer.consumeCh <- RPC{
		From: t.addr,
		Payload: payload,
	}
	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}