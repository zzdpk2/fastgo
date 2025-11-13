package store

import "golang.org/x/net/context"

type PostStore interface {
	Create(ctx context.Context)
}
