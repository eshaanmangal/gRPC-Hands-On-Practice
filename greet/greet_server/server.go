package main

import (
	"context"
	"github.com/eshaanmangal/grpc-go-cource/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type server struct {

}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest)(*greetpb.GreetResponse,error){
	log.Println(req)
	firstName := req.GetGreetRequest().GetFirstName()
	request := &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}
	return request,nil
}

func (*server) GreetEveryone(req *greetpb.GreetEveryoneRequest, stream greetpb.GreetService_GreetEveryoneServer) error{
	log.Println(req)
	for i:=0;i<5;i++{
		resp := &greetpb.GreetEveryoneResponse{Result: req.GreetEveryoneRequest.GetFirstName() + " " +req.GreetEveryoneRequest.GetLastName() + " "+strconv.Itoa(i) }
		stream.Send(resp)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error{
	var result string
	for{
		resp,err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&greetpb.LongGreetResponse{Result: result})
		}
		if err !=nil{
			return err
		}
		result += "Hello " + resp.GetName() + ", "
	}
}

func main(){
	lis, err := net.Listen("tcp","0.0.0.0:50051")
	if err!=nil{
		log.Fatalf("Error Connecting to %v",err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err:=s.Serve(lis); err!=nil{
		log.Fatalf("Failed to Serve %v",err)
	}
}