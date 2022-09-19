package main

import (
	"encoding/json"
	"net/http"
)

type Token struct {
	Idjwt uint
	Token string
}

type FilePath struct {
	Path string
}

func (this *ServerType) getToken(w http.ResponseWriter, r *http.Request) {
	var request FilePath
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		this.logError(r, err, "При попытке считывании запроса")
		this.responseError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	var responseDto = Token{
		Idjwt: 100500,
		Token: "header.payload.signature",
	}

	responseJson, err := json.Marshal(responseDto)
	if err != nil {
		this.logError(r, err, "При попытке маршаллинга статистики")
		this.responseError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(responseJson)
	this.logInfo(r, "Отдал токен")
}

func (this *ServerType) putToken(w http.ResponseWriter, r *http.Request) {
	var request Token
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		this.logError(r, err, "При попытке считывании запроса")
		this.responseError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	var responseDto = FilePath{
		Path:  "some generated path",
	}

	responseJson, err := json.Marshal(responseDto)
	if err != nil {
		this.logError(r, err, "При попытке маршаллинга")
		this.responseError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(responseJson)
	this.logInfo(r, "Положил токен")
}
