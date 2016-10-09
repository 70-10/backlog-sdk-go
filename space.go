package backlog

import (
	"encoding/json"
	"time"
)

func (c *Client) GetSpace() (*Space, error) {
	resp, err := c.Get("/space")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	space := &Space{}
	err = json.NewDecoder(resp.Body).Decode(&space)
	if err != nil {
		return nil, err
	}
	return space, err
}

type Space struct {
	SpaceKey           string    `json:"spaceKey"`
	Name               string    `json:"name"`
	OwnerID            int       `json:"ownerId"`
	Lang               string    `json:"lang"`
	Timezone           string    `json:"timezone"`
	ReportSendTime     string    `json:"reportSendTime"`
	TextFormattingRule string    `json:"textFormattingRule"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
}
