package main

import (
	"fmt"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/config"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/db"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/pb"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/services"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db, h := db.InitDB(c.DB_Url)
	defer db.Close()

	jwt := utils.JwtWrapper{
		SecretKey:       c.JwtSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
