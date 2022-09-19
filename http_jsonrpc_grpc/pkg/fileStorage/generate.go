package fileStorage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

/*	Функция для генерации уникального несуществующего в данной папке имени файла  */
func generateFileNameThatNotExist(pathPrefix, generatedFilePath, fileNamePrefix string) (string, error) {
	/*	Две попытки на создание уникального имени файла  */
	for i := 0; i < 2; i++ {
		fileName := generateFileNameWithPrefix(fileNamePrefix)
		fullPath := filepath.Join(pathPrefix, generatedFilePath, fileName)
		if isFileExists(fullPath) == false {
			return fileName, nil
		}
	}
	return "", errors.New("Не смог сгенерировать уникальное имя файла в папке " + filepath.Join(pathPrefix, generatedFilePath))
}

/*	Генерация хэша как имени файла (если расширение присутствует - его также прилепить)  */
func generateFileNameWithPrefix(fileNamePrefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s_token_%s_%d", fileNamePrefix, now.Format("Jan-02_15-04-05"), now.Nanosecond() % 100)
}

/*	Генерирует путь к файлу без префикса (префикс это путь к папке монтирования файлового хранилища)
**	то есть это путь к файлу без имени файла в самом файловом хранилище  */
func generateFilePathByIdjwt(idjwt uint) string {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	return filepath.Join("api_token_storage", year, splitIdjwtToSubFolders(idjwt))
}

/*	Разделяет id абитуриента на подпапки чтобы полностью исключить возможность создания более 1000 файлов в одной папке  */
func splitIdjwtToSubFolders(idjwt uint) string {
	nbr1 := fmt.Sprintf("%02d", idjwt/100000000)
	nbr2 := fmt.Sprintf("%02d", (idjwt/1000000)%100)
	nbr3 := fmt.Sprintf("%02d", (idjwt/10000)%100)
	nbr4 := fmt.Sprintf("%02d", (idjwt/100)%100)
	nbr5 := fmt.Sprintf("%02d", idjwt%100)
	return filepath.Join(nbr1, nbr2, nbr3, nbr4, nbr5)
}

func isFileExists(fullPath string) bool {
	if _, err := os.Stat(fullPath); err != nil {
		return false
	}
	return true
}

// func generateHash(src []byte) string {
// 	sha1Hash := sha1.Sum(src)
// 	return hex.EncodeToString([]byte(sha1Hash[:]))
// }
