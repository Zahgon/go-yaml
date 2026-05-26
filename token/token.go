package token

import (
	"time"
)

// Character type for character
type Character byte

const (
	// SequenceEntryCharacter character for sequence entry
	SequenceEntryCharacter Character = '-'
	// MappingKeyCharacter character for mapping key
	MappingKeyCharacter Character = '?'
	// MappingValueCharacter character for mapping value
	MappingValueCharacter Character = ':'
	// CollectEntryCharacter character for collect entry
	CollectEntryCharacter Character = ','
	// SequenceStartCharacter character for sequence start
	SequenceStartCharacter Character = '['
	// SequenceEndCharacter character for sequence end
	SequenceEndCharacter Character = ']'
	// MappingStartCharacter character for mapping start
	MappingStartCharacter Character = '{'
	// MappingEndCharacter character for mapping end
	MappingEndCharacter Character = '}'
	// CommentCharacter character for comment
	CommentCharacter Character = '#'
	// AnchorCharacter character for anchor
	AnchorCharacter Character = '&'
	// AliasCharacter character for alias
	AliasCharacter Character = '*'
	// TagCharacter character for tag
	TagCharacter Character = '!'
	// LiteralCharacter character for literal
	LiteralCharacter Character = '|'
	// FoldedCharacter character for folded
	FoldedCharacter Character = '>'
	// SingleQuoteCharacter character for single quote
	SingleQuoteCharacter Character = '\''
	// DoubleQuoteCharacter character for double quote
	DoubleQuoteCharacter Character = '"'
	// DirectiveCharacter character for directive
	DirectiveCharacter Character = '%'
	// SpaceCharacter character for space
	SpaceCharacter Character = ' '
	// LineBreakCharacter character for line break
	LineBreakCharacter Character = '\n'
)

// Type type identifier for token
type Type int

const (
	// UnknownType reserve for invalid type
	UnknownType Type = iota
	// DocumentHeaderType type for DocumentHeader token
	DocumentHeaderType
	// DocumentEndType type for DocumentEnd token
	DocumentEndType
	// SequenceEntryType type for SequenceEntry token
	SequenceEntryType
	// MappingKeyType type for MappingKey token
	MappingKeyType
	// MappingValueType type for MappingValue token
	MappingValueType
	// MergeKeyType type for MergeKey token
	MergeKeyType
	// CollectEntryType type for CollectEntry token
	CollectEntryType
	// SequenceStartType type for SequenceStart token
	SequenceStartType
	// SequenceEndType type for SequenceEnd token
	SequenceEndType
	// MappingStartType type for MappingStart token
	MappingStartType
	// MappingEndType type for MappingEnd token
	MappingEndType
	// CommentType type for Comment token
	CommentType
	// AnchorType type for Anchor token
	AnchorType
	// AliasType type for Alias token
	AliasType
	// TagType type for Tag token
	TagType
	// LiteralType type for Literal token
	LiteralType
	// FoldedType type for Folded token
	FoldedType
	// SingleQuoteType type for SingleQuote token
	SingleQuoteType
	// DoubleQuoteType type for DoubleQuote token
	DoubleQuoteType
	// DirectiveType type for Directive token
	DirectiveType
	// SpaceType type for Space token
	SpaceType
	// NullType type for Null token
	NullType
	// ImplicitNullType type for implicit Null token.
	// This is used when explicit keywords such as null or ~ are not specified.
	// It is distinguished during encoding and output as an empty string.
	ImplicitNullType
	// InfinityType type for Infinity token
	InfinityType
	// NanType type for Nan token
	NanType
	// IntegerType type for Integer token
	IntegerType
	// BinaryIntegerType type for BinaryInteger token
	BinaryIntegerType
	// OctetIntegerType type for OctetInteger token
	OctetIntegerType
	// HexIntegerType type for HexInteger token
	HexIntegerType
	// FloatType type for Float token
	FloatType
	// StringType type for String token
	StringType
	// BoolType type for Bool token
	BoolType
	// InvalidType type for invalid token
	InvalidType
)

// String type identifier to text
func (t Type) String() string { _ = "STUB: not implemented"; return "" }

// CharacterType type for character category
type CharacterType int

const (
	// CharacterTypeIndicator type of indicator character
	CharacterTypeIndicator CharacterType = iota
	// CharacterTypeWhiteSpace type of white space character
	CharacterTypeWhiteSpace
	// CharacterTypeMiscellaneous type of miscellaneous character
	CharacterTypeMiscellaneous
	// CharacterTypeEscaped type of escaped character
	CharacterTypeEscaped
	// CharacterTypeInvalid type for a invalid token.
	CharacterTypeInvalid
)

// String character type identifier to text
func (c CharacterType) String() string { _ = "STUB: not implemented"; return "" }

// Indicator type for indicator
type Indicator int

const (
	// NotIndicator not indicator
	NotIndicator Indicator = iota
	// BlockStructureIndicator indicator for block structure ( '-', '?', ':' )
	BlockStructureIndicator
	// FlowCollectionIndicator indicator for flow collection ( '[', ']', '{', '}', ',' )
	FlowCollectionIndicator
	// CommentIndicator indicator for comment ( '#' )
	CommentIndicator
	// NodePropertyIndicator indicator for node property ( '!', '&', '*' )
	NodePropertyIndicator
	// BlockScalarIndicator indicator for block scalar ( '|', '>' )
	BlockScalarIndicator
	// QuotedScalarIndicator indicator for quoted scalar ( ''', '"' )
	QuotedScalarIndicator
	// DirectiveIndicator indicator for directive ( '%' )
	DirectiveIndicator
	// InvalidUseOfReservedIndicator indicator for invalid use of reserved keyword ( '@', '`' )
	InvalidUseOfReservedIndicator
)

// String indicator to text
func (i Indicator) String() string { _ = "STUB: not implemented"; return "" }

var (
	reservedNullKeywords = []string{
		"null",
		"Null",
		"NULL",
		"~",
	}
	reservedBoolKeywords = []string{
		"true",
		"True",
		"TRUE",
		"false",
		"False",
		"FALSE",
	}
	// For compatibility with other YAML 1.1 parsers
	// Note that we use these solely for encoding the bool value with quotes.
	// go-yaml should not treat these as reserved keywords at parsing time.
	// as go-yaml is supposed to be compliant only with YAML 1.2.
	reservedLegacyBoolKeywords = []string{
		"y",
		"Y",
		"yes",
		"Yes",
		"YES",
		"n",
		"N",
		"no",
		"No",
		"NO",
		"on",
		"On",
		"ON",
		"off",
		"Off",
		"OFF",
	}
	reservedInfKeywords = []string{
		".inf",
		".Inf",
		".INF",
		"-.inf",
		"-.Inf",
		"-.INF",
	}
	reservedNanKeywords = []string{
		".nan",
		".NaN",
		".NAN",
	}
	reservedKeywordMap = map[string]func(string, string, *Position) *Token{}
	// reservedEncKeywordMap contains is the keyword map used at encoding time.
	// This is supposed to be a superset of reservedKeywordMap,
	// and used to quote legacy keywords present in YAML 1.1 or lesser for compatibility reasons,
	// even though this library is supposed to be YAML 1.2-compliant.
	reservedEncKeywordMap = map[string]func(string, string, *Position) *Token{}
)

func reservedKeywordToken(typ Type, value, org string, pos *Position) *Token {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	for _, keyword := range reservedNullKeywords {
		f := func(value, org string, pos *Position) *Token {
			return reservedKeywordToken(NullType, value, org, pos)
		}

		reservedKeywordMap[keyword] = f
		reservedEncKeywordMap[keyword] = f
	}
	for _, keyword := range reservedBoolKeywords {
		f := func(value, org string, pos *Position) *Token {
			return reservedKeywordToken(BoolType, value, org, pos)
		}
		reservedKeywordMap[keyword] = f
		reservedEncKeywordMap[keyword] = f
	}
	for _, keyword := range reservedLegacyBoolKeywords {
		reservedEncKeywordMap[keyword] = func(value, org string, pos *Position) *Token {
			return reservedKeywordToken(BoolType, value, org, pos)
		}
	}
	for _, keyword := range reservedInfKeywords {
		reservedKeywordMap[keyword] = func(value, org string, pos *Position) *Token {
			return reservedKeywordToken(InfinityType, value, org, pos)
		}
	}
	for _, keyword := range reservedNanKeywords {
		reservedKeywordMap[keyword] = func(value, org string, pos *Position) *Token {
			return reservedKeywordToken(NanType, value, org, pos)
		}
	}
}

// ReservedTagKeyword type of reserved tag keyword
type ReservedTagKeyword string

const (
	// IntegerTag `!!int` tag
	IntegerTag ReservedTagKeyword = "!!int"
	// FloatTag `!!float` tag
	FloatTag ReservedTagKeyword = "!!float"
	// NullTag `!!null` tag
	NullTag ReservedTagKeyword = "!!null"
	// SequenceTag `!!seq` tag
	SequenceTag ReservedTagKeyword = "!!seq"
	// MappingTag `!!map` tag
	MappingTag ReservedTagKeyword = "!!map"
	// StringTag `!!str` tag
	StringTag ReservedTagKeyword = "!!str"
	// BinaryTag `!!binary` tag
	BinaryTag ReservedTagKeyword = "!!binary"
	// OrderedMapTag `!!omap` tag
	OrderedMapTag ReservedTagKeyword = "!!omap"
	// SetTag `!!set` tag
	SetTag ReservedTagKeyword = "!!set"
	// TimestampTag `!!timestamp` tag
	TimestampTag ReservedTagKeyword = "!!timestamp"
	// BooleanTag `!!bool` tag
	BooleanTag ReservedTagKeyword = "!!bool"
	// MergeTag `!!merge` tag
	MergeTag ReservedTagKeyword = "!!merge"
)

var (
	// ReservedTagKeywordMap map for reserved tag keywords
	ReservedTagKeywordMap = map[ReservedTagKeyword]func(string, string, *Position) *Token{
		IntegerTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		FloatTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		NullTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		SequenceTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		MappingTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		StringTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		BinaryTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		OrderedMapTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		SetTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		TimestampTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		BooleanTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
		MergeTag: func(value, org string, pos *Position) *Token {
			return &Token{
				Type:          TagType,
				CharacterType: CharacterTypeIndicator,
				Indicator:     NodePropertyIndicator,
				Value:         value,
				Origin:        org,
				Position:      pos,
			}
		},
	}
)

type NumberType string

const (
	NumberTypeDecimal NumberType = "decimal"
	NumberTypeBinary  NumberType = "binary"
	NumberTypeOctet   NumberType = "octet"
	NumberTypeHex     NumberType = "hex"
	NumberTypeFloat   NumberType = "float"
)

type NumberValue struct {
	Type  NumberType
	Value any
	Text  string
}

func ToNumber(value string) *NumberValue { _ = "STUB: not implemented"; return nil }

func isNumber(value string) bool { _ = "STUB: not implemented"; return false }

func toNumber(value string) (*NumberValue, error) { _ = "STUB: not implemented"; return nil, nil }

// This is a subset of the formats permitted by the regular expression
// defined at http://yaml.org/type/timestamp.html. Note that time.Parse
// cannot handle: "2001-12-14 21:59:43.10 -5" from the examples.
var timestampFormats = []string{
	time.RFC3339Nano,
	"2006-01-02t15:04:05.999999999Z07:00", // RFC3339Nano with lower-case "t".
	time.DateTime,
	time.DateOnly,

	// Not in examples, but to preserve backward compatibility by quoting time values.
	"15:4",
}

func isTimestamp(value string) bool { _ = "STUB: not implemented"; return false }

// IsNeedQuoted checks whether the value needs quote for passed string or not
func IsNeedQuoted(value string) bool { _ = "STUB: not implemented"; return false }

// LiteralBlockHeader detect literal block scalar header
func LiteralBlockHeader(value string) string { _ = "STUB: not implemented"; return "" }

// New create reserved keyword token or number token and other string token.
func New(value string, org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Position type for position in YAML document
type Position struct {
	Line        int
	Column      int
	Offset      int
	IndentNum   int
	IndentLevel int
}

// String position to text
func (p *Position) String() string { _ = "STUB: not implemented"; return "" }

// Token type for token
type Token struct {
	// Type is a token type.
	Type Type
	// CharacterType is a character type.
	CharacterType CharacterType
	// Indicator is a indicator type.
	Indicator Indicator
	// Value is a string extracted with only meaningful characters, with spaces and such removed.
	Value string
	// Origin is a string that stores the original text as-is.
	Origin string
	// Error keeps error message for InvalidToken.
	Error string
	// Position is a token position.
	Position *Position
	// Next is a next token reference.
	Next *Token
	// Prev is a previous token reference.
	Prev *Token
}

// PreviousType previous token type
func (t *Token) PreviousType() Type { _ = "STUB: not implemented"; return *new(Type) }

// NextType next token type
func (t *Token) NextType() Type { _ = "STUB: not implemented"; return *new(Type) }

// AddColumn append column number to current position of column
func (t *Token) AddColumn(col int) { _ = "STUB: not implemented"; return }

// Clone copy token ( preserve Prev/Next reference )
func (t *Token) Clone() *Token { _ = "STUB: not implemented"; return nil }

// Dump outputs token information to stdout for debugging.
func (t *Token) Dump() { _ = "STUB: not implemented"; return }

// Tokens type of token collection
type Tokens []*Token

func (t Tokens) InvalidToken() *Token { _ = "STUB: not implemented"; return nil }

func (t *Tokens) add(tk *Token) { _ = "STUB: not implemented"; return }

// Add append new some tokens
func (t *Tokens) Add(tks ...*Token) { _ = "STUB: not implemented"; return }

// Dump dump all token structures for debugging
func (t Tokens) Dump() { _ = "STUB: not implemented"; return }

// String create token for String
func String(value string, org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// SequenceEntry create token for SequenceEntry
func SequenceEntry(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// MappingKey create token for MappingKey
func MappingKey(pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// MappingValue create token for MappingValue
func MappingValue(pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// CollectEntry create token for CollectEntry
func CollectEntry(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// SequenceStart create token for SequenceStart
func SequenceStart(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// SequenceEnd create token for SequenceEnd
func SequenceEnd(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// MappingStart create token for MappingStart
func MappingStart(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// MappingEnd create token for MappingEnd
func MappingEnd(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Comment create token for Comment
func Comment(value string, org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Anchor create token for Anchor
func Anchor(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Alias create token for Alias
func Alias(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Tag create token for Tag
func Tag(value string, org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Literal create token for Literal
func Literal(value string, org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Folded create token for Folded
func Folded(value string, org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// SingleQuote create token for SingleQuote
func SingleQuote(value string, org string, pos *Position) *Token {
	_ = "STUB: not implemented"
	return nil
}

// DoubleQuote create token for DoubleQuote
func DoubleQuote(value string, org string, pos *Position) *Token {
	_ = "STUB: not implemented"
	return nil
}

// Directive create token for Directive
func Directive(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// Space create token for Space
func Space(pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// MergeKey create token for MergeKey
func MergeKey(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// DocumentHeader create token for DocumentHeader
func DocumentHeader(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// DocumentEnd create token for DocumentEnd
func DocumentEnd(org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

func Invalid(err string, org string, pos *Position) *Token { _ = "STUB: not implemented"; return nil }

// DetectLineBreakCharacter detect line break character in only one inside scalar content scope.
func DetectLineBreakCharacter(src string) string { _ = "STUB: not implemented"; return "" }
