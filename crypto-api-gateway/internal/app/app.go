package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/config"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/handler"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/router"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/service/asset"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/service/order"
	assetpb "github.com/Amir-Sadati/crypto-microservice/crypto-proto/proto/asset"
	orderpb "github.com/Amir-Sadati/crypto-microservice/crypto-proto/proto/order"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	config     *config.Config
	grpcServer *grpc.Server
	r          *gin.Engine
	httpServer *http.Server
}

func New() *App {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error in loading config")
	}
	return &App{
		config: cfg,
	}
}

func (a *App) Run() {
	orderGrpcAddr := net.JoinHostPort(a.config.GRPC.OrderHost, a.config.GRPC.OrderPort)
	orderConn, err := grpc.NewClient(orderGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer orderConn.Close()
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	assetGrpcAddr := net.JoinHostPort(a.config.GRPC.AssetHost, a.config.GRPC.AssetPort)
	assetConn, err := grpc.NewClient(assetGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer assetConn.Close()
	assetClient := assetpb.NewAssetServiceClient(assetConn)

	v := validator.New()
	orderService := order.NewService(orderClient)
	orderHandler := handler.NewOrderHandler(orderService, v)
	assetService := asset.NewService(assetClient)
	assetHandler := handler.NewAssetHandler(assetService, v)

	r := router.New(orderHandler, assetHandler)
	a.r = r
	go a.ServeHTTP()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}

}

func (a *App) ServeHTTP() {
	a.httpServer = &http.Server{
		Addr:         net.JoinHostPort(a.config.HTTP.Host, a.config.HTTP.Port),
		Handler:      a.r,
		WriteTimeout: 10 * time.Second,
	}

	err := a.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
