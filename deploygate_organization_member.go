package deploygate

import "fmt"

type GetOrganizationMemberInput struct {
	OrganizationName string
}

type GetOrganizationMemberResponse struct {
	Error   bool      `mapstructure:"error"`
	Members []*Member `mapstructure:"members"`
}

type Member struct {
	Name     string `mapstructure:"name"`
	URL      string `mapstructure:"url"`
	IconURL  string `mapstructure:"icon_url"`
	Type     string `mapstructure:"type"`
	Inviting bool   `mapstructure:"inviting"`
}

func (c *Client) GetOrganizationMember(i *GetOrganizationMemberInput) (*GetOrganizationMemberResponse, error) {
	path := fmt.Sprintf("/organizations/%s/members", i.OrganizationName)

	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var g *GetOrganizationMemberResponse
	if err := decodeJSON(&g, resp.Body); err != nil {
		return nil, err
	}
	return g, nil
}

type AddOrganizationMemberInput struct {
	OrganizationName string `form:"-"`
	UserName         string `form:"username,omitempty"`
	Email            string `form:"email,omitempty"`
}

type AddOrganizationMemberResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
}

func (c *Client) AddOrganizationMember(i *AddOrganizationMemberInput) (*AddOrganizationMemberResponse, error) {
	path := fmt.Sprintf("/organizations/%s/members", i.OrganizationName)

	resp, err := c.PostForm(path, i, nil)

	if err != nil {
		return nil, err
	}

	var a *AddOrganizationMemberResponse
	if err := decodeJSON(&a, resp.Body); err != nil {
		return nil, err
	}
	return a, nil
}

type DeleteOrganizationMemberInput AddOrganizationMemberInput
type DeleteOrganizationMemberResponse AddOrganizationMemberResponse

func (c *Client) DeleteOrganizationMember(i *DeleteOrganizationMemberInput) (*DeleteOrganizationMemberResponse, error) {
	var id string
	if i.UserName != "" {
		id = i.UserName
	} else {
		id = i.Email
	}

	path := fmt.Sprintf("/organizations/%s/members/%s", i.OrganizationName, id)

	resp, err := c.DeleteForm(path, i, nil)

	if err != nil {
		return nil, err
	}

	var d *DeleteOrganizationMemberResponse
	if err := decodeJSON(&d, resp.Body); err != nil {
		return nil, err
	}
	return d, nil
}
