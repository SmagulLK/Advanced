package services

import (
	"context"
	"fmt"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/db"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/models"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/pb"
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/utils"
	"log"
	"net/http"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	H   db.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) mustEmbedUnimplementedAuthServiceServer() {

}
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.H.UserRepository.FindByEmail(req.Email)

	if err == nil && user != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	newUser := models.User{
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
	}
	fmt.Println("NEW: ", newUser)

	err = s.H.UserRepository.CreateUser(&newUser)
	if err != nil {
		fmt.Println("Error while creating user")
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil

}
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.H.UserRepository.FindByEmail(req.Email)
	if err != nil {
		fmt.Println("Error while fetching user from database")
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	if user == nil {
		fmt.Println("2")
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		fmt.Println("User not found")
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(*user)

	if err != nil {
		log.Fatalln("Error while generating token: ", err)
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil

}

func (s *Server) ValidateToken(token string) (*models.User, error) {
	claims, err := s.Jwt.ValidateToken(token)
	if err != nil {
		log.Fatalln("Error while validating token: ", err)
		return nil, err
	}

	user, err := s.H.UserRepository.FindByEmail(claims.Email)
	fmt.Println("user is empty")
	if err != nil {
		log.Fatalln("Error while fetching user from database: ", err)
		return nil, err
	}

	if user == nil {
		log.Fatalln("NO USER ")
		return nil, nil // User not found
	}

	return user, nil

}
func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	user, err := s.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	if user == nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found validator",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: int64(user.ID),
	}, nil
}
