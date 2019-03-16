package main

type getResponse struct {
	Exists bool   `json:"Exists"`
	URL    string `json:"URL"`
}

type shortenResponse struct {
	Success      bool   `json:"Success"`
	ShortenedURL string `json:"ShortenedURL"`
	Error        string `json:"Error"`
}
