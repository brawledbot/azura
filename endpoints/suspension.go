package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// SuspensionEndpoint represents the endpoint for suspension-related API operations.
type SuspensionEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewSuspensionEndpoint creates a new SuspensionEndpoint instance.
func NewSuspensionEndpoint(client *azura.Client) *SuspensionEndpoint {
	return &SuspensionEndpoint{client: client}
}

// GetSuspension retrieves suspension details for a user by their ID.
func (ss *SuspensionEndpoint) GetSuspension(id string) (*types.Suspension, error) {
	endpoint := fmt.Sprintf("/v1/suspension/id/%s", id)
	response, err := ss.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var suspension types.Suspension
	err = parser.ParseDataToType(response.Data, &suspension)
	if err != nil {
		return nil, err
	}

	return &suspension, nil
}

// SuspendUser suspends a user by their ID with a specified reason.
// Returns the suspension details or an error if the suspension fails.
func (ss *SuspensionEndpoint) SuspendUser(id, reason string) (*types.Suspension, error) {
	endpoint := fmt.Sprintf("/v1/suspend/%s?reason=%s", id, reason)
	response, err := ss.client.DoRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var suspension types.Suspension
	err = parser.ParseDataToType(response.Data, &suspension)
	if err != nil {
		return nil, err
	}

	return &suspension, nil
}

// UnsuspendUser unsuspends a user by their ID.
// Returns an error if the unsuspension fails.
func (ss *SuspensionEndpoint) UnsuspendUser(id string) error {
	endpoint := fmt.Sprintf("/v1/unsuspend/%s", id)
	_, err := ss.client.DoRequest("POST", endpoint, nil)
	return err
}
