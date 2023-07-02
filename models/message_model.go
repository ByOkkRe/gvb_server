package models

//聊天室消息Model
type MessageModel struct {
	MODEL
	SendUserID       uint      `json:"send_user_id" gorm:"primarykey"` //发送人id
	SendUserModel    UserModel `json:"-" gorm:"foreignkey:SendUserID"`
	SendUserNickName string    `json:"send_user_nick_name" gorm:"size:42"`
	RevUserID        uint      `json:"rev_user_id" gorm:"primarykey"` //接收人id
	RevUserModel     UserModel `json:"-" gorm:"foreignkey:RevUserID"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	IsRead           bool      `json:"is_read" gorm:"default:false"` //接收方是否查看
	Content          string    `json:"content"`                      //消息内容
}
