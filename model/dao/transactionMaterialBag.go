package dao

type (
	TransactionMaterialBag struct {
		ID                    uint64   `json:"id" gorm:"primary_key;column:id_transact"`
		IDTransactionWeighing uint64   `json:"idTransactionWeighing" gorm:"not null"`
		IDTransactionNFC      uint64   `json:"idTransactionNFC" gorm:"not null"`
		NFCTagNo              string   `json:"nfcTagNo" gorm:"not null"`
		UsedQty               *float64 `json:"usedQty" gorm:"null"`
		WeightPack            float64  `json:"weightPack" gorm:"not null"`
	}
)

func (TransactionMaterialBag) TableName() string {
	return "transaction_material_bag"
}

func (TransactionMaterialBag) ModelName() string {
	return "TransactionMaterialBag"
}
