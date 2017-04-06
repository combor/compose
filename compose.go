package compose

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	apiProtocol = "https"
	apiHost     = "api.compose.io"
	apiVersion  = "2016-07"
)

func NewClient(token string, apiURL *url.URL) *Client {
	return &Client{
		Token:  token,
		ApiURL: apiURL,
	}
}

func GetApiToken() (token string, err error) {
	token = os.Getenv("COMPOSE_TOKEN")
	if token == "" {
		return "", fmt.Errorf("Empty token. Please export $COMPOSE_TOKEN")
	}
	return token, nil
}

func GetApiURL() *url.URL {
	return &url.URL{
		Scheme: apiProtocol,
		Host:   apiHost,
		Path:   "/" + apiVersion,
	}
}

func (c *Client) requestDo(reqtype, path string, body []byte) (resp *http.Response, err error) {
	url := c.ApiURL.String() + path
	req, err := http.NewRequest(reqtype, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	switch reqtype {
	default:
		req.Header.Set("Content-Type", "application/hal+json")
	case "POST":
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 202 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s", string(body))
	}

	return resp, nil
}

func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &v)
	return err
}

func validateResponse(resp *http.Response, err error) (ComposeResponse, error) {
	cr := ComposeResponse{}
	if err != nil {
		return cr, err
	}
	err = decodeResponse(resp, &cr)
	if err != nil {
		return cr, err
	}
	return cr, nil
}

func (c *Client) GetAccounts() ([]Account, error) {
	cr, err := validateResponse(c.requestDo("GET", "/accounts", nil))
	if err != nil {
		return nil, err
	}
	return cr.Embedded.Accounts, nil
}

func (c *Client) GetDeployments() ([]Deployment, error) {
	cr, err := validateResponse(c.requestDo("GET", "/deployments", nil))
	if err != nil {
		return nil, err
	}
	return cr.Embedded.Deployments, nil
}

func (c *Client) GetDeployment(deploymentId string) (Deployment, error) {
	resp, err := c.requestDo("GET", "/deployments/"+deploymentId, nil)
	deplResp := Deployment{}
	if err != nil {
		return deplResp, err
	}
	err = decodeResponse(resp, &deplResp)
	if err != nil {
		return deplResp, err
	}
	return deplResp, nil
}

func (c *Client) CreateDeployment(accountId, name, dbtype, dataCenter, version string, units int, ssl, wiredtiger bool) (Recipe, error) {

	depReq := struct {
		Deployment `json:"deployment"`
	}{
		Deployment{
			AccountId:  accountId,
			Name:       name,
			Type:       dbtype,
			Datacenter: dataCenter,
			Version:    version,
			Units:      units,
			SSL:        ssl,
			WiredTiger: wiredtiger,
		},
	}
	depResp := Recipe{}
	var body []byte
	body, err := json.Marshal(&depReq)
	if err != nil {
		return depResp, err
	}
	resp, err := c.requestDo("POST", "/deployments", body)
	if err != nil {
		return depResp, err
	}
	err = decodeResponse(resp, &depResp)
	if err != nil {
		return depResp, err
	}
	return depResp, nil
}

func (c *Client) DeleteDeployment(deploymentId string) (Recipe, error) {
	resp, err := c.requestDo("DELETE", "/deployments/"+deploymentId, nil)
	delResp := Recipe{}
	if err != nil {
		return delResp, err
	}
	err = decodeResponse(resp, &delResp)
	if err != nil {
		return delResp, err
	}
	return delResp, nil
}

func (c *Client) GetRecipe(recipeId string) (Recipe, error) {
	resp, err := c.requestDo("GET", "/recipes/"+recipeId, nil)
	recResp := Recipe{}
	if err != nil {
		return recResp, err
	}
	err = decodeResponse(resp, &recResp)
	if err != nil {
		return recResp, err
	}
	return recResp, nil
}

func (c *Client) GetDeploymentRecipes(deploymentId string) ([]Recipe, error) {
	cr, err := validateResponse(c.requestDo("GET", "/deployments/"+deploymentId+"/recipes", nil))
	if err != nil {
		return nil, err
	}
	return cr.Embedded.Recipes, nil
}

func (c *Client) GetClusters() ([]Cluster, error) {
	cl, err := validateResponse(c.requestDo("GET", "/clusters", nil))
	if err != nil {
		return nil, err
	}
	return cl.Embedded.Clusters, nil
}

func (c *Client) GetDeploymentVersions(deploymentId string) ([]Version, error) {
	cr, err := validateResponse(c.requestDo("GET", "/deployments/"+deploymentId+"/versions", nil))
	if err != nil {
		return nil, err
	}
	return cr.Embedded.Versions, nil
}

func (c *Client) GetDatacenters() ([]Datacenter, error) {
	cl, err := validateResponse(c.requestDo("GET", "/datacenters", nil))
	if err != nil {
		return nil, err
	}
	return cl.Embedded.Datacenters, nil
}

func (c *Client) GetDeploymentScalings(deploymentId string) (Scalings, error) {
	resp, err := c.requestDo("GET", "/deployments/"+deploymentId+"/scalings", nil)
	scalResp := Scalings{}
	if err != nil {
		return scalResp, err
	}
	err = decodeResponse(resp, &scalResp)
	if err != nil {
		return scalResp, err
	}
	return scalResp, nil
}

func (c *Client) ScaleDeployment(deploymentId string, units int) (Recipe, error) {

	scalingReq := struct {
		Deployment `json:"deployment"`
	}{
		Deployment{
			Units: units,
		},
	}
	scalingResp := Recipe{}
	var body []byte
	body, err := json.Marshal(&scalingReq)
	if err != nil {
		return scalingResp, err
	}
	resp, err := c.requestDo("POST", "/deployments/"+deploymentId+"/scalings", body)
	if err != nil {
		return scalingResp, err
	}
	err = decodeResponse(resp, &scalingResp)
	if err != nil {
		return scalingResp, err
	}
	return scalingResp, nil
}

func (c *Client) UpgradeDeployment(deploymentId string, version string) (Recipe, error) {

	upgradeReq := struct {
		Deployment `json:"deployment"`
	}{
		Deployment{
			Version: version,
		},
	}
	upgradeResp := Recipe{}
	var body []byte
	body, err := json.Marshal(&upgradeReq)
	if err != nil {
		return upgradeResp, err
	}
	resp, err := c.requestDo("PATCH", "/deployments/"+deploymentId+"/versions", body)
	if err != nil {
		return upgradeResp, err
	}
	err = decodeResponse(resp, &upgradeResp)
	if err != nil {
		return upgradeResp, err
	}
	return upgradeResp, nil
}
