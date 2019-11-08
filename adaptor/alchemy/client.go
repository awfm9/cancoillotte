// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package alchemy

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/awishformore/cancoillotte/model"
)

type Client struct {
	url   string
	token string
	email string
}

func NewClient(url string, token string, email string) (*Client, error) {
	c := &Client{
		url:   url,
		token: token,
		email: email,
	}
	_, err := c.Wizards()
	if err != nil {
		return nil, errors.Wrap(err, "could not check status")
	}
	return c, nil
}

func (c *Client) Wizards() (model.WizardList, error) {

	if c.url == "" {
		return model.WizardList{}, nil
	}

	var res WizardResponse
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("x-api-token", c.token).
		SetHeader("x-email", c.email).
		SetResult(&res).
		Get(c.url + "/wizards")
	if err != nil {
		return nil, errors.Wrap(err, "could not get wizards")
	}

	wizards := make(model.WizardList, 0, len(res.Wizards))
	for _, w := range res.Wizards {
		wizards = append(wizards, w.Convert())
	}

	return wizards, nil
}

func (c *Client) Duels() (model.DuelList, error) {

	if c.url == "" {
		return model.DuelList{}, nil
	}

	var res DuelResponse
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("x-api-token", c.token).
		SetHeader("x-email", c.email).
		SetResult(&res).
		Get(c.url + "/duels")
	if err != nil {
		return nil, errors.Wrap(err, "could not get duels")
	}

	duels := make(model.DuelList, 0, len(res.Duels))
	for _, duel := range res.Duels {
		duels = append(duels, duel.Convert())
	}

	return duels, nil
}
