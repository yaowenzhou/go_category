package model

type User struct {
	ID int64 `gorm:"primaryKey;not null;autoIncrement:false"`
}
