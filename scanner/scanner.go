package scanner

import (
	"github.com/goccy/go-yaml/token"
)

// IndentState state for indent
type IndentState int

const (
	// IndentStateEqual equals previous indent
	IndentStateEqual IndentState = iota
	// IndentStateUp more indent than previous
	IndentStateUp
	// IndentStateDown less indent than previous
	IndentStateDown
	// IndentStateKeep uses not indent token
	IndentStateKeep
)

// Scanner holds the scanner's internal state while processing a given text.
// It can be allocated as part of another data structure but must be initialized via Init before use.
type Scanner struct {
	source     []rune
	sourcePos  int
	sourceSize int
	// line number. This number starts from 1.
	line int
	// column number. This number starts from 1.
	column int
	// offset represents the offset from the beginning of the source.
	offset int
	// lastDelimColumn is the last column needed to compare indent is retained.
	lastDelimColumn int
	// indentNum indicates the number of spaces used for indentation.
	indentNum int
	// prevLineIndentNum indicates the number of spaces used for indentation at previous line.
	prevLineIndentNum int
	// indentLevel indicates the level of indent depth. This value does not match the column value.
	indentLevel            int
	isFirstCharAtLine      bool
	isAnchor               bool
	isAlias                bool
	isDirective            bool
	startedFlowSequenceNum int
	startedFlowMapNum      int
	indentState            IndentState
	savedPos               *token.Position
}

func (s *Scanner) pos() *token.Position { _ = "STUB: not implemented"; return nil }

func (s *Scanner) bufferedToken(ctx *Context) *token.Token { _ = "STUB: not implemented"; return nil }

// Since we are in a literal, folded or raw folded
// we can use the indent level from the last token.

// The last token should never be nil here.

func (s *Scanner) progressColumn(ctx *Context, num int) { _ = "STUB: not implemented"; return }

func (s *Scanner) progressOnly(ctx *Context, num int) { _ = "STUB: not implemented"; return }

func (s *Scanner) progressLine(ctx *Context) { _ = "STUB: not implemented"; return }

func (s *Scanner) progress(ctx *Context, num int) { _ = "STUB: not implemented"; return }

func (s *Scanner) isNewLineChar(c rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) newLineCount(src []rune) int { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) updateIndentLevel() { _ = "STUB: not implemented"; return }

func (s *Scanner) updateIndentState(ctx *Context) { _ = "STUB: not implemented"; return }

// If lastDelimColumn and s.column are the same,
// treat as Down state since it is the same column as delimiter.

func (s *Scanner) updateIndent(ctx *Context, c rune) { _ = "STUB: not implemented"; return }

// found tab indent.
// In this case, scanTab returns error.

func (s *Scanner) isChangedToIndentStateDown() bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isChangedToIndentStateUp() bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) addBufferedTokenIfExists(ctx *Context) { _ = "STUB: not implemented"; return }

func (s *Scanner) breakMultiLine(ctx *Context) { _ = "STUB: not implemented"; return }

func (s *Scanner) scanSingleQuote(ctx *Context) (*token.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// '' handle as ' character

func hexToInt(b rune) int { _ = "STUB: not implemented"; return 0 }

func hexRunesToInt(b []rune) int { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) scanDoubleQuote(ctx *Context) (*token.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// \u0000 style must have 5 characters at least.

// handle surrogate pairs.

// \u0000\u0000 style must have 11 characters at least.

// \U00000000 style must have 9 characters at least.

// Skip \n after \r in CRLF sequences

func (s *Scanner) validateDocumentSeparatorMarker(ctx *Context, src []rune) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) foundDocumentSeparatorMarker(src []rune) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Scanner) scanQuote(ctx *Context, ch rune) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Scanner) scanWhiteSpace(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isMergeKey(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanTag(ctx *Context) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// skip '!' character

// progress column before collect-entry for scanning it at scanFlowEntry function.

// progress column before new-line-char for scanning new-line-char at scanNewLine function.

func (s *Scanner) scanComment(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// skip '#' character

// document ends with comment.

func (s *Scanner) scanMultiLine(ctx *Context, c rune) error { _ = "STUB: not implemented"; return nil }

// normalize CR and CRLF to LF

func (s *Scanner) scanNewLine(ctx *Context, c rune) { _ = "STUB: not implemented"; return }

// if the following case, origin buffer has unnecessary two spaces.
// So, `removeRightSpaceFromOriginBuf` remove them, also fix column number too.
// ---
// a:[space][space]
//   b: c

// There is no problem that we ignore CR which followed by LF and normalize it to LF, because of following YAML1.2 spec.
// > Line breaks inside scalar content must be normalized by the YAML processor. Each such line break must be parsed into a single line feed character.
// > Outside scalar content, YAML allows any line break to be used to terminate lines.
// > -- https://yaml.org/spec/1.2/spec.html

func (s *Scanner) isFlowMode() bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanFlowMapStart(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanFlowMapEnd(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanFlowArrayStart(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanFlowArrayEnd(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanFlowEntry(ctx *Context, c rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanMapDelim(ctx *Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// like http://

// mapping value

// If the map key is quote, the buffer does not exist because it has already been cut into tokens.
// Therefore, we need to check the last token.

func (s *Scanner) scanDocumentStart(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanDocumentEnd(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanMergeKey(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanRawFoldedChar(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanSequence(ctx *Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Scanner) scanMultiLineHeader(ctx *Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Scanner) validateMultiLineHeaderOption(opt string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) scanMultiLineHeaderOption(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// skip '|' or '>' character

// process \n in the next iteration

// Exclude \r

func (s *Scanner) scanMapKey(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanDirective(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanAnchor(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanAlias(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanReservedChar(ctx *Context, c rune) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) scanTab(ctx *Context, c rune) error { _ = "STUB: not implemented"; return nil }

// tabs character is allowed in flow mode.

func (s *Scanner) scan(ctx *Context) error { _ = "STUB: not implemented"; return nil }

// First, change the IndentState.
// If the target character is the first character in a line, IndentState is Up/Down/Equal state.
// The second and subsequent letters are Keep.

// If IndentState is down, tokens are split, so the buffer accumulated until that point needs to be cutted as a token.

// If literal/folded content is empty, no string token is added.
// Therefore, add an empty string token.
// But if literal/folded token column is 1, it is invalid at down state.

// tab indent for plain text (yaml-test-suite's spec-example-7-12-plain-lines).

// Init prepares the scanner s to tokenize the text src by setting the scanner at the beginning of src.
func (s *Scanner) Init(text string) { _ = "STUB: not implemented"; return }

func (s *Scanner) clearState() { _ = "STUB: not implemented"; return }

// Scan scans the next token and returns the token collection. The source end is indicated by io.EOF.
func (s *Scanner) Scan() (token.Tokens, error) {
	_ = "STUB: not implemented"
	return *new(token.Tokens), nil
}
