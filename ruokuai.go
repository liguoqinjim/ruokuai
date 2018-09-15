package ruokuai

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

type RuoKuaiApp struct {
	Username string
	Password string
}

func New(username, password string) *RuoKuaiApp {
	return &RuoKuaiApp{
		Username: username,
		Password: password,
	}
}

func (app *RuoKuaiApp) Info() {
	request := gorequest.New()
	resp, body, errs := request.Post("http://api.ruokuai.com/info.json").
		Set("Content-Type", "multipart/form-data; boundary=-------------RK").
		SendString(`
---------------RK
Content-Disposition: form-data; name="username"

liguoqinjim
---------------RK
Content-Disposition: form-data; name="password"

wangxiu
---------------RK--`).
		End()

	if errs != nil {
		log.Fatalf("errors:%v", errs)
	}

	_ = body
	log.Println(resp.StatusCode)
}
