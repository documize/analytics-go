package event

import (
	"errors"
	"strings"
	"time"

	"github.com/documize/analytics-go/transport"
)

type Event struct {
	TrackerID string `json:"trackerId"`
	Date      int64  `json:"date"`
}

// Record event occurrence.
func Record(e Event) (err error) {
	// Validate data
	e.TrackerID = strings.TrimSpace(e.TrackerID)
	if len(e.TrackerID) == 0 {
		return errors.New("missing TrackerID")
	}
	if e.Date == 0 || time.Unix(e.Date, 0).IsZero() {
		return errors.New("missing Date (epoch value)")
	}

	err = transport.SendEvent(e)

	return nil
}
