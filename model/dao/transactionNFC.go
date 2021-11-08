package dao

import "time"

type (
	TransactionNFC struct {
		ID                  uint64     `json:"id" gorm:"primary_key"`
		IDMasterPickingSlip uint64     `json:"idMasterPickingSlip" gorm:"not null"`
		PickSlipNo          string     `json:"pickSlipNo" gorm:"not null"`
		BatchNo             string     `json:"batchNo" gorm:"not null"`
		ItemCode            string     `json:"itemCode" gorm:"not null"`
		ItemDescription     string     `json:"itemDescription" gorm:"not null"`
		StockLocator        string     `json:"stockLocator" gorm:"not null"`
		LotNo               string     `json:"lotNo" gorm:"not null"`
		SupplierLotNo       string     `json:"supplierLotNo" gorm:"not null"`
		NFCTagNo            *string    `json:"nfcTagNo" gorm:"null"`
		BagNo               uint64     `json:"bagNo" gorm:"not null"`
		TotalBag            uint64     `json:"totalBag" gorm:"not null"`
		WeightPack          float64    `json:"weightPack" gorm:"not null"`
		RemainingWeight     float64    `json:"remainingWeight" gorm:"not null"`
		DateRegister        *time.Time `json:"dateRegister" gorm:"null"`
		TimeRegister        *string    `json:"timeRegister" gorm:"null"`
		OperatorRegister    *string    `json:"operatorRegister" gorm:"null"`
		DateGatewayCheck    *time.Time `json:"dateGatewayCheck" gorm:"null"`
		TimeGatewayCheck    *string    `json:"timeGatewayCheck" gorm:"null"`
		IsRegister          bool       `json:"isRegister" gorm:"not null"`
		IsGatewayCheck      bool       `json:"isGatewayCheck" gorm:"not null"`
		IsCompleteBag       bool       `json:"isCompleteBag" gorm:"not null"`
		IsSplitMaterial     bool       `json:"isSplitMaterial" gorm:"not null"`
	}
)

func (TransactionNFC) TableName() string {
	return "transaction_nfc"
}

func (TransactionNFC) ModelName() string {
	return "TransactionNFC"
}
