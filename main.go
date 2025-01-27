package main

import (
	"bytes"
	"context"
	"os"
	"webview-automated-pishock/autoshock"

	"github.com/spf13/viper"

	webview "github.com/webview/webview_go"
)

type shockconfig struct {
	APIKEY   string
	CODES    []string
	UNAME    string
	NICKNAME string
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	var cfg shockconfig
	//var cfgbuf []byte
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	//viper.ReadConfig(bytes.NewReader(cfgbuf))
	viper.Unmarshal(&cfg)

	user := cfg.UNAME
	code := cfg.CODES
	key := cfg.APIKEY
	name := cfg.NICKNAME
	if len(os.Args) > 4 {
		name = os.Args[4]
	}

	var buf bytes.Buffer
	autoshock.Index(key, code, user, name).Render(context.Background(), &buf)

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Autoshock")
	w.SetSize(480, 1000, webview.HintNone)
	w.SetHtml(buf.String())
	w.Run()
}
