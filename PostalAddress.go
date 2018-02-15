package conversationapi

type PostalAddress struct {
	Revision           float64  `json:"revision"`
	RegionCode         string   `json:"regionCode"`
	LanguageCode       string   `json:"LanguageCode"`
	PostalCode         string   `json:"PostalCode"`
	SortingCode        string   `json:"sortingCode"`
	AdministrativeArea string   `json:"administrativeArea"`
	Locality           string   `json:"locality"`
	SubLocality        string   `json:"sublocality"`
	AddressLines       []string `json:"addressLines"`
	Recipients         []string `json:"recipients"`
	Organization       []string `json:"organization"`
}
