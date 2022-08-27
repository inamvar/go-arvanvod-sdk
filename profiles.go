package arvanvod

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get all channel's profiles
func (c *Client) GetChannelProfiles(ctx context.Context, channel, filter string, page, perPage int) (*GetChannelProfilesModel, error) {

	requestURL := fmt.Sprintf("%s/channels/%s/profiles?filter=%s&page=%d&per_page=%d", c.options.BaseUrl,channel, filter, page, perPage)
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

	response := new(GetChannelProfilesModel)
	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}



type  GetChannelProfilesModel struct {

}