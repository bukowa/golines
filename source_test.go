package golines

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var checkDeepEqual = func(t *testing.T, x, desired interface{}) {
	if !reflect.DeepEqual(x, desired) {
		t.Errorf("'%v' IS NOT '%v'", x, desired)
	}
}

func TestSource_Write(t *testing.T) {
	var b = []byte("1\n2\n3")
	s := &Source{}
	s.SetBytes(b)
	_, err := s.Write([]byte("4"))
	if err != nil {
		panic(err)
	}
	data := s.StringLines()
	if len(data) != 4 {
		t.Error(data)
	}
	if data[3] != "4" {
		t.Error(data[3])
	}
}

func TestNewSourceFromBytes(t *testing.T) {
	var b = []byte("1\n2\n3")
	s := NewSourceBytes(b)
	checkDeepEqual(t, s.bytes, b)
}

func TestNewSourceFromString(t *testing.T) {
	var str = "1\n2\n3"
	var desired = []byte(str)
	s := NewSourceString(str)
	checkDeepEqual(t, s.bytes, desired)
}

func TestSource_SetSource(t *testing.T) {
	var desired = "test"
	s := &Source{}
	s.SetSource(desired)
	checkDeepEqual(t, s.Source, desired)
}

func TestSource_SetBytes(t *testing.T) {
	var desired = []byte("byte")
	s := &Source{}
	s.SetBytes(desired)
	checkDeepEqual(t, s.bytes, desired)
}

func TestSource_SetString(t *testing.T) {
	var str = "test"
	var desired = []byte(str)
	s := &Source{}
	s.SetString(str)
	checkDeepEqual(t, s.bytes, desired)
}

//func TestSource_Parse(t *testing.T) {
//	s := &Source{}
//}
func TestSource_CountBytes(t *testing.T) {
	var b = []byte("1\n2\n3")
	var in = []byte("123456")
	s := NewSourceBytes(b)
	c := s.CountBytes(in)
	if c != 3 {
		t.Error(c)
	}
	var in2 = []byte("12")
	c = s.CountBytes(in2)
	if c != 2 {
		t.Error(c)
	}
}

func TestSource_CountString(t *testing.T) {
	var str = "1\n2\n3"
	var in = "123456"
	s := NewSourceString(str)
	c := s.CountString(in)
	if c != 3 {
		t.Error()
	}
	var in2 = "12"
	c = s.CountString(in2)
	if c != 2 {
		t.Error()
	}
}

func TestSource_CountStringMapLineN(t *testing.T) {
	var str = "1\n2\n3"
	var in = "112223334"
	var desired = map[string]int{"1": 2, "2": 3, "3":3}

	s := NewSourceString(str)
	m := s.CountStringMapLineN(in)

	if !reflect.DeepEqual(m, desired) {
		t.Error(m)
	}
}

// this may fail, because order is not guaranteed
// we should sort the keys later
func TestSource_CountStringMapNLines(t *testing.T) {
	var str = "1\n2\n3\n4\n5\n6"
	var in = "112223334"
	var desired = map[int][]string{
		0: {"5", "6"},
		2: {"1"},
		3: {"2", "3"},
		1: {"4"},
	}
	s := NewSourceString(str)
	m := s.CountStringMapNLines(in)
	checkDeepEqual(t, m, desired)
}

func TestSource_CountBytesMapNLines(t *testing.T) {
	var b = []byte("1\n2\n3\n4\n5\n6")
	var in = []byte("112223334")
	var desired = map[int][][]byte{
		0: {[]byte("5"), []byte("6")},
		2: {[]byte("1")},
		3: {[]byte("2"), []byte("3")},
		1: {[]byte("4")},
	}
	s := NewSourceBytes(b)
	m := s.CountBytesMapNLines(in)
	checkDeepEqual(t, m, desired)
}

func TestSourceBytesStringsLines(t *testing.T) {
	s := &Source{
		bytes: []byte("1\n2\n3"),
	}
	if !reflect.DeepEqual(s.Bytes(), []byte("1\n2\n3")) {
		t.Error()
	}
	if !reflect.DeepEqual(s.ByteLines(), [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	if !reflect.DeepEqual(s.StringLines(), []string{"1", "2", "3"}) {
		t.Error()
	}
	intLines, _ := s.IntLines()
	if !reflect.DeepEqual(intLines, []int{1, 2, 3}) {
		t.Error()
	}
}

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
	if !reflect.DeepEqual(s.ByteLines(), [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	if !reflect.DeepEqual(s.StringLines(), []string{"1", "2", "3"}) {
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
	if !reflect.DeepEqual(s.ByteLines(), [][]byte{[]byte("1"), []byte("2"), []byte("3")}) {
		t.Error()
	}
	if s.String() != "1\n2\n3" {
		t.Error()
	}
	if !reflect.DeepEqual(s.StringLines(), []string{"1", "2", "3"}) {
		t.Error()
	}
	intLines, _ := s.IntLines()
	if !reflect.DeepEqual(intLines, []int{1, 2, 3}) {
		t.Error()
	}

	if !reflect.DeepEqual(s.ByteLinesPreSuf([]byte("a"), []byte("b")), [][]byte{[]byte("a1b"), []byte("a2b"), []byte("a3b")}) {
		t.Error()
	}
	if !reflect.DeepEqual(s.StringLinesPreSuf("a", "b"), []string{"a1b", "a2b", "a3b"}) {
		t.Error()
	}

	if !reflect.DeepEqual(s.ByteLinesPreSuf([]byte("a"), nil), [][]byte{[]byte("a1"), []byte("a2"), []byte("a3")}) {
		t.Error()
	}
	if !reflect.DeepEqual(s.StringLinesPreSuf("a", ""), []string{"a1", "a2", "a3"}) {
		t.Error()
	}

	if !reflect.DeepEqual(s.ByteLinesPreSuf(nil, []byte("b")), [][]byte{[]byte("1b"), []byte("2b"), []byte("3b")}) {
		t.Error()
	}
	if !reflect.DeepEqual(s.StringLinesPreSuf("", "b"), []string{"1b", "2b", "3b"}) {
		t.Error()
	}
}
