package data

import (
	"bytes"
	"log"
)

func GetAssetBytes(name string) []byte {
	a, err := Assets.Open(name)
	if err != nil {
		log.Fatalf("can't open asset %s, err: %v", name, err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(a)
	if err != nil {
		log.Fatalf("can't read asset %s, err: %v", name, err)
	}
	return buf.Bytes()
}

func GetAssetStr(name string) string {
	return string(GetAssetBytes(name))
}
