package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/veerlakshya/my-go-all/nats-example/model"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Couldnt connect to nats: %v", err)
	}
	defer nc.Close()

	cnt := 1
	pl := model.Payload{
		Data: "Hello World",
	}
	for {
		pl.Count = cnt
		data, _ := json.Marshal(pl)
		reply, err := nc.Request("intros", data, time.Duration(time.Millisecond*500))
		time.Sleep(time.Second * 1)
		if err != nil {
			log.Printf("Couldnt send message: %v", err)
			continue
		}
		cnt++
		log.Printf("sent message %v, got reply %v", cnt, string(reply.Data))
	}
}
