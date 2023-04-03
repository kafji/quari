package quari

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
)

func MustValue[T any](ctx context.Context, key any) T {
	v, ok := ctx.Value(key).(T)
	if !ok {
		msg := fmt.Sprintf("failed to get value from context with key `%v`", key)
		slog.Error(msg)
		panic(msg)
	}
	return v
}
