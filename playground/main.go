package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	data, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	u := &User{}
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}

	log.Println("user=", u)

	//app := ruokuai.Default(u.Username, u.Password)

	//登录
	//ir, er := app.Info()
	//if ir != nil {
	//	log.Println("ir=", ir)
	//} else {
	//	log.Println("er=", er)
	//}

	//验证码
	//ir, er := app.Create("3040", "../testData/3.png")
	//if ir != nil {
	//	log.Println("ir=", ir)
	//} else {
	//	log.Println("er=", er)
	//}

	//验证码url

	//app.CreateUrl("3040","https://center.abuyun.com/captcha")
}
