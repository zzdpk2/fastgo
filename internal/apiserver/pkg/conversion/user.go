package conversion

import (
	"github.com/onexstack/fastgo/internal/apiserver/model"
	apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
	"github.com/onexstack/onexstack/pkg/core"
)

// UserodelToUserV1 Converts the model layer User object to the Protobuf layer User (v1 object).
func UserodelToUserV1(userModel *model.User) *apiv1.User {
	var protoUser apiv1.User
	_ = core.CopyWithConverters(&protoUser, userModel)
	return &protoUser
}

// UserV1ToUserodel Converts the Protobuf layer User (v1 user object) to the model layer User (user model object).
func UserV1ToUserodel(protoUser *apiv1.User) *model.User {
	var userModel model.User
	_ = core.CopyWithConverters(&userModel, protoUser)
	return &userModel
}
