package main

import (
	"io/ioutil"
	"log"

	"github.com/tisma/go-http2-streaming/http2"
)

func main() {
	waitc := make(chan bool)

	data, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))

	server := new(http2.Server)
	err = server.Initialize()
	if err != nil {
		log.Println(err)
		return
	}

	<-waitc
}
