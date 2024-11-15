package repository

import (
	"database/sql"
	"fmt"
	"jumpStart-backEnd/entities"
	"log"
	"time"

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

func (repo *ShareRepository) FindAllShares() ([]entities.Share, error) {
	numberSharesPerQuery, err := repo.DetermineNumberRowsToSearch()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	query := fmt.Sprintf(`SELECT * FROM tb_share as ts WHERE ts.dateShare = '2024-10-12' ORDER BY ts.dateShare DESC LIMIT %d`, numberSharesPerQuery)
	rows, err := repo.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return buildShare(rows)

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

func (repo *ShareRepository) ListSharesBasedOffSet(offset int) ([]entities.Share, error) {
	numberSharesPerQuery, err := repo.DetermineNumberRowsToSearch()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	offset *= 10;
	query := fmt.Sprintf(`SELECT * FROM tb_share as ts  ORDER BY ts.id DESC LIMIT %d OFFSET %d`, numberSharesPerQuery, offset)
	rows, err := repo.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return buildShare(rows)

}

func buildShare(rows *sql.Rows) ([]entities.Share, error) {
	var shares []entities.Share

	for rows.Next() {
		var dateShare time.Time
		var share entities.Share
		err := rows.Scan(&share.Id, &share.NameShare, &dateShare, &share.OpenShare, &share.HighShare, &share.LowShare, &share.CloseShare, &share.VolumeShare)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		share.DateShare = dateShare.Format("2006-01-02")

		shares = append(shares, share)

	}

	return shares, nil

}
