package client

import (
	"fmt"

	"github.com/evalphobia/httpwrapper/request"
)

// Client is http client for IPQualityScore API.
type Client struct {
	Option
	APIKey string
}

func New() *Client {
	return &Client{}
}

func (c *Client) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

func (c *Client) SetOption(opt Option) {
	c.Option = opt
}

// CallGET sends GET request to `url` with `params` and set reqponse to `result`
func (c *Client) CallGET(path string, params interface{}, result Response) (err error) {
	opt := c.Option
	url := fmt.Sprintf("%s%s", opt.getBaseURL(), path)

	resp, err := request.GET(url, request.Option{
		Query:     params,
		Headers:   map[string]string{"IPQS-KEY": c.APIKey},
		Retry:     opt.Retry,
		Debug:     opt.Debug,
		UserAgent: opt.getUserAgent(),
		Timeout:   opt.getTimeout(),
	})
	if !resp.Ok {
		result.SetStatusCode(resp.StatusCode)
		err = resp.HasStatusCodeError()
	}
	if e := resp.JSON(result); e != nil {
		return e
	}
	return err
}

// CallPOST sends POST request to `url` with `params` and set reqponse to `result`
func (c *Client) CallPOST(path string, params interface{}, result Response) (err error) {
	opt := c.Option
	url := fmt.Sprintf("%s%s", opt.getBaseURL(), path)

	resp, err := request.POST(url, request.Option{
		Payload:     params,
		PayloadType: request.PayloadTypeJSON,
		Headers:     map[string]string{"IPQS-KEY": c.APIKey},
		Retry:       opt.Retry,
		Debug:       opt.Debug,
		UserAgent:   opt.getUserAgent(),
		Timeout:     opt.getTimeout(),
	})
	if err != nil {
		return err
	}
	if !resp.Ok {
		result.SetStatusCode(resp.StatusCode)
		err = resp.HasStatusCodeError()
	}
	if e := resp.JSON(result); e != nil {
		return e
	}
	return err
}

type Response interface {
	SetStatusCode(status int)
}
