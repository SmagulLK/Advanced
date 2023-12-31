package auth

import (
	"fmt"
	"github.com/SmagulLK/APIGateway/pkg/auth/pb"
	"github.com/SmagulLK/APIGateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error in  ServiceClient ", err)
	}
	return pb.NewAuthServiceClient(cc)
}
