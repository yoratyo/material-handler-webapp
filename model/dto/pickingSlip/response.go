package pickingSlip

type (
	GetPendingPickingResponseDTO struct {
		ID                 uint64 `json:"id"`
		PickSlipNo         string `json:"pickSlipNo"`
		BatchNo            string `json:"batchNo"`
		FormulaDescription string `json:"formulaDescription"`
		ItemCode           string `json:"itemCode"`
		ItemDescription    string `json:"itemDescription"`
		StockLocator       string `json:"stockLocator"`
		SupplierLotNo      string `json:"supplierLotNo"`
		Description        string `json:"description"`
		TotalBag           uint64 `json:"totalBag"`
		IsManual           bool   `json:"isManual"`
	}

	PatchCompletePickResponseDTO struct {
		Message string `json:"message"`
	}
)
