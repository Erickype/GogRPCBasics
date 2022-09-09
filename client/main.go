package main

import (
	"context"
	"fmt"
	pb "github.com/Erickype/GogRPCBasics/proto"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("Connection error:  " + err.Error())
	}

	clientService := pb.NewWishlistServiceClient(conn)

	res, err := clientService.CreateWishlist(context.Background(), &pb.CreateWishlistReq{
		Wishlist: &pb.Wishlist{
			Id:   generateRandomID(),
			Name: "Erick",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func generateRandomID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}
