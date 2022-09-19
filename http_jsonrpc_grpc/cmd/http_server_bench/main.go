package main

import (
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/dto"
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/u_conf"
	"fmt"
	"log"
	"os"
)

const (
	RED       = "\033[31m"
	RED_BG    = "\033[41;30m"
	GREEN     = "\033[32m"
	GREEN_BG  = "\033[42;30m"
	YELLOW    = "\033[33m"
	YELLOW_BG = "\033[43;30m"
	NO_COLOR  = "\033[m"
)

var (
	gStorageConf *dto.StorageConfig
	gServerConf  *dto.ServerConfig
)

func main() {
	if err := initConfig(); err != nil {
		log.Printf("Error while initializing configuration: %v", err)
		os.Exit(-1)
	}

	server := New(gServerConf.Ip, gServerConf.Port)
	server.Listen()
}

func initConfig() error {
	fmt.Print("Считываю конфигурационный файл\t\t- ")
	if err := u_conf.SetConfigFile("conf.json"); err != nil {
		fmt.Println(RED + "ошибка" + NO_COLOR)
		return err
	}
	fmt.Println(GREEN + "успешно" + NO_COLOR)

	gStorageConf = &dto.StorageConfig{}
	fmt.Print("Инициализирую инструкцию Storage\t- ")
	if err := u_conf.ParsePackageConfig(gStorageConf, "Storage"); err != nil {
		fmt.Println(RED + "ошибка" + NO_COLOR)
		return err
	}
	fmt.Println(GREEN + "успешно" + NO_COLOR)

	gServerConf = &dto.ServerConfig{}
	fmt.Print("Инициализирую инструкцию Server\t\t- ")
	if err := u_conf.ParsePackageConfig(gServerConf, "Server"); err != nil {
		fmt.Println(RED + "ошибка" + NO_COLOR)
		return err
	}
	fmt.Println(GREEN + "успешно" + NO_COLOR)
	return nil
}
