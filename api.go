package clash

import "net/http"

var (
	UserAgent = "mangosaftlama/clash"
)

type ClientError struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Detail  any    `json:"detail"`
}

type Client struct {
	apiKey string
}

type Label struct {
	// Label
	Name     string `json:"name"`
	Id       int    `json:"id"`
	IconUrls any    `json:"iconUrls"`
}

type LabelList []Label

type Language struct {
	Name         string `json:"name"`
	Id           int    `json:"id"`
	LanguageCode string `json:"languageCode"`
}

func (c Client) DefaultHeader() http.Header {
	header := make(http.Header)
	header.Add("Authorization", "Bearer "+c.apiKey)
	header.Add("Content-Type", "application/json")
	header.Add("User-Agent", UserAgent)
	return header
}

func NewClient(apiKey string) Client {
	return Client{apiKey: apiKey}
}
