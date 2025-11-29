package conversion

import (
	"github.com/onexstack/fastgo/internal/apiserver/model"
	apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
	"github.com/onexstack/onexstack/pkg/core"
)

// PostodelToPostV1 Converts the model layer Post object to the Protobuf layer Post (v1 object).
func PostodelToPostV1(postModel *model.Post) *apiv1.Post {
	var protoPost apiv1.Post
	_ = core.CopyWithConverters(&protoPost, postModel)
	return &protoPost
}

// PostV1ToPostodel Converts the Protobuf layer Post (v1 blog object) to the model layer Post (blog model object).
func PostV1ToPostodel(protoPost *apiv1.Post) *model.Post {
	var postModel model.Post
	_ = core.CopyWithConverters(&postModel, protoPost)
	return &postModel
}
