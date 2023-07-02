package models

// 用户登录数据
type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `json:"ip" gorm:"size:20"`
	NickName  string    `json:"nick_name" gorm:"size:42"`
	Token     string    `json:"token" gorm:"size:256"`
	Device    string    `json:"device" gorm:"size:256"`
	Addr      string    `json:"addr" gorm:"size:64"`
}
