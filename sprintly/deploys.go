package sprintly

import (
	"fmt"
	"net/http"
)

// DeploysService holds all the methods for manipulating Sprintly items.
type DeploysService struct {
	client *Client
}

func newDeploysService(client *Client) *DeploysService {
	return &DeploysService{client}
}

// Deploy represents the Sprintly Deploy resource.
type Deploy struct {
	Environment string `json:"environment,omitempty"`
	Items       []Item `json:"items,omitempty"`
}

type DeployListArgs struct {
	Environment string `json:"environment,omitempty"`
}

// List can be used to list deploys for the given product.
//
// See https://sprintly.uservoice.com/knowledgebase/articles/138392-deploys
func (srv DeploysService) List(productId int, opt *DeployListArgs) ([]Deploy, *http.Response, error) {
	u := fmt.Sprintf("products/%v/deploys.json", productId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := srv.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var deploys []Deploy
	resp, err := srv.client.Do(req, &deploys)
	if err != nil {
		switch resp.StatusCode {
		case 403:
			return nil, nil, &ErrDeploys403{err.(*ErrAPI)}
		case 404:
			return nil, nil, &ErrDeploys404{err.(*ErrAPI)}
		default:
			return nil, resp, err
		}
	}

	return deploys, resp, nil
}

// Create can be used to create a new deployment for the given product.
//
// See https://sprintly.uservoice.com/knowledgebase/articles/138392-deploys
func (srv DeploysService) Create(productId int, args *DeployCreateArgs) ([]Deploy, *http.Response, error) {
	u := fmt.Sprintf("products/%v/deploys.json", productId)
	u, err := addOptions(u, args)
	if err != nil {
		return nil, nil, err
	}

	req, err := srv.client.NewRequest("POST", u, args)
	if err != nil {
		return nil, nil, err
	}

	var deploys []Deploy
	resp, err := srv.client.Do(req, &deploys)
	if err != nil {
		switch resp.StatusCode {
		case 400:
			return nil, nil, &ErrDeploys400{err.(*ErrAPI)}
		case 403:
			return nil, nil, &ErrDeploys403{err.(*ErrAPI)}
		case 404:
			return nil, nil, &ErrDeploys404{err.(*ErrAPI)}
		default:
			return nil, resp, err
		}

	}

	return deploys, resp, nil
}