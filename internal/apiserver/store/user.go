package store

import (
	"errors"
	"github.com/onexstack/fastgo/internal/apiserver/model"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log/slog"
)

type UserStore interface {
	Create(ctx context.Context, obj *model.User) error
	Update(ctx context.Context, obj *model.User) error
	Delete(ctx context.Context, obj *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.User, error)
	List(ctx context.Context, opts *where.Options) ([]*model.User, error)

	UserExpansion
}

type UserExpansion interface{}

type userStore struct {
	store *datastore
}

//var _ UserStore = (*userStore)(nil)

func newUserStore(store *datastore) *userStore {
	return &userStore{store}
}

func (s *userStore) Create(ctx context.Context, obj *model.User) error {
	if err := s.store.DB(ctx).Create(&obj).Error; err != nil {
		slog.Error("Failed to insert user into database", "err", err, "user", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *userStore) Update(ctx context.Context, obj *model.User) error {
	if err := s.store.DB(ctx).Save(obj).Error; err != nil {
		slog.Error("Failed to update user in database", "err", err, "user", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *userStore) Delete(ctx context.Context, opts *where.Options) error {
	err := s.store.DB(ctx).Delete(new(model.User)).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to delete user from database", "err", err, "conditions", opts)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

func (s *userStore) Get(ctx context.Context, opts *where.Options) (*model.User, error) {
	var obj model.User
	if err := s.store.DB(ctx, opts).First(&obj).Error; err != nil {
		slog.Error("Failed to retrieve user from database", "err", err, "conditions", opts)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.ErrUserNotFound
		}
		return nil, errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return &obj, nil
}

// nolint: nonamedreturns
func (s *userStore) List(ctx context.Context, opts *where.Options) (count int64, ret []*model.User, err error) {
	err = s.store.DB(ctx, opts).Order("id desc").Find(&ret).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		slog.Error("Failed to list users from database", "err", err, "conditions", opts)
		err = errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return
}
