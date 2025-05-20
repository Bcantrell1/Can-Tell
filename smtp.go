package verifier

type SMTP struct {
	HasHost     bool `json:"has_host"`
	InboxFull   bool `json:"inbox_full"`
	Deliverable bool `json:"deliverable"`
	Banned      bool `json:"banned"`
}
