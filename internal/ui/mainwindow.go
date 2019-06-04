package ui

import (
	"encoding/json"
	"github.com/zserge/webview"
	_ "github.com/zserge/webview"
	"log"
	"rompebolas/data"
	"rompebolas/internal/controller"
)

type rpcData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func handleRPC(w webview.WebView, data string) {
	log.Printf("got rpc data %s", data)
	json_data := rpcData{}
	err := json.Unmarshal([]byte(data), &json_data)
	if err != nil {
		log.Fatalf("error reading json from rpc")
	}

	switch json_data.Type {
	case "exit":
		w.Terminate()
	case "submit":
		controller.SaveToFile(json_data.Value)
	}
}

func OpenWindow() {
	w := webview.New(webview.Settings{
		Width:                  640,
		Height:                 480,
		Title:                  "Simple window demo",
		Resizable:              true,
		URL:                    "data:text/html,<!doctype html><html><body></body></html>",
		ExternalInvokeCallback: handleRPC,
	})

	w.Dispatch(func() {
		_ = w.Eval(injectHtml("body", data.GetAssetStr("/index.html")))
		_ = w.Eval(injectCss(data.GetAssetStr("/main.css")))
		_ = w.Eval(injectJs(data.GetAssetStr("/main.js")))
	})

	w.SetFullscreen(true)

	defer w.Exit()
	w.Run()
}
