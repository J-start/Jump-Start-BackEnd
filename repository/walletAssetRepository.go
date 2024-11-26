package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"

	_ "github.com/go-sql-driver/mysql"
)

type WalletAssetRepository struct {
	db *sql.DB
}

func NewWalletAssetRepository(db *sql.DB) *WalletAssetRepository {
	return &WalletAssetRepository{db: db}
}

func (war *WalletAssetRepository) FindAssetWallet(assetName string, idWallet int) (entities.WalletAsset, error) {

	var walletAsset entities.WalletAsset
	query := `SELECT * FROM tb_walletAsset WHERE assetName = ? AND idWallet = ?`

	err := war.db.QueryRow(query, assetName,idWallet).Scan(&walletAsset.Id, &walletAsset.AssetName, &walletAsset.AssetType, &walletAsset.AssetQuantity, &walletAsset.IdWallet)

	if err != nil {
		if err == sql.ErrNoRows {
			return entities.WalletAsset{}, errors.New("ativo não existe na carteira")
		} 
		return entities.WalletAsset{}, err
	}

	return walletAsset, nil
}

func (war *WalletAssetRepository) InsertAssetIntoWallet(walletAsset entities.WalletAsset,repositoryService *sql.Tx) error {
	tx := repositoryService

	query := `INSERT INTO tb_walletAsset(assetName,assetType,assetQuantity,idWallet) VALUES (?,?,?,?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(walletAsset.AssetName, walletAsset.AssetType, walletAsset.AssetQuantity, walletAsset.IdWallet)
	if err != nil {
		return err
	}

	return nil
}

func (war *WalletAssetRepository) UpdateAssetIntoWallet(newQuantity float64,idWallet int,repositoryService *sql.Tx) error {
	tx := repositoryService

	query := `UPDATE tb_walletAsset SET assetQuantity = ? WHERE idWalletAsset = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newQuantity, idWallet)
	if err != nil {
		return err
	}

	return nil
}

// func (war *WalletAssetRepository) UpdateAssetIntoWalletWithTx(tx *sql.Tx, newQuantity float64, idWallet int) error {
// 	query := `UPDATE tb_walletAsset SET assetQuantity = ? WHERE idWallet = ?`
// 	_, err := tx.Exec(query, newQuantity, idWallet)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }


func (war *WalletAssetRepository) DeleteAssetWallet(idWallet int)  error {
	query := fmt.Sprintf(`DELETE FROM tb_walletAsset WHERE idWalletAsset = %d `, idWallet)
	_, err := war.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
