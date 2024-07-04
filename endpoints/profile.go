package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// ProfileEndpoint represents the endpoint for profile-related API operations.
type ProfileEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewProfileEndpoint creates a new ProfileEndpoint instance.
func NewProfileEndpoint(client *azura.Client) *ProfileEndpoint {
	return &ProfileEndpoint{client: client}
}

// GetProfile retrieves the profile for a user by their tag.
func (p *ProfileEndpoint) GetProfile(tag string) (*types.Profile, error) {
	endpoint := fmt.Sprintf("/v2/profile/%s", tag)
	response, err := p.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var profile types.Profile
	result, err := parser.ParseDataToType(response.Data, &profile)
	if err != nil {
		return nil, err
	}

	return result.(*types.Profile), nil
}
