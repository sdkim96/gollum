package internal

import (
	"bufio"
	"errors"
	"io"
)

type Lines struct {
	brd  *bufio.Reader
	err  error
	line string
	next NextFunc
	scan ScanFunc
}

type NextFunc func(br *bufio.Reader) (string, error)
type ScanFunc func(line string, dest ...any)

func NewSSE(rd io.Reader, next NextFunc, scan ScanFunc) *Lines {
	b, ok := rd.(*bufio.Reader)
	if !ok {
		b = bufio.NewReader(rd)
	}
	ln := &Lines{
		brd:  b,
		next: next,
		scan: scan,
	}
	return ln
}

func (ln *Lines) Next() bool {
	newline, err := ln.next(ln.brd)
	if errors.Is(err, io.EOF) {
		return false
	}
	if err != nil {
		ln.err = err
		return false
	}
	ln.line = newline
	return true
}

func (ln *Lines) Scan(dest ...any) {
	ln.scan(ln.line, dest...)
}

func (ln *Lines) Err() error {
	return ln.err
}
