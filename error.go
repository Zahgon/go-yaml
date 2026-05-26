package yaml

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/internal/errors"
)

var (
	ErrInvalidQuery               = errors.New("invalid query")
	ErrInvalidPath                = errors.New("invalid path instance")
	ErrInvalidPathString          = errors.New("invalid path string")
	ErrNotFoundNode               = errors.New("node not found")
	ErrUnknownCommentPositionType = errors.New("unknown comment position type")
	ErrInvalidCommentMapValue     = errors.New("invalid comment map value. it must be not nil value")
	ErrDecodeRequiredPointerType  = errors.New("required pointer type value")
	ErrExceededMaxDepth           = errors.New("exceeded max depth")
	FormatErrorWithToken          = errors.FormatError
)

type (
	SyntaxError             = errors.SyntaxError
	TypeError               = errors.TypeError
	OverflowError           = errors.OverflowError
	DuplicateKeyError       = errors.DuplicateKeyError
	UnknownFieldError       = errors.UnknownFieldError
	UnexpectedNodeTypeError = errors.UnexpectedNodeTypeError
	Error                   = errors.Error
)

func ErrUnsupportedHeadPositionType(node ast.Node) error { _ = "STUB: not implemented"; return nil }

func ErrUnsupportedLinePositionType(node ast.Node) error { _ = "STUB: not implemented"; return nil }

func ErrUnsupportedFootPositionType(node ast.Node) error { _ = "STUB: not implemented"; return nil }

// IsInvalidQueryError whether err is ErrInvalidQuery or not.
func IsInvalidQueryError(err error) bool { _ = "STUB: not implemented"; return false }

// IsInvalidPathError whether err is ErrInvalidPath or not.
func IsInvalidPathError(err error) bool { _ = "STUB: not implemented"; return false }

// IsInvalidPathStringError whether err is ErrInvalidPathString or not.
func IsInvalidPathStringError(err error) bool { _ = "STUB: not implemented"; return false }

// IsNotFoundNodeError whether err is ErrNotFoundNode or not.
func IsNotFoundNodeError(err error) bool { _ = "STUB: not implemented"; return false }

// IsInvalidTokenTypeError whether err is ast.ErrInvalidTokenType or not.
func IsInvalidTokenTypeError(err error) bool { _ = "STUB: not implemented"; return false }

// IsInvalidAnchorNameError whether err is ast.ErrInvalidAnchorName or not.
func IsInvalidAnchorNameError(err error) bool { _ = "STUB: not implemented"; return false }

// IsInvalidAliasNameError whether err is ast.ErrInvalidAliasName or not.
func IsInvalidAliasNameError(err error) bool { _ = "STUB: not implemented"; return false }
