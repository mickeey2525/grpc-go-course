package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/mickeey2525/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cloud not connect: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC ...")
	req := &calculatorpb.SumRequest{
		FirstNum:  3,
		SecondNum: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet rpc: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streming RPC ...")
	req := &calculatorpb.PrimeNumberRequest{
		Number: 765,
	}

	stream, err := c.PrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet rpc: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something Happned: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
