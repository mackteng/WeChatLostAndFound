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
		log.Println("NewDatabase", "Database Connected")
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
/*
func (Database *Database) nextOwnerChannel(OpenID string) (int, error) {

	db := Database.SQLDriver
	rows, err := db.Query("select ownerchannel from tag where ownerid = $1", OpenID)
	defer rows.Close()
	if err != nil {
		return -1, err
	} else {
		var table [5]bool
		for rows.Next() {
			var cur int
			err := rows.Scan(&cur)
			if err != nil {
				return -1, err
			} else {
				table[cur-1] = true
			}
		}

		for i := range table {
			if !table[i] {
				return i + 1, nil
			}
		}
	}
	return -1, errors.New(sysmsg.OWNER_LIMIT_REACHED)
}

func (Database *Database) nextFinderChannel(OpenID string) (int, error) {

	//return 6, nil

	db := Database.SQLDriver
	rows, err := db.Query("select finderchannel from tag where finderid = $1", OpenID)
	defer rows.Close()
	if err != nil {
		return 1, err
	} else {
		var table [5]bool
		for rows.Next() {
			var cur int
			err := rows.Scan(&cur)
			if err != nil {
				return 1, err
			} else {
				table[cur-6] = true
			}
		}

		for i := range table {
			if !table[i] {
				return i + 6, nil
			}
		}
	}
	return -1, errors.New(sysmsg.FINDER_LIMIT_REACHED)

}

func (Database *Database) CurrentChannel(OpenID string) (int, error) {

	db := Database.SQLDriver
	var channel int
	err := db.QueryRow("select ActiveChannel from users where OpenID=$1", OpenID).Scan(&channel)

	if err != nil {
		return -1, err
	}

	return channel, nil

}

func (Database *Database) ChangeChannel(OpenID string, NewChannel int) error {

	cur, _ := Database.CurrentChannel(OpenID)

	if cur == NewChannel{
		return errors.New(sysmsg.SAME_CHANNEL)
	}

	db := Database.SQLDriver
	_, err := db.Exec("UPDATE users SET ActiveChannel=$1 WHERE OpenID = $2", NewChannel, OpenID)
	if err != nil {
		return err
	} else {
		log.Println("Successfully Changed ", OpenID, "'s", " to ", NewChannel)
	}
	return nil
}
*/
func (Database *Database) RegisterTag(OpenID string, TagID string, Info structures.ItemInfo) error {

	if exists, _ := Database.itemExists(TagID); exists {
		return  errors.New(sysmsg.ITEM_ALREADY_REGISTERED)
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
		return  errors.New("sysmsg.ITEM_NOT_REGISTERED")
	}

	_, err := db.Exec(`UPDATE tag SET finderid=$1 WHERE tagid=$2`, FinderOpenID, TagID)

	if err != nil {
		return  err
	}
	log.Println("FindTag", FinderOpenID, " found ", TagID)
	return err
}


func (Database *Database) GetActiveTag(OpenID string) (string, error) {

	db := Database.SQLDriver
	var result string
	err := db.QueryRow(`SELECT ActiveTag FROM users WHERE OpenID=$1`, OpenID).Scan(&result)
	return result,err
}


func (Database *Database) ChangeActiveTag(OpenID, NewActiveTag string) error {

        db := Database.SQLDriver
        _, err :=db.Exec(`UPDATE users SET ActiveTag=$1 WHERE OpenID=$2`, NewActiveTag, OpenID)
	return err
}

func (Database *Database) FindCorrespondingUser(OpenID string) (string, string,  error) {

	var corresponding string
	var err error
	var ownerid string
	var finderid string
	db := Database.SQLDriver



	TagID, cerr := Database.GetActiveTag(OpenID)
	log.Println("TagID", TagID)
	if cerr != nil {
		return "", "", cerr
	}
	err = db.QueryRow(`SELECT ownerid,finderid FROM tag WHERE TagID=$1`, TagID).Scan(&ownerid, &finderid)

	log.Println(ownerid, finderid)

	if err!= nil {
		return "", "", err
	}


	if ownerid==OpenID {
		corresponding = finderid;
	} else {
		corresponding = ownerid;
	}

	log.Println("FindCorrespondingUser", corresponding)
	return TagID, corresponding,  err
}
