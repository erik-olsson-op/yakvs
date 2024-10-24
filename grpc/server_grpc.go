package grpc

import (
	"context"
	"fmt"
	"github.com/erik-olsson-op/yakvs/engine"
	"github.com/erik-olsson-op/yakvs/logger"
	"github.com/erik-olsson-op/yakvs/model"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

// interface to pick up registered rpc calls
type server struct {
	model.UnimplementedYakvsGrpcServer
}

func (s *server) Execute(ctx context.Context, q *model.Query) (*model.Response, error) {
	response, err := engine.Parse(q)
	if err != nil {
		logger.Logger.Warning(err)
		return nil, status.Errorf(codes.Internal, "ERR: %v", err)
	}
	return response, nil
}

func Init(wg *sync.WaitGroup) {
	defer wg.Done()
	port := viper.GetString("GRPCS_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Logger.Fatalf("failed to create listener on port:  %s - %v", port, err)
	}
	cert := viper.GetString("SERVER_CERT")
	key := viper.GetString("SERVER_KEY")
	x509, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(x509))
	model.RegisterYakvsGrpcServer(grpcServer, &server{})
	logger.Logger.Infof("GRPCS server running on port: %v", port)
	reflection.Register(grpcServer)
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Logger.Fatalf("failed to start GRPC server on port: %s - %v", port, err)
	}
}
