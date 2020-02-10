package golines

import (
	"reflect"
	"testing"
)

var checkDeepEqual = func(t *testing.T, x, desired interface{}) {
	if !reflect.DeepEqual(x, desired) {
		t.Errorf("'%v' IS NOT '%v'", x, desired)
	}
}

func TestDataTest(t *testing.T) {

	var testData = []byte(`
<HTML>
<TITLE>my website</TITLE>
<H1>This is a Header</H1>
<H2>This is a Medium Header</H2>
<body>
support@example.com
<P> This is a new paragraph!
<P><B>This is a new paragraph!</B>
<BR><B><I>This is a new sentence without a paragraph break, in bold italics.</I></B>
</BODY>
</HTML> 
`)
	var testBytes = []byte(`<HTML>
<TITLE>
@
This is a new paragraph
<body>
nope
bad`)

	s := &Lines{}
	s.Write(testBytes)
	checkDeepEqual(t, s.CountBytes(testData), 5)
	checkDeepEqual(t, s.CountString(string(testData)), 5)
	checkDeepEqual(t, s.CountStringMapLineN(string(testData)), map[string]int{
		"<HTML>":                  1,
		"<TITLE>":                 1,
		"<body>":                  1,
		"@":                       1,
		"This is a new paragraph": 2,
		"bad":                     0,
		"nope":                    0,
	})
	checkDeepEqual(t, s.CountStringMapNLines(string(testData)), map[int][]string{
		1: {"<HTML>", "<TITLE>", "@", "<body>"},
		2: {"This is a new paragraph"},
		0: {"nope", "bad"},
	})
	checkDeepEqual(t, s.CountBytesMapNLines(testData), map[int][][]byte{
		1: {[]byte("<HTML>"), []byte("<TITLE>"), []byte("@"), []byte("<body>")},
		2: {[]byte("This is a new paragraph")},
		0: {[]byte("nope"), []byte("bad")},
	})
}

func TestLines_Add(t *testing.T) {
	var b = []byte("1\n2\n3")
	var desired = "1\n2\n3\n4\n5\n6454\n7\n467\n3\n5\n3\n5\n9\n7\n4\n2\n0"
	s := &Lines{}
	s.Write(b)

	ss := &Lines{}
	ss.WriteString("\n4\n5\n6")

	if ss.String() != "\n4\n5\n6" {
		t.Error(ss.String())
	}

	if _, err := s.Add(ss); err != nil {
		panic(err)
	}
	if _, err := s.Add([]byte("454\n7")); err != nil {
		panic(err)
	}
	if _, err := s.Add([][]byte{[]byte("\n467\n3"), []byte("\n5\n3")}); err != nil {
		panic(err)
	}
	if _, err := s.Add("\n5\n9\n7"); err != nil {
		panic(err)
	}
	if _, err := s.Add([]string{"\n4\n2\n0"}); err != nil {
		panic(err)
	}

	checkDeepEqual(t, s.String(), desired)
}

func TestLines_CountBytes(t *testing.T) {
	var b = []byte("1\n2\n3")
	var in = []byte("1askljhfipouiouj23dsfgfdg45asdasd6asd")
	s := &Lines{}
	s.Write(b)
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

func TestLines_CountString(t *testing.T) {
	var str = "1\n2\n3"
	var in = "123dsfdsf4sdf5sdfsdf6sdfsdf"
	s := &Lines{}
	s.WriteString(str)
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

func TestLines_CountStringMapLineN(t *testing.T) {
	var str = "1\n2\n3"
	var in = "1sdfsdf12sdfsdf223sdfsdf3dsfsdf34"
	var desired = map[string]int{"1": 2, "2": 3, "3": 3}

	s := &Lines{}
	s.WriteString(str)
	m := s.CountStringMapLineN(in)

	if !reflect.DeepEqual(m, desired) {
		t.Error(m)
	}
}

// this may fail, because order is not guaranteed
// we should sort the keys in this test
func TestLines_CountStringMapNLines(t *testing.T) {
	var str = "1\n2\n3\n4\n5\n6"
	var in = "11asd222asd3asd3asd34"
	var desired = map[int][]string{
		0: {"5", "6"},
		2: {"1"},
		3: {"2", "3"},
		1: {"4"},
	}
	s := &Lines{}
	s.WriteString(str)
	m := s.CountStringMapNLines(in)
	checkDeepEqual(t, m, desired)
}

func TestLines_CountBytesMapNLines(t *testing.T) {
	var b = []byte("1\n2\n3\n4\n5\n6")
	var in = []byte("112asd2asd233asd34")
	var desired = map[int][][]byte{
		0: {[]byte("5"), []byte("6")},
		2: {[]byte("1")},
		3: {[]byte("2"), []byte("3")},
		1: {[]byte("4")},
	}
	s := &Lines{}
	s.Write(b)
	m := s.CountBytesMapNLines(in)
	checkDeepEqual(t, m, desired)
}

func TestLinesBytesStringsLines(t *testing.T) {
	var b = []byte("1\n2\n3")
	s := &Lines{}
	s.Write(b)
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

func TestLines_StringAndBytesLinesPreSuf(t *testing.T) {
	s := &Lines{}
	s.WriteString("1\n2\n3")

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
