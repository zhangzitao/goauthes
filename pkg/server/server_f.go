package server

import "os"

// GenerateServer is a factory method
func GenerateServer() (ser Server, err error) {
	serverType := os.Getenv("GOAUTHES_SERVER_TYPE")
	switch serverType {
	case "Standard":
		ser = &StandardServer{Addr: "localhost:" + os.Getenv("GOAUTHES_SERVER_PORT")}
	}
	return ser, err
}
