package services

import "github.com/shadowCow/matchmaker-go/domain/models"

type MatchmakingService struct {
	queuedPlayerId string
}
func NewMatchmakingService() *MatchmakingService {
	return &MatchmakingService{}
}

func (m *MatchmakingService) FindMatch(playerId string) (models.FindMatchOutput, error) {
	if m.queuedPlayerId == "" {
		m.queuedPlayerId = playerId

		return models.FindingMatch{ PlayerId: playerId }, nil
	} else if m.queuedPlayerId == playerId {
		return models.FindingMatch{ PlayerId: playerId }, nil
	} else {
		output := models.MatchMade{
			Player1Id: m.queuedPlayerId,
			Player2Id: playerId,
		}
		
		m.queuedPlayerId = ""

		return output, nil
	}
}

func (m *MatchmakingService) LeaveMatchmaking(playerId string) (models.LeaveMatchmakingOutput, error) {
	if m.queuedPlayerId == playerId {
		m.queuedPlayerId = ""
	}

	return models.LeftMatchmaking{ PlayerId: playerId }, nil
}
