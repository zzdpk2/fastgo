package model

import "gorm.io/gorm"

// AfterCreate is to generate postID after database record initiated
func (m *PostM) AfterCreate(tx *gorm.DB) error {
	m.PostID = rid.PostID.New(uint(m.ID))
	return tx.Save(m).Error
}

// AfterCreate is to generate userID after database record initiated
func (m *UserM) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint(m.ID))
	return tx.Save(m).Error
}
