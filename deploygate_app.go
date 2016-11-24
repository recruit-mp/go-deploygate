package deploygate

import "fmt"

type GetAppCollaboratorInput struct {
	Owner    string
	Platform string
	AppId    string
}

type GetAppCollaboratorResponse struct {
	Error   bool                              `mapstructure:"error"`
	Results *GetAppCollaboratorResponseResult `mapstructure:"results"`
}

type AccountsUsage struct {
	Used uint `mapstructure:"used"`
	Max  uint `mapstructure:"max"`
}

type GetAppCollaboratorResponseResult struct {
	Usage *AccountsUsage  `mapstructure:"usage"`
	Users []*Collaborator `mapstructure:"users"`
	Teams []*Collaborator `mapstructure:"teams"`
}

func (c *Client) GetAppCollaborator(i *GetAppCollaboratorInput) (*GetAppCollaboratorResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", i.Owner, i.Platform, i.AppId)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var g *GetAppCollaboratorResponse
	if err := decodeJSON(&g, resp.Body); err != nil {
		return nil, err
	}
	return g, nil
}

type AddAppCollaboratorInput struct {
	Owner    string `form:"-"`
	Platform string `form:"-"`
	AppId    string `form:"-"`

	Users string `form:"users"`
	Role  int    `form:"role"`
}

type AddAppCollaboratorResponse struct {
	Error   bool                              `mapstructure:"error"`
	Results *AddAppCollaboratorResponseResult `mapstructure:"results"`
}

type Collaborator struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type AddAppCollaboratorResponseResult struct {
	Invite  string          `mapstructure:"invite"`
	Added   []*Collaborator `mapstructure:"added"`
	Invited []*Collaborator `mapstructure:"invited"`
}

// https://docs.deploygate.com/reference#invite

func (c *Client) AddAppCollaborator(i *AddAppCollaboratorInput) (*AddAppCollaboratorResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", i.Owner, i.Platform, i.AppId)

	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var a *AddAppCollaboratorResponse
	if err := decodeJSON(&a, resp.Body); err != nil {
		return nil, err
	}
	return a, nil
}

type DeleteAppCollaboratorInput AddAppCollaboratorInput
type DeleteAppCollaboratorResponse AddAppCollaboratorResponse

func (c *Client) DeleteAppCollaborator(i *DeleteAppCollaboratorInput) (*DeleteAppCollaboratorResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", i.Owner, i.Platform, i.AppId)

	resp, err := c.DeleteForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var a *DeleteAppCollaboratorResponse
	if err := decodeJSON(&a, resp.Body); err != nil {
		return nil, err
	}
	return a, nil
}
