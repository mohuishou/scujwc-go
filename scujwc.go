package scujwc

import (
	"errors"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/qiniu/iconv"
)

const (
	//DOMAIN 教务处ip/域名
	DOMAIN = "http://202.115.47.141"
)

//Jwc 教务处相关操作
type Jwc struct {
	uid      int
	password string
	client   http.Client
	isLogin  int
}

//Init 初始化学号和密码
func (j *Jwc) Init(uid int, password string) {
	j.password = password
	j.uid = uid
	j.initHTTP()
}

//initHTTP 初始化请求客户端
func (j *Jwc) initHTTP() {
	j.client = http.Client{}
	jar, _ := cookiejar.New(nil)
	j.client.Jar = jar
}

//Login 登录教务处
func (j *Jwc) Login() (err error) {

	url := DOMAIN + "/loginAction.do"
	param := "zjh=" + strconv.Itoa(j.uid) + "&mm=" + j.password

	doc, err := j.post(url, param)
	errinfo := doc.Find("font[color=\"#990000\"]").Text()
	if errinfo != "" {
		j.isLogin = 0
		err := errors.New(string(errinfo))
		return err
	}
	j.isLogin = 1

	return nil
}

//post 发出post请求
func (j *Jwc) post(url, param string) (*goquery.Document, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		return nil, err
	}
	req = setHeader(req)
	resp, err := j.client.Do(req)
	if err != nil {
		return nil, err
	}
	cd, err := iconv.Open("utf-8", "gbk") // convert gbk to utf8

	utfBody := iconv.NewReader(cd, resp.Body, 0)

	if err != nil {
		return nil, err
	}

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		return nil, err
	}
	return doc, err
}

//setHeader 设置header
func setHeader(req *http.Request) *http.Request {
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")
	req.Header.Set("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Forwarded-For", randIP())
	return req
}

//randIP 生成随机ip地址
func randIP() (ip string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		ip += strconv.Itoa(rand.Intn(235))
		if i != 3 {
			ip += "."
		}
	}
	return ip
}
