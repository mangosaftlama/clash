package clash

type League struct {
	// JsonLocalizedName
	Name     string `json:"name"`
	Id       int    `json:"id"`
	IconUrls any    `json:"iconUrls"`
}

type BuilderBaseLeague struct {
	// JsonLocalizedName
	Name string `json:"name"`
	Id   int    `json:"id"`
}
