package model

import (
	"github.com/onexstack/fastgo/internal/pkg/rid"
	"gorm.io/gorm"
)

// AfterCreate is to generate postID after database record initiated
func (m *Post) AfterCreate(tx *gorm.DB) error {
	m.PostID = rid.PostID.New(uint64(m.ID))
	return tx.Save(m).Error
}

// AfterCreate is to generate userID after database record initiated
func (m *User) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))
	return tx.Save(m).Error
}
