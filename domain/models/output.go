package models

type FindingMatch struct {
	PlayerId string
}

type MatchMade struct {
	Player1Id string
	Player2Id string
}

type LeftMatchmaking struct {
	PlayerId string
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
