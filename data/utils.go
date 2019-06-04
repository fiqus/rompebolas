package data

import (
	"bytes"
	"log"
)

func GetAssetStr(name string) string {
	a, err := Assets.Open(name)
	if err != nil {
		log.Fatalf("can't open asset %s, err: %v", name, err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(a)
	if err != nil {
		log.Fatalf("can't read asset %s, err: %v", name, err)
	}
	return buf.String()
}
