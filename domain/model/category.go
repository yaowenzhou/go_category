package model

type Category struct {
	ID int64 `gorm:"primaryKey;not null;autoIncrement:false"`
}
