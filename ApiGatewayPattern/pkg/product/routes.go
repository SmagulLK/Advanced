package product

import (
	"github.com/SmagulLK/APIGateway/pkg/auth"
	"github.com/SmagulLK/APIGateway/pkg/config"
	"github.com/SmagulLK/APIGateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, c *config.Config, authService *auth.ServiceClient) {
	//a := auth.InitAuthMiddleware(authService)
	ProductService := &ProductServiceClient{Client: InitProductClient(c)}
	routes := r.Group("/product")
	//routes.Use(a.AuthRequire)
	routes.POST("/", ProductService.CreateProduct)
	routes.GET("/:id", ProductService.FindOne)
}
func (c *ProductServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, c.Client)
}
func (c *ProductServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, c.Client)

}
