package main

import (
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/dto"
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/u_conf"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
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

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (this *HttpConn) Read(p []byte) (n int, err error)  { return this.in.Read(p) }
func (this *HttpConn) Write(d []byte) (n int, err error) { return this.out.Write(d) }
func (this *HttpConn) Close() error                      { return nil }

/*
--- Пример как положить токен в хранилище
curl -v -X POST -H "Content-Type: application/json" -H "X-Auth: default_key" -d '
{
	"jsonrpc":"2.0",
	"id": 1,
	"method": "Storage.SaveRequestToken",
	"params": [
		{"Idjwt":100500, "Token": "base64Header.base64Payload.base64Signature"}
	]
}' http://localhost:8081/rpc

--- Пример как забюрать токен из хранилища
curl -v -X POST -H "Content-Type: application/json" -H "X-Auth: default_key" -d '
{
	"jsonrpc":"2.0",
	"id": 1,
	"method": "Storage.GetToken",
	"params": [
		{"Path":"api_token_storage/2022/00/00/10/05/00/request_token_Sep-07_00-20-46_63"}
	]
}' http://localhost:8081/rpc
*/

type Handler struct {
	rpcServer *rpc.Server
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth key: ", r.Header.Get("X-Auth"))

	serverCodec := jsonrpc.NewServerCodec(&HttpConn{
		in:  r.Body,
		out: w,
	})
	w.Header().Set("Content-type", "application/json")
	if err := h.rpcServer.ServeRequest(serverCodec); err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, `{"error":"cant serve request"}`, 500)
	}
}

func main() {
	if err := initConfig(); err != nil {
		log.Printf("Error while initializing configuration: %v", err)
		os.Exit(-1)
	}
	storage := NewStorage(gStorageConf.PathPrefix)

	server := rpc.NewServer()
	server.Register(storage)

	sessionHandler := &Handler{
		rpcServer: server,
	}
	http.Handle(gServerConf.Endpoint, sessionHandler)

	fmt.Println("starting server at :" + gServerConf.Port)
	http.ListenAndServe(":" + gServerConf.Port, nil)
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
