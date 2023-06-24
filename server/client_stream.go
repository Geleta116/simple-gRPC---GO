package main

import (
	"io"
	"log"

	pb "github.com/Geleta116/grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer ) error {
	var messages []string

	for{
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages});
		}
		

		if err != nil{
			log.Fatalf("error why receiving stream: %v",err)
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, "Hello ", req.Name)

}


}
