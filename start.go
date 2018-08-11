package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mediocregopher/radix.v2/pubsub"
	"github.com/mediocregopher/radix.v2/redis"
)

type service struct {
	lastMessage  string
	pubsubClient *pubsub.SubClient
}

func (s *service) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s.lastMessage)
	log.Println("/")
}

func (s *service) receive() {
	for {
		message := s.pubsubClient.Receive().Message
		s.lastMessage = message
		log.Println("New message", message)
	}
}

func main() {
	client, _ := redis.Dial("tcp", "localhost:6379")

	var myService service
	myService.pubsubClient = pubsub.NewSubClient(client)
	myService.pubsubClient.Subscribe("subscribeChanngel")
	go myService.receive()

	log.Println("Server has started!")
	http.HandleFunc("/api/", myService.handler)
	http.ListenAndServe(":6060", nil)
}
