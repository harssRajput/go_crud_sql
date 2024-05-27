package entity

import "time"

type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"` //event_date contains date and time accurate upto milliseconds in string format "2021-01-01T10:00:00.000Z"
}
