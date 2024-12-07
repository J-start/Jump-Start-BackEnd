package dto


type ListAssetDTO struct {
	NameAsset string `json:"name"`
	AcronymAsset string `json:"acronym"`
	UrlImage string `json:"urlImage"`
	TypeAsset string `json:"typeAsset"`
}