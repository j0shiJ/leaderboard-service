// internal/leaderboard/endpoints.go
package leaderboard

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SubmitScoreEndpoint endpoint.Endpoint
	GetRankEndpoint     endpoint.Endpoint
	ListTopNEndpoint    endpoint.Endpoint
}

func MakeEndpoints(s LeaderboardService) Endpoints {
	return Endpoints{
		SubmitScoreEndpoint: makeSubmitScoreEndpoint(s),
		GetRankEndpoint:     makeGetRankEndpoint(s),
		ListTopNEndpoint:    makeListTopNEndpoint(s),
	}
}

func makeSubmitScoreEndpoint(s LeaderboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SubmitScoreRequest)
		err := s.SubmitScore(ctx, req.User)
		if err != nil {
			return SubmitScoreResponse{Err: err.Error()}, nil
		}
		return SubmitScoreResponse{Err: ""}, nil
	}
}

func makeGetRankEndpoint(s LeaderboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRankRequest)
		rank, err := s.GetRank(ctx, req.UserName, req.Scope)
		if err != nil {
			return GetRankResponse{Rank: 0, Err: err.Error()}, nil
		}
		return GetRankResponse{Rank: rank, Err: ""}, nil
	}
}

func makeListTopNEndpoint(s LeaderboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListTopNRequest)
		users, err := s.ListTopN(ctx, req.N, req.Scope, req.Country, req.State)
		if err != nil {
			return ListTopNResponse{Users: nil, Err: err.Error()}, nil
		}
		return ListTopNResponse{Users: users, Err: ""}, nil
	}
}

type SubmitScoreRequest struct {
	User User
}

type SubmitScoreResponse struct {
	Err string `json:"err,omitempty"`
}

type GetRankRequest struct {
	UserName string
	Scope    string
}

type GetRankResponse struct {
	Rank int    `json:"rank,omitempty"`
	Err  string `json:"err,omitempty"`
}

type ListTopNRequest struct {
	N       int
	Scope   string
	Country string
	State   string
}

type ListTopNResponse struct {
	Users []User `json:"users,omitempty"`
	Err   string `json:"err,omitempty"`
}
