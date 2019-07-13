package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserSimpleProgram(t *testing.T) {
	p := newParser(strings.NewReader(
		`
package main

import log exposing {Println}

main = Println "first flang program"
		`))

	p.parse()

	assert.Equal(t, 5, len(p.nodes))

	for _, n := range p.nodes {
		fmt.Println(n.parent.t, n.parent.lit)
		for _, e := range n.nodes {
			fmt.Println(">", e.t, e.lit)
		}
	}
}
