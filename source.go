package golines

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

type Source struct {
	Parser Parser
	Source string
	bytes  []byte
}

// there must be a better way
func (s *Source) Close() error {
	return nil
}

func (s *Source) Read(b []byte) (n int, err error) {
	rd := bytes.NewReader(s.bytes)
	return rd.Read(b)
}

// appends b to s.bytes with '\n'
// there must be a better way
func (s *Source) Write(b []byte) (n int, err error) {
	s.SetBytes(bytes.Join([][]byte{s.bytes, b}, []byte("\n")))
	return len(b), nil
}

func (s *Source) Add(x interface{}) (err error){
	if t, ok := x.(string); ok {
		_, err = s.Write([]byte(t))
		return
	}
	if t, ok := x.([]byte); ok {
		_, err = s.Write(t)
		return
	}
	if t, ok := x.([][]byte); ok {
		for _, l := range t {
			_, err = s.Write(l)
			if err != nil {
				return err
			}
		}
		return
	}
	if t, ok := x.([]string); ok {
		for _, l := range t {
			_, err = s.Write([]byte(l))
			if err != nil {
				return err
			}
		}
		return
	}
	if t, ok := x.(Source); ok {
		err = t.ForLine(func(i int, l []byte) error {
			_, err = s.Write(l)
			return err
		})
		return
	}
	return err
}

func NewSource(v interface{}) (s *Source, err error) {
	s = &Source{}
	if t, ok := v.(string); ok {
		s.SetString(t)
	}
	if t, ok := v.([]byte); ok {
		s.SetBytes(t)
	}
	if t, ok := v.([][]byte); ok {
		for _, l := range t {
			_, err = s.Write(l)
			if err != nil {
				return s, err
			}
		}
	}
	if t, ok := v.([]string); ok {
		for _, l := range t {
			_, err = s.Write([]byte(l))
			if err != nil {
				return s, err
			}
		}
	}
	if t, ok := v.(Source); ok {
		err = t.ForLine(func(i int, l []byte) error {
			_, err = s.Write(l)
			return err
		})
	}
	return s, nil
}
func NewSourceBytes(b []byte) *Source {
	return &Source{
		bytes: b,
	}
}

func NewSourceString(s string) *Source {
	return &Source{
		bytes: []byte(s),
	}
}

func (s *Source) SetSource(source string) {
	s.Source = source
}

func (s *Source) SetBytes(b []byte) {
	s.bytes = b
}

func (s *Source) SetString(v string) {
	s.bytes = []byte(v)
}

func (s *Source) Parse() (err error) {
	return s.ParseSource(s.Source)
}

// how many lines are in bytes?
func (s *Source) CountBytes(in []byte) (c int) {
	for _, v := range s.ByteLines() {
		if bytes.ContainsAny(v, string(in)) {
			c += 1
		}
	}
	return c
}

// how many lines are in string?
func (s *Source) CountString(in string) (c int) {
	for _, v := range s.StringLines() {
		if strings.ContainsAny(in, v) {
			c += 1
		}
	}
	return c
}

// map of line => count
func (s *Source) CountStringMapLineN(in string) map[string]int {
	var d = make(map[string]int)
	for _, v := range s.StringLines() {
		d[v] = strings.Count(in, v)
	}
	return d
}

// map of count => [lines...] ; keys are unordered
func (s *Source) CountStringMapNLines(in string) map[int][]string {
	var d = make(map[int][]string)
	for _, l := range s.StringLines() {
		c := strings.Count(in, l)
		d[c] = append(d[c], l)
	}
	return d
}

// map of count => [lines...] ; keys are unordered
func (s *Source) CountBytesMapNLines(in []byte) map[int][][]byte {
	var d = make(map[int][][]byte)
	for _, l := range s.ByteLines() {
		c := bytes.Count(in, l)
		d[c] = append(d[c], l)
	}
	return d
}

func (s *Source) ParseSource(source string) (err error) {
	var b []byte
	for key, action := range s.Parser.PrefixMap() {
		if strings.HasPrefix(source, key) {
			b, err = action(source)
			if err == nil {
				s.SetBytes(b)
			}
			break
		}
	}
	return err
}

func (s *Source) ForLine(x func(i int, l []byte) error) (err error) {
	reader := bytes.NewReader(s.Bytes())
	scanner := bufio.NewScanner(reader)

	var pos = 0
	for scanner.Scan() {
		err = x(pos, scanner.Bytes())
		if err != nil {
			break
		}
		pos += 1
	}
	return err
}

func (s *Source) Bytes() []byte {
	return s.bytes
}

func (s *Source) String() string {
	return string(s.bytes)
}

func (s *Source) ByteLines() (data [][]byte) {
	s.ForLine(func(i int, l []byte) error {
		data = append(data, l)
		return nil
	})
	return
}

func (s *Source) ByteLinesPreSuf(pre, suf []byte) (data [][]byte) {
	s.ForLine(func(i int, l []byte) error {
		// there must be a better way
		data = append(data, bytes.Join([][]byte{pre, l, suf}, nil))
		return nil
	})
	return
}

func (s *Source) StringLines() (data []string) {
	s.ForLine(func(i int, l []byte) error {
		data = append(data, string(l))
		return nil
	})
	return
}

func (s *Source) StringLinesPreSuf(pre, suf string) (data []string) {
	s.ForLine(func(i int, l []byte) error {
		data = append(data, pre+string(l)+suf)
		return nil
	})
	return
}

func (s *Source) IntLines() (data []int, err error) {
	err = s.ForLine(func(i int, l []byte) error {
		i, err := strconv.Atoi(string(l))
		if err != nil {
			return err
		}
		data = append(data, i)
		return err
	})
	return
}
