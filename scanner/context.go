package scanner

import (
	"sync"

	"github.com/goccy/go-yaml/token"
)

// Context context at scanning
type Context struct {
	idx                int
	size               int
	notSpaceCharPos    int
	notSpaceOrgCharPos int
	src                []rune
	buf                []rune
	obuf               []rune
	tokens             token.Tokens
	mstate             *MultiLineState
}

type MultiLineState struct {
	opt                              string
	firstLineIndentColumn            int
	prevLineIndentColumn             int
	lineIndentColumn                 int
	lastNotSpaceOnlyLineIndentColumn int
	spaceOnlyIndentColumn            int
	foldedNewLine                    bool
	isRawFolded                      bool
	isLiteral                        bool
	isFolded                         bool
}

var (
	ctxPool = sync.Pool{
		New: func() interface{} {
			return createContext()
		},
	}
)

func createContext() *Context { _ = "STUB: not implemented"; return nil }

func newContext(src []rune) *Context { _ = "STUB: not implemented"; return nil }

func (c *Context) release() { _ = "STUB: not implemented"; return }

func (c *Context) clear() { _ = "STUB: not implemented"; return }

func (c *Context) reset(src []rune) { _ = "STUB: not implemented"; return }

func (c *Context) resetBuffer() { _ = "STUB: not implemented"; return }

func (c *Context) breakMultiLine() { _ = "STUB: not implemented"; return }

func (c *Context) getMultiLineState() *MultiLineState { _ = "STUB: not implemented"; return nil }

func (c *Context) setLiteral(lastDelimColumn int, opt string) { _ = "STUB: not implemented"; return }

func (c *Context) setFolded(lastDelimColumn int, opt string) { _ = "STUB: not implemented"; return }

func (c *Context) setRawFolded(column int) { _ = "STUB: not implemented"; return }

func firstLineIndentColumnByOpt(opt string) int { _ = "STUB: not implemented"; return 0 }

func (s *MultiLineState) lastDelimColumn() int { _ = "STUB: not implemented"; return 0 }

func (s *MultiLineState) updateIndentColumn(column int) { _ = "STUB: not implemented"; return }

func (s *MultiLineState) updateSpaceOnlyIndentColumn(column int) { _ = "STUB: not implemented"; return }

func (s *MultiLineState) validateIndentAfterSpaceOnly(column int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MultiLineState) validateIndentColumn() error { _ = "STUB: not implemented"; return nil }

func (s *MultiLineState) updateNewLineState() { _ = "STUB: not implemented"; return }

func (s *MultiLineState) isIndentColumn(column int) bool { _ = "STUB: not implemented"; return false }

func (s *MultiLineState) addIndent(ctx *Context, column int) { _ = "STUB: not implemented"; return }

// If the first line of the document has already been evaluated, the number is treated as the threshold, since the `firstLineIndentColumn` is a positive number.

// `c.foldedNewLine` is a variable that is set to true for every newline.

// Since addBuf ignore space character, add to the buffer directly.

// updateNewLineInFolded if Folded or RawFolded context and the content on the current line starts at the same column as the previous line,
// treat the new-line-char as a space.
func (s *MultiLineState) updateNewLineInFolded(ctx *Context, column int) {
	_ = "STUB: not implemented"
	return
}

// Folded or RawFolded.

// ---
// >
//  a
//  b

// if previous line is indent-space and new-line-char only, prevLineIndentColumn is zero.
// In this case, last new-line-char is removed.
// ---
// >
//  a
//
//  b

func (s *MultiLineState) hasTrimAllEndNewlineOpt() bool { _ = "STUB: not implemented"; return false }

func (s *MultiLineState) hasKeepAllEndNewlineOpt() bool { _ = "STUB: not implemented"; return false }

func (c *Context) addToken(tk *token.Token) { _ = "STUB: not implemented"; return }

func (c *Context) addBuf(r rune) { _ = "STUB: not implemented"; return }

func (c *Context) addBufWithTab(r rune) { _ = "STUB: not implemented"; return }

func (c *Context) addOriginBuf(r rune) { _ = "STUB: not implemented"; return }

func (c *Context) removeRightSpaceFromBuf() { _ = "STUB: not implemented"; return }

func (c *Context) isEOS() bool { _ = "STUB: not implemented"; return false }

func (c *Context) isNextEOS() bool { _ = "STUB: not implemented"; return false }

func (c *Context) next() bool { _ = "STUB: not implemented"; return false }

func (c *Context) source(s, e int) string { _ = "STUB: not implemented"; return "" }

func (c *Context) previousChar() rune { _ = "STUB: not implemented"; return 0 }

func (c *Context) currentChar() rune { _ = "STUB: not implemented"; return 0 }

func (c *Context) nextChar() rune { _ = "STUB: not implemented"; return 0 }

func (c *Context) repeatNum(r rune) int { _ = "STUB: not implemented"; return 0 }

func (c *Context) progress(num int) { _ = "STUB: not implemented"; return }

func (c *Context) existsBuffer() bool { _ = "STUB: not implemented"; return false }

func (c *Context) isMultiLine() bool { _ = "STUB: not implemented"; return false }

func (c *Context) bufferedSrc() []rune { _ = "STUB: not implemented"; return nil }

// remove end '\n' character and trailing empty lines.
// https://yaml.org/spec/1.2.2/#8112-block-chomping-indicator

// If the '-' flag is specified, all trailing newline characters will be removed.

// Normally, all but one of the trailing newline characters are removed.

// If the text ends with a space character, remove all of them.

// If the content consists only of a newline,
// it can be considered as the document ending without any specified value,
// so it is treated as an empty string.

func (c *Context) bufferedToken(pos *token.Position) *token.Token {
	_ = "STUB: not implemented"
	return nil
}

// clear value's buffer only.

func (c *Context) setTokenTypeByPrevTag(tk *token.Token) { _ = "STUB: not implemented"; return }

func (c *Context) lastToken() *token.Token { _ = "STUB: not implemented"; return nil }
