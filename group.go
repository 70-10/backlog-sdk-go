package backlog

import (
	"encoding/json"
	"time"
)

func (c *Client) GetGroups() ([]Group, error) {
	resp, err := c.Get("/groups")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	groups := make([]Group, 0)
	err = json.NewDecoder(resp.Body).Decode(&groups)
	if err != nil {
		return nil, err
	}

	return groups, err

}

type Group struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Members      []User    `json:"members"`
	DisplayOrder int       `json:"displayOrder"`
	CreatedUser  User      `json:"createdUser"`
	Created      time.Time `json:"created"`
	UpdatedUser  User      `json:"updatedUser"`
	Updated      time.Time `json:"updated"`
}
