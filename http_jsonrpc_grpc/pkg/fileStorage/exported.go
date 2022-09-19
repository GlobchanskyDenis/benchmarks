package fileStorage

import (
)

/*	Порождаешь конструктором нужную сущность и следуешь по пайпу  */

func New(pathPrefix string) IFile {
	return &file{
		pathPrefix: pathPrefix,
	}
}

type IFile interface {
	SaveRequest(idjwt uint, fileBody []byte) (IResult, error)
	SaveResponse(idjwt uint, fileBody []byte) (IResult, error)
	LoadFromStorage(filePathWithName string) (IResult, error)
}

type IResult interface {
	GetFilePath() string
	GetFileBody() []byte
}
