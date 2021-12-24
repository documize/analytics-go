package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	recordEventURL = "http://localhost:4000/api-public/a/r" // https://api.zerabase.com/api-public
)

var client *http.Client

func init() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	client = &http.Client{Transport: tr}
}

// SendEvent to Zerabase.
func SendEvent(i interface{}) (err error) {
	data, err := json.Marshal(i)
	if err != nil {
		fmt.Println("data 1")
		return err
	}

	resp, err := client.Post(recordEventURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return errors.New("not found: TrackerID")
	case http.StatusBadRequest:
		return errors.New("missing: TrackerID value or epoch Date value")
	default:
		return errors.New("unknown error")
	}
}
