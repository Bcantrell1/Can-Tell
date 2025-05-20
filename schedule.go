package verifier

import (
	"time"
)

type schedule struct {
	stopChan  chan struct{}
	jobFunc   interface{}
	jobParams []interface{}
	tick      *time.Ticker
	running   bool
}
