package network

import (
	"testing"
)

func TestLocalTransport(t *testing.T) {
	server7070 := NewLocalTransport(":7070")
	server8080 := NewLocalTransport(":8080")

	if err := server7070.Connect(server8080); err != nil {
		t.Error("error in connecting:", err)
	}

	if err := server8080.Connect(server7070); err != nil {
		t.Error("error in connecting:", err)
	}
	if err := server7070.SendMessage(":8080", []byte("Hello")); err != nil {
		t.Error("error in sending messaga:", err)
	}
	rpc := <- server8080.Consume()
	if rpc.From != ":7070" {
		t.Errorf("expected (%v) got (%v)", ":7070", rpc.From)
	}
	if string(rpc.Payload) != "Hello" {
		t.Errorf("expected (%v) got (%v)", "Hello", string(rpc.Payload))
	}
}