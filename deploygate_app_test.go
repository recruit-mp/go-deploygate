package deploygate

import "testing"

func TestClient_AddAppCollaborator(t *testing.T) {
	t.Parallel()

	var err error
	var resp *AddAppCollaboratorResponse
	// Add collaborator
	record(t, "apps/add_collaborator", func(c *Client) {
		resp, err = c.AddAppCollaborator(&AddAppCollaboratorInput{
			Owner:    "owner",
			Platform: "platform",
			AppId:    "app_id",
			Users:    "testuser",
			Role:     2,
		})
	})

	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}

	if resp.Results.Invite != "member invite success" {
		t.Error("bad invite field")
	}
}
