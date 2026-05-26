package yaml

import (
	"context"
	"io"
	"reflect"
	"time"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/token"
)

// Decoder reads and decodes YAML values from an input stream.
type Decoder struct {
	reader               io.Reader
	referenceReaders     []io.Reader
	anchorNodeMap        map[string]ast.Node
	anchorValueMap       map[string]reflect.Value
	customUnmarshalerMap map[reflect.Type]func(context.Context, interface{}, []byte) error
	commentMaps          []CommentMap
	toCommentMap         CommentMap
	opts                 []DecodeOption
	referenceFiles       []string
	referenceDirs        []string
	isRecursiveDir       bool
	isResolvedReference  bool
	validator            StructValidator
	disallowUnknownField bool
	allowedFieldPrefixes []string
	allowDuplicateMapKey bool
	useOrderedMap        bool
	useJSONUnmarshaler   bool
	parsedFile           *ast.File
	streamIndex          int
	decodeDepth          int
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader, opts ...DecodeOption) *Decoder { _ = "STUB: not implemented"; return nil }

const maxDecodeDepth = 10000

func (d *Decoder) stepIn() { _ = "STUB: not implemented"; return }

func (d *Decoder) stepOut() { _ = "STUB: not implemented"; return }

func (d *Decoder) isExceededMaxDepth() bool { _ = "STUB: not implemented"; return false }

func (d *Decoder) castToFloat(v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// if error occurred, return zero value

func (d *Decoder) mapKeyNodeToString(ctx context.Context, node ast.MapKeyNode) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *Decoder) setToMapValue(ctx context.Context, node ast.Node, m map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) setToOrderedMapValue(ctx context.Context, node ast.Node, m *MapSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) setPathToCommentMap(node ast.Node) { _ = "STUB: not implemented"; return }

func (d *Decoder) addHeadOrLineCommentToMap(node ast.Node) { _ = "STUB: not implemented"; return }

func (d *Decoder) addSequenceNodeCommentToMap(node *ast.SequenceNode) {
	_ = "STUB: not implemented"
	return
}

func (d *Decoder) addFootCommentToMap(node ast.Node) { _ = "STUB: not implemented"; return }

func (d *Decoder) addCommentToMap(path string, comment *Comment) { _ = "STUB: not implemented"; return }

// already added same comment

func (d *Decoder) nodeToValue(ctx context.Context, node ast.Node) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// To handle the case where alias is processed recursively, the result of alias can be set to nil in advance.

// self recursion.

func (d *Decoder) getMapNode(node ast.Node, isMerge bool) (ast.MapNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.MapNode), nil
}

func (d *Decoder) getArrayNode(node ast.Node) (ast.ArrayNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.ArrayNode), nil
}

func (d *Decoder) convertValue(v reflect.Value, typ reflect.Type, src ast.Node) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Special case for "strings -> floats" aka scientific notation
// If the destination type is a float and the source type is a string, check if we can
// use strconv.ParseFloat to convert the string to a float.

// else, fall through to the error below

// cast value to string

// Handle named types, e.g., `type MyString string`

func (d *Decoder) deleteStructKeys(structType reflect.Type, unknownFields map[string]ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) unmarshalableDocument(node ast.Node) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decoder) unmarshalableText(node ast.Node) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type jsonUnmarshaler interface {
	UnmarshalJSON([]byte) error
}

func (d *Decoder) existsTypeInCustomUnmarshalerMap(t reflect.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *Decoder) unmarshalerFromCustomUnmarshalerMap(t reflect.Type) (func(context.Context, interface{}, []byte) error, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (d *Decoder) canDecodeByUnmarshaler(dst reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *Decoder) decodeByUnmarshaler(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	astNodeType = reflect.TypeOf((*ast.Node)(nil)).Elem()
)

func (d *Decoder) decodeValue(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// set nil value to pointer

// handle scientific notation

// couldn't be parsed as float

// handle scientific notation

// couldn't be parsed as float

func (d *Decoder) createDecodableValue(typ reflect.Type) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func (d *Decoder) castToAssignableValue(value reflect.Value, target reflect.Type, src ast.Node) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (d *Decoder) createDecodedNewValue(
	ctx context.Context, typ reflect.Type, defaultVal reflect.Value, node ast.Node,
) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (d *Decoder) keyToNodeMap(ctx context.Context, node ast.Node, ignoreMergeKey bool, getKeyOrValueNode func(*ast.MapNodeIter) ast.Node) (map[string]ast.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decoder) keyToKeyNodeMap(ctx context.Context, node ast.Node, ignoreMergeKey bool) (map[string]ast.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decoder) keyToValueNodeMap(ctx context.Context, node ast.Node, ignoreMergeKey bool) (map[string]ast.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decoder) setDefaultValueIfConflicted(v reflect.Value, fieldMap StructFieldMap) error {
	_ = "STUB: not implemented"
	return nil
}

// if declared same key name, set default value

// This is a subset of the formats allowed by the regular expression
// defined at http://yaml.org/type/timestamp.html.
var allowedTimestampFormats = []string{
	"2006-1-2T15:4:5.999999999Z07:00", // RCF3339Nano with short date fields.
	"2006-1-2t15:4:5.999999999Z07:00", // RFC3339Nano with short date fields and lower-case "t".
	"2006-1-2 15:4:5.999999999",       // space separated with no time zone
	"2006-1-2",                        // date only
}

func (d *Decoder) castToTime(ctx context.Context, src ast.Node) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// invalid format

func (d *Decoder) decodeTime(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) castToDuration(ctx context.Context, src ast.Node) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (d *Decoder) decodeDuration(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// getMergeAliasName support single alias only
func (d *Decoder) getMergeAliasName(src ast.Node) string { _ = "STUB: not implemented"; return "" }

func (d *Decoder) decodeStruct(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// dst value implements ast.Node

// set nil value to pointer

// set nil value to pointer

// Ignore unknown fields when parsing an inline struct (recognized by a nil token).
// Unknown fields are expected (they could be fields from the parent struct).

// TODO: to make FieldError message cutomizable

// A missing required field will not be in the keyToNodeMap
// the error needs to be associated with the parent of the source node

// getParentMapTokenIfExists if the NodeType is a container type such as MappingType or SequenceType,
// it is necessary to return the parent MapNode's colon token to represent the entire container.
func (d *Decoder) getParentMapTokenIfExistsForValidationError(typ ast.NodeType, tk *token.Token) *token.Token {
	_ = "STUB: not implemented"
	return nil
}

// map:
//   key: value
//      ^ current token ( colon )

// map:
//   - value
//   ^ current token ( sequence entry )

func (d *Decoder) decodeArray(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// set nil value to pointer

func (d *Decoder) decodeSlice(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// set nil value to pointer

func (d *Decoder) decodeMapItem(ctx context.Context, dst *MapItem, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) validateDuplicateKey(keyMap map[string]struct{}, key interface{}, keyNode ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) decodeMapSlice(ctx context.Context, dst *MapSlice, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) decodeMap(ctx context.Context, dst reflect.Value, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// set nil value to pointer

// expect nil key

func (d *Decoder) fileToReader(file string) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (d *Decoder) isYAMLFile(file string) bool { _ = "STUB: not implemented"; return false }

func (d *Decoder) readersUnderDir(dir string) ([]io.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decoder) readersUnderDirRecursive(dir string) ([]io.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decoder) resolveReference(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// assign new anchor definition to anchorMap

func (d *Decoder) parse(ctx context.Context, bytes []byte) (*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// try to decode ast.Node to value and map anchor value to anchorMap

func (d *Decoder) isInitialized() bool { _ = "STUB: not implemented"; return false }

func (d *Decoder) decodeInit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) decode(ctx context.Context, v reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// empty document.

// Decode reads the next YAML-encoded value from its input
// and stores it in the value pointed to by v.
//
// See the documentation for Unmarshal for details about the
// conversion of YAML into a Go value.
func (d *Decoder) Decode(v interface{}) error { _ = "STUB: not implemented"; return nil }

// DecodeContext reads the next YAML-encoded value from its input
// and stores it in the value pointed to by v with context.Context.
func (d *Decoder) DecodeContext(ctx context.Context, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// DecodeFromNode decodes node into the value pointed to by v.
func (d *Decoder) DecodeFromNode(node ast.Node, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// DecodeFromNodeContext decodes node into the value pointed to by v with context.Context.
func (d *Decoder) DecodeFromNodeContext(ctx context.Context, node ast.Node, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// resolve references to the anchor on the same file
