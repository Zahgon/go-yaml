package errors

import (
	"errors"
	"reflect"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/token"
)

var (
	As  = errors.As
	Is  = errors.Is
	New = errors.New
)

const (
	defaultFormatColor   = false
	defaultIncludeSource = true
)

type Error interface {
	error
	GetToken() *token.Token
	GetMessage() string
	FormatError(bool, bool) string
}

var (
	_ Error = new(SyntaxError)
	_ Error = new(TypeError)
	_ Error = new(OverflowError)
	_ Error = new(DuplicateKeyError)
	_ Error = new(UnknownFieldError)
	_ Error = new(UnexpectedNodeTypeError)
)

type SyntaxError struct {
	Message string
	Token   *token.Token
}

type TypeError struct {
	DstType         reflect.Type
	SrcType         reflect.Type
	StructFieldName *string
	Token           *token.Token
}

type OverflowError struct {
	DstType reflect.Type
	SrcNum  string
	Token   *token.Token
}

type DuplicateKeyError struct {
	Message string
	Token   *token.Token
}

type UnknownFieldError struct {
	Message string
	Token   *token.Token
}

type UnexpectedNodeTypeError struct {
	Actual   ast.NodeType
	Expected ast.NodeType
	Token    *token.Token
}

// ErrSyntax create syntax error instance with message and token
func ErrSyntax(msg string, tk *token.Token) *SyntaxError { _ = "STUB: not implemented"; return nil }

// ErrOverflow creates an overflow error instance with message and a token.
func ErrOverflow(dstType reflect.Type, num string, tk *token.Token) *OverflowError {
	_ = "STUB: not implemented"
	return nil
}

// ErrTypeMismatch cerates an type mismatch error instance with token.
func ErrTypeMismatch(dstType, srcType reflect.Type, token *token.Token) *TypeError {
	_ = "STUB: not implemented"
	return nil
}

// ErrDuplicateKey creates an duplicate key error instance with token.
func ErrDuplicateKey(msg string, tk *token.Token) *DuplicateKeyError {
	_ = "STUB: not implemented"
	return nil
}

// ErrUnknownField creates an unknown field error instance with token.
func ErrUnknownField(msg string, tk *token.Token) *UnknownFieldError {
	_ = "STUB: not implemented"
	return nil
}

func ErrUnexpectedNodeType(actual, expected ast.NodeType, tk *token.Token) *UnexpectedNodeTypeError {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyntaxError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *SyntaxError) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (e *SyntaxError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *SyntaxError) FormatError(colored, inclSource bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *OverflowError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *OverflowError) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (e *OverflowError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *OverflowError) FormatError(colored, inclSource bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *OverflowError) msg() string { _ = "STUB: not implemented"; return "" }

func (e *TypeError) msg() string { _ = "STUB: not implemented"; return "" }

func (e *TypeError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *TypeError) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (e *TypeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *TypeError) FormatError(colored, inclSource bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *DuplicateKeyError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *DuplicateKeyError) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (e *DuplicateKeyError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *DuplicateKeyError) FormatError(colored, inclSource bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *UnknownFieldError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *UnknownFieldError) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (e *UnknownFieldError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *UnknownFieldError) FormatError(colored, inclSource bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *UnexpectedNodeTypeError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *UnexpectedNodeTypeError) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (e *UnexpectedNodeTypeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *UnexpectedNodeTypeError) FormatError(colored, inclSource bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *UnexpectedNodeTypeError) msg() string { _ = "STUB: not implemented"; return "" }

func FormatError(errMsg string, token *token.Token, colored, inclSource bool) string {
	_ = "STUB: not implemented"
	return ""
}
