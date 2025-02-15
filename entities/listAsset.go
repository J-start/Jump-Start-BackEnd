package entities


type ListAsset struct {
	IdList int `json:"id"`
	NameAsset string `json:"name"`
	AcronymAsset string `json:"acronym"`
	UrlImage string `json:"urlImage"`
	TypeAsset string `json:"typeAsset"`
}
type NewAsset struct {
	NameAsset string `json:"name"`
	AcronymAsset string `json:"acronym"`
	UrlImage string `json:"urlImage"`
	TypeAsset string `json:"typeAsset"`
}

type UpdateUrlImage struct {
	IdAsset int `json:"idAsset"`
	NewUrl string `json:"newUrl"`
}
type CryptoHistory struct {
	Value float64 `json:"value"`
	Date string `json:"date"`
}
