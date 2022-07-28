package dao

import "time"

type (
	MasterPickingSlip struct {
		ID                 uint64     `json:"id" gorm:"primary_key;column:id_transact"`
		PickSlipNo         string     `json:"pickSlipNo" gorm:"not null"`
		BatchNo            string     `json:"batchNo" gorm:"not null"`
		FormulaDescription string     `json:"formulaDescription" gorm:"not null"`
		ItemCode           string     `json:"itemCode" gorm:"not null"`
		ItemDescription    string     `json:"itemDescription" gorm:"not null"`
		StockLocator       string     `json:"stockLocator" gorm:"not null"`
		ActualQty          *float64   `json:"actualQty" gorm:"null"`
		PrimaryQty         float64    `json:"primaryQty" gorm:"not null"`
		PrimaryUomCode     string     `json:"primaryUomCode" gorm:"not null"`
		SecondaryQty       float64    `json:"secondaryQty" gorm:"not null"`
		SecondaryUomCode   string     `json:"secondaryUomCode" gorm:"not null"`
		Subinventory       string     `json:"subinventory" gorm:"not null"`
		ToSubinventory     string     `json:"toSubinventory" gorm:"not null"`
		LotNo              string     `json:"lotNo" gorm:"not null"`
		SupplierLotNo      string     `json:"supplierLotNo" gorm:"not null"`
		QtyBlend           float64    `json:"qtyBlend" gorm:"not null"`
		DtlUm              string     `json:"dtlUm" gorm:"not null"`
		Description        string     `json:"description" gorm:"not null"`
		DateSync           *time.Time `json:"dateSync" gorm:"null"`
		TimeSync           *string    `json:"timeSync" gorm:"null"`
		DateSubmit         *time.Time `json:"dateSubmit" gorm:"null"`
		TimeSubmit         *string    `json:"timeSubmit" gorm:"null"`
		OperatorSubmit     *string    `json:"operatorSubmit" gorm:"null"`
		DatePick           *time.Time `json:"datePick" gorm:"null"`
		TimePick           *string    `json:"timePick" gorm:"null"`
		OperatorPick       *string    `json:"operatorPick" gorm:"null"`
		WeightPack         *float64   `json:"weightPack" gorm:"null"`
		TotalBag           uint64     `json:"totalBag" gorm:"not null"`
		ActualBag          *uint64    `json:"actualBag" gorm:"null"`
		IsActive           bool       `json:"isActive" gorm:"not null"`
		IsReadyForPick     bool       `json:"isReadyForPick" gorm:"not null"`
		IsCompletePick     bool       `json:"isCompletePick" gorm:"not null"`
		IsManual           bool       `json:"isManual" gorm:"not null"`
		IsEdit             bool       `json:"isEdit" gorm:"not null"`
		IsCancel           bool       `json:"isCancel" gorm:"not null"`
	}
)

func (MasterPickingSlip) TableName() string {
	return "master_picking_slip"
}

func (MasterPickingSlip) ModelName() string {
	return "MasterPickingSlip"
}
