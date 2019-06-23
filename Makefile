all: build

setup:
	go get -u github.com/shurcooL/vfsgen
	go get -u github.com/nicksnyder/go-i18n/v2/goi18n
	go get -u github.com/zserge/webview

fmt:
	go fmt ./...

build:
	cd data && go generate -tags dev && cd -
	go build -ldflags '-X main.version=1.0' ./cmd/rompebolas/

run:
	go run ./cmd/rompebolas/

clean:
	go clean
	rm -f data/assets_vfsdata.go

.ONESHELL:
i18n_newlang_init:
	cd data/assets/translations/
	touch translate.$(LANG).json
	goi18n merge -format json active.en.json translate.$(LANG).json
	cd -

i18n_newlang_finish:
	rm data/assets/translations/active.$(LANG).json
	mv data/assets/translations/translate.$(LANG).json data/assets/translations/active.$(LANG).json

i18n_extract:
	goi18n extract -outdir data/assets/translations -format json

i18n_translation_init: i18n_extract
	goi18n merge -outdir /data/assets/translations/ -format json active.*.json

i18n_translation_finish:
	goi18n merge -outdir /data/assets/translations/ -format json active.*.json translate.*.json
