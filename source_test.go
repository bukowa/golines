package golines

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSourceBytesStringsLines(t *testing.T) {
	s := &Source{
		bytes:  []byte("1\n2\n3"),
	}
	if !reflect.DeepEqual(s.Byte(), []byte("1\n2\n3")) {
		t.Error()
	}
	byteLines, _ := s.ByteLines(nil, nil)
	if !reflect.DeepEqual(byteLines, [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	stringLines, _ := s.StringLines("", "")
	if !reflect.DeepEqual(stringLines, []string{"1", "2", "3"}) {
		t.Error()
	}
	intLines, _ := s.IntLines()
	if !reflect.DeepEqual(intLines, []int{1, 2, 3}) {
		t.Error()
	}
}

// file://
func TestSource_ParseFile(t *testing.T) {
	var b = []byte("1\n2\n3")
	var ts = httptest.NewServer(http.FileServer(http.Dir(".")))
	defer ts.Close()

	file, fileName := GetFileTest()
	defer file.Close()
	file.Write(b)

	s := &Source{
		Parser: &BasicParser{},
		Source: "file://"+fileName,
	}
	err := s.Parse()
	if err != nil {panic(err)}

	// duplicate
	if !reflect.DeepEqual(s.Byte(), []byte("1\n2\n3")) {
		t.Error()
	}
	byteLines, _ := s.ByteLines(nil, nil)
	if !reflect.DeepEqual(byteLines, [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	stringLines, _ := s.StringLines("", "")
	if !reflect.DeepEqual(stringLines, []string{"1", "2", "3"}) {
		t.Error()
	}
	intLines, _ := s.IntLines()
	if !reflect.DeepEqual(intLines, []int{1, 2, 3}) {
		t.Error()
	}
}

// http://
func TestSource_ParseHttp(t *testing.T){
	var b = []byte("1\n2\n3")
	var ts = httptest.NewServer(http.FileServer(http.Dir(".")))
	defer ts.Close()

	file, fileName := GetFileTest()
	defer file.Close()
	file.Write(b)

	s := &Source{
		Parser: &BasicParser{},
		Source: fmt.Sprintf("%v/%v", ts.URL, fileName),
	}
	err := s.Parse()
	if err != nil {panic(err)}
	if !reflect.DeepEqual(s.bytes, b) {
		t.Errorf("%v not equal %v", string(s.bytes), string(b))
	}
	// duplicate
	if !reflect.DeepEqual(s.Byte(), []byte("1\n2\n3")) {
		t.Error()
	}
	byteLines, _ := s.ByteLines(nil, nil)
	if !reflect.DeepEqual(byteLines, [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	stringLines, _ := s.StringLines("", "")
	if !reflect.DeepEqual(stringLines, []string{"1", "2", "3"}) {
		t.Error()
	}
	intLines, _ := s.IntLines()
	if !reflect.DeepEqual(intLines, []int{1, 2, 3}) {
		t.Error()
	}
}