package models

// 广告表
type AdvertModel struct {
	MODEL
	Title  string `json:"title" gorm:"size:32"` //显示的标题
	Href   string `json:"href"`                 //跳转链接
	Images string `json:"images"`               //图片
	IsShow bool   `json:"is_show"`              //广告是否展示
}
