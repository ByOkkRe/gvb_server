package models

import "gvb_server/models/ctype"

type UserModel struct {
	MODEL
	NickName       string           `json:"nick_name" gorm:"size:36"`                                                         //昵称
	UserName       string           `json:"user_name" gorm:"size:36"`                                                         //用户名
	Password       string           `json:"password" gorm:"size:128"`                                                         //密码
	Avatar         string           `json:"avatar" gorm:"size:128"`                                                           //头像id
	Email          string           `json:"email" gorm:"size:128"`                                                            //邮箱
	Tel            string           `json:"tel" gorm:"size:18"`                                                               //电话
	Addr           string           `json:"addr" gorm:"size:64"`                                                              //地址
	Token          string           `json:"token" gorm:"size:64"`                                                             //第三方平台唯一id
	IP             string           `json:"ip" gorm:"size:20"`                                                                //IP地址
	Role           ctype.Role       `json:"role" gorm:"size:4,default:1"`                                                     //角色
	SignStatus     ctype.SignStatus `json:"sign_status" gorm:"size:4,default:1"`                                              //注册来源
	ArticleModels  []ArticleModel   `json:"-" gorm:"foreignkey:ArticleId"`                                                    //发布文章列表
	CollectsModels []ArticleModel   `json:"-" gorm:"many2many:auth2_collects;joinForeignkey:AuthId;JoinReferences:ArticleId"` //收藏文章列表
}
