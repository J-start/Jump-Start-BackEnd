package servicerepository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ServiceRepository struct {
	db *sql.DB
}


func NewWServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

 func (sr *ServiceRepository) StartTransaction() (*sql.Tx,error) {
 	tx, err := sr.db.Begin()
 	if err != nil {
 		return nil,err
 	}

	return tx, nil
 }

 func (sr *ServiceRepository) RollbackTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
 	if err != nil {
 		return err
 	}	
	return nil
}

 func (sr *ServiceRepository) CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
 	if err != nil {
 		return err
 	}
	return nil
}