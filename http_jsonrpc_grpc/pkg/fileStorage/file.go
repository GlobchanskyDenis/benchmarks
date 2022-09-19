package fileStorage

import (
	"path/filepath"
)

type file struct {
	pathPrefix  string
	path        string
	name        string
	body        []byte
}

// func (entity *file) GenerateFilePathByIdEntrant(idEntrant uint, tableName string) FileNameGenerater {
// 	entity.path = generateFilePathByIdEntrant(idEntrant, tableName)
// 	return entity
// }

// func (entity *file) GenerateFilePathByIdOrganization(idOrganization uint, tableName string) FileNameGenerater {
// 	entity.path = generateFilePathByIdOrganization(idOrganization, tableName)
// 	return entity
// }

// func (entity *file) GenerateFileName(pathPrefix string, fileExtension *string) (FileExtensionValidityChecker, error) {
// 	name, err := generateFileNameThatNotExist(pathPrefix, entity.path, fileExtension)
// 	if err != nil {
// 		return nil, err
// 	}
// 	entity.pathPrefix = pathPrefix
// 	entity.name = name
// 	entity.extension = fileExtension
// 	return entity, nil
// }

// func (entity *file) CheckExtensionValidity() (FileSaver, error) {
// 	if err := checkFileExtensionValidity(entity.extension); err != nil {
// 		return nil, err
// 	}
// 	return entity, nil
// }

func (this *file) SaveRequest(idjwt uint, fileBody []byte) (IResult, error) {
	return this.save(idjwt, fileBody, "request")
}

func (this *file) SaveResponse(idjwt uint, fileBody []byte) (IResult, error) {
	return this.save(idjwt, fileBody, "response")
}

func (this *file) save(idjwt uint, fileBody []byte, fileNamePrefix string) (IResult, error) {
	this.path = generateFilePathByIdjwt(idjwt)
	fileName, err := generateFileNameThatNotExist(this.pathPrefix, this.path, fileNamePrefix)
	if err != nil {
		return nil, err
	}
	this.name = fileName

	if err := putFile(this.pathPrefix, this.path, this.name, fileBody); err != nil {
		return nil, err
	}
	this.body = fileBody
	return this, nil
}

func (this *file) GetFilePath() string {
	return filepath.Join(this.path, this.name)
}

func (this *file) LoadFromStorage(filePathWithName string) (IResult, error) {
	body, err := getFile(filepath.Join(this.pathPrefix, filePathWithName))
	if err != nil {
		return nil, err
	}
	path, name := filepath.Split(filePathWithName)
	this.path = path
	this.name = name
	this.body = body
	return this, nil
}

// func (entity *file) GetExtension() *string {
// 	return entity.extension
// }

func (entity *file) GetFileBody() []byte {
	return entity.body
}
