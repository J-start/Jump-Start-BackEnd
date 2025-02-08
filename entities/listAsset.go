package entities


type ListAsset struct {
	IdList int `json:"id"`
	NameAsset string `json:"name"`
	AcronymAsset string `json:"acronym"`
	UrlImage string `json:"urlImage"`
	TypeAsset string `json:"typeAsset"`
}


