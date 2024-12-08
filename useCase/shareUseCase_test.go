package usecase

import (
	"database/sql"
	"jumpStart-backEnd/repository"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestFindAllShares_Integration(t *testing.T) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/jumpStartTest?parseTime=true")
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("TRUNCATE TABLE tb_share")
	if err != nil {
		t.Fatalf("failed to truncate table: %v", err)
	}
	_, err = db.Exec(`
			INSERT INTO tb_share (id, nameShare, dateShare, openShare, highShare, lowShare, closeShare, volumeShare) VALUES
			(71,    'PETR4.SA',     '2024-10-12',    37.6,  37.65,  37.32, 37.62,   16343000),
			(72,    'BBAS3.SA',     '2024-10-12',    26.28, 26.46,  26.17, 26.33,   12175400),
			(73,    'ITSA4.SA',     '2024-10-12',    10.52, 10.54,  10.44, 10.47,   11660200),
			(75,    'VALE3.SA',     '2024-10-12',    60.99, 62.27,  60.98, 62.13, 20939400);
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}


	repo := repository.NewShareRepository(db)
	shares, err := repo.FindAllShares()
	assert.NoError(t, err)
	assert.Equal(t, len(shares), 4) 
	assert.Equal(t,shares[0].Id, 72)
	assert.Equal(t,shares[1].Id, 73)
	assert.Equal(t,shares[2].Id, 71)
	assert.Equal(t,shares[3].Id, 75)
}


