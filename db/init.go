package db

import (
	"database/sql"
	"time"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	_ "github.com/go-sql-driver/mysql"
)

const sessionTime = 300
const sessionCleantime = 60

//-- Session DB initialization
func DbInit() {

	dbUserName := util.GetConfig().Get("dbUserName")
	dbPassword := util.GetConfig().Get("dbPassword")
	dbHost := util.GetConfig().Get("dbHost")
	dbPort := util.GetConfig().Get("dbPort")
	databaseName := util.GetConfig().Get("databaseName")

	dataSource := dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + databaseName

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		util.LogErr(err)
	}

	err = db.Ping()
	if err != nil {
		util.LogErr(err)
	}

	errLogIdInit(db)
	go dbCleanPeriod(db)

	cns.Db = db
}

//-- Session delete period
func dbCleanPeriod(db *sql.DB) {
	var sessionId string
	var timestamp string
	timeFormat := "2006-01-02 15:04:05"

	for {
		time.Sleep(time.Second * sessionCleantime)
		queryResult, err := db.Query("select session_id, timestamp from session")
		if err != nil {
			util.LogErr(err)
		}
		defer queryResult.Close()

		for queryResult.Next() {
			err = queryResult.Scan(&sessionId, &timestamp)
			if err != nil {
				util.LogErr(err)
			}

			timeParse, err := time.Parse(timeFormat, timestamp)
			if err != nil {
				util.LogErr(err)
			}

			if time.Now().Sub(timeParse) > time.Second*sessionTime {
				DelSession(db, sessionId)
			}
		}
	}
}

//-- Error log db
func errLogIdInit(db *sql.DB) {
	var index string

	queryResult, _ := db.Query("select index_id from errorLog order by index_id desc limit 1")
	defer queryResult.Close()

	for queryResult.Next() {
		queryResult.Scan(&index)
	}

	if index == "" {
		util.InitLogIndex().NewLogIndex("0")
	} else {
		util.InitLogIndex().NewLogIndex(index)
	}
}
