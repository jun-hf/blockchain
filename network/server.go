package network

import (
	"errors"
	"fmt"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	quitCh chan interface{}
	rpcCh chan RPC
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		quitCh: make(chan interface{}),
		rpcCh: make(chan RPC),
	}
}

func (s *Server) Start() error {
	s.listenTransports()
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("Rpc: %+v\n", rpc)
		case <-s.quitCh:
			return errors.New("quiting server")
		}
	}
}

func (s *Server) Close() {
	close(s.quitCh)
}

func (s *Server) listenTransports() {
	for _, t := range(s.Transports) {
		go func(tran Transport) {
			for rpc := range t.Consume() {
				s.rpcCh <- rpc
			}
		}(t)
	}
}