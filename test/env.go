package test

import (
	"testing"

	"github.com/lfhy/flag"

	glog "log"

	"github.com/lfhy/baidu-pan-client/log"
	"github.com/lfhy/baidu-pan-client/types"
)

type TestLogger struct {
}

func (t *TestLogger) Printf(f string, v ...any) {
	glog.Printf(f, v...)
}

func (t *TestLogger) Println(v ...any) {
	glog.Println(v...)
}

func ReadConfig() {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	f.String("c", "../config.toml", "配置文件")
	f.StringConfigVar(&types.ClientId,
		"client_id",
		"auth", "client_id",
		"",
		"client_id",
	)
	f.StringConfigVar(&types.ClientSecret,
		"client_secret",
		"auth", "client_secret",
		"",
		"client_secret",
	)
	f.StringConfigVar(&types.RedirectUri,
		"redirect_uri",
		"auth", "redirect_uri",
		"",
		"redirect_uri",
	)
	f.StringConfigVar(&types.AccessToken,
		"access_token",
		"auth", "access_token",
		"",
		"access_token",
	)
	f.StringConfigVar(&types.RefreshToken,
		"refresh_token",
		"auth", "refresh_token",
		"",
		"refresh_token",
	)
	f.Parse(nil)
}

func TestSetEnv(t *testing.T) {
	ReadConfig()
	log.SetLogger(&TestLogger{})
}

func PrintRes(res any, err error) {
	if err != nil {
		log.Println("发生错误:", err)
	} else {
		log.Printf("响应结果:%+v", res)
	}
}
