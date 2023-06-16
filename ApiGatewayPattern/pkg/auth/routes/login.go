package routes

import (
	"fmt"
	"github.com/SmagulLK/APIGateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	fmt.Println(b)
	res, err := c.Login(ctx, &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})
	fmt.Println(b)
	if err != nil {
		ctx.AbortWithError(502, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
