package grpcapp

import (
	"context"
	"jwt-service/internal/config"
	"jwt-service/internal/helper"
	"jwt-service/internal/models"
	"jwt-service/protobuffs/jwt-service"
	"strings"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	config config.Config
	Server *grpc.Server
	jwt.UnimplementedJWTServiceServer
}

func New(cnf config.Config) *GrpcServer {
	s := grpc.NewServer()
	jwt.RegisterJWTServiceServer(s, &GrpcServer{config: cnf})
	return &GrpcServer{config: cnf, Server: s}
}

type GRPCServer interface {
	SayHello(context.Context, *jwt.Empty) (*jwt.HelloReply, error)
}

func (s *GrpcServer) SayHello(ctx context.Context, in *jwt.Empty) (*jwt.HelloReply, error) {
	return &jwt.HelloReply{Message: "hello"}, nil
}

func (s *GrpcServer) GenerateToken(ctx context.Context, in *jwt.GenerateTokenRequest) (*jwt.GenerateTokenResponse, error) {
	for _, val := range []string{in.Mail, in.Name} {
		if val == "" {
			return &jwt.GenerateTokenResponse{Message: "Failed", Token: ""}, nil
		}
	}

	signContent := models.JWTContent{
		Mail: in.Mail,
		Name: in.Name,
	}
	token, err := helper.JWTSignContent(signContent, s.config.Secret)
	if err != nil {
		return &jwt.GenerateTokenResponse{
			Message: "Failed",
			Token:   "",
		}, err
	}

	return &jwt.GenerateTokenResponse{
		Message: "Success",
		Token:   "Bearer " + token,
	}, nil
}

func (s *GrpcServer) VerifyToken(ctx context.Context, in *jwt.VerifyTokenRequest) (*jwt.VerifyTokenResponse, error) {
	if strings.Split(in.Token, " ")[0] != "Bearer" {
		return &jwt.VerifyTokenResponse{
			Message:    "Failed",
			JwtContent: nil,
		}, nil
	}

	token := strings.Split(in.Token, " ")[1]

	claims, err := helper.JWTParseToken(token, s.config.Secret)
	if err != nil {
		return &jwt.VerifyTokenResponse{
			Message:    "Failed",
			JwtContent: nil,
		}, err
	}

	return &jwt.VerifyTokenResponse{
		Message: "Success",
		JwtContent: &jwt.JwtContent{
			Mail: claims.Mail,
			Name: claims.Name,
		},
	}, nil
}
