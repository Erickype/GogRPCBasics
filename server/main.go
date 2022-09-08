package main

import (
	"context"
	"fmt"
	pb "github.com/Erickype/GogRPCBasics/proto"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type server struct {
	pb.UnimplementedWishlistServiceServer
}

func generateRandomID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

func (s *server) createWishlist(ctx context.Context, req *pb.CreateWishlistReq) (*pb.CreateWishlistRes, error) {
	fmt.Printf("Creating wishlist: %v", req.Wishlist.Name)
	return &pb.CreateWishlistRes{
		WishlistId: generateRandomID(),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Cannot create tcp connection: " + err.Error())
	}

	serv := grpc.NewServer()

	//Server register
	pb.RegisterWishlistServiceServer(serv, &server{})

	err = serv.Serve(listen)
	if err != nil {
		panic("Cannot initialize server: " + err.Error())
	}
}
