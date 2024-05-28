package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	ent "github.com/harssRajput/go_crud_sql/internal/entity"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

const (
	url = "http://localhost:8080"
)

func TestMain(m *testing.M) {
	// Setup Docker Compose
	log.Printf("starting Docker services...")
	cmd := exec.Command("docker-compose", "-f", "../docker-compose.yml", "up", "-d")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to start docker: %v", err)
	}

	// Wait for services to be ready
	time.Sleep(20 * time.Second) // Adjust as needed
	log.Printf("Docker services are ready")

	// Run tests
	code := m.Run()

	// Teardown Docker Compose
	//cmd = exec.Command("docker-compose", "-f", "../docker-compose.yml", "down")
	cmd = exec.Command("docker", "stop", "mysql", "webapp")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to stop docker-compose : %v", err)
	}
	log.Printf("Docker services stopped")
	time.Sleep(5 * time.Second) // Adjust as needed
	//Exit with the result of the tests
	os.Exit(code)
}

func Test_CreateAccount(t *testing.T) {
	randDocNumber := generateRandomDocNumber()
	type Request struct {
		documentNumber string
	}
	type ExpectedResponse struct {
		account    *ent.Account
		httpStatus int
		err        string
	}
	tests := []struct {
		name   string
		fields Request
		want   ExpectedResponse
	}{
		{
			name: "valid document number",
			fields: Request{
				documentNumber: randDocNumber,
			},
			want: ExpectedResponse{
				account: &ent.Account{
					AccountId:      1,
					DocumentNumber: randDocNumber,
				},
				httpStatus: http.StatusCreated,
				err:        "",
			},
		},
		{
			name: "already exist document number",
			fields: Request{
				documentNumber: randDocNumber,
			},
			want: ExpectedResponse{
				account: &ent.Account{
					AccountId:      1,
					DocumentNumber: randDocNumber,
				},
				httpStatus: http.StatusConflict,
				err:        "Duplicate entry",
			},
		},
		{
			name: "invalid document number: length more than 11",
			fields: Request{
				documentNumber: "00340123456789098",
			},
			want: ExpectedResponse{
				account:    nil,
				httpStatus: http.StatusBadRequest,
				err:        "document number should be of length 11",
			},
		},
		{
			name: "invalid document number: empty",
			fields: Request{
				documentNumber: "",
			},
			want: ExpectedResponse{
				account:    nil,
				httpStatus: http.StatusBadRequest,
				err:        "document number is required",
			},
		},
		{
			name: "invalid document number: non-numeric",
			fields: Request{
				documentNumber: "0034067890a",
			},
			want: ExpectedResponse{
				account:    nil,
				httpStatus: http.StatusBadRequest,
				err:        "document number should be numeric",
			},
		},
		{
			name: "invalid document number: length less than 11",
			fields: Request{
				documentNumber: "00340",
			},
			want: ExpectedResponse{
				account:    nil,
				httpStatus: http.StatusBadRequest,
				err:        "document number should be of length 11",
			},
		},
	}
	createAccountUrl := url + "/accounts/"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountRequest := map[string]string{
				"document_number": tt.fields.documentNumber,
			}
			accountJSON, err := json.Marshal(accountRequest)
			if err != nil {
				t.Fatalf("Error marshalling account request: %v", err)
			}

			resp, err := http.Post(createAccountUrl, "application/json", bytes.NewBuffer(accountJSON))
			if err != nil {
				t.Fatalf("Error sending request to server: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.want.httpStatus {
				t.Fatalf("Assertion failed: httpStatusCode got %v, want: %v", resp.StatusCode, tt.want.httpStatus)
			}

			if tt.want.httpStatus == http.StatusCreated {
				var acc ent.Account
				err = json.NewDecoder(resp.Body).Decode(&acc)
				if err != nil {
					t.Fatalf("Error decoding response: %v", err)
				}
				if acc.DocumentNumber != tt.want.account.DocumentNumber {
					t.Fatalf("Assertion failed: document_number got %v, want: %v", acc.DocumentNumber, tt.want.account.DocumentNumber)
				}
				log.Printf("CreateAccount result: %+v, account_id: %d", acc, acc.AccountId)
			} else {
				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatal("Error reading response body:", err)
				}
				if !strings.Contains(string(respBody), tt.want.err) {
					t.Fatalf("Assertion failed: error got %v, want: %v", string(respBody), tt.want.err)
				}
				log.Printf("CreateAccount result: %s", respBody)
			}

		})
	}
}

func Test_CreateTransaction(t *testing.T) {
	type Request struct {
		accountId     int
		amount        float64
		operationType int
	}
	type ExpectedResponse struct {
		httpStatus int
		err        string
	}
	tests := []struct {
		name   string
		fields Request
		want   ExpectedResponse
	}{
		{
			name: "valid transaction",
			fields: Request{
				accountId:     1,
				amount:        -100.0,
				operationType: 1,
			},
			want: ExpectedResponse{
				httpStatus: http.StatusCreated,
				err:        "",
			},
		},
		{
			name: "invalid account id",
			fields: Request{
				accountId:     222,
				amount:        -100.0,
				operationType: 1,
			},
			want: ExpectedResponse{
				httpStatus: http.StatusBadRequest,
				err:        "Cannot add or update a child row: a foreign key constraint fails",
			},
		},
		{
			name: "invalid amount",
			fields: Request{
				accountId:     1,
				amount:        0,
				operationType: 1,
			},
			want: ExpectedResponse{
				httpStatus: http.StatusBadRequest,
				err:        "Amount cannot be zero",
			},
		},
		{
			name: "invalid operation type",
			fields: Request{
				accountId:     1,
				amount:        -100.0,
				operationType: 0,
			},
			want: ExpectedResponse{
				httpStatus: http.StatusInternalServerError,
				err:        "operation type not found",
			},
		},
		{
			name: "invalid amount: negative amount for credit voucher operation type",
			fields: Request{
				accountId:     1,
				amount:        -100.0,
				operationType: 4,
			},
			want: ExpectedResponse{
				httpStatus: http.StatusBadRequest,
				err:        "invalid amount credit voucher amount should be positive",
			},
		},
		{
			name: "invalid amount: positive amount for purchase operation type",
			fields: Request{
				accountId:     1,
				amount:        100.0,
				operationType: 1,
			},
			want: ExpectedResponse{
				httpStatus: http.StatusBadRequest,
				err:        "invalid amount amount should be positive",
			},
		},
	}
	trxUrl := url + "/transactions/"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTransactionBody := ent.Transaction{
				AccountID:       tt.fields.accountId,
				OperationTypeID: tt.fields.operationType,
				Amount:          tt.fields.amount,
			}
			createTransactionJSON, err := json.Marshal(createTransactionBody)
			if err != nil {
				t.Fatalf("Error marshalling account request: %v", err)
			}

			resp, err := http.Post(trxUrl, "application/json", bytes.NewBuffer(createTransactionJSON))
			if err != nil {
				t.Fatalf("Error sending request to server: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.want.httpStatus {
				t.Fatalf("Assertion failed: httpStatusCode got %v, want: %v", resp.StatusCode, tt.want.httpStatus)
			}

			if tt.want.httpStatus == http.StatusCreated {
				var trx ent.Transaction
				err = json.NewDecoder(resp.Body).Decode(&trx)
				if err != nil {
					t.Fatalf("Error decoding response: %v", err)
				}
				log.Printf("CreateTransaction result: %+v", trx)
			} else {
				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatal("Error reading response body:", err)
				}
				if !strings.Contains(string(respBody), tt.want.err) {
					t.Fatalf("Assertion failed: error got %v, want: %v", string(respBody), tt.want.err)
				}
				log.Printf("CreateTransaction result: %s", respBody)
			}

		})
	}
}

func Test_GetAccount(t *testing.T) {
	type Request struct {
		accountId string
	}
	type ExpectedResponse struct {
		account    *ent.Account
		httpStatus int
		err        string
	}
	tests := []struct {
		name   string
		fields Request
		want   ExpectedResponse
	}{
		{
			name: "valid account id",
			fields: Request{
				accountId: "1",
			},
			want: ExpectedResponse{
				account: &ent.Account{
					AccountId:      1,
					DocumentNumber: "00340678901",
				},
				httpStatus: http.StatusOK,
				err:        "",
			},
		},
		{
			name: "invalid account id",
			fields: Request{
				accountId: "222",
			},
			want: ExpectedResponse{
				account:    nil,
				httpStatus: http.StatusNotFound,
				err:        "Account not found",
			},
		},
		{
			name: "invalid account id: non-int",
			fields: Request{
				accountId: "na",
			},
			want: ExpectedResponse{
				account:    nil,
				httpStatus: http.StatusBadRequest,
				err:        "Invalid account ID",
			},
		},
	}
	accountURL := fmt.Sprintf("%s/accounts/", url)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getAccountURL := fmt.Sprintf("%s%s", accountURL, tt.fields.accountId)
			resp, err := http.Get(getAccountURL)
			if err != nil {
				t.Fatalf("Error calling GET /accounts/%s err: %s", tt.fields.accountId, err.Error())
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.want.httpStatus {
				t.Fatalf("Assertion failed: httpStatusCode got %v, want: %v", resp.StatusCode, tt.want.httpStatus)
			}
			if tt.want.httpStatus == http.StatusOK {
				var acc ent.Account
				err = json.NewDecoder(resp.Body).Decode(&acc)
				if err != nil {
					t.Fatalf("Error decoding response: %v", err)
				}
				if acc.AccountId != tt.want.account.AccountId {
					t.Fatalf("Assertion failed: AccountId got %v, want: %v", acc.AccountId, tt.want.account.AccountId)
				}
			} else {
				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatal("Error reading response body:", err)
				}
				if !strings.Contains(string(respBody), tt.want.err) {
					t.Fatalf("Assertion failed: error got %v, want: %v", string(respBody), tt.want.err)
				}
				log.Printf("Response from GET /accounts/%s response: %s", tt.fields.accountId, string(respBody))
			}
		})
	}
}

func generateRandomDocNumber() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100,000,000
	randNum := rand.Intn(1000000000) + 1
	randDocNumber := fmt.Sprintf("%011d", randNum)
	return randDocNumber
}
