package utilityStore

import "os"

// TODO: put eligible const in a config file. (optional: put separate config file based on envTYpe)
var (
	HTTP_PORT   = os.Getenv("PORT")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_NAME     = os.Getenv("DB_NAME")
)

//
//var (
//	HTTP_PORT   = "8080"
//	DB_USER     = "root"
//	DB_PASSWORD = "root"
//	DB_HOST     = "127.0.0.1"
//	DB_PORT     = "3306"
//	DB_NAME     = "webapp"
//)
