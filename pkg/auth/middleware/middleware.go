package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *auth.ServiceClient
}

func InitAuthMiddleware(svc *auth.ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) UserAuth(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if res.Role == "user" {
		ctx.Set("userId", res.Id)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Next()
}

func (c *AuthMiddlewareConfig) AdminAuth(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if res.Role == "admin" {
		ctx.Set("adminId", res.Id)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Next()
}
