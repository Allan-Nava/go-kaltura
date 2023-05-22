package gokaltura

import (
	"log"
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
	//
	return k
}