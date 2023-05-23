package models

//聊天室消息Model
type MessageModel struct {
	MODEL
	SendUserID       uint      `json:"send_user_id" gorm:"primarykey"`
	SendUserModel    UserModel `json:"-" gorm:"foreignkey:SendUserID"`
	SendUserNickName string    `json:"send_user_nick_name" gorm:"size:42"`
	RevUserID        uint      `json:"rev_user_id" gorm:"primarykey"`
	RevUserModel     UserModel `json:"-" gorm:"foreignkey:RevUserID"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	IsRead           bool      `json:"is_read" gorm:"default:false"`
	Content          string    `json:"content"`
}
