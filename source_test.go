package golines

import (
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