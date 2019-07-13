package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanSimpleStatement(t *testing.T) {
	scn := newScanner(
		strings.NewReader(`main = log.Println "testing flang"`),
	)

	tok, lit := scn.scan()
	assert.EqualValues(t, IDENT, tok)
	assert.Equal(t, "main", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, WS, tok)
	assert.Equal(t, " ", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, EQUAL, tok)
	assert.Equal(t, "=", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, WS, tok)
	assert.Equal(t, " ", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, IDENT, tok)
	assert.Equal(t, "log", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, DOT, tok)
	assert.Equal(t, ".", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, IDENT, tok)
	assert.Equal(t, "Println", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, WS, tok)
	assert.Equal(t, " ", lit)

	tok, lit = scn.scan()
	assert.EqualValues(t, STRING, tok)
	assert.Equal(t, "testing flang", lit)
}
