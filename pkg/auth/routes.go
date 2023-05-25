package auth

import (
	"fmt"

	"github.com/akshayUr04/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/akshayUr04/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.POST("/admin/register", svc.AdminRegister)
	routes.POST("/admin/login", svc.AdminLogin)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) AdminRegister(ctx *gin.Context) {
	fmt.Println("---adminregister r----")
	routes.AdminRegister(ctx, svc.Client)
}

func (svc *ServiceClient) AdminLogin(ctx *gin.Context) {
	fmt.Println("---adminLogin r----")
	routes.AdminLogin(ctx, svc.Client)
}
