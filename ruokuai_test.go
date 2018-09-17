package ruokuai

import (
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
)

func init() {
	//todo 读取用户名和密码，要是为空就先报错
}

func TestRuoKuaiApp_Info(t *testing.T) {
	app := Default(os.Getenv(ENV_USERNAME), os.Getenv(ENV_PASSWORD))

	_, er := app.Info()
	if er != nil {
		if er.ErrorCode == "" {
			t.Errorf("error:%s", er.Error)
		} else {
			t.Errorf("error:%s,errorcode:%s", er.Error, er.ErrorCode)
		}
	}
}
