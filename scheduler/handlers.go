package main

import (
	"net/http"
	"video_server/scheduler/dbops"
	"github.com/julienschmidt/httprouter"
)

func vidDelReHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := 	p.ByName("vid_id")
	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}

