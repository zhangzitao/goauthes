package main

import (
	"log"
	"os"

	"github.com/zhangzitao/goauthes"
	"github.com/zhangzitao/goauthes/config"
)

func main() {
	log.Println("This is simple server")
	config.LoadConfig()
	s, _ := goauthes.GenerateServer()
	log.Println(os.Getenv("GOAUTHES_STORAGE_TYPE"))

	s.Run()
}
