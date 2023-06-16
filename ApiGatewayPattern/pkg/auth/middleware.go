package auth

import (
	"context"
	"fmt"
	"github.com/SmagulLK/APIGateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthMiddleware struct {
	service *ServiceClient
}

func InitAuthMiddleware(service *ServiceClient) *AuthMiddleware {
	return &AuthMiddleware{
		service: service,
	}
}
func (c *AuthMiddleware) AuthRequire(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.AbortWithStatus(401)
		return
	}
	token := strings.Split(authorization, "Bearer")
	fmt.Println("token", len(token))
	if len(token) != 2 {
		ctx.AbortWithStatus(401)
		return
	}
	res, err := c.service.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != 200 {
		fmt.Println(err)
		ctx.AbortWithStatus(502)
		return
	}
	ctx.Set("user_id", res.UserId)
	ctx.Next()

}
