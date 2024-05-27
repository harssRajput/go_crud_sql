package transaction

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type OperationType struct {
	OperationTypeID int    `json:"operation_type_id"`
	Description     string `json:"description"`
}

const (
	OperationTypeNormalPurchase      = 1
	OperationTypeInstallmentPurchase = 2
	OperationTypeWithdraw            = 3
	OperationTypeCreditVoucher       = 4
)

var operationTypes = []OperationType{
	{OperationTypeID: OperationTypeNormalPurchase, Description: "Normal Purchase"},
	{OperationTypeID: OperationTypeInstallmentPurchase, Description: "Installment Purchase"},
	{OperationTypeID: OperationTypeWithdraw, Description: "Withdraw"},
	{OperationTypeID: OperationTypeCreditVoucher, Description: "credit voucher"}, // only this operation type is allowed to have positive amount
}

type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"` //event_date contains date and time accurate upto milliseconds in string format "2021-01-01T10:00:00.000Z"
}

type TransactionService interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
}

type transactionService struct {
	sqldb  *sql.DB
	logger *log.Logger
}

func InitTransactionService(sqldb *sql.DB, log *log.Logger) (TransactionService, error) {
	return &transactionService{sqldb: sqldb, logger: log}, nil
}

func (tx *transactionService) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	// check if operation type exists
	if opType, err := getOperationTypeByID(transaction.OperationTypeID); err != nil {
		return nil, err
	} else if opType == nil {
		return nil, errors.New("operation type not found")
	} else if opType.OperationTypeID == OperationTypeCreditVoucher {
		if transaction.Amount <= 0 {
			return nil, errors.New("credit voucher amount should be positive")
		}
	} else if transaction.Amount >= 0 {
		// amount should be negative for all operation types except credit voucher
		return nil, errors.New("amount should be positive")
	}

	eventDateTime, err := time.Parse(time.RFC3339, time.Now().Format("2006-01-02T15:04:05.000Z"))
	if err != nil {
		tx.logger.Printf("Error parsing date: %v\n", err)
		return nil, err
	}
	transaction.EventDate = eventDateTime
	// insert transaction in a db
	query := `INSERT INTO Transaction (account_id, operation_type_id, amount, event_date) VALUES (?, ?, ?, ?)`
	response, err := tx.sqldb.Exec(query, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate)
	if err != nil {
		tx.logger.Printf("Error inserting transaction: %v\n", err)
		return nil, err
	}
	transactionID, err := response.LastInsertId()
	tx.logger.Println("transaction created with transactionID: ", transactionID)
	transaction.TransactionID = int(transactionID)

	return transaction, nil
}

func getOperationTypeByID(operationTypeID int) (*OperationType, error) {
	for _, operationType := range operationTypes {
		if operationType.OperationTypeID == operationTypeID {
			return &operationType, nil
		}
	}
	return nil, nil
}
