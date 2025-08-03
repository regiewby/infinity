package storage

import "time"

type User struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	SafeID      string    `gorm:"column:safe_id"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Email       string    `gorm:"column:email"`
	PIN         string    `gorm:"column:pin"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:CURRENT_TIMESTAMP;autoUpdateTime:CURRENT_TIMESTAMP"`
}

func (a *User) TableName() string {
	return "users"
}
