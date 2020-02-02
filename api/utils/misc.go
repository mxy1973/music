package utils

import (
	"crypto/rand"
	"io"
	"fmt"
	"avenssi/config"
	"log"
	"net/http"
)

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x", uuid[0:4], uuid[4:6]), err
}

func SendDeleteVideoRequest(vid string)  {
	addr := config.GetLBAddr() + ":9001/"
	url := "http://" + addr + "/video-delete-record/" + vid

	_, err := http.Get(url)
	if err != nil {
		log.Printf("send deleting video error: %s", err)
	}
}


