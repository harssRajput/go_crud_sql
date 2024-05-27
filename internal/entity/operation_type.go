package entity

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

var OperationTypes = []OperationType{
	{OperationTypeID: OperationTypeNormalPurchase, Description: "Normal Purchase"},
	{OperationTypeID: OperationTypeInstallmentPurchase, Description: "Installment Purchase"},
	{OperationTypeID: OperationTypeWithdraw, Description: "Withdraw"},
	{OperationTypeID: OperationTypeCreditVoucher, Description: "credit voucher"}, // only this operation type is allowed to have positive amount
}
