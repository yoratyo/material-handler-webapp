package transactionNFC

type (
	GetRegisterNFCRequestDTO struct {
		BatchNo       *string `form:"batchNo" binding:"omitempty"`
		SupplierLotNo *string `form:"supplierLotNo" binding:"omitempty"`
	}
	PatchRegisterNFCRequestDTO struct {
		ID               *uint64 `form:"idTransaction" binding:"omitempty"`
		BatchNo          *string `form:"batchNo" binding:"omitempty"`
		SupplierLotNo    *string `form:"supplierLotNo" binding:"omitempty"`
		OperatorRegister string
		DataNFC          string
	}
)
