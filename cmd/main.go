package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/harssRajput/go_crud_sql/internal/server"
	"log"
)

func main() {
	log.Println("Starting go-crud application...")
	server.RunServer()
}

//func mains() {
//
//	query := `INSERT INTO Account (document_number) VALUES (?)`
//	response, err := db.Exec(query, 12345678901)
//	if err != nil {
//		panic(err.Error())
//	}
//	id, err := response.LastInsertId()
//	fmt.Println("data inserted", id)
//
//}
