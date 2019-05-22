package main

import (
	"encoding/json"
	"log"
	"net/http"

	"APIserver_demo/token"

	"github.com/gorilla/mux"
)

type RespSuccessData struct {
	Token string `json:token`
}

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
	rounter.HandleFunc("/api/user/signup", signUp).Methods(http.MethodPost)
	rounter.HandleFunc("/api/user/login", login).Methods(http.MethodPost)
	rounter.HandleFunc("/api/user/logout", logout).Methods(http.MethodDelete)
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

	tokenString, err := token.NewToken(accountData.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ResponseError(http.StatusInternalServerError, "Internal error"))
		return
	}

	w.Write(ResponseSuccess(tokenString))
}

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

func ResponseError(errorCode int, message string) []byte {
	resp, err := json.Marshal(RespMessageData{Error: respErrorMessage{Code: errorCode, Message: message}})
	if err != nil {
		log.Println(err)
		return nil
	}

	return resp
}

func ResponseSuccess(data string) []byte {
	resp, err := json.Marshal(&RespSuccessData{Token: data})
	if err != nil {
		log.Println(err)
		return nil
	}

	return resp
}
