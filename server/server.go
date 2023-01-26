package main

import (
	"context"
	pb "grpc-auth/gen/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AuthApiServerImpl struct {
	pb.UnimplementedAuthApiServer
}

func (s *AuthApiServerImpl) GetPublicKey(ctx context.Context, req *pb.KeyRequest) (*pb.KeyResponse, error) {
	publicKey := GetPublicKey()
	return &pb.KeyResponse{PublicKey: &pb.KeyResponse_PublicKey{N: publicKey.N.String(), E: int64(publicKey.E)}}, nil
}

func (s *AuthApiServerImpl) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	decryptedLogin := Decrypt(req.EncLogin)
	decryptedPassword := Decrypt(req.EncPassword)
	log.Println("Login: ", decryptedLogin)
	log.Println("Pass:", decryptedPassword)

	return &pb.AuthResponse{Uid: decryptedLogin}, nil
}

func main() {
	Init()
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthApiServer(s, &AuthApiServerImpl{})
	err = s.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
