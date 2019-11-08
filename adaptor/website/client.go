// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/awishformore/cancoillotte/model"
)

type Client struct {
	url   string
	api   string
	token string
}

func NewClient(url string, api string, token string) (*Client, error) {
	c := &Client{
		url:   url,
		api:   api,
		token: token,
	}
	_, err := c.Player(6139)
	if err != nil {
		return nil, errors.Wrap(err, "could not check status")
	}
	return c, nil
}

func (c *Client) Challenges() (model.ChallengeList, error) {
	data := ChallengesBody
	body, err := c.post(c.api, []byte(data))
	if err != nil {
		return nil, errors.Wrap(err, "could not execute request")
	}
	var res ChallengesResponse
	err = c.process(body, &res)
	if err != nil {
		return nil, errors.Wrap(err, "could not process request")
	}
	var challenges model.ChallengeList
	for _, c := range res.Data.Challenges.Sent {
		challenges = append(challenges, c.Convert())
	}
	for _, c := range res.Data.Challenges.Received {
		challenges = append(challenges, c.Convert())
	}
	for _, c := range res.Data.Challenges.Active {
		challenges = append(challenges, c.Convert())
	}
	return challenges, nil
}

func (c *Client) Suggestions() (model.ChallengeList, error) {
	data := SuggestionBody
	body, err := c.post(c.api, []byte(data))
	if err != nil {
		return nil, errors.Wrap(err, "could not execute request")
	}
	var res SuggestionResponse
	err = c.process(body, &res)
	if err != nil {
		return nil, errors.Wrap(err, "could not process response")
	}
	var challenges model.ChallengeList
	for _, cha := range res.Data.Challenges {
		challenges = append(challenges, cha.Convert())
	}
	return challenges, nil
}

func (c *Client) Issue(wizID1 uint64, wizID2 uint64) (string, error) {
	data := fmt.Sprintf(IssueBody,
		strconv.FormatUint(wizID1, 10),
		strconv.FormatUint(wizID2, 10),
	)
	body, err := c.post(c.api, []byte(data))
	if err != nil {
		return "", errors.Wrap(err, "could not execute request")
	}
	var res IssueResponse
	err = c.process(body, &res)
	if err != nil {
		return "", errors.Wrap(err, "could not decode response")
	}
	return res.Data.Challenge.ID, nil
}

func (c *Client) Commit(wizID uint64, owner string, chaID string, commit [32]byte, sig []byte) (string, error) {
	data := fmt.Sprintf(CommitBody,
		owner,
		strconv.FormatUint(wizID, 10),
		chaID,
		hex.EncodeToString(commit[:]),
		hex.EncodeToString(sig),
	)
	body, err := c.post(c.api, []byte(data))
	if err != nil {
		return "", errors.Wrap(err, "could not execute request")
	}
	var res CommitResponse
	err = c.process(body, &res)
	if err != nil {
		return "", errors.Wrap(err, "could not decode response")
	}
	return res.Data.Challenge.ID, nil
}

func (c *Client) Reveal(wizID uint64, chaID string, salt model.Hash, set model.Set) (string, error) {
	full := set.Full()
	data := fmt.Sprintf(RevealBody,
		strconv.FormatUint(wizID, 10),
		chaID,
		hex.EncodeToString(salt[:]),
		hex.EncodeToString(full[:]),
	)
	body, err := c.post(c.api, []byte(data))
	if err != nil {
		return "", errors.Wrap(err, "could not execute request")
	}
	var res RevealResponse
	err = c.process(body, &res)
	if err != nil {
		return "", errors.Wrap(err, "could not decode response")
	}
	return res.Data.Challenge.ID, nil
}

func (c *Client) Player(wizID uint64) (string, error) {
	id := strconv.FormatUint(wizID, 10)
	body, err := c.get(c.url + "/wizard/" + id)
	if err != nil {
		return "", errors.Wrap(err, "could not get page")
	}
	rx := regexp.MustCompile(`<a class="button___.{5} nicknameButton___.{5}" href="\/profile\/(.{8}-.{4}-.{4}-.{4}-.{12})">`)
	matches := rx.FindSubmatch(body)
	if len(matches) != 2 {
		return "", errors.New("no regex match")
	}
	return string(matches[1]), nil
}

func (c *Client) Online(player string) (bool, error) {
	body, err := c.get(c.url + "/profile/" + player)
	if err != nil {
		return false, errors.Wrap(err, "could not get page")
	}
	rx := regexp.MustCompile(`<div class="userExtraInfo___.{5}">(online|offline)<\/div>`)
	matches := rx.FindSubmatch(body)
	if len(matches) != 2 {
		return false, errors.New("no regex match")
	}
	return string(matches[1]) == "online", nil
}

func (c *Client) process(body []byte, res interface{}) error {
	var ser ErrorResponse
	err := json.Unmarshal(body, &ser)
	if err != nil {
		return errors.Wrap(err, "could not decode error")
	}
	if len(ser.Errors) != 0 {
		return errors.New(ser.Errors[0].Message)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.Wrap(err, "could not decode response")
	}
	return nil
}

func (c *Client) get(url string) ([]byte, error) {
	res, err := c.request().Get(url)
	return res.Body(), err
}

func (c *Client) post(url string, body []byte) ([]byte, error) {
	res, err := c.request().SetBody(body).Post(url)
	return res.Body(), err
}

func (c *Client) request() *resty.Request {
	req := resty.New().R().
		SetHeader("accept", "*/*").
		SetHeader("content-type", "application/json").
		SetHeader("Referer", "https://www.cheezewizards.com/").
		SetHeader("Sec-Fetch-Mode", "cors").
		SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/76.0.3809.100 Chrome/76.0.3809.100 Safari/537.36").
		SetHeader("x-authentication-token", c.token)
	return req
}
