package router

import (
	"net/http"

	_ "github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/docs"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/handler"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func New(
	orderHandler *handler.OrderHandler,
	assetHandler *handler.AssetHandler,
) *gin.Engine {
	r := gin.New()
	r.Use(globalRecover())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // or specific frontend domain
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	v1 := r.Group("/api/v1")
	//************** Order Routes **************
	orderRoutes := v1.Group("/order")
	orderRoutes.POST("/create", orderHandler.CreateOrder)

	//************** Asset Routes **************
	assetRoutes := v1.Group("/asset")
	assetRoutes.POST("/create", assetHandler.CreateAsset)

	//************** swagger Route **************
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}

func globalRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				response.WriteFailNoData(c.Writer, http.StatusInternalServerError, "Internal Server Error", "An unexpected error occurred")
			}
		}()
		c.Next()
	}
}
