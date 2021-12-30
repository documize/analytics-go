package metric

import (
	"errors"
	"strings"
	"time"

	"github.com/documize/zerabase-go/transport"
)

type Event struct {
	TrackerID string `json:"trackerId"`
	Date      int64  `json:"date"`
}

// TrackEvent records metric occurrence.
func TrackEvent(e Event) (err error) {
	// Validate data
	e.TrackerID = strings.TrimSpace(e.TrackerID)
	if len(e.TrackerID) == 0 {
		return errors.New("missing TrackerID")
	}
	if e.Date == 0 || time.Unix(e.Date, 0).IsZero() {
		e.Date = time.Now().UTC().Unix()
	}

	err = transport.SendEvent(e)

	return nil
}
