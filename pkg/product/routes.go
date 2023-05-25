package product

import (
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth/middleware"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/config"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := middleware.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AdminAuth)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
