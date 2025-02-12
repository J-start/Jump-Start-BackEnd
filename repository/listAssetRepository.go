package repository

import (
	"database/sql"
	"errors"
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

func(repo *ListAssetRepository) FetchListAssetsAdm() ([]entities.ListAsset,error){
	query := fmt.Sprintf(`SELECT * FROM list_asset ORDER BY typeAsset`)
	rows, err := repo.db.Query(query)
	if err != nil {
		return []entities.ListAsset{}, err
	}
	defer rows.Close()
	var listAssetRequest []entities.ListAsset
	for rows.Next() {
		var asset entities.ListAsset
		err := rows.Scan(&asset.IdList, &asset.NameAsset, &asset.AcronymAsset, &asset.UrlImage, &asset.TypeAsset)
		if err != nil {
			return []entities.ListAsset{},err
		}
		listAssetRequest = append(listAssetRequest, asset)
	}

	return listAssetRequest,nil
}

func(repo *ListAssetRepository) UpdateAssetImageUrlById(newUrl string,idAsset int) error {
	query := fmt.Sprintf(`UPDATE list_asset SET url_image = '%s' WHERE idList = %d`, newUrl, idAsset)
	_, err := repo.db.Query(query)
	if err != nil {
		return errors.New("erro ao atualizar url, tente novamente")
	}
	return nil
}

func(repo *ListAssetRepository) InsertAsset(asset entities.NewAsset) error {
	query := fmt.Sprintf(`INSERT INTO list_asset (nameAsset,acronymAsset,url_image,typeAsset) VALUES ('%s','%s','%s','%s')`, asset.NameAsset, asset.AcronymAsset, asset.UrlImage, asset.TypeAsset)
	_, err := repo.db.Query(query)
	if err != nil {
		return errors.New("erro ao adicionar asset")
	}
	return nil
}
