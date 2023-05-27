package cart

import (
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth/middleware"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/cart/routes"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := middleware.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("/cart")
	routes.Use(a.UserAuth)
	routes.POST("/add/:productId", svc.AddToCart)
	routes.POST("/remove/:productId", svc.RemoveFromCart)
}

func (svc *ServiceClient) AddToCart(ctx *gin.Context) {
	routes.AddToCart(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveFromCart(ctx *gin.Context) {
	routes.RemoveFromCart(ctx, svc.Client)
}
