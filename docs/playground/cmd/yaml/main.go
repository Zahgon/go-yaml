//go:build js && wasm

package main

import (
	"context"
	"syscall/js"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"github.com/goccy/go-yaml/token"
)

func response(v any, err error) map[string]any { _ = "STUB: not implemented"; return nil }

func decode(this js.Value, args []js.Value) any { _ = "STUB: not implemented"; return *new(any) }

func tokenize(this js.Value, args []js.Value) any { _ = "STUB: not implemented"; return *new(any) }

func parseGroup(this js.Value, args []js.Value) any { _ = "STUB: not implemented"; return *new(any) }

func parse(this js.Value, args []js.Value) any { _ = "STUB: not implemented"; return *new(any) }

func Decode(v string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type Token struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Origin string `json:"origin"`
	Error  string `json:"error"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Offset int    `json:"offset"`
}

func Tokenize(v string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func toToken(tk *token.Token) *Token { _ = "STUB: not implemented"; return nil }

type GroupedToken struct {
	Token       *Token      `json:"token"`
	Group       *TokenGroup `json:"group"`
	LineComment *Token      `json:"lineComment"`
}

type TokenGroup struct {
	Type   string          `json:"type"`
	Tokens []*GroupedToken `json:"tokens"`
}

func toGroupedToken(tk *parser.Token) *GroupedToken { _ = "STUB: not implemented"; return nil }

func toTokenGroup(g *parser.TokenGroup) *TokenGroup { _ = "STUB: not implemented"; return nil }

func ParseGroup(v string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func Parse(ctx context.Context, v string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type NodeRenderer struct {
	id    int
	edges []*Edge
}

type Edge struct {
	start *Node
	end   *Node
}

type Node struct {
	graphName string
	node      *graphviz.Node
}

func (r *NodeRenderer) createID() string { _ = "STUB: not implemented"; return "" }

func (r *NodeRenderer) createNodeGraph(parent *graphviz.Graph, node any, name string) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) createNode(graph *graphviz.Graph, name string) (*graphviz.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) createEdge(fromGraph *graphviz.Graph, fromNode *graphviz.Node, toGraph *graphviz.Graph) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeRenderer) renderFile(graph *graphviz.Graph, file *ast.File) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderDocument(graph *graphviz.Graph, doc *ast.DocumentNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderNode(graph *graphviz.Graph, node ast.Node) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderMappingNode(graph *graphviz.Graph, node *ast.MappingNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderMappingValueNode(graph *graphviz.Graph, node *ast.MappingValueNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderStringNode(graph *graphviz.Graph, n *ast.StringNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderNullNode(graph *graphviz.Graph, n *ast.NullNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderIntegerNode(graph *graphviz.Graph, n *ast.IntegerNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderFloatNode(graph *graphviz.Graph, n *ast.FloatNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderLiteralNode(graph *graphviz.Graph, n *ast.LiteralNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderMergeKeyNode(graph *graphviz.Graph, n *ast.MergeKeyNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderBoolNode(graph *graphviz.Graph, n *ast.BoolNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderInfinityNode(graph *graphviz.Graph, n *ast.InfinityNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderNaNNode(graph *graphviz.Graph, n *ast.NanNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderMappingKeyNode(graph *graphviz.Graph, n *ast.MappingKeyNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderSequenceNode(graph *graphviz.Graph, n *ast.SequenceNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderAnchorNode(graph *graphviz.Graph, n *ast.AnchorNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderAliasNode(graph *graphviz.Graph, n *ast.AliasNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderDirectiveNode(graph *graphviz.Graph, n *ast.DirectiveNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderTagNode(graph *graphviz.Graph, n *ast.TagNode) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderCommentGroupNode(graph *graphviz.Graph, n *ast.CommentGroupNode, pos yaml.CommentPosition) (*graphviz.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeRenderer) renderPath(graph *graphviz.Graph, p string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeRenderer) renderToken(graph *graphviz.Graph, tk *token.Token) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	js.Global().Set("decode", js.FuncOf(decode))
	js.Global().Set("tokenize", js.FuncOf(tokenize))
	js.Global().Set("parseGroup", js.FuncOf(parseGroup))
	js.Global().Set("parse", js.FuncOf(parse))

	<-make(chan struct{})
}
