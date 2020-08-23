package main

import (
	"context"
	"fmt"
	"github.com/eshaanmangal/grpc-go-cource/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main(){
	cc, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect %v",err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
	doServerStream(c)
	doClientStream(c)
}

func doUnary(c greetpb.GreetServiceClient){

	greeting := &greetpb.GreetRequest{
		GreetRequest : &greetpb.Greeting{
			FirstName: "Kashvi",
			LastName:  "Mangal",
		},
	}

	resp,err := c.Greet(context.Background(),greeting)
	if err!=nil{
		log.Fatalf("Error connecting to server %v",err)
	}
	fmt.Printf("Output of Unary API -> %s \n",resp.Result)
}

func doServerStream(c greetpb.GreetServiceClient){
	req := &greetpb.GreetEveryoneRequest{
		GreetEveryoneRequest: &greetpb.Greeting{
			FirstName: "Kashvi",
			LastName:  "Mangal",
		},
	}
	resp,err := c.GreetEveryone(context.Background(),req)
	if err!=nil{
		log.Fatalf("Error connecting to server %v",err)
	}
	for{
		body, err := resp.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("error parsing data %v",err)
		}
		fmt.Printf("Output of Server Stream API -> %s \n",body.Result)
	}
}

func doClientStream(c greetpb.GreetServiceClient){
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Name: "Kashvi Mangal",
		},
		&greetpb.LongGreetRequest{
			Name: "Person I",
		},
		&greetpb.LongGreetRequest{
			Name: "Person II",
		},
		&greetpb.LongGreetRequest{
			Name: "Person III",
		},
		&greetpb.LongGreetRequest{
			Name: "Person IV",
		},
	}

	stream,err := c.LongGreet(context.Background())
	if err!=nil{
		log.Fatalf("Error connecting to server%v",err)
	}

	for _,req := range requests{
		stream.Send(req)
	}

	finalResponse,err := stream.CloseAndRecv();
	if err!=nil{
		fmt.Printf("error recieving stream data %v",err)
	}
	fmt.Printf("Output of Client Stream API -> %v \n",finalResponse.Result)
}