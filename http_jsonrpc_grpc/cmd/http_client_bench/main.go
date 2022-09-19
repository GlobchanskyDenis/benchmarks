package main

import (
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/dto"
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/http_request"
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/u_conf"
	"time"
	"fmt"
	"log"
	"sync"
	"os"
)

var (
	gServerConf  *dto.ServerConfig
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

func makeSomeRequestsAndMarkTheTime(requestAmount uint) error {
	timeStart := time.Now()
	for i := uint(0); i < requestAmount; i++ {
		if _, err := http_request.Send([]byte(`{"Idjwt": 100500, "Token":"header.payload.signature"}`),
			"http://" + gServerConf.Ip + ":" + gServerConf.Port, "/api/token/put"); err != nil {
			log.Printf("Error while request to server: %v", err)
			os.Exit(-1)
		}
	}
	totalTime := time.Now().Sub(timeStart)
	if int(totalTime.Seconds()) > 0 {
		fmt.Printf("Total time for %d requests %d sec %d millisec (%d millisec per request)\n", requestAmount, int(totalTime.Seconds()), totalTime.Milliseconds() % 1000, uint(totalTime.Milliseconds()) / requestAmount)
	} else {
		fmt.Printf("Total time for %d requests %d millisec %d microsec (%d microsec per request)\n", requestAmount, totalTime.Milliseconds(), totalTime.Microseconds() % 1000, uint(totalTime.Microseconds()) / requestAmount)
	}
	return nil
}

func main() {
	if err := initConfig(); err != nil {
		log.Printf("Error while initializing configuration: %v", err)
		os.Exit(-1)
	}

	fmt.Printf("%sTry to send %d requests in 1 stream%s\n", GREEN, 100, NO_COLOR)
	if err := makeSomeRequestsAndMarkTheTime(100); err != nil {
		log.Printf("Error while initializing configuration: %v", err)
		os.Exit(-1)
	}

	fmt.Printf("%sTry to send %d requests in 5 streams%s\n", GREEN, 20, NO_COLOR)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			if err := makeSomeRequestsAndMarkTheTime(100); err != nil {
				log.Printf("Error while initializing configuration: %v", err)
			}
			wg.Done()
		}(wg)
	}

	wg.Wait()

	fmt.Println(GREEN + "Тест окончен успешно" + NO_COLOR)
}

func initConfig() error {
	fmt.Print("Считываю конфигурационный файл\t\t- ")
	if err := u_conf.SetConfigFile("conf.json"); err != nil {
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
