package internal

import (
	"bufio"
	"strings"
)

const (
	ReadEventHeader = iota
	ReadEventBody
	ReadDataHeader
	ReadDataBody
)

const (
	Event = iota
	Data
	EOL
)

type tokenState int
type CurrentState int

func NextFunc(br *bufio.Reader) (string, error) {

	var sb strings.Builder
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			return "", err
		}
		if line == "\n" {
			break
		}
		_, err = sb.WriteString(line)
		if err != nil {
			return "", err
		}

	}

	return sb.String(), nil
}

func ScanFunc(line string, dest ...any) {
	type state struct {
		ts  tokenState
		now int
	}
	var event strings.Builder
	var data strings.Builder

	s := state{ts: ReadEventHeader, now: 0}
	for s.now < len(line) {
		tkn := line[s.now]
		switch s.ts {
		case ReadEventHeader:
			if tkn == ':' {
				s.ts = ReadEventBody
				goto Next
			}
		case ReadEventBody:
			if tkn == '\n' {
				s.ts = ReadDataHeader
				goto Next
			}
			event.WriteByte(tkn)

		case ReadDataHeader:
			if tkn == ':' {
				s.ts = ReadDataBody
				goto Next
			}
		case ReadDataBody:
			if tkn == '\n' {
				s.ts = ReadEventHeader
				goto Next
			}
			data.WriteByte(tkn)
		}

	Next:
		s.now += 1
	}

	*dest[0].(*string) = strings.TrimSpace(event.String())
	*dest[1].(*string) = strings.TrimSpace(data.String())
}
