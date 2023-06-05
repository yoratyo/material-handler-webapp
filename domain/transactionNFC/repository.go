package transactionNFC

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
)

type Repository interface {
	BulkCreate(ctx *gin.Context, req []dao.TransactionNFC) ([]dao.TransactionNFC, error)
	GetByID(ctx *gin.Context, ID uint64) (dao.TransactionNFC, error)
	CheckExistByIDPickingSlip(ctx *gin.Context, IDPickingSlip uint64) (bool, error)
	GetListRegisterNFC(ctx *gin.Context, req transactionNFCDTO.GetRegisterNFCRequestDTO) ([]dao.TransactionNFC, error)
	GetCountPendingToRegister(ctx *gin.Context) (uint64, error)
	GetMonitoringNFCByIP(ctx *gin.Context, ipAddress string) (dao.MonitoringNFC, error)
	GetDefaultMonitoringNFC(ctx *gin.Context) (dao.MonitoringNFC, error)
	GetDistinctOKP(ctx *gin.Context, status string) ([]string, error)
	GetBySupplierLotNoGroupByBatchNo(ctx *gin.Context, supplierLotNo string) ([]dao.TransactionNFC, error)
	GetRegisterNFC(ctx *gin.Context, request transactionNFCDTO.PatchRegisterNFCRequestDTO) (dao.TransactionNFC, error)
	GetByNFCTagNo(ctx *gin.Context, nfcTagNo string) ([]dao.TransactionNFC, error)
	PatchByID(ctx *gin.Context, ID uint64, data map[string]interface{}) error
	PatchGatewayCheck(ctx *gin.Context, Ids []string, data map[string]interface{}) (int64, error)
	GetCountGatewayToday(ctx *gin.Context) (dao.CountGatewayTransactionNFC, error)
	GetCountGatewayTodayPerOKP(ctx *gin.Context) ([]dao.CountGatewayPerOKP, error)
	GetRegisterMonitoringToday(ctx *gin.Context) (dao.RegisterNFCMonitoring, error)
	GetListUnreadNFC(ctx *gin.Context) ([]dao.NfcReader, error)
	PatchReadNFR(ctx *gin.Context, tags []string, data map[string]interface{}) (int64, error)
}
