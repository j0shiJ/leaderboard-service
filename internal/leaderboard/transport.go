// internal/leaderboard/transport.go
package leaderboard

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	r := http.NewServeMux()
	r.Handle("/submit", httptransport.NewServer(
		endpoints.SubmitScoreEndpoint,
		decodeSubmitScoreRequest,
		encodeResponse,
	))
	r.Handle("/get_rank", httptransport.NewServer(
		endpoints.GetRankEndpoint,
		decodeGetRankRequest,
		encodeResponse,
	))
	r.Handle("/list_top_n", httptransport.NewServer(
		endpoints.ListTopNEndpoint,
		decodeListTopNRequest,
		encodeResponse,
	))
	return r
}

func decodeSubmitScoreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req SubmitScoreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetRankRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetRankRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeListTopNRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req ListTopNRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
