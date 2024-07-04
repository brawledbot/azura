package endpoints

import (
	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/types"
)

// VerificationEndpoint represents the endpoint for verification-related API operations.
type VerificationEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewVerificationEndpoint creates a new VerificationEndpoint instance.
func NewVerificationEndpoint(client *azura.Client) *VerificationEndpoint {
	return &VerificationEndpoint{client: client}
}

// CreateVerification creates a new verification request with the provided data.
func (vc *VerificationEndpoint) CreateVerification(verification *types.Verification) error {
	endpoint := "/v1/verify"
	_, err := vc.client.DoRequest("POST", endpoint, verification)
	return err
}
