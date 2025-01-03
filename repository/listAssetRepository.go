package repository

import (
	"database/sql"
	"fmt"
	"jumpStart-backEnd/entities"

	_ "github.com/go-sql-driver/mysql"
)


type ListAssetRepository struct {
	db *sql.DB
}


func NewListAssetRepository(db *sql.DB) *ListAssetRepository {
	return &ListAssetRepository{db: db}
}

func (repo *ListAssetRepository) ListAsset(asset string) ([]entities.ListAsset,error) {
	query := fmt.Sprintf(`SELECT * FROM list_asset WHERE typeAsset =  '%s' `, asset)
	rows, err := repo.db.Query(query)
	if err != nil {
		return []entities.ListAsset{}, err
	}
	defer rows.Close()
	listAsset := []entities.ListAsset{}
	for rows.Next() {
		asset := entities.ListAsset{}

		err := rows.Scan(&asset.IdList, &asset.NameAsset, &asset.AcronymAsset, &asset.UrlImage, &asset.TypeAsset)
		if err != nil {
			return []entities.ListAsset{},err
		}
		listAsset = append(listAsset, asset)
	}



	return listAsset,nil
}

func (repo *ListAssetRepository) ListAssetRequest(asset string) ([]string,error) {
	query := fmt.Sprintf(`SELECT acronymAsset FROM list_asset WHERE typeAsset =  '%s' `, asset)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listAssetRequest []string
	for rows.Next() {
		var asset string
		err := rows.Scan(&asset)
		if err != nil {
			return nil,err
		}
		listAssetRequest = append(listAssetRequest, asset)
	}

	return listAssetRequest,nil
}

