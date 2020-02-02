package dbops

import (
	"database/sql"
	"log"
	"sync"
	"video_server/api/defs"
	"strconv"
)

func InserSession(sid string, ttl int64, uname string)  error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name,) VALUE (?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, uname)
	if err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string

		if er := rows.Scan(&id, &ttlstr, &login_name); er != nil {
			log.Printf("retrieve allsesions error: %s", err)
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{TTL: ttl, Username: login_name}
			m.Store(id, ss)
			log.Printf("session id: %s, ttl: %d", id, ss.TTL)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE  FROM sessions WHERE session_id=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	if _, err := stmtDel.Exec(sid); err != nil {
		return err
	}
	return nil
}