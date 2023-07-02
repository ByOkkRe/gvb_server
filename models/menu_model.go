package models

import "gvb_server/models/ctype"

//菜单表
type MenuModel struct {
	MODEL
	MenuTitle    string        `json:"menu_title" gorm:"size:32"`
	MenuTitleEn  string        `json:"menu_title_en" gorm:"size:32"`
	Slogan       string        `json:"slogan" gorm:"size:64"`
	Abstract     ctype.Array   `json:"abstract" gorm:"type:string"`
	AbstractTime int           `json:"abstract_time"`
	Banners      []BannerModel `json:"banners" gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferer:ImageID"`
	BannerTime   int           `json:"banner_time"`
	Sort         int           `json:"sort" gorm:"size:10"`
}
