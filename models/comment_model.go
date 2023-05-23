package models

type CommentModel struct {
	MODEL
	SubComment         []*CommentModel `json:"sub_comment" gorm:"foreignkey:ParentCommentID"`   //子评论列表
	ParentCommentModel *CommentModel   `json:"comment_model" gorm:"foreignkey:ParentCommentID"` //父级评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               //父级评论id
	Content            string          `json:"content" gorm:"size:256"`                         //评论内容
	DiggCount          int             `json:"digg_count" gorm:"size:10;default:0"`             //评论点赞数
	CommentCount       int             `json:"comment_count" gorm:"size:10;default:0"`          //子评论数
	Article            ArticleModel    `json:"article" gorm:"foreignkey:ArticleID"`             //关联的文章
	ArticleID          uint            `json:"article_id"`                                      // 关联文章id
	User               UserModel       `json:"user"`                                            //关联的用户
	UserID             uint            `json:"user_id"`                                         //用户ID
}
