package db

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/Moonyongjung/cns-gw/util"
)

//-- Session ID and pk are mapped
func UserKeySessionDb(w http.ResponseWriter, db *sql.DB, pk string) http.ResponseWriter {
	sessionId := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionId,
		Path:  "/",
	}
	cookie.MaxAge = sessionTime

	http.SetCookie(w, cookie)
	dbExe, err := db.Prepare("insert into session values (?, ?, ?)")
	if err != nil {
		util.LogErr("db prepare err : ", err)
	}
	defer dbExe.Close()

	result, err := dbExe.Exec(cookie.Value, pk, time.Now())
	if err != nil {
		util.LogErr("db exec err : ", err)
	} else {
		lastInsertId, _ := result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		util.LogGw("DB exec result lastInsertId : ", lastInsertId)
		util.LogGw("DB exec result rowsAffected : ", rowsAffected)
	}

	return w
}

//-- Check session existing
func CheckSession(db *sql.DB, sessionId string) string {
	var pk string

	queryResult, err := db.Query("select pk from session where session_id = ?", sessionId)
	if err != nil {
		util.LogErr("db query err : ", err)
	}
	defer queryResult.Close()

	for queryResult.Next() {
		err = queryResult.Scan(&pk)
		if err != nil {
			util.LogErr("db query scan err : ", err)
		}
	}

	if pk == "" {
		util.LogGw("No session")
		return ""
	}

	return pk
}

//-- Delete session
func DelSession(db *sql.DB, sessionId string) {
	dbExe, err := db.Prepare("delete from session where session_id = ?")
	if err != nil {
		util.LogErr(err)
	}

	result, err := dbExe.Exec(sessionId)
	if err != nil {
		util.LogErr(err)
	} else {
		lastInsertId, _ := result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		util.LogGw("Delete session of DB")
		util.LogGw("DB exec result lastInsertId : ", lastInsertId)
		util.LogGw("DB exec result rowsAffected : ", rowsAffected)
	}
}
