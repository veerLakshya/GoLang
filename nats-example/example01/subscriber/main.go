package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my-go-all/nats-example/model"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Couldnt connect to nats: %v", err)
	}
	defer nc.Close()

	nc.Subscribe("intros", func(msg *nats.Msg) {
		pl := &model.Payload{}
		json.Unmarshal(msg.Data, pl)
		replyData := fmt.Sprintf("ack message # %v", pl.Count)
		msg.Respond([]byte(replyData))
		log.Println("Received message:", string(msg.Data))
	})
	time.Sleep(time.Minute * 1)
}
