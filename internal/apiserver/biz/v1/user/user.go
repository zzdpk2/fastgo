package user

import "golang.org/x/net/context"
import apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1\"

//go:generate mockgen -destination mock_user.go -package user github.com/onexstack/fastgo/internal/apiserver/biz/v1/user UserBiz

type UserBiz interface {
	Create(ctx context.Context, rq *apiv1.C)
}
