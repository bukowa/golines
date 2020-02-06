package golines

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Parser interface {
	PrefixMap() map[string]func(string) ([]byte, error)
}

type BasicParser struct {
	Client *http.Client
	Timeout time.Duration
}

func (p *BasicParser) PrefixMap() map[string]func(string) ([]byte, error) {
	return map[string]func(string) ([]byte, error){
		"http://": p.FromUrl,
		"https://": p.FromUrl,
		"file://": func(s string) (bytes []byte, e error) {
			return p.FromFile(strings.TrimPrefix(s, "file://"))
		},
	}
}

func (p *BasicParser) FromFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func (p *BasicParser) FromUrl(url string) ([]byte, error) {
	client := p.Client

	if client == nil {
		client = &http.Client{
			Timeout:p.Timeout,
		}
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return ioutil.ReadAll(resp.Body)

	} else {
		return nil, errors.New(string(resp.StatusCode))
	}
}
