package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// GenerateEndpoint represents the endpoint for image generation API operations.
type GenerateEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewGenerateEndpoint creates a new GenerateEndpoint instance.
func NewGenerateEndpoint(client *azura.Client) *GenerateEndpoint {
	return &GenerateEndpoint{client: client}
}

// Generate generates an image based on the player tag and render type.
func (g *GenerateEndpoint) Generate(playerTag string, renderType string) (*types.GeneratedImage, error) {
	endpoint := fmt.Sprintf("/v1/generate/%s?renderType=%s", playerTag, renderType)
	response, err := g.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var generatedImage types.GeneratedImage
	err = parser.ParseDataToType(response.Data, &generatedImage)
	if err != nil {
		return nil, err
	}

	return &generatedImage, nil
}

// GenerateWithConfigURL generates an image with custom configuration specified by a URL.
func (g *GenerateEndpoint) GenerateWithConfigURL(playerTag string, renderType string, configURL string) (*types.GeneratedImage, error) {
	endpoint := fmt.Sprintf("/v1/generate/%s?renderType=%s&custom=%s", playerTag, renderType, configURL)
	response, err := g.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var generatedImage types.GeneratedImage
	err = parser.ParseDataToType(response.Data, &generatedImage)
	if err != nil {
		return nil, err
	}

	return &generatedImage, nil
}
