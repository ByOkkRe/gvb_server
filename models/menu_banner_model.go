package models

// 菜单banner表
type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignkey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignkey:BannerID"`
	Sort        int         `json:"sort" gorm:"size:10"`
}
