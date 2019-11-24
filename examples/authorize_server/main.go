package main

import (
	"log"
	"os"

	"github.com/zhangzitao/goauthes/config"
	"github.com/zhangzitao/goauthes/server"
)

func main() {
	log.Println("This is simple server")
	config.LoadConfig()
	s, _ := server.GenerateServer()
	log.Println(os.Getenv("GOAUTHES_STORAGE_TYPE"))

	s.Run()
}
