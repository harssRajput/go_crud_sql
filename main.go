package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/harssRajput/go_crud_sql/internal/server"
	"log"
)

func main() {
	log.Println("Starting go-crud application.....")
	//time.Sleep(15)
	//log.Println("Sleeped well...")
	server.RunServer()
}
