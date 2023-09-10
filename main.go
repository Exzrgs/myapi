package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Exzrgs/myapi/api"
	"github.com/Exzrgs/myapi/apperrors"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser = "docker"
	// dbPassword = os.Getenv("DB_PASSWORD")
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	fmt.Println(apperrors.MyAppError{ErrCode: apperrors.BadParameter, Message: "a", Err: sql.ErrNoRows})

	db, err := connectDB()
	if err != nil {
		log.Println("fail to connect db")
		return
	}

	r := api.NewRouter(db)

	log.Println("running")
	log.Fatal(http.ListenAndServe(":8080", r))
}
