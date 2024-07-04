package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// ClubEndpoint represents the endpoint for club-related API operations.
type ClubEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewClubEndpoint creates a new ClubEndpoint instance.
func NewClubEndpoint(client *azura.Client) *ClubEndpoint {
	return &ClubEndpoint{client: client}
}

// GetClub retrieves club details by their tag.
func (c *ClubEndpoint) GetClub(tag string) (*types.Club, error) {
	endpoint := fmt.Sprintf("/v2/club/%s", tag)
	response, err := c.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var club types.Club
	result, err := parser.ParseDataToType(response.Data, &club)
	if err != nil {
		return nil, err
	}

	return result.(*types.Club), nil
}
