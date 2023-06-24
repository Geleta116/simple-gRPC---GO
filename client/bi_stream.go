package main

import (
	"context"
	"log"
	"io"
	"time"
	pb "github.com/Geleta116/grpc/proto"
)

func callSayHelloBiDirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList){
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while sending request: %v",err)
	}
	waitc := make(chan struct{})
	go func(){
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil{
				log.Fatalf("Error whil receiving:  %v", err)
			}

			log.Printf("%v", res.Message)


		}
		close(waitc)
	
	}()
	for _, name := range names.Names{
		
		log.Printf("sending request for : %v", name)
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil{
			log.Fatalf("Error while sending Request: %v", err)
		}
		time.Sleep(2 * time.Second)
	
}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")

	

	

}