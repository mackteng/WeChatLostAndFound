package database

import (
	"database/sql"
	_ "github.com/lib/pq"

	"errors"
	"log"

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
		log.Println("Database Connected")
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

	log.Println("Adding User: ", OpenID)
	db := Database.SQLDriver
	_, err := db.Exec("INSERT INTO users VALUES($1, $2)", OpenID, nil)
	if err != nil {
		return err
	}
	log.Println("User Added: ", OpenID)
	return nil

}

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

func (Database *Database) RegisterTag(OpenID string, TagID string) (int, error) {

	if exists, _ := Database.itemExists(TagID); exists {
		return -1, errors.New(sysmsg.ITEM_ALREADY_REGISTERED)
	}

	next_channel, err := Database.nextOwnerChannel(OpenID)

	if err != nil {
		return -1, err
	}

	_, err = Database.SQLDriver.Exec(`INSERT INTO tag VALUES($1, $2, $3, $4, $5, $6, $7)`, TagID, "foo", "foo", OpenID, next_channel, nil, nil)

	if err != nil {
		return -1, err
	}
	log.Println("Added item ", TagID, "for owner ", OpenID, "on channel ", next_channel)
	return next_channel, nil
}

func (Database *Database) FindTag(FinderOpenID string, TagID string) (int, error) {

	db := Database.SQLDriver

	if exists, _ := Database.itemExists(TagID); !exists {
		return -1, errors.New("sysmsg.ITEM_NOT_REGISTERED")
	}

	next_channel, err := Database.nextFinderChannel(FinderOpenID)

	if err != nil {
		return -1, err
	}

	_, err = db.Exec(`UPDATE tag SET finderid=$1, finderchannel=$2 WHERE tagid=$3`, FinderOpenID, next_channel, TagID)

	if err != nil {
		return -1, err
	}
	log.Println("Find Tag: ", FinderOpenID, " found ", TagID)
	return next_channel, nil
}

func (Database *Database) FindCorrespondingUser(OpenID string) (string, int, error) {

	var id string
	var channel int
	var err error

	db := Database.SQLDriver
	Channel, cerr := Database.CurrentChannel(OpenID)

	if cerr != nil {
		return "", 0, nil
	}

	if Channel <= 5 {

		err = db.QueryRow("select finderid, finderchannel FROM tag WHERE ownerid=$1 AND ownerchannel=$2", OpenID, Channel).Scan(&id, &channel)
		if err != nil {
			return "", 0, err
		}

	} else {

		err = db.QueryRow("select ownerid, ownerchannel FROM tag WHERE finderid=$1 AND finderchannel=$2", OpenID, Channel).Scan(&id, &channel)
		if err != nil {
			return "", 0, err
		}

	}

	return id, channel, err
}
