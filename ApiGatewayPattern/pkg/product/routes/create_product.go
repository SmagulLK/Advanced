package routes

import (
	"context"
	"github.com/SmagulLK/APIGateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Stock int64  `json:"stock"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Price: body.Price,
		Stock: body.Stock,
	})
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
