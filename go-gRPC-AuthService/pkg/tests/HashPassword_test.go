package tests

import (
	"github.com/SmagulLK/go-gRPC-AuthService/pkg/utils"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash := utils.HashPassword(password)
	if hash == password {
		t.Errorf("HashPassword(%s) = %s; want %s", password, hash, password)
	}
}
