package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type position struct {
	pos  int
	line int
}

type exp struct {
	t   token
	lit string
	pos position
}

type expression struct {
	parent exp
	nodes  []exp
}

type parser struct {
	s       *scanner
	pos     position
	nodes   []expression
	exposed map[string]string
}

func newParser(r io.Reader) *parser {
	return &parser{
		s: newScanner(r),
		pos: position{
			line: 0,
			pos:  0,
		},
		nodes:   make([]expression, 0),
		exposed: make(map[string]string),
	}
}

func (p *parser) parse() {
	for {
		tok, lit := p.s.scan()
		if tok == EOF {
			break
		} else if tok == IDENT {
			if kw, nt, ok := p.isKeyword(lit); ok {
				e := newExpression(kw, lit, p.pos)
				if tok, lit = p.scanUntil(nt); tok == EOF {
					log.Printf("expected %s got end of file", nt)
					break
				}

				if nt == OPENBRACE {
					results, err := p.getExpressionUntil(CLOSEBRACE)
					if err != nil {
						log.Println(err)
						break
					}

					if kw == EXPOSING {
						p.syncExposed(lit, results...)
					}

					e.nodes = append(e.nodes, results...)
				} else {
					e.nodes = append(e.nodes, exp{t: tok, lit: lit, pos: p.pos})
				}
				p.nodes = append(p.nodes, e)
			} else {
				p.ident(tok, lit)
			}
		}
	}
}

func (p *parser) scanUntil(t token) (tok token, lit string) {
	for {
		tok, lit = p.s.scan()
		if tok == EOF || tok == t {
			return
		}
	}
}

func (p *parser) getExpressionUntil(t token) ([]exp, error) {
	var results []exp
	for {
		tok, lit := p.s.scan()
		if tok == EOF {
			return nil, fmt.Errorf("reached end of file was expecting %s", t)
		} else if tok == t {
			break
		} else {
			results = append(results, exp{t: tok, lit: lit, pos: p.pos})
		}
	}
	return results, nil
}

func (p *parser) scanLine() []exp {
	var results []exp
	for {
		if tok, lit := p.s.scan(); tok == EOF {
			return results
		} else if tok == WS && strings.Contains(lit, "\n") {
			return results
		} else {
			results = append(results, exp{t: tok, lit: lit, pos: p.pos})
		}

	}
}

func (p *parser) ident(tok token, lit string) {
	e := newExpression(tok, lit, p.pos)

	results := p.scanLine()

	// is there any assignment
	for i, r := range results {
		if r.t == EQUAL {
			e.parent.t = FUNCTION
			e.nodes = append(e.nodes, results[:i]...)

			stmt := newExpression(IDENT, results[i+1].lit, p.pos)
			stmt.nodes = append(stmt.nodes, results[i+2:]...)

			p.nodes = append(p.nodes, e)
			p.nodes = append(p.nodes, stmt)

			return
		}
	}

	e.nodes = append(e.nodes, results...)
	p.nodes = append(p.nodes, e)
}

func newExpression(t token, lit string, pos position) expression {
	return expression{
		parent: exp{t: t, lit: lit, pos: pos},
	}
}

func (p *parser) isKeyword(k string) (t token, nt token, ok bool) {
	t, nt = ILLEGAL, ILLEGAL
	switch strings.ToUpper(k) {
	case "PACKAGE":
		t, nt, ok = PACKAGE, IDENT, true
	case "EXPOSING":
		t, nt, ok = EXPOSING, OPENBRACE, true
	case "IMPORT":
		t, nt, ok = IMPORT, IDENT, true
	}
	return
}

func (p *parser) syncExposed(pkg string, ident ...exp) {
	for _, e := range ident {
		if _, ok := p.exposed[e.lit]; !ok {
			p.exposed[e.lit] = pkg
		}
	}
}
