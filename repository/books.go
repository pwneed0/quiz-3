package repository

import (
	"database/sql"
	"mini-project/structs"
	"time"
)

func GetBooks(db *sql.DB) (err error, result []structs.Books) {
	sql := `SELECT * FROM books`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var books structs.Books

		err = rows.Scan(&books.Id, &books.Title, &books.Description, &books.ImageUrl, &books.ReleaseYear, &books.Price, &books.TotalPage, &books.Thickness, &books.CreatedAt, &books.UpdatedAt, &books.CategoryId)
		if err != nil {
			panic(err)
		}

		result = append(result, books)
	}
	return

}
func PostBooks(db *sql.DB, books structs.Books) (err error) {
	s := `INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, updated_at, created_at, thickness)
		  VALUES ($1, $2, $3, $4, $5 , $6, $7, $8, $9, $10)`

	errs := db.QueryRow(s,
		books.Title,
		books.Description,
		books.ImageUrl,
		books.ReleaseYear,
		books.Price,
		books.TotalPage,
		books.CategoryId,
		time.Now(),
		time.Now(),
		books.Thickness,
	)

	return errs.Err()

}
func UpdateBooks(db *sql.DB, books structs.Books) (err error) {
	s := `UPDATE books SET title=$1, description=$2, image_url=$3, release_year = $4, price = $5, total_page = $6, category_id = $7 WHERE id = $8`

	errs := db.QueryRow(s,
		books.Title,
		books.Description,
		books.ImageUrl,
		books.ReleaseYear,
		books.Price,
		books.TotalPage,
		books.CategoryId,
		books.Id,
	)

	return errs.Err()
}
func DeleteBooks(db *sql.DB, books structs.Books) (err error) {
	sql := `DELETE FROM category WHERE id=$1`

	errs := db.QueryRow(sql, books.Id)

	return errs.Err()
}
