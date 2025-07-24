package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/config"
	"github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/database/postgres"
	grpchandler "github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/handler/grpc"
	"github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/service/asset"
	assetpb "github.com/Amir-Sadati/crypto-microservice/crypto-proto/proto/asset"
	"google.golang.org/grpc"
)

type App struct {
	config     *config.Config
	grpcServer *grpc.Server
	listener   net.Listener
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
	db, err := postgres.Connect(a.config.Postgres)
	if err != nil {
		log.Fatal("error in connecting to database")
	}
	defer db.Close()

	assetService := asset.NewService(db)
	assetGrpcHandler := grpchandler.NewAssetGrpcHandler(assetService)

	addr := net.JoinHostPort(a.config.GRPC.Host, a.config.GRPC.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	a.listener = listener
	a.grpcServer = grpc.NewServer()
	assetpb.RegisterAssetServiceServer(a.grpcServer, assetGrpcHandler)

	go a.serveGrpc()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")
	a.grpcServer.GracefulStop()

}

func (a *App) serveGrpc() {
	log.Printf("starting gRPC server on %v:%v", a.config.GRPC.Host, a.config.GRPC.Port)
	err := a.grpcServer.Serve(a.listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
