package transactionNFC

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
)

type Service interface {
	BulkCreate(ctx *gin.Context, pickingSlip dao.MasterPickingSlip, actualBag uint64) ([]dao.TransactionNFC, error)
	GetByID(ctx *gin.Context, ID uint64) (dao.TransactionNFC, error)
	GetListRegisterNFC(ctx *gin.Context, req transactionNFCDTO.GetRegisterNFCRequestDTO) (uint64, []dao.TransactionNFC, error)
	SetTickerMonitoringNFC(ctx *gin.Context, conn *websocket.Conn) error
	GetPendingOKP(ctx *gin.Context, status string) ([]string, error)
	GetBySupplierLotNo(ctx *gin.Context, supplierLotNo string) ([]dao.TransactionNFC, error)
	GetRegisterNFC(ctx *gin.Context, request transactionNFCDTO.PatchRegisterNFCRequestDTO) (dao.TransactionNFC, error)
	GetMonitoringNFC(ctx *gin.Context) (dao.MonitoringNFC, error)
	PatchCompleteRegister(ctx *gin.Context, ID uint64, data transactionNFCDTO.PatchRegisterNFCRequestDTO) error
	BulkPatchGatewayCheck(ctx *gin.Context, request transactionNFCDTO.PatchGatewayCheckRequestDTO) error
	GetMonitoringGateway(ctx *gin.Context, conn *websocket.Conn) error
	GetTodayRegisterMonitoring(ctx *gin.Context) (dao.RegisterNFCMonitoring, error)
	GetTodayGatewayMonitoring(ctx *gin.Context) (dao.GatewayMonitoring, error)
}
