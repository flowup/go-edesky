package edesky

import (
	"github.com/parnurzeal/gorequest"
	"encoding/xml"
	"errors"
)

const (
	APIUrl           = "https://edesky.cz/api/v1"
	DocumentsAPIUrl  = APIUrl + "/documents"
	DashboardsAPIUrl = APIUrl + "/dashboards"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}

func (c *Client) FindDocuments(options DocumentOptions) ([]Document, []error) {
	if options.Keywords == "" {
		return nil, []error{errors.New("Keywords option is required")}
	}

	request := gorequest.New()
	//request.SetDebug(true)

	_, body, errs := request.Get(DocumentsAPIUrl).Query(options).End()
	if len(errs) > 0 {
		return nil, errs
	}

	documents := DocumentRoot{}
	err := xml.Unmarshal([]byte(body), &documents)
	if err != nil {
		return nil, []error{err}
	}

	return documents.Documents, nil
}

func (c *Client) FindDashboards(options DashboardOptions) ([]Dashboard, []error) {
	request := gorequest.New()
	//request.SetDebug(true)

	_, body, errs := request.Get(DashboardsAPIUrl).Query(options).End()
	if len(errs) > 0 {
		return nil, errs
	}

	dashboards := DashboardRoot{}
	err := xml.Unmarshal([]byte(body), &dashboards)
	if err != nil {
		return nil, []error{err}
	}

	return dashboards.Dashboards, nil
}
