package main

import (
	"bytes"
	"context"
	"os"
	"webview-automated-pishock/autoshock"

	webview "github.com/webview/webview_go"
)

func main() {

	if len(os.Args) < 4 {
		println("Not enough arguments")
		return
	}
	var user = os.Args[1]
	var code = os.Args[2]
	var key = os.Args[3]
	var name = "Autoshock-Default"
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
