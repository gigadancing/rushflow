package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rushflow/video_server/scheduler/dbops"
)

//
func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}
	if err := dbops.AddVideoDeletionRecord(vid); err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}
