package golines

import "testing"

func TestFilterEmptyValuesValid(t *testing.T) {
	l := &Lines{}
	l.WriteString("1")
	err := FilterEmptyLines(l)
	if err != nil {
		panic(err)
	}
}

func TestFilterEmptyValuesInvalid(t *testing.T) {
	l := &Lines{}
	l.WriteString("1\n\n2\n")
	err := FilterEmptyLines(l)
	if err == nil {
		panic(err)
	}
}
