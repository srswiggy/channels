package main

import (
	"fmt"
	"net"
)

type Server struct {
	listenArr string
	ln        net.Listener
	quitch    chan struct{}
}

func NewServer(listenArr string) *Server {
	return &Server{listenArr: listenArr, quitch: make(chan struct{})}
}
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenArr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.aceptLoop()
	<-s.quitch
	return nil
}

func (s *Server) aceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Connect Accept error: ", err)
			continue
		}

		fmt.Println("New Connection to the server: ", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error: ", err)
			continue
		}

		msg := string(buf[:n])
		fmt.Println(conn.RemoteAddr().String() + " " + msg)
	}
}
func main() {
	server := NewServer(":8000")
	server.Start()
}
