package model

type Image struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Random bool   `json:random`
}

type ImageResponse struct {
	Error  string  `json:"error"`
	Images []Image `json:"images"`
}
