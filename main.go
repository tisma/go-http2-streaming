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
	// log.Println(string(data))

	go func() {
		client := new(http2.Client)
		client.Dial()
		client.Post(data)
	}()

	server := new(http2.Server)
	err = server.Initialize()
	if err != nil {
		log.Println(err)
		return
	}

	<-waitc
}
