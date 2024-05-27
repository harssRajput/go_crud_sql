package utilityStore

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/service"
	"log"
	"os"
)

type UtilityStore struct {
	Logger       *log.Logger
	SqlDB        *sql.DB
	ServiceStore *service.ServiceStore
	HttpRouter   *mux.Router
}

func InitUtilityStore() (*UtilityStore, error) {
	logger, err := initLogger()
	if err != nil {
		logger.Fatal("Error initializing logger", err)
	}

	sqldb, err := initSQLDB(logger)
	if err != nil {
		logger.Fatal("Error connecting to database", err)
	}

	serviceStore, err := service.InitServiceStore(sqldb, logger)
	if err != nil {
		logger.Fatalf("Error initializing services: %v\n", err)
	}
	return &UtilityStore{
		Logger:       logger,
		SqlDB:        sqldb,
		ServiceStore: serviceStore,
		HttpRouter:   mux.NewRouter(),
	}, nil
}

func initSQLDB(logger *log.Logger) (*sql.DB, error) {
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
