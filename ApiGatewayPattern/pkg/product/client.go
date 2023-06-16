package product

import (
	"fmt"
	"github.com/SmagulLK/APIGateway/pkg/config"
	"github.com/SmagulLK/APIGateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductClient(c *config.Config) pb.ProductServiceClient {

	conn, err := grpc.Dial(c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error in InitProductClient", err)
	}
	return pb.NewProductServiceClient(conn)
}
