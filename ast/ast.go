package ast

import (
	"errors"
	"io"

	"github.com/goccy/go-yaml/token"
)

var (
	ErrInvalidTokenType  = errors.New("invalid token type")
	ErrInvalidAnchorName = errors.New("invalid anchor name")
	ErrInvalidAliasName  = errors.New("invalid alias name")
)

// NodeType type identifier of node
type NodeType int

const (
	// UnknownNodeType type identifier for default
	UnknownNodeType NodeType = iota
	// DocumentType type identifier for document node
	DocumentType
	// NullType type identifier for null node
	NullType
	// BoolType type identifier for boolean node
	BoolType
	// IntegerType type identifier for integer node
	IntegerType
	// FloatType type identifier for float node
	FloatType
	// InfinityType type identifier for infinity node
	InfinityType
	// NanType type identifier for nan node
	NanType
	// StringType type identifier for string node
	StringType
	// MergeKeyType type identifier for merge key node
	MergeKeyType
	// LiteralType type identifier for literal node
	LiteralType
	// MappingType type identifier for mapping node
	MappingType
	// MappingKeyType type identifier for mapping key node
	MappingKeyType
	// MappingValueType type identifier for mapping value node
	MappingValueType
	// SequenceType type identifier for sequence node
	SequenceType
	// SequenceEntryType type identifier for sequence entry node
	SequenceEntryType
	// AnchorType type identifier for anchor node
	AnchorType
	// AliasType type identifier for alias node
	AliasType
	// DirectiveType type identifier for directive node
	DirectiveType
	// TagType type identifier for tag node
	TagType
	// CommentType type identifier for comment node
	CommentType
	// CommentGroupType type identifier for comment group node
	CommentGroupType
)

// String node type identifier to text
func (t NodeType) String() string { _ = "STUB: not implemented"; return "" }

// String node type identifier to YAML Structure name
// based on https://yaml.org/spec/1.2/spec.html
func (t NodeType) YAMLName() string { _ = "STUB: not implemented"; return "" }

// Node type of node
type Node interface {
	io.Reader
	// String node to text
	String() string
	// GetToken returns token instance
	GetToken() *token.Token
	// Type returns type of node
	Type() NodeType
	// AddColumn add column number to child nodes recursively
	AddColumn(int)
	// SetComment set comment token to node
	SetComment(*CommentGroupNode) error
	// Comment returns comment token instance
	GetComment() *CommentGroupNode
	// GetPath returns YAMLPath for the current node
	GetPath() string
	// SetPath set YAMLPath for the current node
	SetPath(string)
	// MarshalYAML
	MarshalYAML() ([]byte, error)
	// already read length
	readLen() int
	// append read length
	addReadLen(int)
	// clean read length
	clearLen()
}

// MapKeyNode type for map key node
type MapKeyNode interface {
	Node
	IsMergeKey() bool
	// String node to text without comment
	stringWithoutComment() string
}

// ScalarNode type for scalar node
type ScalarNode interface {
	MapKeyNode
	GetValue() interface{}
}

type BaseNode struct {
	Path    string
	Comment *CommentGroupNode
	read    int
}

func addCommentString(base string, node *CommentGroupNode) string {
	_ = "STUB: not implemented"
	return ""
}

func (n *BaseNode) readLen() int { _ = "STUB: not implemented"; return 0 }

func (n *BaseNode) clearLen() { _ = "STUB: not implemented"; return }

func (n *BaseNode) addReadLen(len int) {
	_ = "STUB: not implemented"

	// GetPath returns YAMLPath for the current node.
	return
}

func (n *BaseNode) GetPath() string { _ = "STUB: not implemented"; return "" }

// SetPath set YAMLPath for the current node.
func (n *BaseNode) SetPath(path string) { _ = "STUB: not implemented"; return }

// GetComment returns comment token instance
func (n *BaseNode) GetComment() *CommentGroupNode {
	_ = "STUB: not implemented"

	// SetComment set comment token
	return nil
}

func (n *BaseNode) SetComment(node *CommentGroupNode) error { _ = "STUB: not implemented"; return nil }

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

func readNode(p []byte, node Node) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func checkLineBreak(t *token.Token) bool { _ = "STUB: not implemented"; return false }

// if the previous type is sequence entry use the previous type for that

// as well as switching to previous type count any new lines in origin to account for:
// -
//   b: c

// Remove any line breaks included in multiline string

// Due to the way that comment parsing works its assumed that when a null value does not have new line in origin
// it was squashed therefore difference is ignored.
// foo:
//  bar:
//  # comment
//  baz: 1
// becomes
// foo:
//  bar: null # comment
//
//  baz: 1

// Null create node for null value
func Null(tk *token.Token) *NullNode { _ = "STUB: not implemented"; return nil }

// Bool create node for boolean value
func Bool(tk *token.Token) *BoolNode { _ = "STUB: not implemented"; return nil }

// Integer create node for integer value
func Integer(tk *token.Token) *IntegerNode { _ = "STUB: not implemented"; return nil }

// Float create node for float value
func Float(tk *token.Token) *FloatNode { _ = "STUB: not implemented"; return nil }

// Infinity create node for .inf or -.inf value
func Infinity(tk *token.Token) *InfinityNode { _ = "STUB: not implemented"; return nil }

// Nan create node for .nan value
func Nan(tk *token.Token) *NanNode { _ = "STUB: not implemented"; return nil }

// String create node for string value
func String(tk *token.Token) *StringNode { _ = "STUB: not implemented"; return nil }

// Comment create node for comment
func Comment(tk *token.Token) *CommentNode { _ = "STUB: not implemented"; return nil }

func CommentGroup(comments []*token.Token) *CommentGroupNode { _ = "STUB: not implemented"; return nil }

// MergeKey create node for merge key ( << )
func MergeKey(tk *token.Token) *MergeKeyNode { _ = "STUB: not implemented"; return nil }

// Mapping create node for map
func Mapping(tk *token.Token, isFlowStyle bool, values ...*MappingValueNode) *MappingNode {
	_ = "STUB: not implemented"
	return nil
}

// MappingValue create node for mapping value
func MappingValue(tk *token.Token, key MapKeyNode, value Node) *MappingValueNode {
	_ = "STUB: not implemented"
	return nil
}

// MappingKey create node for map key ( '?' ).
func MappingKey(tk *token.Token) *MappingKeyNode { _ = "STUB: not implemented"; return nil }

// Sequence create node for sequence
func Sequence(tk *token.Token, isFlowStyle bool) *SequenceNode {
	_ = "STUB: not implemented"
	return nil
}

func Anchor(tk *token.Token) *AnchorNode { _ = "STUB: not implemented"; return nil }

func Alias(tk *token.Token) *AliasNode { _ = "STUB: not implemented"; return nil }

func Document(tk *token.Token, body Node) *DocumentNode { _ = "STUB: not implemented"; return nil }

func Directive(tk *token.Token) *DirectiveNode { _ = "STUB: not implemented"; return nil }

func Literal(tk *token.Token) *LiteralNode { _ = "STUB: not implemented"; return nil }

func Tag(tk *token.Token) *TagNode { _ = "STUB: not implemented"; return nil }

// File contains all documents in YAML file
type File struct {
	Name string
	Docs []*DocumentNode
}

// Read implements (io.Reader).Read
func (f *File) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// String all documents to text
func (f *File) String() string { _ = "STUB: not implemented"; return "" }

// DocumentNode type of Document
type DocumentNode struct {
	*BaseNode
	Start *token.Token // position of DocumentHeader ( `---` )
	End   *token.Token // position of DocumentEnd ( `...` )
	Body  Node
}

// Read implements (io.Reader).Read
func (d *DocumentNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns DocumentNodeType
		nil
}

func (d *DocumentNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (d *DocumentNode) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

// AddColumn add column number to child nodes recursively
func (d *DocumentNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String document to text
func (d *DocumentNode) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (d *DocumentNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NullNode type of null node
type NullNode struct {
	*BaseNode
	Token *token.Token
}

// Read implements (io.Reader).Read
func (n *NullNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns NullType
		nil
}

func (n *NullNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *NullNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *NullNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns nil value
func (n *NullNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// String returns `null` text
	return nil
}

func (n *NullNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *NullNode) stringWithoutComment() string {
	_ = "STUB: not implemented"

	// MarshalYAML encodes to a YAML text
	return ""
}

func (n *NullNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *NullNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// IntegerNode type of integer node
	return false
}

type IntegerNode struct {
	*BaseNode
	Token *token.Token
	Value interface{} // int64 or uint64 value
}

// Read implements (io.Reader).Read
func (n *IntegerNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns IntegerType
		nil
}

func (n *IntegerNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *IntegerNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *IntegerNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns int64 value
func (n *IntegerNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// String int64 to text
	return nil
}

func (n *IntegerNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *IntegerNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *IntegerNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *IntegerNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// FloatNode type of float node
	return false
}

type FloatNode struct {
	*BaseNode
	Token     *token.Token
	Precision int
	Value     float64
}

// Read implements (io.Reader).Read
func (n *FloatNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns FloatType
		nil
}

func (n *FloatNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *FloatNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *FloatNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns float64 value
func (n *FloatNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// String float64 to text
	return nil
}

func (n *FloatNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *FloatNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *FloatNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *FloatNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// StringNode type of string node
	return false
}

type StringNode struct {
	*BaseNode
	Token *token.Token
	Value string
}

// Read implements (io.Reader).Read
func (n *StringNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns StringType
		nil
}

func (n *StringNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *StringNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *StringNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns string value
func (n *StringNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// IsMergeKey returns whether it is a MergeKey node.
	return nil
}

func (n *StringNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// escapeSingleQuote escapes s to a single quoted scalar.
	// https://yaml.org/spec/1.2.2/#732-single-quoted-style
	return false
}

func escapeSingleQuote(s string) string { _ = "STUB: not implemented"; return "" }

// s includes also one ' from the doubled pair
// opening and closing '
// ' added by ReplaceAll

// String string value to text with quote or literal header if required
func (n *StringNode) String() string { _ = "STUB: not implemented"; return "" }

// This block assumes that the line breaks in this inside scalar content and the Outside scalar content are the same.
// It works mostly, but inconsistencies occur if line break characters are mixed.

func (n *StringNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// This block assumes that the line breaks in this inside scalar content and the Outside scalar content are the same.
// It works mostly, but inconsistencies occur if line break characters are mixed.

// MarshalYAML encodes to a YAML text
func (n *StringNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// LiteralNode type of literal node
type LiteralNode struct {
	*BaseNode
	Start *token.Token
	Value *StringNode
}

// Read implements (io.Reader).Read
func (n *LiteralNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns LiteralType
		nil
}

func (n *LiteralNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *LiteralNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *LiteralNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns string value
func (n *LiteralNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// String literal to text
	return nil
}

func (n *LiteralNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *LiteralNode) stringWithoutComment() string {
	_ = "STUB: not implemented"

	// MarshalYAML encodes to a YAML text
	return ""
}

func (n *LiteralNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *LiteralNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// MergeKeyNode type of merge key node
	return false
}

type MergeKeyNode struct {
	*BaseNode
	Token *token.Token
}

// Read implements (io.Reader).Read
func (n *MergeKeyNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns MergeKeyType
		nil
}

func (n *MergeKeyNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *MergeKeyNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// GetValue returns '<<' value
	return nil
}

func (n *MergeKeyNode) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

// String returns '<<' value
func (n *MergeKeyNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *MergeKeyNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// AddColumn add column number to child nodes recursively
func (n *MergeKeyNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// MarshalYAML encodes to a YAML text
func (n *MergeKeyNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *MergeKeyNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// BoolNode type of boolean node
	return false
}

type BoolNode struct {
	*BaseNode
	Token *token.Token
	Value bool
}

// Read implements (io.Reader).Read
func (n *BoolNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns BoolType
		nil
}

func (n *BoolNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *BoolNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *BoolNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns boolean value
func (n *BoolNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// String boolean to text
	return nil
}

func (n *BoolNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *BoolNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *BoolNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *BoolNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// InfinityNode type of infinity node
	return false
}

type InfinityNode struct {
	*BaseNode
	Token *token.Token
	Value float64
}

// Read implements (io.Reader).Read
func (n *InfinityNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns InfinityType
		nil
}

func (n *InfinityNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *InfinityNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *InfinityNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns math.Inf(0) or math.Inf(-1)
func (n *InfinityNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// String infinity to text
	return nil
}

func (n *InfinityNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *InfinityNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *InfinityNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *InfinityNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// NanNode type of nan node
	return false
}

type NanNode struct {
	*BaseNode
	Token *token.Token
}

// Read implements (io.Reader).Read
func (n *NanNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns NanType
		nil
}

func (n *NanNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *NanNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *NanNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// GetValue returns math.NaN()
func (n *NanNode) GetValue() interface{} {
	_ = "STUB: not implemented"

	// String returns .nan
	return nil
}

func (n *NanNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *NanNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *NanNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *NanNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// MapNode interface of MappingValueNode / MappingNode
	return false
}

type MapNode interface {
	MapRange() *MapNodeIter
}

// MapNodeIter is an iterator for ranging over a MapNode
type MapNodeIter struct {
	values []*MappingValueNode
	idx    int
}

const (
	startRangeIndex = -1
)

// Next advances the map iterator and reports whether there is another entry.
// It returns false when the iterator is exhausted.
func (m *MapNodeIter) Next() bool { _ = "STUB: not implemented"; return false }

// Key returns the key of the iterator's current map node entry.
func (m *MapNodeIter) Key() MapKeyNode { _ = "STUB: not implemented"; return *new(MapKeyNode) }

// Value returns the value of the iterator's current map node entry.
func (m *MapNodeIter) Value() Node { _ = "STUB: not implemented"; return *new(Node) }

// KeyValue returns the MappingValueNode of the iterator's current map node entry.
func (m *MapNodeIter) KeyValue() *MappingValueNode { _ = "STUB: not implemented"; return nil }

// MappingNode type of mapping node
type MappingNode struct {
	*BaseNode
	Start       *token.Token
	End         *token.Token
	IsFlowStyle bool
	Values      []*MappingValueNode
	FootComment *CommentGroupNode
}

func (n *MappingNode) startPos() *token.Position { _ = "STUB: not implemented"; return nil }

// Merge merge key/value of map.
func (n *MappingNode) Merge(target *MappingNode) { _ = "STUB: not implemented"; return }

// SetIsFlowStyle set value to IsFlowStyle field recursively.
func (n *MappingNode) SetIsFlowStyle(isFlow bool) { _ = "STUB: not implemented"; return }

// Read implements (io.Reader).Read
func (n *MappingNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns MappingType
		nil
}

func (n *MappingNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *MappingNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *MappingNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

func (n *MappingNode) flowStyleString(commentMode bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (n *MappingNode) blockStyleString(commentMode bool) string {
	_ = "STUB: not implemented"
	return ""
}

// String mapping values to text
func (n *MappingNode) String() string { _ = "STUB: not implemented"; return "" }

// MapRange implements MapNode protocol
func (n *MappingNode) MapRange() *MapNodeIter { _ = "STUB: not implemented"; return nil }

// MarshalYAML encodes to a YAML text
func (n *MappingNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MappingKeyNode type of tag node
type MappingKeyNode struct {
	*BaseNode
	Start *token.Token
	Value Node
}

// Read implements (io.Reader).Read
func (n *MappingKeyNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns MappingKeyType
		nil
}

func (n *MappingKeyNode) Type() NodeType {
	_ = "STUB: not implemented"
	return *

	// GetToken returns token instance
	new(NodeType)
}

func (n *MappingKeyNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *MappingKeyNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String tag to text
func (n *MappingKeyNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *MappingKeyNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *MappingKeyNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *MappingKeyNode) IsMergeKey() bool { _ = "STUB: not implemented"; return false }

// MappingValueNode type of mapping value
type MappingValueNode struct {
	*BaseNode
	Start        *token.Token // delimiter token ':'.
	CollectEntry *token.Token // collect entry token ','.
	Key          MapKeyNode
	Value        Node
	FootComment  *CommentGroupNode
	IsFlowStyle  bool
}

// Replace replace value node.
func (n *MappingValueNode) Replace(value Node) error { _ = "STUB: not implemented"; return nil }

// Read implements (io.Reader).Read
func (n *MappingValueNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns MappingValueType
		nil
}

func (n *MappingValueNode) Type() NodeType {
	_ = "STUB: not implemented"
	return *

	// GetToken returns token instance
	new(NodeType)
}

func (n *MappingValueNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *MappingValueNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// SetIsFlowStyle set value to IsFlowStyle field recursively.
func (n *MappingValueNode) SetIsFlowStyle(isFlow bool) { _ = "STUB: not implemented"; return }

// String mapping value to text
func (n *MappingValueNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *MappingValueNode) toString() string { _ = "STUB: not implemented"; return "" }

// implicit null value.

// For flow-style values indented on the next line, we need to add the proper indentation

// MapRange implements MapNode protocol
func (n *MappingValueNode) MapRange() *MapNodeIter { _ = "STUB: not implemented"; return nil }

// MarshalYAML encodes to a YAML text
func (n *MappingValueNode) MarshalYAML() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// ArrayNode interface of SequenceNode
}

type ArrayNode interface {
	ArrayRange() *ArrayNodeIter
}

// ArrayNodeIter is an iterator for ranging over a ArrayNode
type ArrayNodeIter struct {
	values []Node
	idx    int
}

// Next advances the array iterator and reports whether there is another entry.
// It returns false when the iterator is exhausted.
func (m *ArrayNodeIter) Next() bool { _ = "STUB: not implemented"; return false }

// Value returns the value of the iterator's current array entry.
func (m *ArrayNodeIter) Value() Node {
	_ = "STUB: not implemented"
	return *

	// Len returns length of array
	new(Node)
}

func (m *ArrayNodeIter) Len() int { _ = "STUB: not implemented"; return 0 }

// SequenceNode type of sequence node
type SequenceNode struct {
	*BaseNode
	Start             *token.Token
	End               *token.Token
	IsFlowStyle       bool
	Values            []Node
	ValueHeadComments []*CommentGroupNode
	Entries           []*SequenceEntryNode
	FootComment       *CommentGroupNode
}

// Replace replace value node.
func (n *SequenceNode) Replace(idx int, value Node) error { _ = "STUB: not implemented"; return nil }

// Merge merge sequence value.
func (n *SequenceNode) Merge(target *SequenceNode) { _ = "STUB: not implemented"; return }

// SetIsFlowStyle set value to IsFlowStyle field recursively.
func (n *SequenceNode) SetIsFlowStyle(isFlow bool) { _ = "STUB: not implemented"; return }

// Read implements (io.Reader).Read
func (n *SequenceNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns SequenceType
		nil
}

func (n *SequenceNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *SequenceNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *SequenceNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

func (n *SequenceNode) flowStyleString() string { _ = "STUB: not implemented"; return "" }

func (n *SequenceNode) blockStyleString() string { _ = "STUB: not implemented"; return "" }

// If multi-line string, the space characters for indent have already been added, so delete them.

// this line is \n or white space only

// String sequence to text
func (n *SequenceNode) String() string { _ = "STUB: not implemented"; return "" }

// ArrayRange implements ArrayNode protocol
func (n *SequenceNode) ArrayRange() *ArrayNodeIter { _ = "STUB: not implemented"; return nil }

// MarshalYAML encodes to a YAML text
func (n *SequenceNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SequenceEntryNode is the sequence entry.
type SequenceEntryNode struct {
	*BaseNode
	HeadComment *CommentGroupNode // head comment.
	LineComment *CommentGroupNode // line comment e.g.) - # comment.
	Start       *token.Token      // entry token.
	Value       Node              // value node.
}

// String node to text
func (n *SequenceEntryNode) String() string {
	_ = "STUB: not implemented"
	// TODO

	// GetToken returns token instance
	return ""
}

func (n *SequenceEntryNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// Type returns type of node
	return nil
}

func (n *SequenceEntryNode) Type() NodeType {
	_ = "STUB: not implemented"
	return *

	// AddColumn add column number to child nodes recursively
	new(NodeType)
}

func (n *SequenceEntryNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// SetComment set line comment.
func (n *SequenceEntryNode) SetComment(cm *CommentGroupNode) error {
	_ = "STUB: not implemented"
	return nil
}

// Comment returns comment token instance
func (n *SequenceEntryNode) GetComment() *CommentGroupNode { _ = "STUB: not implemented"; return nil }

// MarshalYAML
func (n *SequenceEntryNode) MarshalYAML() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *SequenceEntryNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// SequenceEntry creates SequenceEntryNode instance.
		nil
}

func SequenceEntry(start *token.Token, value Node, headComment *CommentGroupNode) *SequenceEntryNode {
	_ = "STUB: not implemented"
	return nil
}

// SequenceMergeValue creates SequenceMergeValueNode instance.
func SequenceMergeValue(values ...MapNode) *SequenceMergeValueNode {
	_ = "STUB: not implemented"
	return nil
}

// SequenceMergeValueNode is used to convert the Sequence node specified for the merge key into a MapNode format.
type SequenceMergeValueNode struct {
	values []MapNode
}

// MapRange returns MapNodeIter instance.
func (n *SequenceMergeValueNode) MapRange() *MapNodeIter { _ = "STUB: not implemented"; return nil }

// AnchorNode type of anchor node
type AnchorNode struct {
	*BaseNode
	Start *token.Token
	Name  Node
	Value Node
}

func (n *AnchorNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

func (n *AnchorNode) SetName(name string) error { _ = "STUB: not implemented"; return nil }

// Read implements (io.Reader).Read
func (n *AnchorNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns AnchorType
		nil
}

func (n *AnchorNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *AnchorNode) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (n *AnchorNode) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

// AddColumn add column number to child nodes recursively
func (n *AnchorNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String anchor to text
func (n *AnchorNode) String() string { _ = "STUB: not implemented"; return "" }

// implicit null value.

// MarshalYAML encodes to a YAML text
func (n *AnchorNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *AnchorNode) IsMergeKey() bool { _ = "STUB: not implemented"; return false }

// AliasNode type of alias node
type AliasNode struct {
	*BaseNode
	Start *token.Token
	Value Node
}

func (n *AliasNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

func (n *AliasNode) SetName(name string) error { _ = "STUB: not implemented"; return nil }

// Read implements (io.Reader).Read
func (n *AliasNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns AliasType
		nil
}

func (n *AliasNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *AliasNode) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

func (n *AliasNode) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

// AddColumn add column number to child nodes recursively
func (n *AliasNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String alias to text
func (n *AliasNode) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *AliasNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *AliasNode) IsMergeKey() bool {
	_ = "STUB: not implemented"

	// DirectiveNode type of directive node
	return false
}

type DirectiveNode struct {
	*BaseNode
	// Start is '%' token.
	Start *token.Token
	// Name is directive name e.g.) "YAML" or "TAG".
	Name Node
	// Values is directive values e.g.) "1.2" or "!!" and "tag:clarkevans.com,2002:app/".
	Values []Node
}

// Read implements (io.Reader).Read
func (n *DirectiveNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns DirectiveType
		nil
}

func (n *DirectiveNode) Type() NodeType {
	_ = "STUB: not implemented"
	return *

	// GetToken returns token instance
	new(NodeType)
}

func (n *DirectiveNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *DirectiveNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String directive to text
func (n *DirectiveNode) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *DirectiveNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// TagNode type of tag node
type TagNode struct {
	*BaseNode
	Directive *DirectiveNode
	Start     *token.Token
	Value     Node
}

func (n *TagNode) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

func (n *TagNode) stringWithoutComment() string { _ = "STUB: not implemented"; return "" }

// Read implements (io.Reader).Read
func (n *TagNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns TagType
		nil
}

func (n *TagNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *TagNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *TagNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String tag to text
func (n *TagNode) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *TagNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsMergeKey returns whether it is a MergeKey node.
func (n *TagNode) IsMergeKey() bool { _ = "STUB: not implemented"; return false }

func (n *TagNode) ArrayRange() *ArrayNodeIter { _ = "STUB: not implemented"; return nil }

// CommentNode type of comment node
type CommentNode struct {
	*BaseNode
	Token *token.Token
}

// Read implements (io.Reader).Read
func (n *CommentNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns TagType
		nil
}

func (n *CommentNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *CommentNode) GetToken() *token.Token {
	_ = "STUB: not implemented"

	// AddColumn add column number to child nodes recursively
	return nil
}

func (n *CommentNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String comment to text
func (n *CommentNode) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *CommentNode) MarshalYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// CommentGroupNode type of comment node
type CommentGroupNode struct {
	*BaseNode
	Comments []*CommentNode
}

// Read implements (io.Reader).Read
func (n *CommentGroupNode) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Type returns TagType
		nil
}

func (n *CommentGroupNode) Type() NodeType {
	_ = "STUB: not implemented"

	// GetToken returns token instance
	return *new(NodeType)
}

func (n *CommentGroupNode) GetToken() *token.Token { _ = "STUB: not implemented"; return nil }

// AddColumn add column number to child nodes recursively
func (n *CommentGroupNode) AddColumn(col int) { _ = "STUB: not implemented"; return }

// String comment to text
func (n *CommentGroupNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *CommentGroupNode) StringWithSpace(col int) string { _ = "STUB: not implemented"; return "" }

// MarshalYAML encodes to a YAML text
func (n *CommentGroupNode) MarshalYAML() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Visitor has Visit method that is invokded for each node encountered by Walk.
	// If the result visitor w is not nil, Walk visits each of the children of node with the visitor w,
	// followed by a call of w.Visit(nil).
}

type Visitor interface {
	Visit(Node) Visitor
}

// Walk traverses an AST in depth-first order: It starts by calling v.Visit(node); node must not be nil.
// If the visitor w returned by v.Visit(node) is not nil,
// Walk is invoked recursively with visitor w for each of the non-nil children of node,
// followed by a call of w.Visit(nil).
func Walk(v Visitor, node Node) { _ = "STUB: not implemented"; return }

func walkComment(v Visitor, base *BaseNode) { _ = "STUB: not implemented"; return }

type filterWalker struct {
	typ     NodeType
	results []Node
}

func (v *filterWalker) Visit(n Node) Visitor { _ = "STUB: not implemented"; return *new(Visitor) }

type parentFinder struct {
	target Node
}

func (f *parentFinder) walk(parent, node Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// Parent get parent node from child node.
func Parent(root, child Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// Filter returns a list of nodes that match the given type.
func Filter(typ NodeType, node Node) []Node { _ = "STUB: not implemented"; return nil }

// FilterFile returns a list of nodes that match the given type.
func FilterFile(typ NodeType, file *File) []Node { _ = "STUB: not implemented"; return nil }

type ErrInvalidMergeType struct {
	dst Node
	src Node
}

func (e *ErrInvalidMergeType) Error() string { _ = "STUB: not implemented"; return "" }

// Merge merge document, map, sequence node.
func Merge(dst Node, src Node) error { _ = "STUB: not implemented"; return nil }
