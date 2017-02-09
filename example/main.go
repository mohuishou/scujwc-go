package main

import (
	"log"
	"net/http"

	"encoding/json"

	"fmt"

	"../"
)

var s scujwc.Jwc

func login() {

}

func gpa() {

}

func gpaAll() {

}

func gpaNotPass() {

}

func isLoginHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}
}

func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type jsonData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func successReturn(w http.ResponseWriter, msg string, data interface{}) {
	jsonReturn(w, 1, msg, data)
}

func errorRetrun(w http.ResponseWriter, e error, data interface{}) {
	if e != nil {
		jsonReturn(w, 0, e.Error(), data)
	}
}

func jsonReturn(w http.ResponseWriter, status int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	d := jsonData{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	dataByte, err := json.Marshal(d)
	if err != nil {
		fmt.Fprint(w, err)
	}
	w.Write(dataByte)
}
