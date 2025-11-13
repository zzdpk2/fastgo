package store

import (
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	S    *datastore
)

type IStore interface {
	DB(ctx context.Context, wheres ...where.Where) *gorm.DB
	TX(ctx context.Context, fn func(ctx context.Context) error)

	User() UserStore
	Post() PostStore
}

// transactionKey is for saving transaction in context.Context
type transactionKey struct{}

// datastore is the implementation of IStore
type datastore struct {
	core *gorm.DB

	// we can add more DB instances
	// fake *gorm.DB
}

// make sure datastore implement the IStore interface
var _ IStore = (*datastore)(nil)

// NewStore is to create an instance of IStore
func NewStore(db *gorm.DB) *datastore {
	// make sure S can only be initialized once
	once.Do(
		func() {
			S = &datastore{db}
		})
	return S
}

func (store *datastore) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
	db := store.core
	// Extract transaction instance from context
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}

	// Traverse all conditions and apply to queries
	for _, whr := range wheres {
		db = whr.Where(db)
	}
	return db
}

// TX to return a new transaction instance
// nolint: fatcontext
func (store *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return store.core.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// User Users return an implementation of UserStore Interface
func (store *datastore) User() UserStore {
	return newUserStore(store)
}

func (store *datastore) Post() PostStore {
	return newPostStore(store)
}
