package structs

import "time"

type Category struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Books struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryId  int       `json:"category_id"`
}
