package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"restful/entity"
)

var postgresDatasource *sql.DB

func ConnectToPostgres() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "nasrin", "1234", "restful") //net search

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Connected to postgres database.")
	postgresDatasource = db
	return nil
}

func FetchAllTestTableRecs(ctx context.Context) ([]entity.TestTableEntity, error) {
	entities := make([]entity.TestTableEntity, 0)

	rows, err := postgresDatasource.QueryContext(ctx, "SELECT * FROM restful.restful_persistence_1.test_table") //time out
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := entity.TestTableEntity{}
		err := rows.Scan(&temp.Col1, &temp.Col2)
		if err != nil {
			return nil, err
		}
		entities = append(entities, temp)
	}
	return entities, nil
}
