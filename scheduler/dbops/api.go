package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 1. user -> api service -> delete video
// 2. api service -> scheduler -> write video deletion record
// 3. timer
// 4. timer -> runner -> read write video deletion record -> delete video from folder

//
//
func AddVideoDeletionRecord(vid string) error {
	var (
		stmtIns *sql.Stmt
		err     error
	)

	if stmtIns, err = dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)"); err != nil {
		return err
	}

	if _, err = stmtIns.Exec(vid); err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
