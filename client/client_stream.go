package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Geleta116/grpc/proto"
)



func callSayHelloClientStream(client pb.GreetServiceClient, namelist *pb.NameList){
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}

	

	for _, name := range namelist.Names{

		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil{
			log.Fatalf("Erro while sending %v", err)
		}
		log.Printf("sent the request with name: %v", name)
		time.Sleep(2 * time.Second)

	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err != nil{
		log.Fatalf("Error while recieving a response %v", err)
	}
	log.Printf("%v", res.Messages)
	

}

