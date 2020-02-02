package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("testPageHandler in")
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}

func streamHeadler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Entered the streamHandler")
	targetUrl := "https://lb-videos-1258876329.cos.ap-chengdu.myqcloud.com/videos/" + p.ByName("vid_id")
	http.Redirect(w, r, targetUrl, 301)
}

func uploadHeadler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE);  err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Get formfile fied")
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	fn := p.ByName("vid_id")
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	cosfn := "videos/" + fn
	path := "./videos" +fn
	ret :=  UploadToCos(cosfn)
	if !ret {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	os.Remove(path)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"Uploaded successfully")
}

