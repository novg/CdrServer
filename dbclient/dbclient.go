package dbclient

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db   *sql.DB
	stmt *sql.Stmt
)

const insertQuery = `INSERT INTO calls(datetime, duration, seg, sop, dest, numin, numout, str1, str2)
							VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

// InitDatabase create table of database
func InitDatabase(pdb **sql.DB) {
	db = *pdb
	makeTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS calls (" +
			"id BIGSERIAL PRIMARY KEY, " +
			"datetime TIMESTAMP NOT NULL, " +
			"duration INTERVAL NOT NULL, " +
			"seg VARCHAR, " +
			"sop VARCHAR, " +
			"dest VARCHAR, " +
			"numin VARCHAR NOT NULL, " +
			"numout VARCHAR NOT NULL, " +
			"str1 VARCHAR, " +
			"str2 VARCHAR" +
			");")

	_, err := db.Exec(makeTable)
	checkErr(err)

	stmt, err = db.Prepare(insertQuery)
	checkErr(err)
}

// Run working with database
func sendToDB(info *callInfo) {
	_, err := stmt.Exec(info.datetime, info.duration, info.seg, info.sop, info.dest,
		info.numin, info.numout, info.str1, info.str2)

	if err != nil {
		log.Println(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
