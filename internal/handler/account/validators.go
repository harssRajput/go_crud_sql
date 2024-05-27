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
