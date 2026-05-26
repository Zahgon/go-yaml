package yaml

import (
	"context"
)

type (
	ctxMergeKey  struct{}
	ctxAnchorKey struct{}
)

func withMerge(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func isMerge(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func withAnchor(ctx context.Context, name string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getAnchorMap(ctx context.Context) map[string]struct{} { _ = "STUB: not implemented"; return nil }
