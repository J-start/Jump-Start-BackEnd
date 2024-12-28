package share_repository

import (
	"database/sql"
	"fmt"
	"jumpStart-backEnd/entities"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ShareRepository struct {
	db *sql.DB
}

func NewShareRepository(db *sql.DB) *ShareRepository {
	return &ShareRepository{db: db}
}

func (repo *ShareRepository) FindAllShares() ([]entities.Share, error) {

	rows, err := repo.db.Query(`SELECT ts.* FROM tb_share ts WHERE ts.id IN ( SELECT MAX(id) FROM tb_share GROUP BY nameShare ) ORDER BY ts.nameShare`)

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

func (repo *ShareRepository) FindShareById(shareName string) (entities.Share, error) {
	var share []entities.Share
	query := fmt.Sprintf(`SELECT * FROM tb_share WHERE nameShare = '%s' ORDER BY dateShare DESC LIMIT 1`, shareName)
	rows, err := repo.db.Query(query)

	if err != nil {
		return entities.Share{}, err
	}

	 defer rows.Close()

	share, err = buildShare(rows)
	if err != nil {
		return entities.Share{}, err
	}

	if len(share) == 0 {
		return entities.Share{}, nil
	}
	return share[0],nil
}

func (repo *ShareRepository) ShareNameList() ([]entities.Share, error) {

	rows, err := repo.db.Query(`SELECT DISTINCT * FROM tb_share`)

	if err != nil {
		log.Fatal(err)
	}
	return buildShare(rows)
}

func (repo *ShareRepository) ShareList(shareName string) ([]entities.Share, error) {
	var share []entities.Share
    const LIMIT = 40
	query := fmt.Sprintf(`SELECT * FROM tb_share WHERE nameShare = '%s' AND id IN (SELECT MIN(id) FROM tb_share WHERE nameShare = '%s' GROUP BY dateShare) ORDER BY DateShare LIMIT %d`, shareName,shareName,LIMIT)
	rows, err := repo.db.Query(query)

	if err != nil {
		return []entities.Share{}, err
	}

	 defer rows.Close()

	share, err = buildShare(rows)
	if err != nil {
		return []entities.Share{}, err
	}

	if len(share) == 0 {
		return []entities.Share{}, nil
	}
	return share,nil
}

func buildShare(rows *sql.Rows) ([]entities.Share, error) {
	var shares []entities.Share

	for rows.Next() {
		var dateShare time.Time
		var share entities.Share
		err := rows.Scan(&share.Id, &share.NameShare, &dateShare, &share.OpenShare, &share.HighShare, &share.LowShare, &share.CloseShare, &share.VolumeShare)
		if err != nil {
			return nil, err
		}
		share.DateShare = dateShare.Format("02-01-2006")

		shares = append(shares, share)

	}

	return shares, nil

}


