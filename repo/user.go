package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

// repository pattern
type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, pass string) (*User, error)
}

// property
type userRepo struct {
	users []User
}

// constractor or Constractor function
func NewUserRepo() UserRepo {
	return &userRepo{}
}

// POST create product store to database
func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}

// Post user api
func (r *userRepo) Find(email, pass string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email && user.Password == pass {
			return &user, nil
		}
	}
	return nil, nil
}
