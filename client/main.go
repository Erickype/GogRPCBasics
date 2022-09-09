package main

import (
	pb "github.com/Erickype/GogRPCBasics/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("Connection error:  " + err.Error())
	}

	pb.NewWishlistServiceClient(conn)
}
