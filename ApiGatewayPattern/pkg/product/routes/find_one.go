package routes

import (
	"context"
	"fmt"
	"github.com/SmagulLK/APIGateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error", err)
		ctx.AbortWithError(400, err)

		return
	}
	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: id,
	})
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
