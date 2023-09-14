package services_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/Exzrgs/myapi/services"
	_ "github.com/go-sql-driver/mysql"
)

var s *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := "docker"
	// dbPassword = os.Getenv("DB_PASSWORD")
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println(err)
		return
	}

	s = services.NewMyAppService(db)

	m.Run()
}
