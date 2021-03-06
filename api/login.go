package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mw "github.com/NpoolPlatform/third-login-gateway/pkg/middleware/login"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/thirdlogingateway"
)

func (s *Server) Login(ctx context.Context, in *npool.LoginRequest) (*npool.LoginResponse, error) {
	if in.GetCode() == "" {
		logger.Sugar().Error("auth login error code is empty")
		return &npool.LoginResponse{}, status.Error(codes.InvalidArgument, "code empty")
	}

	if _, err := uuid.Parse(in.GetThirdPartyID()); err != nil {
		logger.Sugar().Errorf("invalid request third party id: %v", err)
		return &npool.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorf("invalid request app id: %v", err)
		return &npool.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	resp, err := mw.Login(ctx, in.GetCode(), in.GetAppID(), in.GetThirdPartyID())
	if err != nil {
		logger.Sugar().Errorw("auth login error: %v", err)
		return &npool.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.LoginResponse{
		Info: resp,
	}, nil
}
