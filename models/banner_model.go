package models

type BannerModel struct {
	MODEL
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //判断是否重复上传图片hash
	Name string `json:"name" gorm:"size:38"` //图片名称
}
