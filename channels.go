package arvanvod

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// store a newly channel
func (c *Client) CreateChannel(ctx context.Context, model CreateUpdateChannelModel) (*CreateChannelResp, error) {
	jsonBody, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("%s/channels", c.options.BaseUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(CreateChannelResp)
	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}
	return response, nil

}

// update a channel
func (c *Client) UpdateChannel(ctx context.Context, channel string, model CreateUpdateChannelModel) error {
	jsonBody, err := json.Marshal(model)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("%s/channels/%s", c.options.BaseUrl, channel)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, requestURL, bodyReader)
	if err != nil {
		return err
	}
	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return getErrorByStatus(res.StatusCode)

}

// Get all user's channels
func (c *Client) GetChannels(ctx context.Context, filter string, page, perPage int) (*GetChannelsModel, error) {

	requestURL := fmt.Sprintf("%s/channels?filter=%s&page=%d&per_page=%d", c.options.BaseUrl, filter, page, perPage)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(GetChannelsModel)
	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Get a specifed channel
func (c *Client) GetChannel(ctx context.Context, channel string) (*GetChannelResp, error) {

	requestURL := fmt.Sprintf("%s/channels/%s", c.options.BaseUrl, channel)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(GetChannelResp)
	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// delete a specifed channel
func (c *Client) DeleteChannel(ctx context.Context, channel string) error {

	requestURL := fmt.Sprintf("%s/channels/%s", c.options.BaseUrl, channel)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL, nil)
	if err != nil {
		return err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return getErrorByStatus(res.StatusCode)

}

type ChannelPresentType string

const (
	ChannelPeresentAuto  ChannelPresentType = "auto"
	ChannelPresentManual ChannelPresentType = "manual"
)

type ChannelModel struct {
	Id                string             `json:"id"`
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	SecureLinkEnabled int                `json:"secure_link_enabled"`
	SecureLinkKey     string             `json:"secure_link_key"`
	SecureLinkWithIp  bool               `json:"secure_link_with_ip"`
	AdsEnabled        int                `json:"ads_enabled"`
	PresentType       ChannelPresentType `json:"present_type"`
	CampaignId        string             `json:"campaign_id"`
}

type CreateUpdateChannelModel struct {
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	SecureLinkEnabled int                `json:"secure_link_enabled"`
	SecureLinkKey     string             `json:"secure_link_key"`
	SecureLinkWithIp  bool               `json:"secure_link_with_ip"`
	AdsEnabled        int                `json:"ads_enabled"`
	PresentType       ChannelPresentType `json:"present_type"`
	CampaignId        string             `json:"campaign_id"`
}
type GetChannelsModel struct {
	Data  []ChannelModel `json:"data"`
	Links *Links         `json:"links"`
	Meta  *Meta          `json:"meta"`
}

type CreateChannelResp struct {
	Data    ChannelModel `json:"data"`
	Message string       `json:"message"`
}

type GetChannelResp struct {
	Data ChannelModel `json:"data"`
}
