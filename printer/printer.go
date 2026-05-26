package printer

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/token"
)

// Property additional property set for each the token
type Property struct {
	Prefix string
	Suffix string
}

// PrintFunc returns property instance
type PrintFunc func() *Property

// Printer create text from token collection or ast
type Printer struct {
	LineNumber       bool
	LineNumberFormat func(num int) string
	MapKey           PrintFunc
	Anchor           PrintFunc
	Alias            PrintFunc
	Bool             PrintFunc
	String           PrintFunc
	Number           PrintFunc
	Comment          PrintFunc
}

func defaultLineNumberFormat(num int) string { _ = "STUB: not implemented"; return "" }

func (p *Printer) property(tk *token.Token) *Property { _ = "STUB: not implemented"; return nil }

// PrintTokens create text from token collection
func (p *Printer) PrintTokens(tokens token.Tokens) string { _ = "STUB: not implemented"; return "" }

// PrintNode create text from ast.Node
func (p *Printer) PrintNode(node ast.Node) []byte { _ = "STUB: not implemented"; return nil }

func (p *Printer) setDefaultColorSet() { _ = "STUB: not implemented"; return }

func (p *Printer) PrintErrorMessage(msg string, isColored bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Printer) removeLeftSideNewLineChar(src string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Printer) removeRightSideNewLineChar(src string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Printer) removeRightSideWhiteSpaceChar(src string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Printer) newLineCount(s string) int { _ = "STUB: not implemented"; return 0 }

func (p *Printer) isNewLineLastChar(s string) bool { _ = "STUB: not implemented"; return false }

func (p *Printer) printBeforeTokens(tk *token.Token, minLine, extLine int) token.Tokens {
	_ = "STUB: not implemented"
	return *new(token.Tokens)
}

// add white spaces to minTk by prev token

// add suffix to header of next token

func (p *Printer) printAfterTokens(tk *token.Token, maxLine int) token.Tokens {
	_ = "STUB: not implemented"
	return *new(token.Tokens)
}

func (p *Printer) setupErrorTokenFormat(annotateLine int, isColored bool) {
	_ = "STUB: not implemented"
	return
}

func (p *Printer) PrintErrorToken(tk *token.Token, isColored bool) string {
	_ = "STUB: not implemented"
	return ""
}

// if last character ( exclude white space ) is new line character, ignore it.
