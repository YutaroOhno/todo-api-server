package ports

type TodoInputPort struct {
	Title string
	Text string
}

type TodoOutputPort struct {
	Title string `json:"title"`
	Text string `json:"text"`
}