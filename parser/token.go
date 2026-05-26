package parser

import (
	"github.com/goccy/go-yaml/token"
)

type TokenGroupType int

const (
	TokenGroupNone TokenGroupType = iota
	TokenGroupDirective
	TokenGroupDirectiveName
	TokenGroupDocument
	TokenGroupDocumentBody
	TokenGroupAnchor
	TokenGroupAnchorName
	TokenGroupAlias
	TokenGroupLiteral
	TokenGroupFolded
	TokenGroupScalarTag
	TokenGroupMapKey
	TokenGroupMapKeyValue
)

func (t TokenGroupType) String() string { _ = "STUB: not implemented"; return "" }

type Token struct {
	Token       *token.Token
	Group       *TokenGroup
	LineComment *token.Token
}

func (t *Token) RawToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (t *Token) Type() token.Type { _ = "STUB: not implemented"; return *new(token.Type) }

func (t *Token) GroupType() TokenGroupType { _ = "STUB: not implemented"; return *new(TokenGroupType) }

func (t *Token) Line() int { _ = "STUB: not implemented"; return 0 }

func (t *Token) Column() int { _ = "STUB: not implemented"; return 0 }

func (t *Token) SetGroupType(typ TokenGroupType) { _ = "STUB: not implemented"; return }

func (t *Token) Dump() { _ = "STUB: not implemented"; return }

func (t *Token) dump(ctx *groupTokenRenderContext) { _ = "STUB: not implemented"; return }

type groupTokenRenderContext struct {
	num int
}

type TokenGroup struct {
	Type   TokenGroupType
	Tokens []*Token
}

func (g *TokenGroup) First() *Token { _ = "STUB: not implemented"; return nil }

func (g *TokenGroup) Last() *Token { _ = "STUB: not implemented"; return nil }

func (g *TokenGroup) dump(ctx *groupTokenRenderContext) { _ = "STUB: not implemented"; return }

func (g *TokenGroup) RawToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (g *TokenGroup) Line() int { _ = "STUB: not implemented"; return 0 }

func (g *TokenGroup) Column() int { _ = "STUB: not implemented"; return 0 }

func (g *TokenGroup) TokenType() token.Type { _ = "STUB: not implemented"; return *new(token.Type) }

func CreateGroupedTokens(tokens token.Tokens) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTokens(tks token.Tokens) []*Token { _ = "STUB: not implemented"; return nil }

func createLineCommentTokenGroups(tokens []*Token) []*Token { _ = "STUB: not implemented"; return nil }

func createLiteralAndFoldedTokenGroups(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createAnchorAndAliasTokenGroups(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createScalarTagTokenGroups(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// secondary tag.

func createAnchorWithScalarTagTokenGroups(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMapKeyTokenGroups(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMapKeyByMappingKey(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMapKeyByMappingValue(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMapKeyValueTokenGroups(tokens []*Token) []*Token { _ = "STUB: not implemented"; return nil }

func createDirectiveTokenGroups(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDocumentTokens(tokens []*Token) ([]*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if current token is last token, add DocumentHeader only tokens to ret.

func isScalarType(tk *Token) bool { _ = "STUB: not implemented"; return false }

func isNotMapKeyType(tk *Token) bool { _ = "STUB: not implemented"; return false }

func isFlowType(tk *Token) bool { _ = "STUB: not implemented"; return false }
