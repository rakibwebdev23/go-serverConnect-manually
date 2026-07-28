package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

// repository pattern
type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, pass string) (*User, error)
}

// property
type userRepo struct {
	db *sqlx.DB
}

// constractor or Constractor function
func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

// POST create product store to database
func (r *userRepo) Create(user User) (*User, error) {

	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES (
		:first_name,
		:last_name,
		:email,
		:password,
		:is_shop_owner
		)
		RETURNING id
	`

	var userID int

	rows, err := r.db.NamedQuery(query, user)
	if err != nil{
		fmt.Println("server error", err)
		return nil, err
	}

	if rows.Next(){
		rows.Scan(&userID)
	}

	user.ID = userID
	return &user, nil
}

// Post user api
func (r *userRepo) Find(email, pass string) (*User, error) {

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		FROM users
		WHERE email = $1
		  AND password = $2
	`

	var user User

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		return nil, err
	}

	return &user, nil
}