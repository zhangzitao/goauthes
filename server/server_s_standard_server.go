package server

import (
	"log"
	"net/http"
)

// StandardServer is
type StandardServer struct {
	Addr string
}

// Run is
func (s *StandardServer) Run() {
	http.HandleFunc("/token", handlerToken)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}
