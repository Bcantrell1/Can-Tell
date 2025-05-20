package verifier

import (
	"time"
)

type Verifier struct {
	useSMTP   bool
	fromEmail string
	fromName  string
	schedule  *schedule

	connectTimeout   time.Duration
	operationTimeout time.Duration
}

type Result struct {
	Email      string    `json:"email"`
	Obtainable string    `json:"obtainable"`
	Structure  Structure `json:"structure"`
	SMTP       *SMTP     `json:"smtp"`
	Disposable bool      `json:"disposable"`
	Free       bool      `json:"free"`
	MxRecords  bool      `json:"mx_records"`
}

func NewVerifier() *Verifier {
	return &Verifier{
		fromEmail:        defaultEmail,
		fromName:         defaultName,
		connectTimeout:   10 * time.Second,
		operationTimeout: 10 * time.Second,
	}
}

func (ver *Verifier) Validate(email string) (*Result, error) {

	res := Result{
		Email:      email,
		Obtainable: obtainableUnknown,
	}

	structure := ver.ParseEmail(email)
	res.Structure = structure
	if !structure.Valid {
		return &res, nil
	}

	return &res, nil
}
