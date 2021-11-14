package model

type Color struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Random bool   `json:"random"`
}

type ColorResponse struct {
	Error  string  `json:"error"`
	Colors []Color `json:"colors"`
}
