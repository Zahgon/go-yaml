package parser

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/token"
)

type Mode uint

const (
	ParseComments Mode = 1 << iota // parse comments and add them to AST
)

// ParseBytes parse from byte slice, and returns ast.File
func ParseBytes(bytes []byte, mode Mode, opts ...Option) (*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse parse from token instances, and returns ast.File
func Parse(tokens token.Tokens, mode Mode, opts ...Option) (*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse parse from filename, and returns ast.File
func ParseFile(filename string, mode Mode, opts ...Option) (*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type YAMLVersion string

const (
	YAML10 YAMLVersion = "1.0"
	YAML11 YAMLVersion = "1.1"
	YAML12 YAMLVersion = "1.2"
	YAML13 YAMLVersion = "1.3"
)

var yamlVersionMap = map[string]YAMLVersion{
	"1.0": YAML10,
	"1.1": YAML11,
	"1.2": YAML12,
	"1.3": YAML13,
}

type parser struct {
	tokens                []*Token
	pathMap               map[string]ast.Node
	yamlVersion           YAMLVersion
	allowDuplicateMapKey  bool
	secondaryTagDirective *ast.DirectiveNode
}

func newParser(tokens token.Tokens, mode Mode, opts []Option) (*parser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// keep prev/next reference between tokens containing comments
// https://github.com/goccy/go-yaml/issues/254

func (p *parser) parse(ctx *context) (*ast.File, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *parser) parseDocument(ctx *context, docGroup *TokenGroup) (*ast.DocumentNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// clear yaml version value if DocumentEnd token (...) is specified.

func (p *parser) parseDocumentBody(ctx *context) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (p *parser) parseToken(ctx *context, tk *Token) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// SequenceEndType is always validated in parseFlowSequence.
// Therefore, if this is found in other cases, it is treated as a syntax error.

// MappingEndType is always validated in parseFlowMap.
// Therefore, if this is found in other cases, it is treated as a syntax error.

func (p *parser) parseScalarValue(ctx *context, tk *Token) (ast.ScalarNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.ScalarNode), nil
}

// this case applies when it is a scalar tag and its value does not exist.
// Examples of cases where the value does not exist include cases like `key: !!str,` or `!!str : value`.

func (p *parser) parseFlowMap(ctx *context) (*ast.MappingNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip MappingStart token

// this case is here: "{ elem, }".
// In this case, ignore the last element and break mapping parsing.

// set line comment if exists. e.g.) } # comment

// skip mapping end token.

func (p *parser) isFlowMapDelim(tk *Token) bool { _ = "STUB: not implemented"; return false }

func (p *parser) parseMap(ctx *context) (*ast.MappingNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [
// key: value
// ] <=

// a: {
//  b: c
// } <=

// If the comment is in the same or deeper column as the last element column in map value,
// treat it as a footer comment for the last element.

func (p *parser) validateMapKeyValueNextToken(ctx *context, keyTk, tk *Token) error {
	_ = "STUB: not implemented"
	return nil
}

// a: b
//  c <= this token is invalid.

func (p *parser) isMapToken(tk *Token) bool { _ = "STUB: not implemented"; return false }

func (p *parser) parseMapKeyValue(ctx *context, g *TokenGroup, entryTk *Token) (*ast.MappingValueNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseMapKey(ctx *context, g *TokenGroup) (ast.MapKeyNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.MapKeyNode), nil
}

// skip mapping key token

func (p *parser) validateMapKey(ctx *context, tk *token.Token, keyPath string, colonTk *Token) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) removeLeftWhiteSpace(src string) string {
	_ = "STUB: not implemented"
	// CR or LF or CRLF
	return ""
}

func (p *parser) removeRightWhiteSpace(src string) string {
	_ = "STUB: not implemented"
	// CR or LF or CRLF
	return ""
}

func (p *parser) existsNewLineCharacter(src string) bool { _ = "STUB: not implemented"; return false }

func (p *parser) newLineCharacterNum(src string) int { _ = "STUB: not implemented"; return 0 }

func (p *parser) mapKeyText(n ast.Node) string { _ = "STUB: not implemented"; return "" }

func (p *parser) parseMapValue(ctx *context, key ast.MapKeyNode, colonTk *Token) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// a: b:
//    ^
//
// a: b: c
//    ^

// in this case,
// ----
// key: <value does not defined>
// next

// in this case,
// ----
// key: &anchor
// next

// key: <value does not defined>
// &anchor

// key: <value does not defined>
// !!tag

// in this case,
// ----
//   key: <value does not defined>
// next

// in this case,
// ----
//   key: &anchor
// next

func (p *parser) validateAnchorValueInMapOrSeq(value ast.Node, col int) error {
	_ = "STUB: not implemented"
	return nil
}

// key:
//   &anchor !!tag
//
// - &anchor !!tag

// key: &anchor
// !!tag
//
// - &anchor
// !!tag

func (p *parser) parseAnchor(ctx *context, g *TokenGroup) (*ast.AnchorNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseAnchorName(ctx *context) (*ast.AnchorNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseAlias(ctx *context) (*ast.AliasNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseLiteral(ctx *context) (*ast.LiteralNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip literal/folded token

func (p *parser) parseScalarTag(ctx *context) (*ast.TagNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseTag(ctx *context) (*ast.TagNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseTagValue(ctx *context, tagRawTk *token.Token, tk *Token) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (p *parser) parseFlowSequence(ctx *context) (*ast.SequenceNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip SequenceStart token

// this case is here: "[ elem, ]".
// In this case, ignore the last element and break sequence parsing.

// set line comment if exists. e.g.) ] # comment

// skip sequence end token.

func (p *parser) parseSequence(ctx *context) (*ast.SequenceNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip sequence entry token

// If the comment is in the same or deeper column as the last element column in sequence value,
// treat it as a footer comment for the last element.

func (p *parser) parseSequenceValue(ctx *context, seqTk *Token) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// in this case,
// ----
// - <value does not defined>
// -

// in this case,
// ----
// - &anchor
// -

// - <value does not defined>
// &anchor

// - <value does not defined>
// !!tag

// in this case,
// ----
//   - <value does not defined>
// next

// in this case,
// ----
//   - &anchor
// next

func (p *parser) parseDirective(ctx *context, g *TokenGroup) (*ast.DirectiveNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseDirectiveName(ctx *context) (*ast.DirectiveNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseComment(ctx *context) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (p *parser) parseHeadComment(ctx *context) *ast.CommentGroupNode {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseFootComment(ctx *context, col int) *ast.CommentGroupNode {
	_ = "STUB: not implemented"
	return nil
}
