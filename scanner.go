package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

const eof = rune(0)

type scanner struct {
	r *bufio.Reader
}

func newScanner(rd io.Reader) *scanner {
	return &scanner{r: bufio.NewReader(rd)}
}

func (s *scanner) next() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *scanner) back() {
	if err := s.r.UnreadRune(); err != nil {
		log.Fatal("unable to unread", err)
	}
}

func (s *scanner) scan() (tok token, lit string) {
	ch := s.next()

	if isWhitespace(ch) {
		s.back()
		return s.scanWhitesapce()
	} else if isLetter(ch) {
		s.back()
		return s.scanIdent()
	} else if isString(ch) {
		return s.scanString()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '=':
		return EQUAL, string(ch)
	case '.':
		return DOT, string(ch)
	case '{':
		return OPENBRACE, string(ch)
	case '}':
		return CLOSEBRACE, string(ch)
	}
	return ILLEGAL, ""
}

func (s *scanner) scanWhitesapce() (tok token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.next())

	for {
		if ch := s.next(); ch == eof {
			break
		} else if isWhitespace(ch) == false {
			s.back()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func (s *scanner) scanIdent() (tok token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.next())

	for {
		if ch := s.next(); ch == eof {
			break
		} else if isLetter(ch) == false {
			s.back()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return IDENT, buf.String()
}

func (s *scanner) scanString() (tok token, lit string) {
	var buf bytes.Buffer

	for {
		if ch := s.next(); ch == eof {
			break
		} else if isString(ch) {
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return STRING, buf.String()
}
