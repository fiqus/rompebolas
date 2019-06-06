package main

import (
	"flag"
	"rompebolas/internal/ui"
)

func main() {
	lang := flag.String("lang", "", "Override ui language")

	flag.Parse()

	ui.SetLanguage(*lang)
	ui.OpenWindow()
}
