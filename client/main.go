package main

import (
	"log"

	pb "github.com/Geleta116/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	names := &pb.NameList{
		Names: []string{"GELETA", "GUTU"},
	}

	client := pb.NewGreetServiceClient(conn)

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBiDirectionalStreaming(client,names)
}
