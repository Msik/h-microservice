package app

import (
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	httpServer *http.Server
}

func getGrpcServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Reqistr(server)

	return server
}

func getHttpServer() *http.Server {
	return nil
}

func NewApp() (*App, error) {
	return &App{
		grpcServer: getGrpcServer(),
		httpServer: getHttpServer(),
	}, nil
}

func (a *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		return nil
	}

	if err = a.grpcServer.Serve(lis); err != nil {
		return nil
	}

	return nil
}
