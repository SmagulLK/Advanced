package tests

import "github.com/SmagulLK/go-gRPC-AuthService/pkg/models"

func mockfindemail(s string) *models.User {
	switch s {
	case "Alkey":
		return &models.User{
			Email:    "smagul.alkey@gmail.com",
			Password: "123456",
		}
	case "Tima":
		return &models.User{
			Email:    "tima@gmail.com",
			Password: "123456",
		}
	case "beka":
		return &models.User{
			Email:    "beka@gmail.com",
			Password: "123456",
		}
	}
	return nil
}
