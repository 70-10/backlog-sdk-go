package backlog

import "encoding/json"

func (c *Client) GetUsers() ([]User, error) {
	resp, err := c.Get("/users")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, err
}

type User struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	RoleType    int    `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}
