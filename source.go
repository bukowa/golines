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

func NewSource() *Source {
	return &Source{
		Parser: &BasicParser{},
	}
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

func (s *Source) ForLine(x func(int, []byte) error) (err error) {
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

func (s *Source) ByteLines(pre, suf []byte) (data [][]byte, err error) {
	err = s.ForLine(func(i int, l []byte) error {
		data = append(data, bytes.Join([][]byte{pre, l, suf}, nil))
		return nil
	})
	return data, err
}

func (s *Source) StringLines(pre, suf string) (data []string, err error) {
	err = s.ForLine(func(i int, l []byte) error {
		data = append(data, pre+string(l)+suf)
		return nil
	})
	return data, err
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
	return data, err
}
