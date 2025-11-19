package store

import (
	"context"
	"errors"
	"github.com/onexstack/fastgo/internal/apiserver/model"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"
	"log/slog"
)

type PostStore interface {
	Create(ctx context.Context, obj *model.Post) error
	Update(ctx context.Context, obj *model.Post) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.Post, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.Post, error)
	PostExpansion
}

// PostExpansion defines the ops on tips
type PostExpansion interface{}

type postStore struct {
	store *datastore
}

// var _ PostStore = (*postStore)(nil)

func newPostStore(store *datastore) *postStore {
	return &postStore{store}
}

// Create is to create a record of a tip
func (s *postStore) Create(ctx context.Context, obj *model.Post) error {
	if err := s.store.DB(ctx).Save(obj).Error; err != nil {
		slog.Error("Failed to update post in database", "err", err, "post", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

// Delete is to delete a tip from a record
func (s *postStore) Delete(ctx context.Context, opts *where.Options) error {
	err := s.store.DB(ctx, opts).Delete(new(model.Post)).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to delete post from database", "err", err, "conditions", opts)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *postStore) Get(ctx context.Context, opts *where.Options) (*model.Post, error) {
	var obj model.Post
	if err := s.store.DB(ctx, opts).First(&obj).Error; err != nil {
		slog.Error("Failed to retrieve post from database", "err", err, "conditions", opts)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.ErrPostNotFound
		}
		return nil, errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return &obj, nil
}

// List is to return the list of tips and its total amount
// nolint: nonamedreturns
func (s *postStore) List(ctx context.Context, opts *where.Options) (count int64, ret []*model.Post, err error) {
	err = s.store.DB(ctx, opts).Order("id desc").Find(&ret).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		slog.Error("Failed to list posts from database", "err", err, "conditions", opts)
		err = errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return
}
