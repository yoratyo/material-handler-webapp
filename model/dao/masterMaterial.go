package dao

import "time"

type (
	MasterMaterial struct {
		ItemCode        string     `json:"itemCode" gorm:"primary_key"`
		ItemDescription string     `json:"itemDescription" gorm:"not null"`
		ItemType        string     `json:"itemType" gorm:"not null"`
		UomType         string     `json:"uomType" gorm:"not null"`
		WeightTolerance float64    `json:"weightTolerance" gorm:"not null"`
		WeightPack      float64    `json:"weightPack" gorm:"not null"`
		TotalQty        float64    `json:"totalQty" gorm:"not null"`
		CreatedDate     time.Time  `json:"createdDate" gorm:"not null"`
		ModifiedDate    *time.Time `json:"modifiedDate" gorm:"null"`
		CreatedBy       string     `json:"createdBy" gorm:"not null"`
		ModifiedBy      *string    `json:"modifiedBy" gorm:"null"`
	}
)

func (MasterMaterial) TableName() string {
	return "master_material"
}

func (MasterMaterial) ModelName() string {
	return "MasterMaterial"
}
