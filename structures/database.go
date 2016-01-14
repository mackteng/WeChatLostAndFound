package structures

/*
import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseAccessInfo struct {
	Database *sql.DB
}

func NewDatabase() *DatabaseAccessInfo {

	url := "postgresql://mackygood:seetopdf87abA1@lostandfound.cq7rqo2imtyn.ap-northeast-1.rds.amazonaws.com:5432/lostandfound"
	db, err := sql.Open("postgres", url)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database Connected")
		return &DatabaseAccessInfo{Database: db}
	}
	return nil
}
*/

type DatabaseInteractor interface {
	AddUser(OpenID string) error
	RegisterTag(OpenID string, TagID string) (int, error)
	FindTag(FinderOpenID string, TagID string) (int, error)
	CurrentChannel(OpenID string) (int, error)
	ChangeChannel(OpenID string, Channel int) error
	FindCorrespondingUser(OpenID string) (string, int, error)
}
