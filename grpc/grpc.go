package grpc

import (
	"book/config"
	"book/genproto/book_service"
	"book/grpc/client"
	"book/grpc/service"
	"book/pkg/logger"
	"book/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	book_service.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
