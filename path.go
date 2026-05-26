package yaml

import (
	"fmt"
	"io"

	"github.com/goccy/go-yaml/ast"
)

// PathString create Path from string
//
// YAMLPath rule
// $     : the root object/element
// .     : child operator
// ..    : recursive descent
// [num] : object/element of array by number
// [*]   : all objects/elements for array.
//
// If you want to use reserved characters such as `.` and `*` as a key name,
// enclose them in single quotation as follows ( $.foo.'bar.baz-*'.hoge ).
// If you want to use a single quote with reserved characters, escape it with `\` ( $.foo.'bar.baz\'s value'.hoge ).
func PathString(s string) (*Path, error) { _ = "STUB: not implemented"; return nil, nil }

func parsePathRecursive(b *PathBuilder, buf []rune, cursor int) (*PathBuilder, []rune, int, error) {
	_ = "STUB: not implemented"
	return nil, nil,

		// skip .. characters
		0, nil
}

func parsePathDot(b *PathBuilder, buf []rune, cursor int) (*PathBuilder, []rune, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// skip . character

// if started single quote, looking for end single quote char

func parseQuotedKey(b *PathBuilder, buf []rune, cursor int) (*PathBuilder, []rune, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// skip single quote

func parsePathIndex(b *PathBuilder, buf []rune, cursor int) (*PathBuilder, []rune, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// skip '[' character

// Path represent YAMLPath ( like a JSONPath ).
type Path struct {
	node pathNode
}

// String path to text.
func (p *Path) String() string { _ = "STUB: not implemented"; return "" }

// Read decode from r and set extracted value by YAMLPath to v.
func (p *Path) Read(r io.Reader, v interface{}) error { _ = "STUB: not implemented"; return nil }

// ReadNode create AST from r and extract node by YAMLPath.
func (p *Path) ReadNode(r io.Reader) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// Filter filter from target by YAMLPath and set it to v.
func (p *Path) Filter(target, v interface{}) error { _ = "STUB: not implemented"; return nil }

// FilterFile filter from ast.File by YAMLPath.
func (p *Path) FilterFile(f *ast.File) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// For simplicity, directives cannot be the target of operations

// FilterNode filter from node by YAMLPath.
func (p *Path) FilterNode(node ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// MergeFromReader merge YAML text into ast.File.
func (p *Path) MergeFromReader(dst *ast.File, src io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// MergeFromFile merge ast.File into ast.File.
func (p *Path) MergeFromFile(dst *ast.File, src *ast.File) error {
	_ = "STUB: not implemented"
	return nil
}

// MergeFromNode merge ast.Node into ast.File.
func (p *Path) MergeFromNode(dst *ast.File, src ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceWithReader replace ast.File with io.Reader.
func (p *Path) ReplaceWithReader(dst *ast.File, src io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceWithFile replace ast.File with ast.File.
func (p *Path) ReplaceWithFile(dst *ast.File, src *ast.File) error {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceNode replace ast.File with ast.Node.
func (p *Path) ReplaceWithNode(dst *ast.File, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// For simplicity, directives cannot be the target of operations

// AnnotateSource add annotation to passed source ( see section 5.1 in README.md ).
func (p *Path) AnnotateSource(source []byte, colored bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PathBuilder represent builder for YAMLPath.
type PathBuilder struct {
	root *rootNode
	node pathNode
}

// Root add '$' to current path.
func (b *PathBuilder) Root() *PathBuilder { _ = "STUB: not implemented"; return nil }

// IndexAll add '[*]' to current path.
func (b *PathBuilder) IndexAll() *PathBuilder { _ = "STUB: not implemented"; return nil }

// Recursive add '..selector' to current path.
func (b *PathBuilder) Recursive(selector string) *PathBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PathBuilder) containsReservedPathCharacters(path string) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *PathBuilder) enclosedSingleQuote(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *PathBuilder) normalizeSelectorName(name string) string {
	_ = "STUB: not implemented"
	return ""
}

// already escaped name

func (b *PathBuilder) child(name string) *PathBuilder { _ = "STUB: not implemented"; return nil }

// Child add '.name' to current path.
func (b *PathBuilder) Child(name string) *PathBuilder { _ = "STUB: not implemented"; return nil }

// Index add '[idx]' to current path.
func (b *PathBuilder) Index(idx uint) *PathBuilder { _ = "STUB: not implemented"; return nil }

// Build build YAMLPath.
func (b *PathBuilder) Build() *Path { _ = "STUB: not implemented"; return nil }

type pathNode interface {
	fmt.Stringer
	chain(pathNode) pathNode
	filter(ast.Node) (ast.Node, error)
	replace(ast.Node, ast.Node) error
}

type basePathNode struct {
	child pathNode
}

func (n *basePathNode) chain(node pathNode) pathNode {
	_ = "STUB: not implemented"
	return *new(pathNode)
}

type rootNode struct {
	*basePathNode
}

func newRootNode() *rootNode { _ = "STUB: not implemented"; return nil }

func (n *rootNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *rootNode) filter(node ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (n *rootNode) replace(node ast.Node, target ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

type selectorNode struct {
	*basePathNode
	selector string
}

func newSelectorNode(selector string) *selectorNode { _ = "STUB: not implemented"; return nil }

func (n *selectorNode) filter(node ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (n *selectorNode) replaceMapValue(value *ast.MappingValueNode, target ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *selectorNode) replace(node ast.Node, target ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *selectorNode) String() string { _ = "STUB: not implemented"; return "" }

type indexNode struct {
	*basePathNode
	selector uint
}

func newIndexNode(selector uint) *indexNode { _ = "STUB: not implemented"; return nil }

func (n *indexNode) filter(node ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (n *indexNode) replace(node ast.Node, target ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *indexNode) String() string { _ = "STUB: not implemented"; return "" }

type indexAllNode struct {
	*basePathNode
}

func newIndexAllNode() *indexAllNode { _ = "STUB: not implemented"; return nil }

func (n *indexAllNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *indexAllNode) filter(node ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (n *indexAllNode) replace(node ast.Node, target ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

type recursiveNode struct {
	*basePathNode
	selector string
}

func newRecursiveNode(selector string) *recursiveNode { _ = "STUB: not implemented"; return nil }

func (n *recursiveNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *recursiveNode) filterNode(node ast.Node) (*ast.SequenceNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *recursiveNode) filter(node ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (n *recursiveNode) replaceNode(node ast.Node, target ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *recursiveNode) replace(node ast.Node, target ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}
