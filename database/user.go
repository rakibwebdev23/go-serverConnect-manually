package database 

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

var users []User

// POST create product store to database
func UserStore(u User) User {
	if u.ID != 0 {
		return  u
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

//Post user api 
func FindUser (email, pass string) *User {
	for _, u := range users {
		if u.Email == email && u.Password == pass {
			return &u
		}
	}
	return  nil
}