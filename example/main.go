package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"../"

	"encoding/json"

	"fmt"

	"errors"
	"net/url"
	"strconv"
)

var signKey = []byte("fyscu_gpa_v2")

func login(w http.ResponseWriter, r *http.Request) (scujwc.Jwc, error) {
	var s scujwc.Jwc
	if r.Method != "POST" {
		return s, errors.New("请求方式错误")
	}
	result, _ := ioutil.ReadAll(r.Body)
	v, err := url.ParseQuery(string(result))
	if err != nil {
		return s, err
	}
	uid, err := strconv.Atoi(v.Get("uid"))
	if err != nil {
		return s, err
	}
	password := v.Get("password")
	err = s.Init(uid, password)
	if err != nil {
		return s, err
	}
	return s, nil
}

func gpa(w http.ResponseWriter, r *http.Request) {
	s, err := login(w, r)
	if err != nil {
		errorRetrun(w, err, nil)
		return
	}
	grade, err := s.GPA()
	if err != nil {
		errorRetrun(w, err, nil)
		return
	}
	successReturn(w, "本学期成绩获取成功", grade)
}

func gpaAll(w http.ResponseWriter, r *http.Request) {
	s, err := login(w, r)
	if err != nil {
		errorRetrun(w, err, nil)
		return
	}
	grade, err := s.GPAAll()
	if err != nil {
		errorRetrun(w, err, nil)
		return
	}
	successReturn(w, "所有成绩获取成功", grade)
}

func gpaNotPass(w http.ResponseWriter, r *http.Request) {
	s, err := login(w, r)
	if err != nil {
		errorRetrun(w, err, nil)
		return
	}
	grade, err := s.GPANotPass()
	if err != nil {
		errorRetrun(w, err, nil)
		return
	}
	successReturn(w, "不及格成绩获取成功", grade)
}

func main() {
	http.HandleFunc("/gpa", gpa)
	http.HandleFunc("/gpa/all", gpaAll)
	http.HandleFunc("/gpa/not-pass", gpaNotPass)

	if err := http.ListenAndServe("0.0.0.0:6627", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("服务启动成功...端口：6627")
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
	jsonReturn(w, 0, e.Error(), data)
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
	if status == 1 {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
	w.Write(dataByte)
	return
}
