package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mickeey2525/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResonse, error) {
	fmt.Printf("Sum function was invoked with %v", req)
	firstNumber := req.FirstNum
	secondNumber := req.SecondNum
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResonse{
		SumResult: result,
	}

	return res, nil
}

func (*server) PrimeNumber(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberServer) error {
	number := req.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberResponse{
				PrimeFactor: divisor,
			})

			number = number / divisor

		} else {
			divisor++
			fmt.Println("Divisor has increased to: %v", divisor)
		}

	}
	return nil
}

func main() {
	fmt.Println("Hello,world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
