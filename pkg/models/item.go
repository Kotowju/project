package models

// Item represents an item in the system.
type Item struct {
    Sector      string `json:"sector" db:"sector"` // Sector cannot be changed
    Name        string `json:"name" db:"name"`     // Name cannot be changed
    Description string `json:"description" db:"description"`
}
