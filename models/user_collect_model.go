package models

import "time"

// 记录用户什么时候收藏了什么文章
// 用户收藏文章表
type UserCollectModel struct {
	UserID       uint         `gorm:"primarykey"`
	UserModel    UserModel    `gorm:"foreignkey:UserID"`
	ArticleID    uint         `gorm:"primarykey"`
	ArticleModel ArticleModel `gorm:"foreignkey:ArticleID"`
	CreatedAt    time.Time
}
