package models

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsAdmin   bool
}
