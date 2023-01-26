package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	pb "grpc-auth/gen/proto"
	"log"
	"math/big"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewAuthApiClient(conn)
	keyResp, err := client.GetPublicKey(context.Background(), &pb.KeyRequest{})
	n := new(big.Int)
	fmt.Sscan(keyResp.PublicKey.N, n)
	publicKey := rsa.PublicKey{
		N: n,
		E: int(keyResp.PublicKey.E),
	}

	encLogin, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte("Hello world"), nil)
	if err != nil {
		log.Println(err)
	}
	encPassword, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte("simplepassword"), nil)
	if err != nil {
		log.Println(err)
	}

	_, err = client.Login(context.Background(), &pb.AuthRequest{EncLogin: encLogin, EncPassword: encPassword})
}
