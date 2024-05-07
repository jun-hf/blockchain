package network

import "testing"

func TestLocalTransport(t *testing.T) {
	server7070 := NewLocalTransport(":7070")
	server8080 := NewLocalTransport(":8080")

	if err := server7070.Connect(server8080); err != nil {
		t.Error("error in connecting:", err)
	}

	if err := server8080.Connect(server7070); err != nil {
		t.Error("error in connecting:", err)
	}
}