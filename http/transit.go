package http

import (
	"bytes"
	"log"
	"net/http"
)

func XRequest(client *http.Client, url string) string {
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error requesting: %s", err.Error())
		return RenderErrDto(err.Error())
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return RenderDataDto(buf.String())
}
