package fileStorage

import (
	// "io/fs" // TODO обновиться до версии 1.16 как минимум для того чтобы добавить данный импорт на io/fs - fs.FileMode вместо os.FileMode
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func putFile(pathPrefix, generatedFilePath, fileName string, fileBody []byte) error {
	/*	Считываю разрешиния у папочки из конфигурационника. У остальных файла и остальных папок будет такое же разрешение  */
	mode, err := getFileMode(pathPrefix)
	if err != nil {
		return err
	}

	/*	Создаю рекурсивно все папочки из переменной generatedFilePath  */
	if err := createDirectories(mode, pathPrefix, generatedFilePath); err != nil {
		return err
	}

	/*	Открываю файл для записи. Возвращается файловый дескриптор  */
	file, err := createFileDescriptor(mode, pathPrefix, generatedFilePath, fileName)
	if err != nil {
		return err
	}

	/*	Записываю данные в файл  */
	if err := writeInFile(file, fileBody); err != nil {
		_ = closeFile(file)
		return err
	}

	/*	Закрываем файл  */
	if err := closeFile(file); err != nil {
		return err
	}

	return nil
}

/*	Считываю разрешиния у папочки из конфигурационника. У остальных файла и остальных папок будет такое же разрешение  */
func getFileMode(pathPrefix string) (os.FileMode, error) {
	/*	Нахожу разрешение файла  */
	permission, err := os.Lstat(pathPrefix)
	if err != nil {
		return 0, errors.New("Не смог считать разрешение для файла. Судя по всему нет прав на папку '" + pathPrefix + "'. Проверь конфигурационник и разрешение папок " + err.Error())
	}
	return permission.Mode(), nil
}

/*	Создаю рекурсивно все папочки из переменной generatedFilePath  */
func createDirectories(mode os.FileMode, pathPrefix, generatedFilePath string) error {
	dirPath := filepath.Join(pathPrefix, generatedFilePath)

	/*	Создаю все папки в пути файла  */
	if err := os.MkdirAll(dirPath, mode); err != nil {
		return errors.New("Не смог создать папки для дальнейшего размещения файла: " + err.Error())
	}
	return nil
}

/*	Открываю файл для записи. Возвращается файловый дескриптор  */
func createFileDescriptor(mode os.FileMode, pathPrefix, generatedFilePath, fileName string) (*os.File, error) {
	fullPath := filepath.Join(pathPrefix, generatedFilePath, fileName)

	fileDescriptor, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return nil, errors.New("Не смог создать файл для записи: " + err.Error())
	}
	return fileDescriptor, nil
}

/*	Записываю данные в файл  */
func writeInFile(file *os.File, fileBody []byte) error {
	if _, err := file.Write(fileBody); err != nil {
		return errors.New("Не смог записать данные в файл: " + err.Error())
	}
	return nil
}

/*	Закрытие файлового дескриптора  */
func closeFile(file *os.File) error {
	if err := file.Close(); err != nil {
		return errors.New("Не смог закрыть файл: " + err.Error())
	}
	return nil
}

// TODO при переходе на новую версию (1.17, 1.18) отказаться от ioutil так как этот пакет deprecated
/*	Считываем из хранилища файл  */
func getFile(fullPath string) ([]byte, error) {
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, errors.New("не смог считать файл с диска " + err.Error())
	}
	return data, nil
}
