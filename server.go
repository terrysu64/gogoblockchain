package main

import (
	"io"
	"net/http"
	"log"
	"strconv"
)

type Server struct {
	port uint16
}

func NewServer(port uint16) *Server {
	s := new(Server)
	s.port = port
	return s
}

func (s *Server) Port() uint16 {
	return s.port
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func (s *Server) Start() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(int(s.port)), nil))
	log.Printf("Server started on port %d", s.port)
}