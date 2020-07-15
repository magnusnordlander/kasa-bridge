package main

import (
	"os"
//	"fmt"
	"net/http"
    "mime"
    "mime/multipart"
    "io/ioutil"
//    "time"
//    "strconv"
//    "log"
)


func check(e error) {
    if e != nil {
    	os.Stderr.WriteString("Panic?")
        panic(e)
    }
}

func generateImage() {
	req, _ := http.NewRequest("GET", "http://127.0.0.1:42243/https/stream/mixed?video=mpjpeg&audio=aac&resolution=hd", nil)
	resp, _ := http.DefaultClient.Do(req)
	_, params, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	mr := multipart.NewReader(resp.Body, params["boundary"])
	for part, err := mr.NextPart(); err == nil; part, err = mr.NextPart() {
//		log.Println("Part. Type: "+part.Header.Get("Content-Type"))

		if part.Header.Get("Content-Type") != "video/x-h264" {
//			log.Println("Not h264")
			continue
		}

		value, _ := ioutil.ReadAll(part)
		_, err := os.Stdout.Write(value)
    	check(err)
	}
}


func main() {
//    fmt.Println("running")
    generateImage()
//    http.HandleFunc("/", handler)
//    log.Fatal(http.ListenAndServe(":8080", nil))
}
