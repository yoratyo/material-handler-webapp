package dao

import "time"

type (
	UserLogin struct {
		IDUser       uint64     `json:"idUser" gorm:"primary_key"`
		Username     string     `json:"username" gorm:"not null"`
		Password     string     `json:"password" gorm:"not null"`
		UserLevel    string     `json:"userLevel" gorm:"not null"`
		IsAssign     bool       `json:"isAssign" gorm:"not null"`
		IsActive     bool       `json:"isActive" gorm:"not null"`
		CreatedDate  time.Time  `json:"createdDate" gorm:"not null"`
		ModifiedDate *time.Time `json:"modifiedDate" gorm:"null"`
		CreatedBy    string     `json:"createdBy" gorm:"not null"`
		ModifiedBy   *string    `json:"modifiedBy" gorm:"null"`
	}
)

func (UserLogin) TableName() string {
	return "user_login"
}

func (UserLogin) ModelName() string {
	return "UserLogin"
}
