package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shadowCow/matchmaker-go/domain/services"
)

type ApiHttp struct {
	matchmakingService *services.MatchmakingService
}

func NewApiHttp(matchmakingService *services.MatchmakingService) *ApiHttp {
	return &ApiHttp{
		matchmakingService,
	}
}

func (a *ApiHttp) Start(port int) error {
	http.HandleFunc("/find-match", a.handleFindMatch)
	http.HandleFunc("/leave-matchmaking", a.handleLeaveMatchmaking)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (a *ApiHttp) handleFindMatch(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var f findMatchRequest

	err := json.NewDecoder(r.Body).Decode(&f)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	result, err := a.matchmakingService.FindMatch(f.PlayerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

type findMatchRequest struct {
	PlayerId string `json:"playerId"`
}

func (a *ApiHttp) handleLeaveMatchmaking(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var l leaveMatchmakingRequest

	err := json.NewDecoder(r.Body).Decode(&l)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	result, err := a.matchmakingService.LeaveMatchmaking(l.PlayerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

type leaveMatchmakingRequest struct {
	PlayerId string `json:"playerId"`
}
