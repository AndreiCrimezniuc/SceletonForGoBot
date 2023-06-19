package entities

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Message struct {
	From      User   `json:"from"`
	Date      int64  `json:"date"`
	MessageID int64  `json:"message_id"`
	Text      string `json:"Text"`
}

type Chat struct {
	ID int64 `json:"id"`
}
