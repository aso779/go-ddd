package data

import "context"

type Counter interface {
	Count(
		ctx context.Context,
		spec Specifier,
	) (int, error)
}
