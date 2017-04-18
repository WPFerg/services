package structs

type Message struct {
	ID      int    `json:"id,omitempty"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
}
