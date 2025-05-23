package storage

var (
	MapByID     = map[int64]Quotes{}
	MapByAuthor = map[string]Quotes{}
)

type Quotes struct {
	ID     int64
	Author string `json: "author"`
	Quote  string `json: "quote"`
}
