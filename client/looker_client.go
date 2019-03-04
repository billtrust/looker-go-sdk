// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/bmccarthy/looker-go-sdk/client/api_auth"
	"github.com/bmccarthy/looker-go-sdk/client/auth"
	"github.com/bmccarthy/looker-go-sdk/client/config"
	"github.com/bmccarthy/looker-go-sdk/client/connection"
	"github.com/bmccarthy/looker-go-sdk/client/content"
	"github.com/bmccarthy/looker-go-sdk/client/data_action"
	"github.com/bmccarthy/looker-go-sdk/client/datagroup"
	"github.com/bmccarthy/looker-go-sdk/client/group"
	"github.com/bmccarthy/looker-go-sdk/client/homepage"
	"github.com/bmccarthy/looker-go-sdk/client/integration"
	"github.com/bmccarthy/looker-go-sdk/client/look"
	"github.com/bmccarthy/looker-go-sdk/client/lookml_model"
	"github.com/bmccarthy/looker-go-sdk/client/project"
	"github.com/bmccarthy/looker-go-sdk/client/query"
	"github.com/bmccarthy/looker-go-sdk/client/render_task"
	"github.com/bmccarthy/looker-go-sdk/client/role"
	"github.com/bmccarthy/looker-go-sdk/client/running_queries"
	"github.com/bmccarthy/looker-go-sdk/client/scheduled_plan"
	"github.com/bmccarthy/looker-go-sdk/client/session"
	"github.com/bmccarthy/looker-go-sdk/client/space"
	"github.com/bmccarthy/looker-go-sdk/client/sql_query"
	"github.com/bmccarthy/looker-go-sdk/client/user"
	"github.com/bmccarthy/looker-go-sdk/client/user_attribute"
	"github.com/bmccarthy/looker-go-sdk/client/workspace"
)

// Default looker HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "billtrustdev.looker.com:19999"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/3.0"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new looker HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Looker {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new looker HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Looker {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new looker client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Looker {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Looker)
	cli.Transport = transport

	cli.APIAuth = api_auth.New(transport, formats)

	cli.Auth = auth.New(transport, formats)

	cli.Config = config.New(transport, formats)

	cli.Connection = connection.New(transport, formats)

	cli.Content = content.New(transport, formats)

	cli.DataAction = data_action.New(transport, formats)

	cli.Datagroup = datagroup.New(transport, formats)

	cli.Group = group.New(transport, formats)

	cli.Homepage = homepage.New(transport, formats)

	cli.Integration = integration.New(transport, formats)

	cli.Look = look.New(transport, formats)

	cli.LookmlModel = lookml_model.New(transport, formats)

	cli.Project = project.New(transport, formats)

	cli.Query = query.New(transport, formats)

	cli.RenderTask = render_task.New(transport, formats)

	cli.Role = role.New(transport, formats)

	cli.RunningQueries = running_queries.New(transport, formats)

	cli.ScheduledPlan = scheduled_plan.New(transport, formats)

	cli.Session = session.New(transport, formats)

	cli.Space = space.New(transport, formats)

	cli.SQLQuery = sql_query.New(transport, formats)

	cli.User = user.New(transport, formats)

	cli.UserAttribute = user_attribute.New(transport, formats)

	cli.Workspace = workspace.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Looker is a client for looker
type Looker struct {
	APIAuth *api_auth.Client

	Auth *auth.Client

	Config *config.Client

	Connection *connection.Client

	Content *content.Client

	DataAction *data_action.Client

	Datagroup *datagroup.Client

	Group *group.Client

	Homepage *homepage.Client

	Integration *integration.Client

	Look *look.Client

	LookmlModel *lookml_model.Client

	Project *project.Client

	Query *query.Client

	RenderTask *render_task.Client

	Role *role.Client

	RunningQueries *running_queries.Client

	ScheduledPlan *scheduled_plan.Client

	Session *session.Client

	Space *space.Client

	SQLQuery *sql_query.Client

	User *user.Client

	UserAttribute *user_attribute.Client

	Workspace *workspace.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Looker) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.APIAuth.SetTransport(transport)

	c.Auth.SetTransport(transport)

	c.Config.SetTransport(transport)

	c.Connection.SetTransport(transport)

	c.Content.SetTransport(transport)

	c.DataAction.SetTransport(transport)

	c.Datagroup.SetTransport(transport)

	c.Group.SetTransport(transport)

	c.Homepage.SetTransport(transport)

	c.Integration.SetTransport(transport)

	c.Look.SetTransport(transport)

	c.LookmlModel.SetTransport(transport)

	c.Project.SetTransport(transport)

	c.Query.SetTransport(transport)

	c.RenderTask.SetTransport(transport)

	c.Role.SetTransport(transport)

	c.RunningQueries.SetTransport(transport)

	c.ScheduledPlan.SetTransport(transport)

	c.Session.SetTransport(transport)

	c.Space.SetTransport(transport)

	c.SQLQuery.SetTransport(transport)

	c.User.SetTransport(transport)

	c.UserAttribute.SetTransport(transport)

	c.Workspace.SetTransport(transport)

}
