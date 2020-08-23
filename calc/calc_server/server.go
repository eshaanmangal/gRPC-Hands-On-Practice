package main

import (
	"context"
	"github.com/eshaanmangal/grpc-go-cource/calc/calc_proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {

}

func (*server) SumOfTwoNumbers(ctx context.Context, req *calc_proto.SumRequest)(res *calc_proto.SumResponse,err error){
	log.Println(req)
	firstNo := req.GetFirstNumber()
	secondNo := req.GetSecondNumber()

	response := &calc_proto.SumResponse{
		Result: firstNo+secondNo,
	}
	return response,nil
}

func (*server) ArithmeticOperations(req *calc_proto.TwoNumbers, stream calc_proto.SumService_ArithmeticOperationsServer) error{
	log.Println(req)
	firstNo := req.GetFirstNumber()
	secondNo := req.GetSecondNumber()

	for i:=0;i<4;i++{
		var outputNumber float32
		if i==0{outputNumber=firstNo+secondNo} else if i==1 {outputNumber=firstNo-secondNo} else if i==2 {outputNumber=firstNo*secondNo} else{outputNumber=firstNo/secondNo}
		response := &calc_proto.AirthmeticOperationsResponse{
			Result: outputNumber,
		}
		stream.Send(response)
		time.Sleep(1000*time.Millisecond)
	}
	return nil
}

func main(){
	req,err := net.Listen("tcp","0.0.0.0:50051")
	if err!= nil{
		log.Fatalf("Error Connecting to Server : %v",err)
	}

	s := grpc.NewServer()
	calc_proto.RegisterSumServiceServer(s,&server{})
	if err:=s.Serve(req); err!=nil{
		log.Fatalf("Error connecting to server %v",err)
	}
}
