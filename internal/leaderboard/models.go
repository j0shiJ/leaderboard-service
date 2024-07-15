// internal/leaderboard/models.go
package leaderboard

// import "gorm.io/gorm"

type User struct {
	// gorm.Model
	UserName string  `gorm:"unique;not null"`
	Country  string  `gorm:"not null"`
	State    string  `gorm:"not null"`
	Score    float64 `gorm:"not null"`
	Game     string  `gorm:"not null"`
}
