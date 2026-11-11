package server

import "net"

type Server struct {
	addr string
	listener net.Listener
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
	}
}