package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RespMessageData struct {
	Error respErrorMessage `json:"error,omitempty"`
}

type respErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserData struct {
	Email    string
	Password string
}

func APIServer() *mux.Router {
	rounter := mux.NewRouter()
	rounter.HandleFunc("api/user/signup", signUp).Methods(http.MethodPost)
	rounter.HandleFunc("api/user/login", login).Methods(http.MethodPost)
	return rounter
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ResponseError(http.StatusInternalServerError, "Internal error"))
		return
	}

	accountData := UserData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	fmt.Println(accountData)
}

func login(w http.ResponseWriter, r *http.Request) {}

func ResponseError(errorCode int, message string) []byte {
	resp, err := json.Marshal(RespMessageData{Error: respErrorMessage{Code: errorCode, Message: message}})
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp
}

func ResponseSuccess() {

}
