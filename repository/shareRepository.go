package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"jumpStart-backEnd/dto"

	_ "github.com/go-sql-driver/mysql"
)

type Share struct {
	Id          int
	NameShare   string
	DateShare   string
	OpenShare   float64
	HighShare   float64
	LowShare    float64
	CloseShare  float64
	VolumeShare float64
}

type ShareRepository struct {
	db *sql.DB
}

func NewShareRepository(db *sql.DB) *ShareRepository {
	return &ShareRepository{db: db}
}

func (repo *ShareRepository) FindAllShares() ([]dto.ShareDTO, error) {
	numberSharesPerQuery, err := repo.DetermineNumberRowsToSearch()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	query := fmt.Sprintf(`SELECT * FROM tb_share as ts WHERE ts.dateShare = DATE_FORMAT(NOW(), '%%Y-%%m-%%d') ORDER BY ts.dateShare DESC LIMIT %d`, numberSharesPerQuery)
	rows, err := repo.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var shares []dto.ShareDTO

	for rows.Next() {
		var dateShare time.Time
		var share Share
		err := rows.Scan(&share.Id, &share.NameShare, &dateShare, &share.OpenShare, &share.HighShare, &share.LowShare, &share.CloseShare, &share.VolumeShare)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		share.DateShare = dateShare.Format("2006-01-02")
        shareDTO := dto.NewShareDTO(share.NameShare, share.DateShare, share.OpenShare, share.HighShare, share.LowShare, share.CloseShare, share.VolumeShare)
		shares = append(shares, *shareDTO)

	}

	return shares, nil

}

func (repo *ShareRepository) DetermineNumberRowsToSearch() (int, error) {
	numberElements, err := repo.db.Query("SELECT COUNT(ts.id) FROM tb_share as ts WHERE ts.dateShare = DATE_FORMAT(NOW(), '%Y-%m-%d') ORDER BY ts.dateShare DESC LIMIT 10")

	if err != nil {
		log.Fatal(err)
	}
	defer numberElements.Close()

	var numberRows int

	for numberElements.Next() {
		err := numberElements.Scan(&numberRows)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}

	}

	return numberRows, nil
}
