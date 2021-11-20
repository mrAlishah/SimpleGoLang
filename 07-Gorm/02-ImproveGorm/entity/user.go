package entity

type User struct {
	//gorm.Model
	ID        uint64 `json:"id"  gorm:"primaryKey,autoIncrement"`
	UserName  string `json:"username" gorm:"type:varchar(100)"`
	FirstName string `json:"firstname" gorm:"type:varchar(100)"`
	LastName  string `json:"lastname" gorm:"type:varchar(100)"`
}
