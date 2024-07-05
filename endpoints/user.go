package endpoints

import (
	"fmt"

	"github.com/brawledbot/azura"
	"github.com/brawledbot/azura/parser"
	"github.com/brawledbot/azura/types"
)

// UserEndpoint represents the endpoint for user-related API operations.
type UserEndpoint struct {
	client *azura.Client // client is the HTTP client used to make requests.
}

// NewUserEndpoint creates a new UserEndpoint instance.
func NewUserEndpoint(client *azura.Client) *UserEndpoint {
	return &UserEndpoint{client: client}
}

// GetUserByID retrieves a user by their ID.
func (u *UserEndpoint) GetUserByID(id string) (*types.User, error) {
	endpoint := fmt.Sprintf("/v1/user/id/%s", id)
	response, err := u.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var user types.User
	err = parser.ParseDataToType(response.Data, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByTag retrieves a user by their tag.
func (u *UserEndpoint) GetUserByTag(tag string) (*types.User, error) {
	endpoint := fmt.Sprintf("/v1/user/tag/%s", tag)
	response, err := u.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var user types.User
	err = parser.ParseDataToType(response.Data, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user with the provided data.
func (u *UserEndpoint) CreateUser(user *types.User) error {
	endpoint := "/v1/user"
	_, err := u.client.DoRequest("POST", endpoint, user)
	return err
}

// DeleteUser deletes a user by their ID.
func (u *UserEndpoint) DeleteUser(id string) error {
	endpoint := fmt.Sprintf("/v1/user/id/%s", id)
	_, err := u.client.DoRequest("DELETE", endpoint, nil)
	return err
}

// SetPremium sets a user's premium status by their ID.
func (u *UserEndpoint) SetPremium(id string, state bool) error {
	endpoint := fmt.Sprintf("/v1/user/id/%s/premium?state=%v", id, state)
	_, err := u.client.DoRequest("PUT", endpoint, nil)
	return err
}

// VerifyUser verifies a user by their tag.
func (u *UserEndpoint) VerifyUser(tag string) error {
	endpoint := fmt.Sprintf("/v1/user/tag/%s/verify", tag)
	_, err := u.client.DoRequest("PUT", endpoint, nil)
	return err
}

// UpdateFreeUses updates a user's remaining free uses by their ID.
func (u *UserEndpoint) UpdateFreeUses(id string, freeUses int) error {
	endpoint := fmt.Sprintf("/v1/user/id/%s/free-uses?freeUses=%d", id, freeUses)
	_, err := u.client.DoRequest("PUT", endpoint, nil)
	return err
}

// GetUsers retrieves all users.
func (u *UserEndpoint) GetUsers() ([]types.User, error) {
	endpoint := "/v1/users"
	response, err := u.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var users []types.User
	err = parser.ParseDataToType(response.Data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
