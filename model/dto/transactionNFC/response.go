package transactionNFC

type (
	PatchGatewayCheckResponseDTO struct {
		Message string `json:"message"`
	}
	GetMonitoringGatewayResponseDTO struct {
		Header  HeaderMonitoringGateway  `json:"header"`
		ListOKP []ListOKPMonitoringGateway `json:"listOkp"`
	}
	HeaderMonitoringGateway struct {
		TodayDate       string `json:"todayDate"`
		TotalChecked    uint64 `json:"totalChecked"`
		TotalRegistered uint64 `json:"totalRegistered"`
	}
	ListOKPMonitoringGateway struct {
		BatchNo            string `json:"batchNo" gorm:"not null"`
		FormulaDescription string `json:"formulaDescription" gorm:"not null"`
		TotalChecked       uint64 `json:"totalChecked"`
		TotalRegistered    uint64 `json:"totalRegistered"`
	}
)
