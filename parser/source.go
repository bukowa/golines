package parser

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var ErrPrefixNotHandled = errors.New("prefix is not handled")

var SourcePrefixMap = map[string]func(string) ([]byte, error){
	"http://":  ParseFromUrl,
	"https://": ParseFromUrl,
	"file://": func(s string) (bytes []byte, e error) {
		return ParseFromFile(strings.TrimPrefix(s, "file://"))
	},
}

func ParseSource(source string) (b []byte, err error) {
	var match bool

	for key, action := range SourcePrefixMap {
		log.Print(source, " ", key)
		if strings.HasPrefix(source, key) {
			match = true
			b, err = action(source)
			if err != nil {
				return
			}
		}
	}
	if !match {
		err = ErrPrefixNotHandled
	}
	return
}

func ParseFromFile(path string) (b []byte, err error) {
	var f *os.File
	f, err = os.Open(path)
	if err != nil {
		return
	}
	return ioutil.ReadAll(f)
}

func ParseFromUrl(url string) (b []byte, err error) {
	var r *http.Response
	var c = &http.Client{}

	r, err = c.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()
	return ioutil.ReadAll(r.Body)
}
