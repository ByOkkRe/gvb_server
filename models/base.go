package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"` //主键
	CreatedAt time.Time `json:"created_at"`            //创建时间
	UpdatedAt time.Time `json:"-"`                     //更新时间
}
