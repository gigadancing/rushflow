package dbops

import (
	"database/sql"
	"log"
)

// api -> videoId -> mysql
// dispatcher -> mysql -> videoId -> datachannel
// executor -> datachannel ->videoId -> delete video

//
func ReadVideoDeleteRecord(count int) ([]string, error) {
	var (
		ids     []string
		stmtOut *sql.Stmt
		rows    *sql.Rows
		err     error
	)
	if stmtOut, err = dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?"); err != nil {
		return ids, err
	}

	if rows, err = stmtOut.Query(count); err != nil {
		log.Printf("query videoDeletionRecord error:%v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	defer stmtOut.Close()
	return ids, nil
}

//
func DelVideoDeletionRecord(vid string) error {
	var (
		stmtDel *sql.Stmt
		err     error
	)
	if stmtDel, err = dbConn.Prepare("SELECT FROM video_del_rec WHERE video_id=?"); err != nil {
		return err
	}

	if _, err = stmtDel.Exec(vid); err != nil {
		log.Printf("deleting videoDeletionRecord error:%v", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}
