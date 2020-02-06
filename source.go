package golines

import (
	"bufio"
	"bytes"
	"strconv"
)

type Source struct {
	Source string
	bytes []byte
}

func (s *Source) ForLine(x func([]byte) error) (err error) {
	scanner := bufio.NewScanner(
		bytes.NewReader(s.bytes))
	for scanner.Scan() {
		err = x(scanner.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Source) Byte() []byte {
	return s.bytes
}

func (s *Source) String() string {
	return string(s.bytes)
}

func (s *Source) ByteLines(pre, suf []byte) (data [][]byte, err error) {
	err = s.ForLine(func(l []byte) error {
		data = append(data, bytes.Join([][]byte{pre, l, suf}, nil))
		return nil
	})
	return data, err
}

func (s *Source) StringLines(pre, suf string) (data []string, err error) {
	err = s.ForLine(func(l []byte) error {
		data = append(data, pre + string(l) + suf)
		return nil
	})
	return data, err
}

func (s *Source) IntLines() (data []int, err error) {
	err = s.ForLine(func(l []byte) error {
		i, err := strconv.Atoi(string(l))
		if err != nil {
			return err
		}
		data = append(data, i)
		return nil
	})
	return data, err
}
//for _, b := range s.bytes {
//data = append(data, string(b))
//}
//return data
//func (s *Source) BytesLines() []