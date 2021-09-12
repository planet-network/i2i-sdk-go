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

func (c *Client) AddProfile(profile string) error {
	input := DMeProfileInput{
		ProfileName: profile,
	}
	var response interface{}
	_, err := c.query(&query{
		query:     mutationDMeProfileAdd,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 2,
		response:  &response,
	})
	return err
}