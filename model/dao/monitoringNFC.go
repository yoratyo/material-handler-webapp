package dao

import "time"

type (
	MonitoringNFC struct {
		ID        uint64    `json:"id" gorm:"primary_key"`
		IpAddress *string   `json:"ipAddress" gorm:"null"`
		NfcData   *string   `json:"nfcData" gorm:"null"`
		IsDefault bool      `json:"isDefault" gorm:"not null"`
		IsActive  bool      `json:"isActive" gorm:"not null"`
		CreatedAt time.Time `json:"createdAt" gorm:"not null"`
		UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
	}
)

func (MonitoringNFC) TableName() string {
	return "monitoring_nfc"
}

func (MonitoringNFC) ModelName() string {
	return "MonitoringNFC"
}
