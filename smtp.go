package verifier

import (
	"time"
)

type SMTP struct {
	HasHost     bool `json:"has_host"`
	InboxFull   bool `json:"inbox_full"`
	Deliverable bool `json:"deliverable"`
	Banned      bool `json:"banned"`
}

func (v *Verifier) CheckWithSMTP(domainName, username string) (*SMTP, error) {
	if !v.useSMTP {
		return nil, nil
	}
	var smtp SMTP
	var err error

	// Check any SMTP server that is available
	client, err := newSMTPClient(domainName, v.connectTimeout, v.operationTimeout)
	if err != nil {
		return nil, err
	}

	defer client.Close()

	smtp.HasHost = true

	if username == "" {
		return &smtp, nil
	}

	return &smtp, nil
}

func newSMTPClient(domainName string, connectTimeout, operationTimeout time.Duration) (*smtpClient, error) {
	return nil, nil
}
