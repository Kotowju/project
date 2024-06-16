package models

// User represents a user in the system.
type User struct {
    Username string `json:"username" db:"username"` // Username is unique and cannot be changed
    Email    string `json:"email" db:"email"`
}
