package mail

type Mail struct {
	To       []string `json:"to"`
	Cc       []string `json:"cc,omitempty"`
	From     string   `json:"from"`
	Sender   string   `json:"sender"`
	Subject  string   `json:"subject"`
	Template string   `json:"template"`
}
