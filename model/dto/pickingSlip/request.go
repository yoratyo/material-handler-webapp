package pickingSlip

type (
	GetPendingPickingRequestDTO struct {
		BatchNo       *string `form:"batchNo" binding:"omitempty"`
		ItemCode      *string `form:"itemCode" binding:"omitempty"`
		SupplierLotNo *string `form:"supplierLotNo" binding:"omitempty"`
	}

	PatchCompletePickRequestDTO struct {
		ID           uint64  `form:"idPickingSlip" binding:"required,gt=0"`
		ActualBag    uint64  `form:"actualBag" binding:"gte=0"`
		BatchNo      *string `form:"batchNo" binding:"omitempty"`
		ItemCode     *string `form:"itemCode" binding:"omitempty"`
		OperatorPick string
	}
)
