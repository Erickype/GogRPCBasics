package main

import (
	"context"
	pb "github.com/Erickype/GogRPCBasics/proto"
)

type server struct {
	pb.UnimplementedWishlistServiceServer
}

func (s *server) createWishlist(ctx context.Context, req *pb.CreateWishlistReq) (*pb.CreateWishlistRes, error) {
	return nil, nil
}

func main() {

}
