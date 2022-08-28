// Abr Arvan VOD API client enabling Go programs to interact with Arvan VOD service in a simple and uniform way
package arvanvod

// options for the client instance
// set parameters like base url and api key
type ClientOptions struct {
	ApiKey  string
	BaseUrl string
}

// Client is a struct  that makes APi calls to arvan vod
type Client struct {
	options *ClientOptions
}

// create new Client instanse with given options
func NewClient(options *ClientOptions) *Client {
	return &Client{
		options: options,
	}
}
