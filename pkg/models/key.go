package models

// Key represents an access key associated with a user.
type Key struct {
    Name     string `json:"name" db:"name"`
    Secret   string `json:"secret"` // Secret should not be stored in DB, handled in services
    Username string `json:"username" db:"username"` // Foreign key to User
}
