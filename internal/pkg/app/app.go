package app

import (
	"context"
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
	serv := grpc.NewServer()
	reflection.Reqistr(serv)

	netListener, err := net.Listen("tcp", ":1111")

	gatewayConn, err := grpc.DialContext(
		context.Background(),
		netListener.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	grpcGwMux := http.NewServeMux()

	return &http.Server{
		Addr:    "8082",
		Handler: grpcGwMux,
	}
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

func (a *App) runHTTPServer() error {
	return a.httpServer.ListenAndServe()
}

func (a *App) Run() error {
	err := a.runGRPCServer()
	if err != nil {
		return err
	}

	err = a.runHTTPServer()
	return err
}
