package yaml

import (
	"context"
	"io"

	"github.com/goccy/go-yaml/ast"
)

// DecodeOption functional option type for Decoder
type DecodeOption func(d *Decoder) error

// ReferenceReaders pass to Decoder that reference to anchor defined by passed readers
func ReferenceReaders(readers ...io.Reader) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// ReferenceFiles pass to Decoder that reference to anchor defined by passed files
func ReferenceFiles(files ...string) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// ReferenceDirs pass to Decoder that reference to anchor defined by files under the passed dirs
func ReferenceDirs(dirs ...string) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// RecursiveDir search yaml file recursively from passed dirs by ReferenceDirs option
func RecursiveDir(isRecursive bool) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// Validator set StructValidator instance to Decoder
func Validator(v StructValidator) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// Strict enable DisallowUnknownField
func Strict() DecodeOption { _ = "STUB: not implemented"; return *new(DecodeOption) }

// DisallowUnknownField causes the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func DisallowUnknownField() DecodeOption { _ = "STUB: not implemented"; return *new(DecodeOption) }

// AllowFieldPrefixes, when paired with [DisallowUnknownField], allows fields
// with the specified prefixes to bypass the unknown field check.
func AllowFieldPrefixes(prefixes ...string) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// AllowDuplicateMapKey ignore syntax error when mapping keys that are duplicates.
func AllowDuplicateMapKey() DecodeOption { _ = "STUB: not implemented"; return *new(DecodeOption) }

// UseOrderedMap can be interpreted as a map,
// and uses MapSlice ( ordered map ) aggressively if there is no type specification
func UseOrderedMap() DecodeOption { _ = "STUB: not implemented"; return *new(DecodeOption) }

// UseJSONUnmarshaler if neither `BytesUnmarshaler` nor `InterfaceUnmarshaler` is implemented
// and `UnmashalJSON([]byte)error` is implemented, convert the argument from `YAML` to `JSON` and then call it.
func UseJSONUnmarshaler() DecodeOption { _ = "STUB: not implemented"; return *new(DecodeOption) }

// CustomUnmarshaler overrides any decoding process for the type specified in generics.
//
// NOTE: If RegisterCustomUnmarshaler and CustomUnmarshaler of DecodeOption are specified for the same type,
// the CustomUnmarshaler specified in DecodeOption takes precedence.
func CustomUnmarshaler[T any](unmarshaler func(*T, []byte) error) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// CustomUnmarshalerContext overrides any decoding process for the type specified in generics.
// Similar to CustomUnmarshaler, but allows passing a context to the unmarshaler function.
func CustomUnmarshalerContext[T any](unmarshaler func(context.Context, *T, []byte) error) DecodeOption {
	_ = "STUB: not implemented"
	return *new(DecodeOption)
}

// EncodeOption functional option type for Encoder
type EncodeOption func(e *Encoder) error

// Indent change indent number
func Indent(spaces int) EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// IndentSequence causes sequence values to be indented the same value as Indent
func IndentSequence(indent bool) EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// UseSingleQuote determines if single or double quotes should be preferred for strings.
func UseSingleQuote(sq bool) EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// Flow encoding by flow style
func Flow(isFlowStyle bool) EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// WithSmartAnchor when multiple map values share the same pointer,
// an anchor is automatically assigned to the first occurrence, and aliases are used for subsequent elements.
// The map key name is used as the anchor name by default.
// If key names conflict, a suffix is automatically added to avoid collisions.
// This is an experimental feature and cannot be used simultaneously with anchor tags.
func WithSmartAnchor() EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// UseLiteralStyleIfMultiline causes encoding multiline strings with a literal syntax,
// no matter what characters they include
func UseLiteralStyleIfMultiline(useLiteralStyleIfMultiline bool) EncodeOption {
	_ = "STUB: not implemented"
	return *new(EncodeOption)
}

// JSON encode in JSON format
func JSON() EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// MarshalAnchor call back if encoder find an anchor during encoding
func MarshalAnchor(callback func(*ast.AnchorNode, interface{}) error) EncodeOption {
	_ = "STUB: not implemented"
	return *new(EncodeOption)
}

// UseJSONMarshaler if neither `BytesMarshaler` nor `InterfaceMarshaler`
// nor `encoding.TextMarshaler` is implemented and `MarshalJSON()([]byte, error)` is implemented,
// call `MarshalJSON` to convert the returned `JSON` to `YAML` for processing.
func UseJSONMarshaler() EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// CustomMarshaler overrides any encoding process for the type specified in generics.
//
// NOTE: If type T implements MarshalYAML for pointer receiver, the type specified in CustomMarshaler must be *T.
// If RegisterCustomMarshaler and CustomMarshaler of EncodeOption are specified for the same type,
// the CustomMarshaler specified in EncodeOption takes precedence.
func CustomMarshaler[T any](marshaler func(T) ([]byte, error)) EncodeOption {
	_ = "STUB: not implemented"
	return *new(EncodeOption)
}

// CustomMarshalerContext overrides any encoding process for the type specified in generics.
// Similar to CustomMarshaler, but allows passing a context to the marshaler function.
func CustomMarshalerContext[T any](marshaler func(context.Context, T) ([]byte, error)) EncodeOption {
	_ = "STUB: not implemented"
	return *new(EncodeOption)
}

// AutoInt automatically converts floating-point numbers to integers when the fractional part is zero.
// For example, a value of 1.0 will be encoded as 1.
func AutoInt() EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// OmitEmpty behaves in the same way as the interpretation of the omitempty tag in the encoding/json library.
// set on all the fields.
// In the current implementation, the omitempty tag is not implemented in the same way as encoding/json,
// so please specify this option if you expect the same behavior.
func OmitEmpty() EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// OmitZero forces the encoder to assume an `omitzero` struct tag is
// set on all the fields. See `Marshal` commentary for the `omitzero` tag logic.
func OmitZero() EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// CommentPosition type of the position for comment.
type CommentPosition int

const (
	CommentHeadPosition CommentPosition = CommentPosition(iota)
	CommentLinePosition
	CommentFootPosition
)

func (p CommentPosition) String() string { _ = "STUB: not implemented"; return "" }

// LineComment create a one-line comment for CommentMap.
func LineComment(text string) *Comment { _ = "STUB: not implemented"; return nil }

// HeadComment create a multiline comment for CommentMap.
func HeadComment(texts ...string) *Comment { _ = "STUB: not implemented"; return nil }

// FootComment create a multiline comment for CommentMap.
func FootComment(texts ...string) *Comment { _ = "STUB: not implemented"; return nil }

// Comment raw data for comment.
type Comment struct {
	Texts    []string
	Position CommentPosition
}

// CommentMap map of the position of the comment and the comment information.
type CommentMap map[string][]*Comment

// WithComment add a comment using the location and text information given in the CommentMap.
func WithComment(cm CommentMap) EncodeOption { _ = "STUB: not implemented"; return *new(EncodeOption) }

// CommentToMap apply the position and content of comments in a YAML document to a CommentMap.
func CommentToMap(cm CommentMap) DecodeOption { _ = "STUB: not implemented"; return *new(DecodeOption) }
