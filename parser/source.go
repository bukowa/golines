package parser

import (
	"errors"
	"io/ioutil"
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
	for key, action := range SourcePrefixMap {
		if strings.HasPrefix(source, key) {
			b, err = action(source)
			if err == nil {
				return
			}
		}
	}
	return nil, ErrPrefixNotHandled
}

func ParseFromFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func ParseFromUrl(url string) ([]byte, error) {
	c := &http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
