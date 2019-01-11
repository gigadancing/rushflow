package taskrunner

import (
	"github.com/kataras/iris/core/errors"
	"log"
	"os"
	"rushflow/video_server/scheduler/dbops"
	"sync"
)

//
func deleteVideo(vid string) error {
	if err := os.Remove(VIDEO_PATH + vid); err != nil && !os.IsNotExist(err) {
		log.Printf("deleting video error:%v", err)
		return err
	}
	return nil
}

//
func VideoClearDispatcher(dc dataChan) error {
	var (
		res []string
		err error
	)
	if res, err = dbops.ReadVideoDeleteRecord(3); err != nil {
		log.Printf("video clear dispatcher error:%v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("all tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

//
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var (
		err error
	)
ForLoop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(vid.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break ForLoop
		}
	}

	errMap.Range(func(k, v interface{}) bool {
		if err = v.(error); err != nil {
			return false
		}
		return true
	})

	return err
}
