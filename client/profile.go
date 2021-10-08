package client

import "time"

type DMeProfileInput struct {
	ProfileName   string  `json:"profile_name"`
	AvatarFileID  *string `json:"avatar_file_id"`
	HideFirstName *bool   `json:"hide_first_name"`
	HideSurname   *bool   `json:"hide_surname"`
	Pseudonym     *string `json:"pseudonym"`
	Bio           *string `json:"bio"`
}

// DMeProfileAdd creates new profile for DMe type of i2i
// The i2i must be previously initialized as DMe.
func (c *Client) DMeProfileAdd(profile *DMeProfileInput) (*Profile, error) {
	response := struct {
		Profile Profile `json:"dMeProfileAdd"`
	}{}

	_, err := c.query(&query{
		query:     mutationDMeProfileAdd,
		variables: map[string]interface{}{"input": profile},
		timeout:   time.Second * 2,
		response:  &response,
	})
	return &response.Profile, err
}

// ProfileList lists configure DMe profiles
func (c *Client) ProfileList() ([]*Profile, error) {
	response := struct {
		Profiles []*Profile `json:"profileList"`
	}{}

	_, err := c.query(&query{
		query:    queryProfileList,
		timeout:  time.Second * 2,
		response: &response,
	})
	return response.Profiles, err
}
