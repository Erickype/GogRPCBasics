package server

import (
	"context"
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
	return nil, nil
}

func main() {

}
