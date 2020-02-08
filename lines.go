package golines

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type Lines struct {
	bytes.Buffer
	Parser Parser
	Source string
}

// Add wraps `bytes.Buffer.Write` handling
// `string`, `[]byte`, `[][]byte`, `[]string`, `*Lines`
// unless unhandled `x` is passed, error is always `nil`
// while passing slice, `i` will be incremented by each `Write`
func (s *Lines) Add(x interface{}) (i int, err error) {
	if t, ok := x.([]byte); ok {
		return s.Write(t)
	}
	if t, ok := x.(string); ok {
		return s.WriteString(t)
	}
	if t, ok := x.(*Lines); ok {
		return s.Write(t.Bytes())
	}
	if t, ok := x.([][]byte); ok {
		for _, l := range t {
			ii, _ := s.Write(l)
			i += ii
		}
		return
	}
	if t, ok := x.([]string); ok {
		for _, l := range t {
			ii, _ := s.WriteString(l)
			i += ii
		}
		return
	}
	// TODO
	return i, errors.New("type not handled")
}

func (s *Lines) Parse() (err error) {
	return s.ParseSource(s.Source)
}

// how many lines are in bytes?
func (s *Lines) CountBytes(in []byte) (c int) {
	for _, v := range s.ByteLines() {
		if bytes.Contains(in, v) {
			c += 1
		}
	}
	return c
}

// how many lines are in string?
func (s *Lines) CountString(in string) (c int) {
	for _, v := range s.StringLines() {
		if strings.Contains(in, v) {
			c += 1
		}
	}
	return c
}

// map of line => count
func (s *Lines) CountStringMapLineN(in string) map[string]int {
	var d = make(map[string]int)
	for _, v := range s.StringLines() {
		d[v] = strings.Count(in, v)
	}
	return d
}

// map of count => [lines...]
func (s *Lines) CountStringMapNLines(in string) map[int][]string {
	var d = make(map[int][]string)
	for _, l := range s.StringLines() {
		c := strings.Count(in, l)
		d[c] = append(d[c], l)
	}
	return d
}

// map of count => [lines...]
func (s *Lines) CountBytesMapNLines(in []byte) map[int][][]byte {
	var d = make(map[int][][]byte)
	for _, l := range s.ByteLines() {
		c := bytes.Count(in, l)
		d[c] = append(d[c], l)
	}
	return d
}

func (s *Lines) ParseSource(source string) (err error) {
	if s.Source == "" {
		// TODO
		return errors.New("source is empty")
	}
	var b []byte
	for key, action := range s.Parser.PrefixMap() {
		if strings.HasPrefix(source, key) {
			b, err = action(source)
			if err == nil {
				s.Write(b)
			}
			return
		}
	}
	// TODO
	return errors.New("prefix is not handled")
}

func (s *Lines) ForLine(x func(i int, l []byte) error) (err error) {
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

func (s *Lines) ByteLines() (data [][]byte) {
	s.ForLine(func(i int, l []byte) error {
		data = append(data, l)
		return nil
	})
	return
}

func (s *Lines) ByteLinesPreSuf(pre, suf []byte) (data [][]byte) {
	s.ForLine(func(i int, l []byte) error {
		// there must be a better way
		data = append(data, bytes.Join([][]byte{pre, l, suf}, nil))
		return nil
	})
	return
}

func (s *Lines) StringLines() (data []string) {
	s.ForLine(func(i int, l []byte) error {
		data = append(data, string(l))
		return nil
	})
	return
}

func (s *Lines) StringLinesPreSuf(pre, suf string) (data []string) {
	s.ForLine(func(i int, l []byte) error {
		data = append(data, pre+string(l)+suf)
		return nil
	})
	return
}

func (s *Lines) IntLines() (data []int, err error) {
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
