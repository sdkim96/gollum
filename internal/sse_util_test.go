package internal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func mockOpenAISSE() io.Reader {
	return strings.NewReader(
		`event: response.created
data: {"type":"response.created","response":{"id":"resp_123"}}

event: response.output_text.delta
data: {"type":"response.output_text.delta","delta":"Hello"}

event: response.output_text.delta
data: {"type":"response.output_text.delta","delta":" world"}

event: response.completed
data: {"type":"response.completed"}

data: [DONE]

`)
}

func nextEscapeNFunc(br *bufio.Reader) (string, error) {

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

//event: response.created\ndata: {"type":"response.created","response":{"id":"resp_123"}}

func scanFunc(line string, dest ...any) {
	type state struct {
		ts  tokenState
		now int
	}
	var event strings.Builder
	var data strings.Builder

	s := state{ts: ReadEventHeader, now: 0}
	for s.now < len(line) {
		tkn := line[s.now]
		// fmt.Printf("Token: %c\n, Event: [%s], Data: [%s]", tkn, event.String(), data.String())
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

func TestSSE_Debug(t *testing.T) {
	r := mockOpenAISSE()
	ln := NewSSE(r, nextEscapeNFunc, scanFunc)

	for ln.Next() {

		var ev, data string

		ln.Scan(&ev, &data)
		fmt.Printf("Event: [%s], Data: [%s]\n", ev, data)
	}

	if err := ln.Err(); err != nil {
		t.Fatal(err)
	}
}
