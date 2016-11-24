package deploygate

import "testing"

func TestClient_AddOrganizationMember(t *testing.T) {
	t.Parallel()

	var err error
	var resp *AddOrganizationMemberResponse
	// Add collaborator
	record(t, "organizations/add_member", func(c *Client) {
		resp, err = c.AddOrganizationMember(&AddOrganizationMemberInput{
			OrganizationName: "organization_name",
			UserName:         "testuser",
		})
	})

	if err != nil {
		t.Fatal(err)
	}

	if resp.Error != false {
		t.Error("response caused error")
	}

	if resp.Message != "testuser was invited to organization_name" {
		t.Error("bad message field")
	}
}
