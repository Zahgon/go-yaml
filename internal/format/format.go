package format

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/token"
)

func FormatNodeWithResolvedAlias(n ast.Node, anchorNodeMap map[string]ast.Node) string {
	_ = "STUB: not implemented"
	return ""
}

func FormatNode(n ast.Node) string { _ = "STUB: not implemented"; return "" }

func FormatFile(file *ast.File) string { _ = "STUB: not implemented"; return "" }

func hasCommentFile(f *ast.File) bool { _ = "STUB: not implemented"; return false }

func hasComment(n ast.Node) bool { _ = "STUB: not implemented"; return false }

func getFirstToken(n ast.Node) *token.Token { _ = "STUB: not implemented"; return nil }

type Formatter struct {
	existsComment    bool
	tokenToOriginMap map[*token.Token]string
	anchorNodeMap    map[string]ast.Node
}

func newFormatter(tk *token.Token, existsComment bool) *Formatter {
	_ = "STUB: not implemented"
	return nil
}

func getIndentNumByFirstLineToken(tk *token.Token) int { _ = "STUB: not implemented"; return 0 }

// key: value
//    ^
//   next

// If the current token is the sequence entry.
// the indent is calculated from the column value of the current token.

// key: value
//    ^
//   next

// If the current token is the key in the mapping-value,
// the indent is calculated from the column value of the current token.

// key: value
//    ^
//   prev

// If the current token is the value in the mapping-value,
// the indent is calculated from the column value of the key two steps back.

// - value
// ^
// prev

// If the value is not a mapping-value and the previous token was a sequence entry,
// the indent is calculated using the column value of the sequence entry token.

func (f *Formatter) format(n ast.Node) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatFile(file *ast.File) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) origin(tk *token.Token) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatDocument(n *ast.DocumentNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatNull(n *ast.NullNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatString(n *ast.StringNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatInteger(n *ast.IntegerNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatFloat(n *ast.FloatNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatBool(n *ast.BoolNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatInfinity(n *ast.InfinityNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatNan(n *ast.NanNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatLiteral(n *ast.LiteralNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatMergeKey(n *ast.MergeKeyNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatMappingValue(n *ast.MappingValueNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatDirective(n *ast.DirectiveNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatMapping(n *ast.MappingNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatTag(n *ast.TagNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatMappingKey(n *ast.MappingKeyNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatSequence(n *ast.SequenceNode) string {
	_ = "STUB: not implemented"
	return ""
}

// add head comment.

func (f *Formatter) formatSequenceEntry(n *ast.SequenceEntryNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatAnchor(n *ast.AnchorNode) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatAlias(n *ast.AliasNode) string { _ = "STUB: not implemented"; return "" }

// If formatted text contains newline characters, indentation needs to be considered.

// If the first character is not a newline, the first line should be output without indentation.

func (f *Formatter) formatNode(n ast.Node) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) formatCommentGroup(g *ast.CommentGroupNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) formatComment(n *ast.CommentNode) string { _ = "STUB: not implemented"; return "" }

// nolint: unused
func (f *Formatter) formatIndent(col int) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) trimNewLineCharPrefix(v string) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) trimSpacePrefix(v string) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) trimIndentSpace(trimIndentNum int, v string) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Formatter) addIndentSpace(indentNum int, v string, isIgnoredFirstLine bool) string {
	_ = "STUB: not implemented"
	return ""
}

// normalizeNewLineChars normalize CRLF and CR to LF.
func normalizeNewLineChars(v string) string { _ = "STUB: not implemented"; return "" }
