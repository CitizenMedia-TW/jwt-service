package grpcapp

import (
	"auth-service/internal/config"
	"auth-service/internal/helper"
	"auth-service/internal/models"
	"auth-service/proto/auth-service"
	"context"
	"strings"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	config config.Config
	Server *grpc.Server
	auth.UnimplementedAuthServiceServer
}

func New(cnf config.Config) *GrpcServer {
	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &GrpcServer{config: cnf})
	return &GrpcServer{config: cnf, Server: s}
}

type GRPCServer interface {
	SayHello(context.Context, *auth.Empty) (*auth.HelloReply, error)
}

func (s *GrpcServer) SayHello(ctx context.Context, in *auth.Empty) (*auth.HelloReply, error) {
	return &auth.HelloReply{Message: "hello"}, nil
}

func (s *GrpcServer) GenerateToken(ctx context.Context, in *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error) {
	for _, val := range []string{in.Id, in.Mail, in.Name} {
		if val == "" {
			return &auth.GenerateTokenResponse{Message: "Failed", Token: ""}, nil
		}
	}

	signContent := models.JWTContent{
		Id:   in.Id,
		Mail: in.Mail,
		Name: in.Name,
	}
	token, err := helper.JWTSignContent(signContent, s.config.Secret)
	if err != nil {
		return &auth.GenerateTokenResponse{
			Message: "Failed",
			Token:   "",
		}, err
	}

	return &auth.GenerateTokenResponse{
		Message: "Success",
		Token:   "Bearer " + token,
	}, nil
}

func (s *GrpcServer) VerifyToken(ctx context.Context, in *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	if strings.Split(in.Token, " ")[0] != "Bearer" {
		return &auth.VerifyTokenResponse{
			Message:    "Failed",
			JwtContent: nil,
		}, nil
	}

	token := strings.Split(in.Token, " ")[1]

	claims, err := helper.JWTParseToken(token, s.config.Secret)
	if err != nil {
		return &auth.VerifyTokenResponse{
			Message:    "Failed",
			JwtContent: nil,
		}, err
	}

	return &auth.VerifyTokenResponse{
		Message: "Success",
		JwtContent: &auth.JwtContent{
			Id:   claims.Id,
			Mail: claims.Mail,
			Name: claims.Name,
		},
	}, nil
}
