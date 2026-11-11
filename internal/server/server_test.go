package server

import "testing"

func TestNew(t *testing.T) {
	addr := "127.0.0.6379"

	srv := New(addr)

	if srv.addr != addr {
		t.Errorf("expected address %s, got %s", addr, srv.addr)
	}

	if srv.listener != nil {
		t.Errorf("expected listner to be nil before start")
	}
}