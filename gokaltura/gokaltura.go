package gokaltura

import (
	"github.com/Allan-Nava/go-kaltura/configuration"
	"github.com/go-resty/resty/v2"
)

type IGoKaltura interface {
	//
}

type gokaltura struct {
	configuration *configuration.Configuration
	restClient    *resty.Client
}


func NewGoKaltura(configuration *configuration.Configuration) IGoKaltura {
	k := &gokaltura{
		configuration: configuration,
	}
	k.restClient = resty.New()
	k.restClient.SetBaseURL(configuration.BaseUrl)
	//
	return k
}



// Resty Methods

func (o *gokaltura) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *gokaltura) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}