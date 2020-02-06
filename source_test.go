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
		bytes: []byte("1\n2\n3"),
	}
	if !reflect.DeepEqual(s.Bytes(), []byte("1\n2\n3")) {
		t.Error()
	}
	byteLines, _ := s.ByteLines()
	if !reflect.DeepEqual(byteLines, [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	stringLines, _ := s.StringLines()
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
		Source: "file://" + fileName,
	}
	err := s.Parse()
	if err != nil {
		panic(err)
	}

	// duplicate
	if !reflect.DeepEqual(s.Bytes(), []byte("1\n2\n3")) {
		t.Error()
	}
	byteLines, _ := s.ByteLines()
	if !reflect.DeepEqual(byteLines, [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	stringLines, _ := s.StringLines()
	if !reflect.DeepEqual(stringLines, []string{"1", "2", "3"}) {
		t.Error()
	}
	intLines, _ := s.IntLines()
	if !reflect.DeepEqual(intLines, []int{1, 2, 3}) {
		t.Error()
	}
}

// http://
func TestSource_ParseHttp(t *testing.T) {
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
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(s.bytes, b) {
		t.Errorf("%v not equal %v", string(s.bytes), string(b))
	}
	// duplicate
	if !reflect.DeepEqual(s.Bytes(), []byte("1\n2\n3")) {
		t.Error()
	}
	byteLines, _ := s.ByteLines()
	if !reflect.DeepEqual(byteLines, [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	stringLines, _ := s.StringLines()
	if !reflect.DeepEqual(stringLines, []string{"1", "2", "3"}) {
		t.Error()
	}
	intLines, _ := s.IntLines()
	if !reflect.DeepEqual(intLines, []int{1, 2, 3}) {
		t.Error()
	}

	byteLinesPreSuf, _ := s.ByteLinesPreSuf([]byte("a"), []byte("b"))
	if !reflect.DeepEqual(byteLinesPreSuf, [][]byte{[]byte("a1b"), []byte("a2b"), []byte("a3b")}) {
		t.Error()
	}
	stringLinesPreSuf, _ := s.StringLinesPreSuf("a", "b")
	if !reflect.DeepEqual(stringLinesPreSuf, []string{"a1b", "a2b", "a3b"}) {
		t.Error()
	}

	byteLinesPreSuf2, _ := s.ByteLinesPreSuf([]byte("a"), nil)
	if !reflect.DeepEqual(byteLinesPreSuf2, [][]byte{[]byte("a1"), []byte("a2"), []byte("a3")}) {
		t.Error()
	}
	stringLinesPreSuf2, _ := s.StringLinesPreSuf("a", "")
	if !reflect.DeepEqual(stringLinesPreSuf2, []string{"a1", "a2", "a3"}) {
		t.Error()
	}

	byteLinesPreSuf3, _ := s.ByteLinesPreSuf(nil, []byte("b"))
	if !reflect.DeepEqual(byteLinesPreSuf3, [][]byte{[]byte("1b"), []byte("2b"), []byte("3b")}) {
		t.Error()
	}
	stringLinesPreSuf3, _ := s.StringLinesPreSuf("", "b")
	if !reflect.DeepEqual(stringLinesPreSuf3, []string{"1b", "2b", "3b"}) {
		t.Error()
	}
}

func TestSource_Write(t *testing.T) {
	var b = []byte("1\n2\n3")
	s := NewSource()
	s.SetBytes(b)
	_, err := s.Write([]byte("4"))
	if err != nil {
		panic(err)
	}
	data, err := s.StringLines()
	if err != nil {
		panic(err)
	}
	if len(data) != 4 {
		t.Error(data)
	}
	if data[3] != "4" {
		t.Error(data[3])
	}
}