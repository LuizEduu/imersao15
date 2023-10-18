package grpc

import (
	"context"

	"github.com/luizeduu/imersao/codepix-go/application/grpc/pb"
	"github.com/luizeduu/imersao/codepix-go/application/usecase"
)

type PixGrpcService struct {
	findKeyUseCase     usecase.FindKeyUseCase
	registerKeyUseCase usecase.RegisterKeyUseCase
	pb.UnimplementedPixServiceServer
}

func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.registerKeyUseCase.Execute(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "created",
	}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.findKeyUseCase.Execute(in.Key, in.Kind)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.ID,
		Kind: pixKey.Kind,
		Key:  pixKey.Key,
		Account: &pb.Account{
			AccountId:     pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

func NewPixGrpcService(registerKeyUseCase usecase.RegisterKeyUseCase, findKeyUseCase usecase.FindKeyUseCase) *PixGrpcService {
	return &PixGrpcService{
		registerKeyUseCase: registerKeyUseCase,
		findKeyUseCase:     findKeyUseCase,
	}
}
