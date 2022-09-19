package main

import ()

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
	*out = FilePath{
		Path: "success",
	}	
	return nil
}
