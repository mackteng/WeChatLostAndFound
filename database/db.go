package database

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ItemExists(dbconfig *structures.DatabaseAccessInfo, TagID string) bool {

	db := dbconfig.Database
	var result string
	err := db.QueryRow("select TagID from tag where TagID = $1", TagID).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	} else {
		return true
	}
	return true
}

func userExists(dbconfig *structures.DatabaseAccessInfo, OpenID string) bool {

	db := dbconfig.Database
	var result string
	err := db.QueryRow("select OpenID from users where OpenID = $1", OpenID).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	} else {
		return true
	}
	return true
}

func addUser(dbconfig *structures.DatabaseAccessInfo, OpenID string) error {

	log.Println("AddUser", OpenID)

	db := dbconfig.Database
	_, err := db.Exec("INSERT INTO users VALUES($1, $2)", OpenID, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("UserAdded", OpenID)
	return err

}

func NextFinderChannel(dbconfig *structures.DatabaseAccessInfo, OpenID string) int {

	db := dbconfig.Database
	rows, err := db.Query("select finderchannel from tag where finderid = $1", OpenID)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		var table [5]bool
		for rows.Next() {
			var cur int
			err := rows.Scan(&cur)
			if err != nil {
				log.Fatal(err)
			} else {
				table[cur-6] = true
			}
		}

		for i := range table {
			if !table[i] {
				return i + 6
			}
		}
	}
	return -1

}

func NextOwnerChannel(dbconfig *structures.DatabaseAccessInfo, OpenID string) int {

	db := dbconfig.Database
	rows, err := db.Query("select ownerchannel from tag where ownerid = $1", OpenID)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		var table [5]bool
		for rows.Next() {
			var cur int
			err := rows.Scan(&cur)
			if err != nil {
				log.Fatal(err)
			} else {
				table[cur-1] = true
			}
		}

		for i := range table {
			if !table[i] {
				return i + 1
			}
		}
	}
	return -1
}

func CurrentChannel(dbconfig *structures.DatabaseAccessInfo, OpenID string) int {

	db := dbconfig.Database
	var channel int
	err := db.QueryRow("select ActiveChannel from users where OpenID=$1", OpenID).Scan(&channel)

	if err != nil {
		log.Fatal(err)
	}

	return channel

}

func ChangeChannel(dbconfig *structures.DatabaseAccessInfo, OpenID string, NewChannel int) error {

	db := dbconfig.Database
	_, err := db.Exec("UPDATE users SET ActiveChannel=$1 WHERE OpenID = $2", NewChannel, OpenID)
	// flush message queue when channel is changed
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func AddItemOwner(dbconfig *structures.DatabaseAccessInfo, OpenID string, info *structures.ItemInfo) error {

	log.Println("AddItemOwner: ", info)

	db := dbconfig.Database

	if !userExists(dbconfig, OpenID) {
		addUser(dbconfig, OpenID)
	}

	// extract from info and insert into database

	next_channel := NextOwnerChannel(dbconfig, OpenID)

	_, err := db.Exec(`INSERT INTO tag VALUES($1, $2, $3, $4, $5, $6, $7)`, info.TagID, info.Name, info.Description, OpenID, next_channel, nil, nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Adding item ", info.TagID, "for owner ", OpenID, "on channel ", next_channel)
	ChangeChannel(dbconfig, OpenID, next_channel)
	return err
}

func AddItemFinder(dbconfig *structures.DatabaseAccessInfo, FinderOpenID string, TagID string) error {

	//db:= dbconfig.Database

	return nil

}
