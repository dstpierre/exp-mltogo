package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTranspilerPackage(t *testing.T) {
	p := newParser(strings.NewReader(
		`
package main

import log exposing {Println}

main = Println "first flang program"

		`))

	p.parse()

	tp := newTranspiler(p.nodes)

	buf := bytes.NewBuffer([]byte{})
	if _, err := buf.ReadFrom(tp); err != nil {
		t.Error(err)
	}

	fmt.Println(buf.String())
}
