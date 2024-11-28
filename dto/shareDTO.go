package dto


type ShareDTO struct {
	NameShare   string
	DateShare   string
	OpenShare   float64
	HighShare   float64
	LowShare    float64
	CloseShare  float64
	VolumeShare float64
}

func NewShareDTO(nameShare string, dateShare string, openShare float64, highShare float64, lowShare float64, closeShare float64, volumeShare float64) *ShareDTO {
	return &ShareDTO{
		NameShare:   nameShare,
		DateShare:   dateShare,
		OpenShare:   openShare,
		HighShare:   highShare,
		LowShare:    lowShare,
		CloseShare:  closeShare,
		VolumeShare: volumeShare,
	}
}
