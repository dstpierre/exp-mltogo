package main

type token string

const (
	ILLEGAL    token = "ILLEGAL"
	EOF              = "EOF"
	WS               = "WS"
	OPENBRACE        = "OPENBRACE"
	CLOSEBRACE       = "CLOSEBRACE"
	IDENT            = "IDENT"
	EQUAL            = "EQUAL"
	ARROW            = "ARROW"
	DOT              = "."
	ARGUMENT         = "ARGUMENT"
	FUNCTION         = "FUNCTION"
	STRING           = "STRING"

	// AST
	PACKAGE  = "PACKAGE"
	IMPORT   = "IMPORT"
	EXPOSING = "EXPOSING"
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isString(ch rune) bool {
	return ch == '"'
}
