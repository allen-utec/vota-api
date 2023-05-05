package infrastructure

type poll struct {
	ID       string       `json:"id"`
	Question string       `json:"question"`
	Options  []pollOption `json:"options"`
	UserID   uint         `json:"user_id"`
}

type pollOption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
