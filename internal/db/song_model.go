package db

type Song struct {
	ID       int64  `json:"ID"`
	URL      string `json:"URL"`
	Category string `json:"Category"`
}
