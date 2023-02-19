package repository

import (
	"database/sql"
	"mini-project/structs"
	"time"
)

func getTime() string {
	currentTime := time.Now()

	sqlTime := currentTime.Format("2006-01-02 15:04:05")

	return sqlTime
}
func GetCat(db *sql.DB) (err error, results []structs.Category) {
	s := `SELECT * FROM category`

	rows, err := db.Query(s)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}
	return
}
func PostCat(db *sql.DB, category structs.Category) (err error) {
	time := getTime()
	sql := `INSERT INTO category (name, created_at, updated_at)
			VALUES ($1, $2, $3)`

	errs := db.QueryRow(sql, category.Name, time, time)

	return errs.Err()
}
func UpdateCat(db *sql.DB, category structs.Category) (err error) {
	sql := `UPDATE category SET name = $1, updated_at =$2 WHERE id= $3`

	errs := db.QueryRow(sql, category.Name, getTime(), category.Id)

	return errs.Err()
}
func DeleteCat(db *sql.DB, category structs.Category) (err error) {
	sql := `DELETE FROM category WHERE id=$1`

	errs := db.QueryRow(sql, category.Id)

	return errs.Err()
}
func GetCatId(db *sql.DB, books structs.Books) (err error, result []structs.Books) {
	sql := `SELECT * FROM books WHERE category_id=$1`

	rows, err := db.Query(sql, books.CategoryId)
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
