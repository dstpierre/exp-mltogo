package main

import (
	"fmt"
	"io"
	"log"
)

type transpiler struct {
	index     int
	nodes     []expression
	substitue map[string]string
	exposed   map[string]string
}

func newTranspiler(nodes []expression) *transpiler {
	return &transpiler{
		index:     0,
		nodes:     nodes,
		substitue: make(map[string]string),
		exposed:   make(map[string]string),
	}
}

func (t *transpiler) Read(buf []byte) (n int, err error) {
	if t.index >= len(t.nodes) {
		return 0, io.EOF
	}

	node := t.nodes[t.index]

	switch node.parent.t {
	case PACKAGE:
		n = copy(buf, t.pkg(node))
	case IMPORT:
		n = copy(buf, t.imp(node))
	case EXPOSING:
		t.sub(node)
	case FUNCTION:
		n = copy(buf, t.fun(node))
	case IDENT:
		v := t.ident(node)
		if pkg, ok := t.exposed[string(v)]; ok {
			n = copy(buf, fmt.Sprintf("%s.%s(", pkg, v))
		} else {
			n = copy(buf, v)
		}
	}

	t.index++
	return n, nil
}

func (t *transpiler) pkg(n expression) []byte {
	if len(n.nodes) != 1 || n.nodes[0].t != IDENT {
		log.Fatal("expected identifier for package at ", n.parent.pos)
	}
	return []byte(fmt.Sprintf("package %s\n\n", n.nodes[0].lit))
}

func (t *transpiler) imp(n expression) []byte {
	if len(n.nodes) != 1 || n.nodes[0].t != IDENT {
		log.Fatal("expected identifier for import at ", n.parent.pos)
	}
	return []byte(fmt.Sprintf("import %s\n", n.nodes[0].lit))
}

func (t *transpiler) sub(n expression) {
	if len(n.nodes) != 1 || n.nodes[0].t != IDENT {
		log.Fatal("expected identifier for exposing")
	}

	t.substitue[n.nodes[0].lit] = "log.Println"
}

func (t *transpiler) fun(n expression) []byte {
	return []byte(fmt.Sprintf("func %s() {\n", n.parent.lit))
}

func (t *transpiler) ident(n expression) []byte {
	var b []byte
	for _, i := range n.nodes {
		switch i.t {
		case IDENT:
			b = append(b, t.ns(i.lit)...)
		case STRING:
			tmp := []byte(fmt.Sprintf(`"%s"`, i.lit))
			b = append(b, tmp...)
		}
	}
	return b
}

func (t *transpiler) ns(s string) []byte {
	if sub, ok := t.substitue[s]; ok {
		return []byte(sub)
	}
	return []byte(s)
}
