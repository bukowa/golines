package golines

import "errors"

var ErrEmptyLine = errors.New("line has empty value")

func FilterEmptyLines(l *Lines) (err error) {
	err = l.ForLine(func(i int, l []byte) (err error) {
		if len(l) < 1 {
			return ErrEmptyLine
		}
		return
	})
	return
}
