package transaction

import (
	"errors"
	"github.com/harssRajput/go_crud_sql/internal/model/account"
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
	TransactionID   int     `json:"transaction_id"`
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	EventDate       string  `json:"event_date"` //event_date contains date and time accurate upto milliseconds in string format "2021-01-01T10:00:00.000Z"
}

// create seed data in memory
var transactions = []Transaction{
	{TransactionID: 1, AccountID: 1, OperationTypeID: OperationTypeNormalPurchase, Amount: 100.0, EventDate: "2021-01-01T10:00:00.000Z"},
}

func CreateTransaction(transaction *Transaction) (*Transaction, error) {
	// check if account exists
	accountObj, err := account.GetAccountByID(transaction.AccountID)
	if err != nil {
		return nil, err
	} else if accountObj == nil {
		return nil, errors.New("account not found")
	}

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

	//generate transaction id. In real scenario, it's auto-generated DB field
	transaction.TransactionID = len(transactions) + 1
	transaction.EventDate = time.Now().Format("2006-01-02T15:04:05.000Z")
	//transactions array is not thread-safe as eventually it will be replaced by DB.
	transactions = append(transactions, *transaction)

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
