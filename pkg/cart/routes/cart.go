package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/akshayUr04/go-grpc-api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

func AddToCart(ctx *gin.Context, c pb.CartServiceClient) {
	paramId := ctx.Param("productId")
	productId, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	userId, _ := ctx.Get("userId")
	res, err := c.AddToCart(context.Background(), &pb.AddToCartRequest{
		ProductId: productId,
		UserId:    userId.(int64),
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func RemoveFromCart(ctx *gin.Context, c pb.CartServiceClient) {
	paramId := ctx.Param("productId")
	productId, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	userId, _ := ctx.Get("userId")
	res, err := c.RemoveFromCart(context.Background(), &pb.RemoveFromCartRequest{
		ProductId: productId,
		UsreId:    userId.(int64),
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
