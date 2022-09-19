package main

import (
	"net/http"
	"fmt"
)

type ServerType struct {
	ip   string
	port string
}

func New(ip, port string) *ServerType {
	return &ServerType{
		ip:   ip,
		port: port,
	}
}

func (this ServerType) getRouteMux() http.Handler {
	mux := http.NewServeMux()

	// GET
	mux.Handle("/api/token/get", this.getMethodMiddleWare(
		http.HandlerFunc(this.getToken)))
	mux.Handle("/api/token/put", this.postMethodMiddleWare(
		http.HandlerFunc(this.putToken)))

	serveMux := this.panicMiddleWare(mux)
	return serveMux
}

func (this *ServerType) getMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			this.logDebug(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "GET" {
			this.logWarning(r, nil, "wrong request method. Should be GET method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		this.logDebug(r, "request from client was received")
		next.ServeHTTP(w, r)
	})
}

func (this *ServerType) postMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			this.logDebug(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "POST" {
			this.logWarning(r, nil, "wrong request method. Should be POST method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		this.logDebug(r, "request from client was received")
		next.ServeHTTP(w, r)
	})
}

func (this *ServerType) panicMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if ok {
					this.logError(r, err, "Паника")
				} else {
					this.logError(r, nil, "Паника: %#v", rec)
				}
				this.responseError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (this *ServerType) Listen() {
	fmt.Printf("Стартую сервер на %s:%s", this.ip, this.port)
	http.ListenAndServe(this.ip+":"+this.port, this.getRouteMux())
	fmt.Printf("Кажется ip / порт заняты")
}

func (server ServerType) responseError(w http.ResponseWriter, responseStatus int, message string) {
	w.WriteHeader(responseStatus)
	w.Write([]byte(`{"error": "` + message + `"}`))
}

func (this ServerType) logDebug(r *http.Request, format string, args ...interface{}) {
	fmt.Printf("[ DEBUG ] RemoteAddr %s \tMethod %s \tURL.Path %s %s\n", r.RemoteAddr, r.Method, r.URL.Path, fmt.Sprintf(format, args...))
}

func (this ServerType) logInfo(r *http.Request, format string, args ...interface{}) {
	fmt.Printf("[ INFO ] RemoteAddr %s \tMethod %s \tURL.Path %s %s\n", r.RemoteAddr, r.Method, r.URL.Path, fmt.Sprintf(format, args...))
}

func (this ServerType) logWarning(r *http.Request, err error, format string, args ...interface{}) {
	fmt.Printf("[ WARNING ] RemoteAddr %s \tMethod %s \tURL.Path %s %s\n", r.RemoteAddr, r.Method, r.URL.Path, fmt.Sprintf(format, args...))
}

func (this ServerType) logError(r *http.Request, err error, format string, args ...interface{}) {
	fmt.Printf("[ ERROR ] RemoteAddr %s \tMethod %s \tURL.Path %s %s\n", r.RemoteAddr, r.Method, r.URL.Path, fmt.Sprintf(format, args...))
}

