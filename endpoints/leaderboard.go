package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// LeaderboardEndpoint represents the endpoint for leaderboard-related API operations.
type LeaderboardEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewLeaderboardEndpoint creates a new LeaderboardEndpoint instance.
func NewLeaderboardEndpoint(client *azura.Client) *LeaderboardEndpoint {
	return &LeaderboardEndpoint{client: client}
}

// GetGlobalLeaderboard retrieves the global leaderboard with the top specified number of entries.
func (lb *LeaderboardEndpoint) GetGlobalLeaderboard(top int) ([]types.Leaderboard, error) {
	endpoint := fmt.Sprintf("/v2/leaderboard/?mode=global&top=%d", top)
	response, err := lb.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var leaderboardEntries []types.Leaderboard
	err = parser.ParseDataToType(response.Data, &leaderboardEntries)
	if err != nil {
		return nil, err
	}

	return leaderboardEntries, nil
}

// GetGuildLeaderboard retrieves the guild-specific leaderboard for the specified guild ID
// with the top specified number of entries.
func (lb *LeaderboardEndpoint) GetGuildLeaderboard(guildID string, top int) ([]types.Leaderboard, error) {
	endpoint := fmt.Sprintf("/v2/leaderboard/?mode=guild&guildId=%s&top=%d", guildID, top)
	response, err := lb.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var leaderboardEntries []types.Leaderboard
	err = parser.ParseDataToType(response.Data, &leaderboardEntries)
	if err != nil {
		return nil, err
	}

	return leaderboardEntries, nil
}
