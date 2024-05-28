package transaction

import (
	"database/sql"
	"errors"
	ent "github.com/harssRajput/go_crud_sql/internal/entity"
	"log"
	"time"
)

type TransactionService interface {
	CreateTransaction(transaction *ent.Transaction) (*ent.Transaction, error)
}

type transactionService struct {
	sqldb  *sql.DB
	logger *log.Logger
}

func InitTransactionService(sqldb *sql.DB, log *log.Logger) (TransactionService, error) {
	return &transactionService{sqldb: sqldb, logger: log}, nil
}

func (tx *transactionService) CreateTransaction(transaction *ent.Transaction) (*ent.Transaction, error) {
	// check if operation type exists
	if opType, err := getOperationTypeByID(transaction.OperationTypeID); err != nil {
		return nil, err
	} else if opType == nil {
		return nil, errors.New("operation type not found")
	} else if opType.OperationTypeID == ent.OperationTypeCreditVoucher {
		if transaction.Amount <= 0 {
			return nil, errors.New("invalid amount credit voucher amount should be positive")
		}
	} else if transaction.Amount >= 0 {
		// amount should be negative for all operation types except credit voucher
		return nil, errors.New("invalid amount amount should be positive")
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

func getOperationTypeByID(operationTypeID int) (*ent.OperationType, error) {
	for _, operationType := range ent.OperationTypes {
		if operationType.OperationTypeID == operationTypeID {
			return &operationType, nil
		}
	}
	return nil, nil
}
