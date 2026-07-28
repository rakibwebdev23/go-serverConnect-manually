package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"img_url" db:"img_url"`
}

var ErrNotFound = sql.ErrNoRows

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

// property
type productRepo struct {
	db *sqlx.DB
}

// constractor or Constractor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
			INSERT INTO products (
				title,
				description,
				price,
				img_url
			)
			VALUES ($1, $2, $3, $4)
			RETURNING id;
		`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var productList []*Product

	query := `
		SELECT
			id, 
			title,
			description,
			price,
			img_url
		from products
	`

	err := r.db.Select(&productList, query)
	if err != nil {
		return nil, err
	}
	return productList, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var product Product

	query := `
		SELECT
			id, 
			title,
			description,
			price,
			img_url
		from products
		WHERE id = $1
	`

	err := r.db.Get(&product, query, id)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	query := `
		UPDATE products 
		SET title=$1, description=$2, price=$3, img_url=$4
		WHERE id = $5
	`

	res, err := r.db.Exec(query, product.Title, product.Description, product.Price, product.ImgUrl, product.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrNotFound
	}

	return &product, nil
}

func (r *productRepo) Delete(id int) error {
	query := `
		DELETE FROM products WHERE id = $1
	`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
