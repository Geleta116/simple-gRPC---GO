package main
import (
	"io"
	"log"

	pb "github.com/Geleta116/grpc/proto"
)

func (s *helloServer) SayHelloBiDirectionalStreaming( stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	
	for{
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("recieved : %v", req.Name)

		res :=  &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}

		if err := stream.Send(res); err != nil{
			return err
		}


	}
}