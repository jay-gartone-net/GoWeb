package models

//User is a cool person
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers is used to return all users
func GetUsers() []*User {
	return users
}

// AddUser adds a user
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
