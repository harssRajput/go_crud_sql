package account

import (
	"errors"
	"log"
	"strconv"
)

// when validators becomes many, we can create validator interface to organise validator hierarchy for each http request.

func validateAccountId(accountIdStr string) error {
	if accountIdStr == "" {
		log.Println("Empty accountId", accountIdStr)
		return errors.New("account ID is required")
	}

	_, err := strconv.Atoi(accountIdStr)
	if err != nil {
		log.Println("Invalid account ID", accountIdStr)
		return errors.New("invalid account ID")
	}
	return nil
}

func validateDocumentNumber(documentNumber string) error {
	if documentNumber == "" {
		log.Println("Empty document number", documentNumber)
		return errors.New("document number is required")
	} else if len(documentNumber) != 11 {
		log.Println("document number should be of length 11", documentNumber)
		return errors.New("document number should be of length 11")
	} else if _, err := strconv.Atoi(documentNumber); err != nil {
		log.Println("document number should be numeric", documentNumber)
		return errors.New("document number should be numeric")
	}

	return nil
}
