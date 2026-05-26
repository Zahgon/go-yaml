package yaml

import (
	"reflect"
)

const (
	// StructTagName tag keyword for Marshal/Unmarshal
	StructTagName = "yaml"
)

// StructField information for each the field in structure
type StructField struct {
	FieldName    string
	RenderName   string
	AnchorName   string
	AliasName    string
	IsAutoAnchor bool
	IsAutoAlias  bool
	IsOmitEmpty  bool
	IsOmitZero   bool
	IsFlow       bool
	IsInline     bool
}

func getTag(field reflect.StructField) string {
	_ = "STUB: not implemented"
	// If struct tag `yaml` exist, use that. If no `yaml`
	// exists, but `json` does, use that and try the best to
	// adhere to its rules
	return ""
}

func structField(field reflect.StructField) *StructField { _ = "STUB: not implemented"; return nil }

func isIgnoredStructField(field reflect.StructField) bool { _ = "STUB: not implemented"; return false }

// private field

type StructFieldMap map[string]*StructField

func (m StructFieldMap) isIncludedRenderName(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (m StructFieldMap) hasMergeProperty() bool { _ = "STUB: not implemented"; return false }

func structFieldMap(structType reflect.Type) (StructFieldMap, error) {
	_ = "STUB: not implemented"
	return *new(StructFieldMap), nil
}
