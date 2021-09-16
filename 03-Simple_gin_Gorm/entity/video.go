package entity

type Video struct {
	ID          uint64 `json:"id"  gorm:"primaryKey,autoIncrement"`
	Title       string `json:"title" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:varchar(100)"`
	Url         string `json:"url" binding:"required" gorm:"type:varchar(100)"`
}
