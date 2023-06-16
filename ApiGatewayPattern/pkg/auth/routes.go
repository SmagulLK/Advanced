package auth

import (
	"github.com/SmagulLK/APIGateway/pkg/auth/routes"
	"github.com/SmagulLK/APIGateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, c *config.Config) *ServiceClient {
	service := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routesAuth := r.Group("/auth")
	routesAuth.POST("/register", service.Register)
	routesAuth.POST("/login", service.Login)
	return service
}
func (s *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, s.Client)
}
func (s *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, s.Client)
}
