package dao

import "time"

type (
	NfcReader struct {
		ID         uint64    `json:"id" gorm:"primary_key;column:intReaderID"`
		NFCTagNo   string    `json:"nfcTagNo" gorm:"null;column:txtTid"`
		AntennaNo  uint64    `json:"antennaNo" gorm:"null;column:intAntenna"`
		IsComplete []uint8   `json:"isComplete" gorm:"null;column:bitActive"`
		TrxDate    time.Time `json:"trxDate" gorm:"null;column:dtmTransactionDate"`
		CreatedBy  string    `json:"createdBy" gorm:"null;column:txtInsertedBy"`
		CreatedAt  time.Time `json:"createdAt" gorm:"null;column:dtmInsertedDate"`
	}
)

func (NfcReader) TableName() string {
	return "nfc_reader"
}

func (NfcReader) ModelName() string {
	return "NfcReader"
}
