package transaction

import (
	"database/sql"
	"errors"
	"fmt"
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
	// validation: check if operation type exists
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
		return nil, errors.New("invalid amount amount should be negative for purchase and withdrawal operations")
	}

	// if amount is positive, handle discharge of negative balance
	if transaction.Amount > 0 {
		err := tx.dischargeNegativeBalance(transaction)
		if err != nil {
			return nil, err
		}
	}

	eventDateTime, err := time.Parse(time.RFC3339, time.Now().Format("2006-01-02T15:04:05.000Z"))
	if err != nil {
		tx.logger.Printf("Error parsing date: %v\n", err)
		return nil, err
	}
	transaction.EventDate = eventDateTime
	// insert transaction in a db
	query := `INSERT INTO Transaction (account_id, operation_type_id, amount, balance, event_date) VALUES (?, ?, ?, ?, ?)`
	response, err := tx.sqldb.Exec(query, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.Balance, transaction.EventDate)
	if err != nil {
		tx.logger.Printf("Error inserting transaction: %v\n", err)
		return nil, err
	}
	transactionID, err := response.LastInsertId()
	tx.logger.Println("transaction created with transactionID: ", transactionID)
	transaction.TransactionID = int(transactionID)

	return transaction, nil
}

//TODO: think for db call optimisations
/*
- 1st optimisation: can come up with rowCounts number to fetch at a time. (instead of fetching all rows which may hit system if rows in thousands for that accountid)
ex -> average rows discharged for 1 positive amount is 10
so fetch 10+3(3 as a padding)=13 rows at a time. then mostly 1 db call required and in rare cases 2 db calls.
 - 2nd optimisation (done already): update all negative rows at the end in single db call.
	-	during system failure, 1 update/db call may end up partial updates.
*/
func (tx *transactionService) dischargeNegativeBalance(reqTrx *ent.Transaction) error {
	negBalTrxs := []*ent.Transaction{}
	//fetch all rows with negative balance
	query := `SELECT transaction_id, balance FROM Transaction WHERE account_id = ? AND balance < 0`
	rows, _ := tx.sqldb.Query(query, reqTrx.AccountID)
	for rows.Next() {
		negBalTrx := ent.Transaction{}
		err := rows.Scan(&negBalTrx.TransactionID, &negBalTrx.Balance)
		if err != nil {
			tx.logger.Printf("Error scanning row: %v\n", err)
		}
		negBalTrxs = append(negBalTrxs, &negBalTrx)
	}

	if len(negBalTrxs) == 0 {
		return nil
	}

	negTrxUpdateAllquery := `UPDATE Transaction SET balance = CASE`
	dischargeCount := 0
	//iteratively until req.balance is > 0 AND there are rows with negative balance
	for _, negBalTrx := range negBalTrxs {
		if reqTrx.Balance > -1*negBalTrx.Balance {
			// if +ve bal > -ve bal
			reqTrx.Balance = reqTrx.Balance + negBalTrx.Balance
			negBalTrx.Balance = 0
		} else {
			// if +ve bal < -ve bal
			negBalTrx.Balance = negBalTrx.Balance + reqTrx.Balance
			reqTrx.Balance = 0
		}
		//make entry in update all negTrxUpdateAllquery
		negTrxUpdateAllquery += ` WHEN transaction_id = ` + fmt.Sprintf("%d", negBalTrx.TransactionID) + ` THEN ` + fmt.Sprintf("%f", negBalTrx.Balance)
		dischargeCount++

		//if req.balance is zero break
		if reqTrx.Balance == 0 {
			break
		}
	}
	//add transaction id of negative discharged balance to update all query
	negTrxUpdateAllquery += ` ELSE balance END WHERE transaction_id IN (`
	for itr := 0; itr < dischargeCount; itr++ {
		negTrxUpdateAllquery += fmt.Sprintf("%d", negBalTrxs[itr].TransactionID) + `,`
	}
	negTrxUpdateAllquery = negTrxUpdateAllquery[:len(negTrxUpdateAllquery)-1] + `)`
	//update negBalanceTrxs to db whose balance gets discharged
	tx.logger.Printf("Discharging transaction update query: %v\n", negTrxUpdateAllquery)
	_, err := tx.sqldb.Exec(negTrxUpdateAllquery)
	if err != nil {
		tx.logger.Printf("Error discharging transaction: %v\n", err)
		return err
	}

	return nil
}

func getOperationTypeByID(operationTypeID int) (*ent.OperationType, error) {
	for _, operationType := range ent.OperationTypes {
		if operationType.OperationTypeID == operationTypeID {
			return &operationType, nil
		}
	}
	return nil, nil
}
