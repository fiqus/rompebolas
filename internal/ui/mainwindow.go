package ui

import (
	"bytes"
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zserge/webview"
	_ "github.com/zserge/webview"
	"html/template"
	"log"
	"rompebolas/data"
	"rompebolas/internal/controller"
)

type rpcData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func handleRPC(w webview.WebView, data string) {
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

func RenderIndexTemplate() string {
	t, err := template.New("foo").Parse(data.GetAssetStr("/index.html"))
	if err != nil {
		log.Fatal("couldn't create index template")
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, struct {
		Title string
		Save  string
		Exit  string
		Ok    string
		No    string
	}{
		Title: Loc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "Title", Other: "What did you do today?"}}),
		Save:  Loc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "Save", Other: "Save"}}),
		Exit:  Loc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "Exit", Other: "Exit"}}),
		Ok:    Loc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "Ok", Other: "Ok"}}),
		No:    Loc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "No", Other: "No"}}),
	})
	if err != nil {
		log.Fatal("couln't render index template")
	}
	return buf.String()
}

func OpenWindow() {
	w := webview.New(webview.Settings{
		Width:                  640,
		Height:                 480,
		Title:                  "Rompebolas",
		Resizable:              true,
		URL:                    "data:text/html,<!doctype html><html><body></body></html>",
		ExternalInvokeCallback: handleRPC,
	})

	w.Dispatch(func() {
		_ = w.Eval(injectHtml("body", RenderIndexTemplate()))
		_ = w.Eval(injectCss(data.GetAssetStr("/main.css")))
		_ = w.Eval(injectJs(data.GetAssetStr("/main.js")))
	})

	w.SetFullscreen(true)

	defer w.Exit()
	w.Run()
}
