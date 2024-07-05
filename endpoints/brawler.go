package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// BrawlerEndpoint represents the endpoint for brawler-related API operations.
type BrawlerEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewBrawlerEndpoint creates a new BrawlerEndpoint instance.
func NewBrawlerEndpoint(client *azura.Client) *BrawlerEndpoint {
	return &BrawlerEndpoint{client: client}
}

// GetBrawlerByID retrieves a brawler by their ID.
func (b *BrawlerEndpoint) GetBrawlerByID(id int) (*types.Brawler, error) {
	endpoint := fmt.Sprintf("/v1/brawler/id/%d", id)
	response, err := b.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var brawler types.Brawler
	err = parser.ParseDataToType(response.Data, &brawler)
	if err != nil {
		return nil, err
	}

	return &brawler, nil
}

// GetBrawlerByName retrieves a brawler by their name.
func (b *BrawlerEndpoint) GetBrawlerByName(name string) (*types.Brawler, error) {
	endpoint := fmt.Sprintf("/v1/brawler/name/%s", name)
	response, err := b.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var brawler types.Brawler
	err = parser.ParseDataToType(response.Data, &brawler)
	if err != nil {
		return nil, err
	}

	return &brawler, nil
}

// GetAllBrawlers retrieves all available brawlers.
func (b *BrawlerEndpoint) GetAllBrawlers() ([]types.Brawler, error) {
	endpoint := "/v1/brawlers"
	response, err := b.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var brawlers []types.Brawler
	err = parser.ParseDataToType(response.Data, &brawlers)
	if err != nil {
		return nil, err
	}

	return brawlers, nil
}
