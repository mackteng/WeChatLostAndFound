package structures

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseAccessInfo struct {
	Database *sql.DB
}

func NewDatabase() *DatabaseAccessInfo {

	url := "postgresql://mackygood:seetopdf87abA1@lostandfound.cq7rqo2imtyn.ap-northeast-1.rds.amazonaws.com:5432/lostfound"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database Connected")
		return &DatabaseAccessInfo{Database: db}
	}
	return nil
}
