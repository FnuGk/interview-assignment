package model

// User is the base model for an item in the users table
type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}
