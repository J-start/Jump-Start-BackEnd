package news_repository

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestFindAllShares(t *testing.T) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/jumpStartTest?parseTime=true")
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("TRUNCATE TABLE tb_news")
	if err != nil {
		t.Fatalf("failed to truncate table: %v", err)
	}
	_, err = db.Exec(`
			INSERT INTO tb_news (id,news,dateNews,datePublished,isApproved) VALUES
  			(1,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1),
    		(2,'{"COIN": {"description": "descricao2", "url": "url2"}}', '2024-12-23','2024-12-19', 1),
    		(3,'{"CRYPTO": {"description": "descricao3", "url": "url3"}}', '2024-12-23','2024-12-19', 1),
    		(4,'{"SHARE": {"description": "descricao4", "url": "url4"}}', '2024-12-23','2024-12-19', 1),
			(5,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1),
    		(6,'{"COIN": {"description": "descricao2", "url": "url2"}}', '2024-12-23','2024-12-19', 1),
    		(7,'{"CRYPTO": {"description": "descricao3", "url": "url3"}}', '2024-12-23','2024-12-19', 1),
    		(8,'{"SHARE": {"description": "descricao4", "url": "url4"}}', '2024-12-23','2024-12-19', 1),
			(9,'{"SHARE": {"description": "descricao4", "url": "url4"}}', '2024-12-23','2024-12-19', 1),
			(10,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1),
			(11,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1);
    		
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}


	repo := NewNewsRepository(db)
	listNews, err := repo.FindAllNews(0)
	assert.NoError(t, err)
	assert.Equal(t, len(listNews), 10) 
	assert.Equal(t,listNews[0].Id, 1)
	assert.Equal(t,listNews[1].Id, 2)
	assert.Equal(t,listNews[2].Id, 3)
	assert.Equal(t,listNews[3].Id, 4)
	assert.Equal(t,listNews[9].Id, 10)
}

func TestFindAllSharesWithOffSet(t *testing.T) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/jumpStartTest?parseTime=true")
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("TRUNCATE TABLE tb_news")
	if err != nil {
		t.Fatalf("failed to truncate table: %v", err)
	}
	_, err = db.Exec(`
		INSERT INTO tb_news (id,news,dateNews,datePublished,isApproved) VALUES
			(1,'{"COIN": {"description": "descricao2", "url": "url2"}}', '2024-12-23','2024-12-19', 1),
    		(2,'{"CRYPTO": {"description": "descricao3", "url": "url3"}}', '2024-12-23','2024-12-19', 1),
    		(3,'{"SHARE": {"description": "descricao4", "url": "url4"}}', '2024-12-23','2024-12-19', 1),
			(4,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1),
    		(5,'{"COIN": {"description": "descricao2", "url": "url2"}}', '2024-12-23','2024-12-19', 1),
    		(6,'{"CRYPTO": {"description": "descricao3", "url": "url3"}}', '2024-12-23','2024-12-19', 1),
    		(7,'{"SHARE": {"description": "descricao4", "url": "url4"}}', '2024-12-23','2024-12-19', 1),
			(8,'{"SHARE": {"description": "descricao4", "url": "url4"}}', '2024-12-23','2024-12-19', 1),
			(9,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1),
			(10,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1),
    		(11,'{"COIN": {"description": "descricao2", "url": "url2"}}', '2024-12-24','2024-12-19', 1),
    		(12,'{"CRYPTO": {"description": "descricao3", "url": "url3"}}', '2024-12-24','2024-12-19', 1),
    		(13,'{"SHARE": {"description": "descricao4", "url": "url4"}}', '2024-12-24','2024-12-19', 1);
    		
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}


	repo := NewNewsRepository(db)
	listNews, err := repo.FindAllNews(1)
	fmt.Println(listNews)
	assert.NoError(t, err)
	assert.Equal(t, len(listNews), 3) 
	assert.Equal(t,listNews[0].Id, 11)
}

func TestDeleteNews(t *testing.T) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/jumpStartTest?parseTime=true")
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("TRUNCATE TABLE tb_news")
	if err != nil {
		t.Fatalf("failed to truncate table: %v", err)
	}
	_, err = db.Exec(`
			INSERT INTO tb_news (id,news,dateNews,datePublished,isApproved) VALUES
  			(1,'{"SHARE": {"description": "descricao1", "url": "url1"}}', '2024-12-23','2024-12-19', 1),
    		(2,'{"COIN": {"description": "descricao2", "url": "url2"}}', '2024-12-23','2024-12-19', 1);
    		
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}


	repo := NewNewsRepository(db)
	err2 := repo.DeleteNews(1)
	assert.NoError(t, err2)

}