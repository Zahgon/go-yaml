package parser

// context context at parsing
type context struct {
	tokenRef *tokenRef
	path     string
	isFlow   bool
}

type tokenRef struct {
	tokens []*Token
	size   int
	idx    int
}

var pathSpecialChars = []string{
	"$", "*", ".", "[", "]",
}

func containsPathSpecialChar(path string) bool { _ = "STUB: not implemented"; return false }

func normalizePath(path string) string { _ = "STUB: not implemented"; return "" }

func (c *context) currentToken() *Token { _ = "STUB: not implemented"; return nil }

func (c *context) isComment() bool { _ = "STUB: not implemented"; return false }

func (c *context) nextToken() *Token { _ = "STUB: not implemented"; return nil }

func (c *context) nextNotCommentToken() *Token { _ = "STUB: not implemented"; return nil }

func (c *context) isTokenNotFound() bool { _ = "STUB: not implemented"; return false }

func (c *context) withGroup(g *TokenGroup) *context { _ = "STUB: not implemented"; return nil }

func (c *context) withChild(path string) *context { _ = "STUB: not implemented"; return nil }

func (c *context) withIndex(idx uint) *context { _ = "STUB: not implemented"; return nil }

func (c *context) withFlow(isFlow bool) *context { _ = "STUB: not implemented"; return nil }

func newContext() *context { _ = "STUB: not implemented"; return nil }

func (c *context) goNext() { _ = "STUB: not implemented"; return }

func (c *context) next() bool { _ = "STUB: not implemented"; return false }

func (c *context) insertNullToken(tk *Token) *Token { _ = "STUB: not implemented"; return nil }

func (c *context) addNullValueToken(tk *Token) *Token { _ = "STUB: not implemented"; return nil }

// add space for map or sequence value.

func (c *context) createImplicitNullToken(base *Token) *Token {
	_ = "STUB: not implemented"
	return nil
}

func (c *context) insertToken(tk *Token) { _ = "STUB: not implemented"; return }

func (c *context) addToken(tk *Token) { _ = "STUB: not implemented"; return }
