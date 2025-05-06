package user

// public
type User struct {
	// public
	Name string

	// private (потому что lower case)
	age int
}

// public
func (u *User) ChangeAge(newAge int) {
	u.age = newAge
}

// private (потому что lower case)
func (u *User) changeName(newName string) {
	u.Name = newName
}
