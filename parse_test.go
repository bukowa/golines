package golines

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func GetFileTest() (*os.File, string) {
	s := "testfile76656453454355"
	os.Remove(s)
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return f, s
}

func TestBasicParser_PrefixMap(t *testing.T) {
	var b = []byte("1\n2\n3")
	var ts = httptest.NewServer(http.FileServer(http.Dir(".")))
	defer ts.Close()

	f, s := GetFileTest()
	defer f.Close()
	f.Write(b)
	p := &BasicParser{}
	for k, f := range p.PrefixMap() {
		if k == "file://" {
			data, err := f("file://"+s)
			if err != nil {panic(err)}
			if !reflect.DeepEqual(data, b) {
				t.Errorf("%v not equal %v", data, b)
			}
		}
		if k == "http://" {
			data, err := f(fmt.Sprintf("%v/%v", ts.URL, s))
			if err != nil {panic(err)}
			if !reflect.DeepEqual(data, b) {
				t.Error(b)
			}
		}
	}
	f.Close()
	os.Remove(s)
}