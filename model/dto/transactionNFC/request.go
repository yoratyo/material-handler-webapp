package transactionNFC

type (
	GetRegisterNFCRequestDTO struct {
		BatchNo       *string `form:"batchNo" binding:"omitempty"`
		ItemCode      *string `form:"itemCode" binding:"omitempty"`
		SupplierLotNo *string `form:"supplierLotNo" binding:"omitempty"`
	}
	PatchRegisterNFCRequestDTO struct {
		ID               *uint64 `form:"idTransaction" binding:"omitempty"`
		BatchNo          *string `form:"batchNo" binding:"omitempty"`
		SupplierLotNo    *string `form:"supplierLotNo" binding:"omitempty"`
		OperatorRegister string
		DataNFC          string `form:"dataNFC" binding:"omitempty"`
	}
	PatchGatewayCheckRequestDTO struct {
		Items []DetailPatchGatewayCheck `json:"items" binding="required"`
	}
	DetailPatchGatewayCheck struct {
		TagNFC string `json:"nfc_id" binding="required"`
	}
)

func (g PatchGatewayCheckRequestDTO) ToUpdateGatewayCheck() []string {
	var ids []string

	for _, i := range g.Items {
		ids = append(ids, i.TagNFC)
	}

	return ids
}
