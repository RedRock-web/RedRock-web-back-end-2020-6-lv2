package app

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

// 用于并发
var ch1 chan int
var ch2 chan int


func Start() {
	for i := 2019211100; i < 2019215111; i++ {
		if StudentIsExist(i) {
			go GetAllClassInfo(i)
		}
	}
	<-ch1
	<-ch2
}

// 获取每页的 body 信息
func GetBody(url string) string {
	userAgent := `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36`
	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", userAgent)
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Failed to get the website information")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}
