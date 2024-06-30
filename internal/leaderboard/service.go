// internal/leaderboard/service.go
package leaderboard

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type LeaderboardService interface {
	SubmitScore(ctx context.Context, user User) error
	GetRank(ctx context.Context, userName, scope string) (int, error)
	ListTopN(ctx context.Context, n int, scope, country, state string) ([]User, error)
}

type leaderboardService struct {
	db *gorm.DB
}

func NewLeaderboardService(db *gorm.DB) LeaderboardService {
	return &leaderboardService{db: db}
}

func (s *leaderboardService) SubmitScore(ctx context.Context, user User) error {
	if err := s.db.Create(&user).Error; err != nil {
		log.Printf("Error in SubmitScore: %v", err)
		return err
	}
	return nil
}

func (s *leaderboardService) GetRank(ctx context.Context, userName, scope string) (int, error) {
	var users []User
	var user User
	var err error

	if err = s.db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		log.Printf("Error finding user in GetRank: %v", err)
		return 0, err
	}

	switch scope {
	case "global":
		err = s.db.Order("score desc").Find(&users).Error
	case "country":
		err = s.db.Order("score desc").Where("country = ?", user.Country).Find(&users).Error
	case "state":
		err = s.db.Order("score desc").Where("state = ?", user.State).Find(&users).Error
	}

	if err != nil {
		log.Printf("Error fetching users in GetRank: %v", err)
		return 0, err
	}

	for i, u := range users {
		if u.UserName == userName {
			return i + 1, nil
		}
	}

	return 0, nil
}

func (s *leaderboardService) ListTopN(ctx context.Context, n int, scope, country, state string) ([]User, error) {
	var users []User
	var err error

	switch scope {
	case "global":
		err = s.db.Order("score desc").Limit(n).Find(&users).Error
	case "country":
		err = s.db.Order("score desc").Where("country = ?", country).Limit(n).Find(&users).Error
	case "state":
		err = s.db.Order("score desc").Where("state = ?", state).Limit(n).Find(&users).Error
	}

	if err != nil {
		log.Printf("Error fetching top N users in ListTopN: %v", err)
	}

	return users, err
}
