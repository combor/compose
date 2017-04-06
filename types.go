package compose

import "net/url"

type ComposeResponse struct {
	Embedded struct {
		Deployments []Deployment `json:"deployments"`
		Accounts    []Account    `json:"accounts"`
		Recipes     []Recipe     `json:"recipes"`
		Clusters    []Cluster    `json:"clusters"`
		Versions    []Version    `json:"transitions"`
		Datacenters []Datacenter `json:"datacenters"`
	} `json:"_embedded"`
}

type Client struct {
	Token  string
	ApiURL *url.URL
}

type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug,omitempty"`
}

type Recipe struct {
	Id           string `json:"id"`
	Template     string `json:"template"`
	Status       string `json:"status"`
	StatusDetail string `json:"status_detail"`
	AccountId    string `json:"account_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeploymentId string `json:"deployment_id"`
	Name         string `json:"name"`
}

type Deployment struct {
	Id                  string             `json:"id,omitempty"`
	AccountId           string             `json:"account_id,omitempty"`
	Name                string             `json:"name,omitempty"`
	CreatedAt           string             `json:"created_at,omitempty"`
	Type                string             `json:"type,omitempty"`
	Datacenter          string             `json:"datacenter,omitempty"`
	Version             string             `json:"version,omitempty"`
	Units               int                `json:"units,omitempty"`
	SSL                 bool               `json:"ssl,omitempty"`
	WiredTiger          bool               `json:"wired_tiger,omitempty"`
	Links               *Links             `json:"_links,omitempty"`
	ProvisionRecipeId   string             `json:"provision_recipe_id,omitempty"`
	CaCertificateBase64 string             `json:"ca_certificate_base64,omitempty"`
	ConnectionStrings   *ConnectionStrings `json:"connection_strings,omitempty"`
}

type ConnectionStrings struct {
	Health   []string `json:"health,omitempty"`
	SSH      []string `json:"ssh,omitempty"`
	Admin    []string `json:"admin,omitempty"`
	SSHAdmin []string `json:"ssh_admin,omitempty"`
	Cli      []string `json:"cli,omitempty"`
	Direct   []string `json:"direct,omitempty"`
}

type Links struct {
	ComposeWebUI struct {
		Href      string `json:"href,omitempty"`
		Templated bool   `json:"templated,omitempty"`
	} `json:"compose_web_ui,omitempty"`
}

type Cluster struct {
	Id          string `json"id"`
	AccountId   string `json:"account_id"`
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Provider    string `json:"provider,omitempty"`
	Region      string `json:"region,omitempty"`
	Multitenant string `json:"multitenant,omitempty"`
	AccountSlug string `json:"account_slug,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Subdomain   string `json:"subdomain,omitempty"`
}

type Version struct {
	Application string `json:"application":`
	Method      string `json:"method"`
	FromVersion string `json:"from_version"`
	ToVersion   string `json:"to_version"`
}

type Datacenter struct {
	Region   string `json:"region"`
	Provider string `json:"provider"`
	Slug     string `json:"slug"`
}

type Scalings struct {
	AllocatedUnits int `json:"allocated_units"`
	UsedUnits      int `json:"used_units"`
	StartingUnits  int `json:"starting_units"`
	MinimumUnits   int `json:"minimum_units"`
}
