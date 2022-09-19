package main

import (
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/grpc_storage"
	"context"
)

type Storage struct {
	pathPrefix string
}

func NewStorage(pathPrefix string) *Storage {
	return &Storage{
		pathPrefix: pathPrefix,
	}
}

func (this *Storage) SaveToken(ctx context.Context, in *grpc_storage.Token) (*grpc_storage.FilePath, error) {
	return &grpc_storage.FilePath{
		Path: "generated path",
	}, nil
}

func (this *Storage) GetToken(ctx context.Context, in *grpc_storage.FilePath) (*grpc_storage.Token, error) {
	return &grpc_storage.Token{
		Token: "header.payload.signature",
	}, nil
}
