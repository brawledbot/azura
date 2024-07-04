package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// GuildEndpoint represents the endpoint for guild-related API operations.
type GuildEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewGuildEndpoint creates a new GuildEndpoint instance.
func NewGuildEndpoint(client *azura.Client) *GuildEndpoint {
	return &GuildEndpoint{client: client}
}

// GetGuild retrieves a guild by its ID.
func (g *GuildEndpoint) GetGuild(id string) (*types.Guild, error) {
	endpoint := fmt.Sprintf("/v1/guild/%s", id)
	response, err := g.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var guild types.Guild
	result, err := parser.ParseDataToType(response.Data, &guild)
	if err != nil {
		return nil, err
	}

	return result.(*types.Guild), nil
}

// CreateGuild creates a new guild with the provided data.
func (g *GuildEndpoint) CreateGuild(guild *types.Guild) error {
	endpoint := "/v1/guild"
	_, err := g.client.DoRequest("POST", endpoint, guild)
	return err
}

// UpdateGuild updates an existing guild with the provided data.
func (g *GuildEndpoint) UpdateGuild(id string, guild *types.Guild) error {
	endpoint := fmt.Sprintf("/v1/guild/%s", id)
	_, err := g.client.DoRequest("PUT", endpoint, guild)
	return err
}

// DeleteGuild deletes a guild by its ID.
func (g *GuildEndpoint) DeleteGuild(id string) error {
	endpoint := fmt.Sprintf("/v1/guild/%s", id)
	_, err := g.client.DoRequest("DELETE", endpoint, nil)
	return err
}

// GetAllGuilds retrieves all guilds.
func (g *GuildEndpoint) GetAllGuilds() ([]types.Guild, error) {
	endpoint := "/v1/guilds"
	response, err := g.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var guilds []types.Guild
	result, err := parser.ParseDataToType(response.Data, &guilds)
	if err != nil {
		return nil, err
	}

	return result.([]types.Guild), nil
}
