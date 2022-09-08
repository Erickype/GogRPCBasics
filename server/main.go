package server

import (
	"context"
	"fmt"
	pb "github.com/Erickype/GogRPCBasics/proto"
	"math/rand"
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

}
