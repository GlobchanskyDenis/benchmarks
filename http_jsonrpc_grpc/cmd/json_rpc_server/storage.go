package main

import (
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/fileStorage"
)

type Storage struct {
	pathPrefix string
}

type Token struct {
	Idjwt uint
	Token string
}

type FilePath struct {
	Path string
}

func NewStorage(pathPrefix string) *Storage {
	return &Storage{
		pathPrefix: pathPrefix,
	}
}

func (this *Storage) SaveRequestToken(in *Token, out *FilePath) error {
	result, err := fileStorage.New(this.pathPrefix).SaveRequest(in.Idjwt, []byte(in.Token))
	if err != nil {
		return err
	}

	*out = FilePath{
		Path: result.GetFilePath(),
	}
	
	return nil
}

func (this *Storage) SaveResponseToken(in *Token, out *FilePath) error {
	result, err := fileStorage.New(this.pathPrefix).SaveResponse(in.Idjwt, []byte(in.Token))
	if err != nil {
		return err
	}

	*out = FilePath{
		Path: result.GetFilePath(),
	}
	
	return nil
}

func (this *Storage) GetToken(in *FilePath, out *Token) error {
	result, err := fileStorage.New(this.pathPrefix).LoadFromStorage(in.Path)
	if err != nil {
		return err
	}

	*out = Token{
		Token: string(result.GetFileBody()),
	}
	return nil
}
