package entity

type Video struct {
	ID          uint64 `json:"id" bson:"-"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Url         string `json:"url" bson:"description" binding:"required"`
}
