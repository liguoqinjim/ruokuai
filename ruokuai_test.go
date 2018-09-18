package ruokuai

import (
	"log"
	"os"
	"testing"
)

//env
const (
	ENV_USERNAME = "RUOKUAI_USERNAME"
	ENV_PASSWORD = "RUOKUAI_PASSWORD"
)

var (
	username string
	password string
	app      *RuoKuaiApp
)

func init() {
	username = os.Getenv(ENV_USERNAME)
	password = os.Getenv(ENV_PASSWORD)

	if username == "" {
		log.Fatalf("env not set RUOKUAI_USERNAME")
	}
	if password == "" {
		log.Fatalf("env not set RUOKUAI_PASSWORD")
	}

	app = Default(username, password)
}

func TestRuoKuaiApp_Info(t *testing.T) {
	_, er := app.Info()
	if er != nil {
		if er.ErrorCode == "" {
			t.Errorf("error:%s", er.Error)
		} else {
			t.Errorf("error:%s,errorcode:%s", er.Error, er.ErrorCode)
		}
	}
}

func TestRuoKuaiApp_Create(t *testing.T) {
	cr, er := app.Create("3040", "testData/1.png")
	if er != nil {
		if er.ErrorCode == "" {
			t.Errorf("erros:%s", er.Error)
		} else {
			t.Errorf("error:%s,errorcode:%s", er.Error, er.ErrorCode)
		}
	}

	if cr.Result != "wwf6" {
		t.Errorf("captcha want wwf6 but got %s", cr.Result)
	}
}
