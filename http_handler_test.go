package gorevisit

import (
	"bytes"
	"image"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func echoService(img image.Image) (image.Image, error) {
	return img, nil
}

func TestRevisitHandlerPost(t *testing.T) {
	msg, err := NewAPIMsgFromFiles("./fixtures/bob.jpg", "./fixtures/scream.ogg")
	if err != nil {
		t.Fatal(err)
	}

	jsonBytes, err := msg.JSON()
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "http://whatever", bytes.NewReader(jsonBytes))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonBytes)))

	w := httptest.NewRecorder()

	service := NewRevisitService(echoService)
	service.TransformationHandler(w, req)

	if w.Code != http.StatusAccepted {
		t.Errorf("expected %d status, received %d", http.StatusAccepted, w.Code)
	}
}
