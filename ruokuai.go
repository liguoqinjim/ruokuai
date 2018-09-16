package ruokuai

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

type RuoKuaiApp struct {
	Username string
	Password string
	SoftId   string
	SoftKey  string
	Timeout  int
}

func New(username, password string) *RuoKuaiApp {
	return &RuoKuaiApp{
		Username: username,
		Password: password,
	}
}

func Defualt(username, password string) *RuoKuaiApp {
	app := New(username, password).
		SetSoftId("112477", "84d9d802e6454936bb0119869d166220").
		SetTimeout(60)
	return app
}

//设置softId和softKey
func (app *RuoKuaiApp) SetSoftId(id, key string) *RuoKuaiApp {
	app.SoftId, app.SoftKey = id, key
	return app
}

//设置timeout
func (app *RuoKuaiApp) SetTimeout(timeout int) *RuoKuaiApp {
	app.Timeout = timeout
	return app
}

//账号信息
func (app *RuoKuaiApp) Info() {
	request := gorequest.New()

	resp, body, errs := request.Post("http://api.ruokuai.com/info.json").
		Type("multipart").
		SendMap(map[string]interface{}{"username": app.Username, "password": app.Password}).
		End()
	if errs != nil {
		log.Fatalf("errors:%v", errs)
	}

	log.Println(body)
	log.Println(resp.StatusCode)
}

func (app *RuoKuaiApp) Create(typeId, imagePath string) {
	request := gorequest.New()

	resp, body, errs := request.Post("http://api.ruokuai.com/create.json").
		Type("multipart").
		SendMap(map[string]interface{}{"username": app.Username, "password": app.Password, "typeid": typeId, "timeout": app.Timeout, "softid": app.SoftId, "softkey": app.SoftKey}).
		SendFile("../testData/1.png", "2.png", "image", ).
		End()
	if errs != nil {
		log.Fatalf("errors:%v", errs)
	}

	log.Println(body)
	log.Println(resp.StatusCode)
}

func (app *RuoKuaiApp) Create
