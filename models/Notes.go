package models

import (
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	Title      string `gorm:"size:255" json:"title"`
	Content    string `gorm:"size:600" json:"content"`
	IsRemoved  bool   `gorm:"default:false" json:"is_removed"`
	IsPinned   bool   `gorm:"default:false" json:"is_pinned"`
	IsArchived bool   `gorm:"default:false" json:"is_archived"`
	CreatedBy  User   `gorm:"foreignKey:UserRefer;association_autoupdate:false;association_autocreate:false"`
	UserRefer  uint
}
