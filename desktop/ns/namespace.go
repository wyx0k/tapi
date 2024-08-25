package ns

import (
	"context"
	"errors"
)

type namespaceKey struct{}

var (
	ErrNamespaceNotFound = errors.New("namespace not found")
)

func Namespace(ctx context.Context) string {
	ns, ok := ctx.Value(namespaceKey{}).(string)
	if !ok {
		return ""
	}
	return ns
}

func ChangeNamespace(ctx context.Context, ns string) context.Context {
	return context.WithValue(ctx, namespaceKey{}, ns)
}

func NamespaceMust(ctx context.Context) (string, error) {
	ns := Namespace(ctx)
	if ns == "" {
		return "", ErrNamespaceNotFound
	}
	return ns, nil
}
