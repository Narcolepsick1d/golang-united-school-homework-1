package structs

type User struct {
	firstName string
	lastName string
}
type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func (u *User) SetFirstName(name string) {
	u.firstName = name
}
func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}
func (u *User) FullName() string {
	return u.lastName + " " + u.firstName
}
func NewUser() User {
	return User{}
}
func ResetUser(input *User) {
	input.SetFirstName("")
	input.SetLastName("")
}
func IsUser(input interface{}) bool {
	if input == nil  {
		return false
	}
	switch input.(type) {
	case User:
		return true
	default:
		return false
	}
}
func ProcessUser(u UserInterface) string {
	u.SetFirstName("John")
	u.SetLastName("Doe")
	
	return u.FullName()
}
