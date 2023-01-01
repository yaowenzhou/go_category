package model

const UserTableName = "user"

type User struct {
	ID int64 `gorm:"primaryKey;not null;autoIncrement:false"`
}

// TableName 获取表名
func (u *User) TableName() string {
	return UserTableName
}
