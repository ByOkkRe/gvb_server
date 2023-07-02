package models

type TagModel struct {
	MODEL
	Title    string         `json:"title" gorm:"size:18"`
	Articles []ArticleModel `json:"-" gorm:"many2many:articles_tag_model"`
}
