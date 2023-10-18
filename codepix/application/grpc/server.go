package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/luizeduu/imersao/codepix-go/application/grpc/pb"
	"github.com/luizeduu/imersao/codepix-go/application/usecase"
	"github.com/luizeduu/imersao/codepix-go/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	registerKeyUseCase := usecase.RegisterKeyUseCase{PixKeyRepository: &pixRepository}
	findKeyUseCase := usecase.FindKeyUseCase{PixKeyRepository: &pixRepository}

	pixGrpcService := NewPixGrpcService(registerKeyUseCase, findKeyUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start grpc")
	}

}
