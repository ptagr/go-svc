package main

import (
	"net/http"
	"encoding/json"
	"github.com/cloudflare/cfssl/log"
)

type StatusResponse struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", StatusPage)
	mux.HandleFunc("/k8s", HealthCheck)

	log.Info("Start server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	StatusPage(w,r)
}

func StatusPage(w http.ResponseWriter, r *http.Request) {
	resp := StatusResponse{
		Message: "Hello World",
	}

	respJson, err := json.Marshal(resp)
	Panic(err)

	w.Header().Set("Content-type", "application/json")
	w.Write(respJson)
}

func Panic(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}