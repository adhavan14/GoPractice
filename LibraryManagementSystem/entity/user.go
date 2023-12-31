package entity


type User struct {

	username string
	password string
	name string
}

func NewUser(username , password, name string) *User {
	return &User{
		username: username,
		password: password,
		name: name,
	}
}

func (user *User) AuthenticateUser(username, password string) bool {
	return user.username == username && user.password == password
}

func (user *User) Equals(currentUser *User) bool {
	return user.username == currentUser.username
}

func (user *User) ToStringInUser() string {
	return "Username : " + user.username + " Name : " + user.name
}