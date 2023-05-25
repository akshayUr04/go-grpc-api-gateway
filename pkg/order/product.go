package order

import (
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth/middleware"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/config"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := middleware.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.UserAuth)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
