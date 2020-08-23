package main

import (
	"context"
	"fmt"
	"github.com/eshaanmangal/grpc-go-cource/calc/calc_proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main(){
	cc, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err!= nil{
		log.Fatalf("Error COnnecting to server %v",err)
	}
	defer cc.Close()

	c := calc_proto.NewSumServiceClient(cc)
	doUnary(c)
	doMutliServerResp(c)
}

func doUnary(c calc_proto.SumServiceClient){
	requestBody := &calc_proto.SumRequest{
		FirstNumber:  10.21,
		SecondNumber: 256.278,
	}
	resp, err := c.SumOfTwoNumbers(context.Background(),requestBody)

	if err!=nil{
		log.Fatalf("Error connecting to server %v",err)
	}
	fmt.Printf("Output of 2 inputed Numbers is %f \n",resp.Result)
}


func doMutliServerResp(c calc_proto.SumServiceClient){
	requestBody := &calc_proto.TwoNumbers{
		FirstNumber:  10.21,
		SecondNumber: 256.278,
	}
	stream, err := c.ArithmeticOperations(context.Background(),requestBody)
	if err!=nil{
		log.Fatalf("Error connecting to server %v",err)
	}

	for {
		resp, err := stream.Recv()
		if err==io.EOF{
			fmt.Printf("EOF reached %v \n",err)
			break
		}
		if err!=nil{
			log.Fatalf("Error connecting to server %v \n",err)
		}
		fmt.Printf("Result %f \n",resp.GetResult())
	}
}