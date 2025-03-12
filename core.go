package clash

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
