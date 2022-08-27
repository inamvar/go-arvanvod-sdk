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
func (c *Client) CreateChannel(ctx context.Context, model *CreateChannelModel) error {
	jsonBody, err := json.Marshal(model)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("%s/channels", c.options.BaseUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return err
	}
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

type ChannelPresentType string

const (
	ChannelPeresentAuto  ChannelPresentType = "auto"
	ChannelPresentManual ChannelPresentType = "manual"
)

type CreateChannelModel struct {
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	SecureLinkEnabled bool               `json:"secure_link_enabled"`
	SecureLinkKey     string             `json:"secure_link_key"`
	SecureLinkWithIp  bool               `json:"secure_link_with_ip"`
	AdsEnabled        bool               `json:"ads_enabled"`
	PresentType       ChannelPresentType `json:"present_type"`
	CampaignId        string             `json:"campaign_id"`
}

type GetChannelsModel struct {
}
