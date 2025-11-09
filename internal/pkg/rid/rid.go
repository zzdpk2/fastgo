package rid

import (
	"github.com/onexstack/onexstack/pkg/id"
)

const defaultABC = "abcdefghijklmnopqrstuvwxyz1234567890"

type ResourceID string

const (
	UserID ResourceID = "user"
	PostID ResourceID = "post"
)

// Convert rid to string
func (rid ResourceID) String() string {
	return string(rid)
}

// New Create the unique ID of prefix
func (rid ResourceID) New(counter uint64) string {
	uniqueStr := id.NewCode(
		counter,
		id.WithCodeChars([]rune(defaultABC)),
		id.WithCodeL(6),
		id.WithCodeSalt(Salt()),
	)
	return rid.String() + "-" + uniqueStr
}
