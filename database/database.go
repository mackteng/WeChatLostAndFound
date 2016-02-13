package database

import (
	"database/sql"
	_ "github.com/lib/pq"

	"errors"
	"log"

	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures/sysmsg"
)

type Database struct {
	SQLDriver *sql.DB
}

func NewDatabase() *Database {

	url := "postgresql://mackygood:seetopdf87abA1@lostandfound.cq7rqo2imtyn.ap-northeast-1.rds.amazonaws.com:5432/lostandfound"
	db, err := sql.Open("postgres", url)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("NewDatabase", "DatabaseConnected")
		return &Database{SQLDriver: db}
	}
	return nil
}

func (Database *Database) itemExists(TagID string) (bool, error) {

	db := Database.SQLDriver
	var result string
	err := db.QueryRow("select TagID from tag where TagID = $1", TagID).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (Database *Database) userExists(OpenID string) (bool, error) {

	db := Database.SQLDriver
	var result string
	err := db.QueryRow("select OpenID from users where OpenID = $1", OpenID).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (Database *Database) userOwns(OpenID string, TagID string) (bool, error) {

	db := Database.SQLDriver

	var ownerid string
	err:= db.QueryRow(`SELECT ownerid FROM tag WHERE tagid=$1`, TagID).Scan(&ownerid)

	if err!=nil {
		return false, err
	}

	if OpenID == ownerid {
		return true , err
	} else {
		return false, err
	}
}

func (Database *Database) AddUser(OpenID string) error {

	exists, existserr := Database.userExists(OpenID)

	if existserr != nil {
		return existserr
	}

	if exists {
		return nil
	}

	db := Database.SQLDriver
	_, err := db.Exec("INSERT INTO users VALUES($1, $2)", OpenID, nil)
	if err != nil {
		return err
	}
	log.Println("AddUser", OpenID)
	return nil

}
func (Database *Database) RegisterTag(OpenID string, TagID string, Info structures.ItemInfo) error {

	if exists, _ := Database.itemExists(TagID); exists {
		return errors.New(sysmsg.ITEM_ALREADY_REGISTERED)
	}

	_, err := Database.SQLDriver.Exec(`INSERT INTO tag VALUES($1, $2, $3, $4, $5)`, TagID, Info.Name, Info.Description, OpenID, nil)

	if err != nil {
		return err
	}
	log.Println("RegisterTag", TagID, "for owner ", OpenID)
	return err
}

func (Database *Database) FindTag(FinderOpenID string, TagID string) error {

	db := Database.SQLDriver

	if exists, _ := Database.itemExists(TagID); !exists {
		return errors.New("sysmsg.ITEM_NOT_REGISTERED")
	}

	_, err := db.Exec(`UPDATE tag SET finderid=$1 WHERE tagid=$2`, FinderOpenID, TagID)

	if err != nil {
		return err
	}
	log.Println("FindTag", FinderOpenID, " found ", TagID)
	return err
}

func (Database *Database) GetActiveTag(OpenID string) (string, error) {

	db := Database.SQLDriver
	var result string
	err := db.QueryRow(`SELECT ActiveTag FROM users WHERE OpenID=$1`, OpenID).Scan(&result)
	return result, err
}

func (Database *Database) ChangeActiveTag(OpenID, NewActiveTag string) error {

	db := Database.SQLDriver

	owns, err := Database.userOwns(OpenID, NewActiveTag)

	if err!=nil {
		return err
	}
	if owns {
		_, err = db.Exec(`UPDATE users SET ActiveTag=$1 WHERE OpenID=$2`, NewActiveTag, OpenID)
	}
	return err
}

func (Database *Database) FindCorrespondingUser(OpenID string) (string, string, error) {

	var corresponding string
	var err error
	var ownerid string
	var finderid string
	db := Database.SQLDriver

	TagID, cerr := Database.GetActiveTag(OpenID)
	if cerr != nil {
		return "", "", cerr
	}
	err = db.QueryRow(`SELECT ownerid,finderid FROM tag WHERE TagID=$1`, TagID).Scan(&ownerid, &finderid)


	if err != nil {
		return "", "", err
	}

	if ownerid == OpenID {
		corresponding = finderid
	} else {
		corresponding = ownerid
	}

	log.Println("FindCorrespondingUser", corresponding)
	return TagID, corresponding, err
}

	
func (Database *Database) GetAllOwnedItems(OpenID string) ([]structures.ItemInfo,  error) {
	
	db:=Database.SQLDriver
	var items []structures.ItemInfo	
	var (
		tagid string 
		name  string 
		description string
		finderid sql.NullString
	)

	rows, err := db.Query(`SELECT tagid, name, description, finderid FROM tag WHERE ownerid=$1`, OpenID)
	defer rows.Close()

	if err!= nil {
		return nil, err
	}


	if err!=nil {
		return nil, err
	}


	for rows.Next() {
		err:=rows.Scan(&tagid, &name, &description, &finderid)

		if err!=nil {
			return nil, err
		}
		items = append(items,structures.ItemInfo {
			TagID: tagid,
			Name: name,
			Description: description,
		})
		log.Println(tagid, name, description, finderid) }
	return items, nil
}

func (Database *Database) DeleteTag(OpenID, TagID string) error {
	log.Println("DeleteTag", OpenID, TagID)
	db := Database.SQLDriver

	var ownerid string
	var  finderid sql.NullString
	

	exists, err := Database.itemExists(TagID)

	if !exists {
		return errors.New("Item Does Not Exist")
	}
	
	err = db.QueryRow(`SELECT ownerid, finderid FROM tag WHERE tagid= $1`, TagID).Scan(&ownerid, &finderid)

	if err!=nil {
		return err
	}

	if ownerid == OpenID {
		_,err = db.Exec(`DELETE FROM tag WHERE tagid=$1`, TagID)
	} else if finderid.Valid && finderid.String == OpenID {
		_,err = db.Exec(`UPDATE tag SET finderid=$1 WHERE tagid=$2`, nil, TagID)
	}

	return err
}

func (Database *Database) DeleteFinder(OpenID, TagID string) error {
	log.Println("DeleteFinder", OpenID, TagID)
	db := Database.SQLDriver

	exists, err := Database.itemExists(TagID)

	if !exists || err!=nil{
		return errors.New("Delete Finder")
	}

	owns, err := Database.userOwns(OpenID, TagID)

	if err!=nil {
		return err
	}

	if owns {
		db.Exec(`UPDATE tag SET finderid=$1 WHERE tagid=$2`, nil, TagID)
	}

	return err
}
