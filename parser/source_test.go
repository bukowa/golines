package parser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func testSource(t *testing.T, b []byte){
	d := strings.Split(string(b), "\n")
	if d[0] != "package parser" {
		t.Error()
	}
	if d[len(d)-2] != "}" {
		t.Error()
	}
}

func TestParseFromFile(t *testing.T) {
	b, err := ParseFromFile("source.go")
	if err != nil {
		panic(err)
	}
	testSource(t, b)
}

func TestParseFromUrl(t *testing.T) {
	var ts = httptest.NewServer(http.FileServer(http.Dir(".")))
	defer ts.Close()
	b, err := ParseFromUrl(fmt.Sprintf("%v/%v", ts.URL, "source.go"))
	if err != nil {
		panic(err)
	}
	testSource(t, b)
}

func TestParseSourceUrl(t *testing.T) {
	var ts = httptest.NewServer(http.FileServer(http.Dir(".")))
	defer ts.Close()

	b, err := ParseSource(fmt.Sprintf("%v/%v", ts.URL, "source.go"))
	if err != nil {
		panic(err)
	}
	testSource(t, b)
}

func TestParseSourceFile(t *testing.T) {
	var ts = httptest.NewServer(http.FileServer(http.Dir(".")))
	defer ts.Close()

	b, err := ParseSource("file://source.go")
	if err != nil {
		panic(err)
	}
	testSource(t, b)
}

func TestParseSourceErr(t *testing.T) {
	b, err := ParseSource("bad://source.go")
	if !reflect.DeepEqual(err, ErrPrefixNotHandled) {
		t.Error(err)
	}
	if b != nil {
		t.Error(b)
	}
}