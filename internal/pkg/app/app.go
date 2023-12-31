package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/Msik/h-microservice/internal/app/service"
	desc "github.com/Msik/h-microservice/pkg/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	httpServer *http.Server
}

var (
	newImpl  = service.NewImplementation()
	httpPort = ":8082"
	grpcPort = ":8080"
)

func getGrpcServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(server)
	desc.RegisterApiServer(server, newImpl)

	return server
}

func getHttpServer() (*http.Server, error) {
	serv := grpc.NewServer()
	reflection.Register(serv)

	netListener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	gatewayConn, err := grpc.DialContext(
		ctx,
		netListener.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	grpcGwMux := runtime.NewServeMux()
	err = desc.RegisterApiHandler(ctx, grpcGwMux, gatewayConn)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    httpPort,
		Handler: grpcGwMux,
	}, nil
}

func NewApp() (*App, error) {
	httpSrv, err := getHttpServer()
	if err != nil {
		return nil, err
	}

	return &App{
		grpcServer: getGrpcServer(),
		httpServer: httpSrv,
	}, nil
}

func (a *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", grpcPort)
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

func (a *App) Run() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.runGRPCServer()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.runHTTPServer()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}
