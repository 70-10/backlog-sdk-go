package backlog

import (
	"encoding/json"
	"fmt"
)

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

func (c *Client) GetUser(id int) (*User, error) {
	resp, err := c.Get(fmt.Sprintf("/users/%d", id))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (c *Client) GetMyself() (*User, error) {
	resp, err := c.Get("/users/myself")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, err
}

type User struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	RoleType    int    `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}
