package server

import (
	"errors"
	"github.com/harssRajput/go_crud_sql/internal/handler"
	us "github.com/harssRajput/go_crud_sql/internal/utilityStore"
	"log"
	"net/http"
)

func RunServer() {
	utilityStore, err := us.InitUtilityStore()
	if err != nil {
		log.Fatalf("Error initializing utility store: %v\n", err)
	}
	defer utilityStore.SqlDB.Close()

	initHttpServer(utilityStore)
}

func initHttpServer(utilityStore *us.UtilityStore) {
	err := handler.InitHandlers(utilityStore)
	if err != nil {
		utilityStore.Logger.Fatalf("Error initializing handlers: %v\n", err)
	}

	//start listening
	err = http.ListenAndServe(":"+us.HTTP_PORT, utilityStore.HttpRouter)
	if errors.Is(err, http.ErrServerClosed) {
		//TODO: put graceful shutdon if possible. (optional: not part of test)
		utilityStore.Logger.Fatalf("server closed\n")
	} else if err != nil {
		utilityStore.Logger.Fatalf("error starting server: %s\n", err)
	}
}
