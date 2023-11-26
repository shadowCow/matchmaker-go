package services_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/shadowCow/matchmaker-go/domain/models"
	"github.com/shadowCow/matchmaker-go/domain/services"
)

func TestMakeMatch(t *testing.T) {
	is := is.New(t)

	// Given
	player1Id := "p1"
	player2Id := "p2"
	service := services.NewMatchmakingService()

	// When
	gotP1, err := service.FindMatch(player1Id)

	// Then
	wantP1 := models.FindingMatch{ PlayerId: player1Id }
	is.NoErr(err)
	is.Equal(gotP1, wantP1)

	// When
	gotP2, err := service.FindMatch(player2Id)

	// Then
	wantP2 := models.MatchMade{
		Player1Id: player1Id,
		Player2Id: player2Id,
	}
	is.NoErr(err)
	is.Equal(gotP2, wantP2)
}

func TestFindMatchForSamePlayer(t *testing.T) {
	is := is.New(t)

	// Given
	player1Id := "p1"
	service := services.NewMatchmakingService()

	// When
	got1, err := service.FindMatch(player1Id)

	// Then
	want := models.FindingMatch{ PlayerId: player1Id }
	is.NoErr(err)
	is.Equal(got1, want)

	// When
	got2, err := service.FindMatch(player1Id)

	// Then
	is.NoErr(err)
	is.Equal(got2, want)
}

func TestLeaveMatchmaking(t *testing.T) {
	is := is.New(t)

	// Given
	player1Id := "p1"
	player2Id := "p2"
	service := services.NewMatchmakingService()

	// When
	gotFindMatchP1, err := service.FindMatch(player1Id)

	// Then
	wantFindMatchP1 := models.FindingMatch{ PlayerId: player1Id }
	is.NoErr(err)
	is.Equal(gotFindMatchP1, wantFindMatchP1)

	// When
	gotLeaveMatchmakingP1, err := service.LeaveMatchmaking(player1Id)

	// Then
	wantLeaveMatchmakingP1 := models.LeftMatchmaking{ PlayerId: player1Id }
	is.NoErr(err)
	is.Equal(gotLeaveMatchmakingP1, wantLeaveMatchmakingP1)

	// When
	gotFindMatchP2, err := service.FindMatch(player2Id)

	// Then
	wantFindMatchP2 := models.FindingMatch{ PlayerId: player2Id }
	is.NoErr(err)
	is.Equal(gotFindMatchP2, wantFindMatchP2)
}

func TestLeaveMatchmakingForDifferentPlayer(t *testing.T) {
	is := is.New(t)

	// Given
	player1Id := "p1"
	player2Id := "p2"
	service := services.NewMatchmakingService()

	// When
	gotFindMatchP1, err := service.FindMatch(player1Id)

	// Then
	wantFindMatchP1 := models.FindingMatch{ PlayerId: player1Id }
	is.NoErr(err)
	is.Equal(gotFindMatchP1, wantFindMatchP1)

	// When
	gotLeaveMatchmakingP2, err := service.LeaveMatchmaking(player2Id)

	// Then
	wantLeaveMatchmakingP2 := models.LeftMatchmaking{ PlayerId: player2Id }
	is.NoErr(err)
	is.Equal(gotLeaveMatchmakingP2, wantLeaveMatchmakingP2)

	// When
	gotFindMatchP2, err := service.FindMatch(player2Id)

	// Then
	wantFindMatchP2 := models.MatchMade{
		Player1Id: player1Id,
		Player2Id: player2Id,
	}
	is.NoErr(err)
	is.Equal(gotFindMatchP2, wantFindMatchP2)
}
