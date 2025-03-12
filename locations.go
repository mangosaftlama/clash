package clash

type Location struct {
	LocalizedName string `json:"localizedName"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	IsCountry     bool   `json:"isCountry"`
	CountryCode   string `json:"countryCode"`
}
