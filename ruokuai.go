package ruokuai

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"log"
	"strings"
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

func Default(username, password string) *RuoKuaiApp {
	app := New(username, password).
		SetSoftIdKey("112477", "84d9d802e6454936bb0119869d166220").
		SetTimeout(60)
	return app
}

//设置softId和softKey
func (app *RuoKuaiApp) SetSoftIdKey(id, key string) *RuoKuaiApp {
	app.SoftId, app.SoftKey = id, key
	return app
}

//设置timeout
func (app *RuoKuaiApp) SetTimeout(timeout int) *RuoKuaiApp {
	app.Timeout = timeout
	return app
}

//账号信息
func (app *RuoKuaiApp) Info() (*InfoResult, *ErrorResult) {
	request := gorequest.New()

	_, body, errs := request.Post("http://api.ruokuai.com/info.json").
		Type("multipart").
		SendMap(map[string]interface{}{"username": app.Username, "password": app.Password}).
		End()
	if errs != nil {
		return nil, &ErrorResult{Error: errs[0].Error()}
	}

	infoResult := &InfoResult{}
	err := json.Unmarshal([]byte(body), infoResult)
	if err != nil {
		return nil, &ErrorResult{Error: err.Error()}
	} else {
		if infoResult.Score == "" {
			errorResult := &ErrorResult{}
			err := json.Unmarshal([]byte(body), errorResult)
			if err != nil {
				return nil, &ErrorResult{Error: err.Error()}
			} else {
				return nil, errorResult
			}
		}

		return infoResult, nil
	}
}

func (app *RuoKuaiApp) Create(typeId, imagePath string) (*CreateResult, *ErrorResult) {
	request := gorequest.New()

	paths := strings.Split(imagePath, "/")
	imageName := paths[len(paths)-1]

	_, body, errs := request.Post("http://api.ruokuai.com/create.json").
		Type("multipart").
		SendMap(map[string]interface{}{"username": app.Username, "password": app.Password, "typeid": typeId, "timeout": app.Timeout, "softid": app.SoftId, "softkey": app.SoftKey}).
		SendFile(imagePath, imageName, "image", ).
		End()
	if errs != nil {
		return nil, &ErrorResult{Error: errs[0].Error()}
	}

	createResult := &CreateResult{}
	err := json.Unmarshal([]byte(body), createResult)
	if err != nil {
		return nil, &ErrorResult{Error: err.Error()}
	} else {
		if createResult.Result == "" {
			errorResult := &ErrorResult{}
			err := json.Unmarshal([]byte(body), errorResult)
			if err != nil {
				return nil, &ErrorResult{Error: err.Error()}
			} else {
				return nil, errorResult
			}
		}

		return createResult, nil
	}
}

func (app *RuoKuaiApp) CreateUrl(typeId, imageUrl string) (*CreateResult, *ErrorResult) {
	request := gorequest.New()

	_, body, errs := request.Post("http://api.ruokuai.com/create.json").
		Type("multipart").
		SendMap(map[string]interface{}{"username": app.Username, "password": app.Password, "typeid": typeId, "timeout": app.Timeout, "softid": app.SoftId, "softkey": app.SoftKey, "imageurl": imageUrl}).
		End()
	if errs != nil {
		log.Fatalf("errors:%v", errs)
	}

	createResult := &CreateResult{}
	err := json.Unmarshal([]byte(body), createResult)
	if err != nil {
		return nil, &ErrorResult{Error: err.Error()}
	} else {
		if createResult.Result == "" {
			errorResult := &ErrorResult{}
			err := json.Unmarshal([]byte(body), errorResult)
			if err != nil {
				return nil, &ErrorResult{Error: err.Error()}
			} else {
				return nil, errorResult
			}
		}

		return createResult, nil
	}
}

func (app *RuoKuaiApp) ReportError(id string) (*ReportErrorResult, *ErrorResult) {
	request := gorequest.New()

	_, body, errs := request.Post("http://api.ruokuai.com/create.json").
		Type("multipart").
		SendMap(map[string]interface{}{"username": app.Username, "password": app.Password, "softid": app.SoftId, "softkey": app.SoftKey, "id": id}).
		End()
	if errs != nil {
		log.Fatalf("errors:%v", errs)
	}

	reportErrorResult := &ReportErrorResult{}
	err := json.Unmarshal([]byte(body), reportErrorResult)
	if err != nil {
		return nil, &ErrorResult{Error: err.Error()}
	} else {
		if reportErrorResult.Result == "" {
			errorResult := &ErrorResult{}
			err := json.Unmarshal([]byte(body), errorResult)
			if err != nil {
				return nil, &ErrorResult{Error: err.Error()}
			} else {
				return nil, errorResult
			}
		}

		return reportErrorResult, nil
	}
}

type InfoResult struct {
	Score        string `json:"Score"`
	HistoryScore string `json:"HistoryScore"`
	TotalScore   string `json:"TotalScore"`
	TotalTopic   string `json:"TotalTopic"`
}

type CreateResult struct {
	Result string `json:"Result"`
	ID     string `json:"Id"`
}

type ReportErrorResult struct {
	Result string `json:"Result"`
}

type ErrorResult struct {
	Error     string `json:"Error"`
	ErrorCode string `json:"Error_Code"`
	Request   string `json:"Request"`
}
