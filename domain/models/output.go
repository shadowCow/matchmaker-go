package models

type FindingMatch struct {
	PlayerId string `json:"playerId"`
}

type MatchMade struct {
	Player1Id string `json:"player1Id"`
	Player2Id string `json:"player2Id"`
}

type LeftMatchmaking struct {
	PlayerId string `json:"playerId"`
}

type FindMatchOutput interface {
	findMatchOutput()
}
func (f FindingMatch) findMatchOutput() {}
func (m MatchMade) findMatchOutput() {}

type LeaveMatchmakingOutput interface {
	leaveMatchmakingOutput()
}
func (l LeftMatchmaking) leaveMatchmakingOutput() {}
