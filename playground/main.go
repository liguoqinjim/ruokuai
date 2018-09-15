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
}
