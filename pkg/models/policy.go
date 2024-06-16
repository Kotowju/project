package models

// Policy represents a policy that can be assigned to multiple users.
type Policy struct {
    Name       string `json:"name" db:"name"`
    Definition string `json:"definition" db:"definition"` // Define the operations allowed
}
