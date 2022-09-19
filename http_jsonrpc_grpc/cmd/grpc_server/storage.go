package main

import (
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/fileStorage"
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
	result, err := fileStorage.New(this.pathPrefix).SaveRequest(uint(in.Idjwt), []byte(in.Token))
	if err != nil {
		return nil, err
	}

	return &grpc_storage.FilePath{
		Path: result.GetFilePath(),
	}, nil
}

func (this *Storage) GetToken(ctx context.Context, in *grpc_storage.FilePath) (*grpc_storage.Token, error) {
	result, err := fileStorage.New(this.pathPrefix).LoadFromStorage(in.Path)
	if err != nil {
		return nil, err
	}

	return &grpc_storage.Token{
		Token: string(result.GetFileBody()),
	}, nil
}
