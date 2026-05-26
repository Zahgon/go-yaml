package parser

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/token"
)

func newMappingNode(ctx *context, tk *Token, isFlow bool, values ...*ast.MappingValueNode) (*ast.MappingNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newMappingValueNode(ctx *context, colonTk, entryTk *Token, key ast.MapKeyNode, value ast.Node) (*ast.MappingValueNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// originally key was commented, but now that null value has been added, value must be commented.

// set line comment by colonTk or entryTk.

// set line comment by colonTk or entryTk.

func newMappingKeyNode(ctx *context, tk *Token) (*ast.MappingKeyNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAnchorNode(ctx *context, tk *Token) (*ast.AnchorNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAliasNode(ctx *context, tk *Token) (*ast.AliasNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDirectiveNode(ctx *context, tk *Token) (*ast.DirectiveNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newMergeKeyNode(ctx *context, tk *Token) (*ast.MergeKeyNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newNullNode(ctx *context, tk *Token) (*ast.NullNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newBoolNode(ctx *context, tk *Token) (*ast.BoolNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newIntegerNode(ctx *context, tk *Token) (*ast.IntegerNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newFloatNode(ctx *context, tk *Token) (*ast.FloatNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newInfinityNode(ctx *context, tk *Token) (*ast.InfinityNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newNanNode(ctx *context, tk *Token) (*ast.NanNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newStringNode(ctx *context, tk *Token) (*ast.StringNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLiteralNode(ctx *context, tk *Token) (*ast.LiteralNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTagNode(ctx *context, tk *Token) (*ast.TagNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSequenceNode(ctx *context, tk *Token, isFlow bool) (*ast.SequenceNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTagDefaultScalarValueNode(ctx *context, tag *token.Token) (ast.ScalarNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.ScalarNode), nil
}

func setLineComment(ctx *context, node ast.Node, tk *Token) error {
	_ = "STUB: not implemented"
	return nil
}

func setHeadComment(cm *ast.CommentGroupNode, value ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}
