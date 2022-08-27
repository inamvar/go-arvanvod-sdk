package arvanvod

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Set subdomain for VOD service
func (c *Client) SetDomain(ctx context.Context, subdomain *SetSubDomainModel) error {

	jsonBody, err := json.Marshal(subdomain)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("%s/domain", c.options.BaseUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bodyReader)
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

// Get subdomain
func (c *Client) GetDomain(ctx context.Context) (*GetSubdomainModel, error) {

	requestURL := fmt.Sprintf("%s/domain", c.options.BaseUrl)
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

	response := new(GetSubdomainModel)
	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type SetSubDomainModel struct {
	Subdomain string `json:"subdomain"`
}

type GetSubdomainModel struct {
	Data struct {
		Subdomain string     `json:"subdomain"`
		Domain    string     `json:"domain"`
		CreatedAT *time.Time `json:"created_at"`
	} `json:"data"`
}
