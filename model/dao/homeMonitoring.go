package dao

type (
	PickingMonitoring struct {
		ProgressCount uint64
		TotalCount    uint64
		Percentage    float64
	}
	RegisterNFCMonitoring struct {
		ProgressCount uint64
		TotalCount    uint64
		Percentage    float64
	}
	GatewayMonitoring struct {
		ProgressCount uint64
		TotalCount    uint64
		Percentage    float64
	}
)
