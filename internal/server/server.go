package server

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/handler"
	"log"
	"net/http"
	"os"
)

// TODO: put it in a config file. (optional: put separate config file based on envTYpe)
const (
	HTTP_PORT   = "3333"
	DB_USERNAME = "root"
	DB_PASSWORD = "root"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_NAME     = "webapp"
)

// TODO: introduce utilityStore (optional)
func RunServer() {
	logger, err := initLogger()
	if err != nil {
		logger.Fatal("Error initializing logger", err)
	}

	sqldb, err := initDB(logger)
	if err != nil {
		logger.Fatal("Error connecting to database", err)
	}
	defer sqldb.Close()

	initHttpServer(sqldb, logger)
}

func initHttpServer(sqldb *sql.DB, logger *log.Logger) {
	r := mux.NewRouter()
	err := handler.InitHandlers(r, sqldb, logger)
	if err != nil {
		logger.Fatalf("Error initializing handlers: %v\n", err)
	}

	//start listening
	err = http.ListenAndServe(":"+HTTP_PORT, r)
	if errors.Is(err, http.ErrServerClosed) {
		//TODO: put graceful shutdon if possible. (optional: not part of test)
		logger.Fatalf("server closed\n")
	} else if err != nil {
		logger.Fatalf("error starting server: %s\n", err)
	}
}

func initDB(logger *log.Logger) (*sql.DB, error) {
	dbUsername := DB_USERNAME
	dbPassword := DB_PASSWORD
	dbHost := DB_HOST
	dbPort := DB_PORT
	dbName := DB_NAME
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Connected to the database!")

	return db, err
}

func initLogger() (*log.Logger, error) {
	logger := log.New(os.Stdout, "webapp: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger, nil
}
