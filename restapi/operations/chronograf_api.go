package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewChronografAPI creates a new Chronograf instance
func NewChronografAPI(spec *loads.Document) *ChronografAPI {
	return &ChronografAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
	}
}

/*ChronografAPI API endpoints for Chronograf */
type ChronografAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// DeleteLayoutsIDHandler sets the operation handler for the delete layouts ID operation
	DeleteLayoutsIDHandler DeleteLayoutsIDHandler
	// DeleteSourcesIDHandler sets the operation handler for the delete sources ID operation
	DeleteSourcesIDHandler DeleteSourcesIDHandler
	// DeleteSourcesIDKapacitorsKapaIDHandler sets the operation handler for the delete sources ID kapacitors kapa ID operation
	DeleteSourcesIDKapacitorsKapaIDHandler DeleteSourcesIDKapacitorsKapaIDHandler
	// DeleteSourcesIDKapacitorsKapaIDProxyHandler sets the operation handler for the delete sources ID kapacitors kapa ID proxy operation
	DeleteSourcesIDKapacitorsKapaIDProxyHandler DeleteSourcesIDKapacitorsKapaIDProxyHandler
	// DeleteSourcesIDRolesRoleIDHandler sets the operation handler for the delete sources ID roles role ID operation
	DeleteSourcesIDRolesRoleIDHandler DeleteSourcesIDRolesRoleIDHandler
	// DeleteSourcesIDUsersUserIDHandler sets the operation handler for the delete sources ID users user ID operation
	DeleteSourcesIDUsersUserIDHandler DeleteSourcesIDUsersUserIDHandler
	// DeleteSourcesIDUsersUserIDExplorationsExplorationIDHandler sets the operation handler for the delete sources ID users user ID explorations exploration ID operation
	DeleteSourcesIDUsersUserIDExplorationsExplorationIDHandler DeleteSourcesIDUsersUserIDExplorationsExplorationIDHandler
	// GetHandler sets the operation handler for the get operation
	GetHandler GetHandler
	// GetLayoutsHandler sets the operation handler for the get layouts operation
	GetLayoutsHandler GetLayoutsHandler
	// GetLayoutsIDHandler sets the operation handler for the get layouts ID operation
	GetLayoutsIDHandler GetLayoutsIDHandler
	// GetMappingsHandler sets the operation handler for the get mappings operation
	GetMappingsHandler GetMappingsHandler
	// GetSourcesHandler sets the operation handler for the get sources operation
	GetSourcesHandler GetSourcesHandler
	// GetSourcesIDHandler sets the operation handler for the get sources ID operation
	GetSourcesIDHandler GetSourcesIDHandler
	// GetSourcesIDKapacitorsHandler sets the operation handler for the get sources ID kapacitors operation
	GetSourcesIDKapacitorsHandler GetSourcesIDKapacitorsHandler
	// GetSourcesIDKapacitorsKapaIDHandler sets the operation handler for the get sources ID kapacitors kapa ID operation
	GetSourcesIDKapacitorsKapaIDHandler GetSourcesIDKapacitorsKapaIDHandler
	// GetSourcesIDKapacitorsKapaIDProxyHandler sets the operation handler for the get sources ID kapacitors kapa ID proxy operation
	GetSourcesIDKapacitorsKapaIDProxyHandler GetSourcesIDKapacitorsKapaIDProxyHandler
	// GetSourcesIDMonitoredHandler sets the operation handler for the get sources ID monitored operation
	GetSourcesIDMonitoredHandler GetSourcesIDMonitoredHandler
	// GetSourcesIDPermissionsHandler sets the operation handler for the get sources ID permissions operation
	GetSourcesIDPermissionsHandler GetSourcesIDPermissionsHandler
	// GetSourcesIDRolesHandler sets the operation handler for the get sources ID roles operation
	GetSourcesIDRolesHandler GetSourcesIDRolesHandler
	// GetSourcesIDRolesRoleIDHandler sets the operation handler for the get sources ID roles role ID operation
	GetSourcesIDRolesRoleIDHandler GetSourcesIDRolesRoleIDHandler
	// GetSourcesIDUsersHandler sets the operation handler for the get sources ID users operation
	GetSourcesIDUsersHandler GetSourcesIDUsersHandler
	// GetSourcesIDUsersUserIDHandler sets the operation handler for the get sources ID users user ID operation
	GetSourcesIDUsersUserIDHandler GetSourcesIDUsersUserIDHandler
	// GetSourcesIDUsersUserIDExplorationsHandler sets the operation handler for the get sources ID users user ID explorations operation
	GetSourcesIDUsersUserIDExplorationsHandler GetSourcesIDUsersUserIDExplorationsHandler
	// GetSourcesIDUsersUserIDExplorationsExplorationIDHandler sets the operation handler for the get sources ID users user ID explorations exploration ID operation
	GetSourcesIDUsersUserIDExplorationsExplorationIDHandler GetSourcesIDUsersUserIDExplorationsExplorationIDHandler
	// GetTokenHandler sets the operation handler for the get token operation
	GetTokenHandler GetTokenHandler
	// PatchSourcesIDHandler sets the operation handler for the patch sources ID operation
	PatchSourcesIDHandler PatchSourcesIDHandler
	// PatchSourcesIDKapacitorsKapaIDHandler sets the operation handler for the patch sources ID kapacitors kapa ID operation
	PatchSourcesIDKapacitorsKapaIDHandler PatchSourcesIDKapacitorsKapaIDHandler
	// PatchSourcesIDKapacitorsKapaIDProxyHandler sets the operation handler for the patch sources ID kapacitors kapa ID proxy operation
	PatchSourcesIDKapacitorsKapaIDProxyHandler PatchSourcesIDKapacitorsKapaIDProxyHandler
	// PatchSourcesIDRolesRoleIDHandler sets the operation handler for the patch sources ID roles role ID operation
	PatchSourcesIDRolesRoleIDHandler PatchSourcesIDRolesRoleIDHandler
	// PatchSourcesIDUsersUserIDHandler sets the operation handler for the patch sources ID users user ID operation
	PatchSourcesIDUsersUserIDHandler PatchSourcesIDUsersUserIDHandler
	// PatchSourcesIDUsersUserIDExplorationsExplorationIDHandler sets the operation handler for the patch sources ID users user ID explorations exploration ID operation
	PatchSourcesIDUsersUserIDExplorationsExplorationIDHandler PatchSourcesIDUsersUserIDExplorationsExplorationIDHandler
	// PostLayoutsHandler sets the operation handler for the post layouts operation
	PostLayoutsHandler PostLayoutsHandler
	// PostSourcesHandler sets the operation handler for the post sources operation
	PostSourcesHandler PostSourcesHandler
	// PostSourcesIDKapacitorsHandler sets the operation handler for the post sources ID kapacitors operation
	PostSourcesIDKapacitorsHandler PostSourcesIDKapacitorsHandler
	// PostSourcesIDKapacitorsKapaIDProxyHandler sets the operation handler for the post sources ID kapacitors kapa ID proxy operation
	PostSourcesIDKapacitorsKapaIDProxyHandler PostSourcesIDKapacitorsKapaIDProxyHandler
	// PostSourcesIDProxyHandler sets the operation handler for the post sources ID proxy operation
	PostSourcesIDProxyHandler PostSourcesIDProxyHandler
	// PostSourcesIDRolesHandler sets the operation handler for the post sources ID roles operation
	PostSourcesIDRolesHandler PostSourcesIDRolesHandler
	// PostSourcesIDUsersHandler sets the operation handler for the post sources ID users operation
	PostSourcesIDUsersHandler PostSourcesIDUsersHandler
	// PostSourcesIDUsersUserIDExplorationsHandler sets the operation handler for the post sources ID users user ID explorations operation
	PostSourcesIDUsersUserIDExplorationsHandler PostSourcesIDUsersUserIDExplorationsHandler
	// PutLayoutsIDHandler sets the operation handler for the put layouts ID operation
	PutLayoutsIDHandler PutLayoutsIDHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ChronografAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ChronografAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ChronografAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ChronografAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ChronografAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ChronografAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ChronografAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ChronografAPI
func (o *ChronografAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DeleteLayoutsIDHandler == nil {
		unregistered = append(unregistered, "DeleteLayoutsIDHandler")
	}

	if o.DeleteSourcesIDHandler == nil {
		unregistered = append(unregistered, "DeleteSourcesIDHandler")
	}

	if o.DeleteSourcesIDKapacitorsKapaIDHandler == nil {
		unregistered = append(unregistered, "DeleteSourcesIDKapacitorsKapaIDHandler")
	}

	if o.DeleteSourcesIDKapacitorsKapaIDProxyHandler == nil {
		unregistered = append(unregistered, "DeleteSourcesIDKapacitorsKapaIDProxyHandler")
	}

	if o.DeleteSourcesIDRolesRoleIDHandler == nil {
		unregistered = append(unregistered, "DeleteSourcesIDRolesRoleIDHandler")
	}

	if o.DeleteSourcesIDUsersUserIDHandler == nil {
		unregistered = append(unregistered, "DeleteSourcesIDUsersUserIDHandler")
	}

	if o.DeleteSourcesIDUsersUserIDExplorationsExplorationIDHandler == nil {
		unregistered = append(unregistered, "DeleteSourcesIDUsersUserIDExplorationsExplorationIDHandler")
	}

	if o.GetHandler == nil {
		unregistered = append(unregistered, "GetHandler")
	}

	if o.GetLayoutsHandler == nil {
		unregistered = append(unregistered, "GetLayoutsHandler")
	}

	if o.GetLayoutsIDHandler == nil {
		unregistered = append(unregistered, "GetLayoutsIDHandler")
	}

	if o.GetMappingsHandler == nil {
		unregistered = append(unregistered, "GetMappingsHandler")
	}

	if o.GetSourcesHandler == nil {
		unregistered = append(unregistered, "GetSourcesHandler")
	}

	if o.GetSourcesIDHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDHandler")
	}

	if o.GetSourcesIDKapacitorsHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDKapacitorsHandler")
	}

	if o.GetSourcesIDKapacitorsKapaIDHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDKapacitorsKapaIDHandler")
	}

	if o.GetSourcesIDKapacitorsKapaIDProxyHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDKapacitorsKapaIDProxyHandler")
	}

	if o.GetSourcesIDMonitoredHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDMonitoredHandler")
	}

	if o.GetSourcesIDPermissionsHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDPermissionsHandler")
	}

	if o.GetSourcesIDRolesHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDRolesHandler")
	}

	if o.GetSourcesIDRolesRoleIDHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDRolesRoleIDHandler")
	}

	if o.GetSourcesIDUsersHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDUsersHandler")
	}

	if o.GetSourcesIDUsersUserIDHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDUsersUserIDHandler")
	}

	if o.GetSourcesIDUsersUserIDExplorationsHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDUsersUserIDExplorationsHandler")
	}

	if o.GetSourcesIDUsersUserIDExplorationsExplorationIDHandler == nil {
		unregistered = append(unregistered, "GetSourcesIDUsersUserIDExplorationsExplorationIDHandler")
	}

	if o.GetTokenHandler == nil {
		unregistered = append(unregistered, "GetTokenHandler")
	}

	if o.PatchSourcesIDHandler == nil {
		unregistered = append(unregistered, "PatchSourcesIDHandler")
	}

	if o.PatchSourcesIDKapacitorsKapaIDHandler == nil {
		unregistered = append(unregistered, "PatchSourcesIDKapacitorsKapaIDHandler")
	}

	if o.PatchSourcesIDKapacitorsKapaIDProxyHandler == nil {
		unregistered = append(unregistered, "PatchSourcesIDKapacitorsKapaIDProxyHandler")
	}

	if o.PatchSourcesIDRolesRoleIDHandler == nil {
		unregistered = append(unregistered, "PatchSourcesIDRolesRoleIDHandler")
	}

	if o.PatchSourcesIDUsersUserIDHandler == nil {
		unregistered = append(unregistered, "PatchSourcesIDUsersUserIDHandler")
	}

	if o.PatchSourcesIDUsersUserIDExplorationsExplorationIDHandler == nil {
		unregistered = append(unregistered, "PatchSourcesIDUsersUserIDExplorationsExplorationIDHandler")
	}

	if o.PostLayoutsHandler == nil {
		unregistered = append(unregistered, "PostLayoutsHandler")
	}

	if o.PostSourcesHandler == nil {
		unregistered = append(unregistered, "PostSourcesHandler")
	}

	if o.PostSourcesIDKapacitorsHandler == nil {
		unregistered = append(unregistered, "PostSourcesIDKapacitorsHandler")
	}

	if o.PostSourcesIDKapacitorsKapaIDProxyHandler == nil {
		unregistered = append(unregistered, "PostSourcesIDKapacitorsKapaIDProxyHandler")
	}

	if o.PostSourcesIDProxyHandler == nil {
		unregistered = append(unregistered, "PostSourcesIDProxyHandler")
	}

	if o.PostSourcesIDRolesHandler == nil {
		unregistered = append(unregistered, "PostSourcesIDRolesHandler")
	}

	if o.PostSourcesIDUsersHandler == nil {
		unregistered = append(unregistered, "PostSourcesIDUsersHandler")
	}

	if o.PostSourcesIDUsersUserIDExplorationsHandler == nil {
		unregistered = append(unregistered, "PostSourcesIDUsersUserIDExplorationsHandler")
	}

	if o.PutLayoutsIDHandler == nil {
		unregistered = append(unregistered, "PutLayoutsIDHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ChronografAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ChronografAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ChronografAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ChronografAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ChronografAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the chronograf API
func (o *ChronografAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ChronografAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/layouts/{id}"] = NewDeleteLayoutsID(o.context, o.DeleteLayoutsIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/sources/{id}"] = NewDeleteSourcesID(o.context, o.DeleteSourcesIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/sources/{id}/kapacitors/{kapa_id}"] = NewDeleteSourcesIDKapacitorsKapaID(o.context, o.DeleteSourcesIDKapacitorsKapaIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/sources/{id}/kapacitors/{kapa_id}/proxy"] = NewDeleteSourcesIDKapacitorsKapaIDProxy(o.context, o.DeleteSourcesIDKapacitorsKapaIDProxyHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/sources/{id}/roles/{role_id}"] = NewDeleteSourcesIDRolesRoleID(o.context, o.DeleteSourcesIDRolesRoleIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/sources/{id}/users/{user_id}"] = NewDeleteSourcesIDUsersUserID(o.context, o.DeleteSourcesIDUsersUserIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/sources/{id}/users/{user_id}/explorations/{exploration_id}"] = NewDeleteSourcesIDUsersUserIDExplorationsExplorationID(o.context, o.DeleteSourcesIDUsersUserIDExplorationsExplorationIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/"] = NewGet(o.context, o.GetHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/layouts"] = NewGetLayouts(o.context, o.GetLayoutsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/layouts/{id}"] = NewGetLayoutsID(o.context, o.GetLayoutsIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/mappings"] = NewGetMappings(o.context, o.GetMappingsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources"] = NewGetSources(o.context, o.GetSourcesHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}"] = NewGetSourcesID(o.context, o.GetSourcesIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/kapacitors"] = NewGetSourcesIDKapacitors(o.context, o.GetSourcesIDKapacitorsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/kapacitors/{kapa_id}"] = NewGetSourcesIDKapacitorsKapaID(o.context, o.GetSourcesIDKapacitorsKapaIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/kapacitors/{kapa_id}/proxy"] = NewGetSourcesIDKapacitorsKapaIDProxy(o.context, o.GetSourcesIDKapacitorsKapaIDProxyHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/monitored"] = NewGetSourcesIDMonitored(o.context, o.GetSourcesIDMonitoredHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/permissions"] = NewGetSourcesIDPermissions(o.context, o.GetSourcesIDPermissionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/roles"] = NewGetSourcesIDRoles(o.context, o.GetSourcesIDRolesHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/roles/{role_id}"] = NewGetSourcesIDRolesRoleID(o.context, o.GetSourcesIDRolesRoleIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/users"] = NewGetSourcesIDUsers(o.context, o.GetSourcesIDUsersHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/users/{user_id}"] = NewGetSourcesIDUsersUserID(o.context, o.GetSourcesIDUsersUserIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/users/{user_id}/explorations"] = NewGetSourcesIDUsersUserIDExplorations(o.context, o.GetSourcesIDUsersUserIDExplorationsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/{id}/users/{user_id}/explorations/{exploration_id}"] = NewGetSourcesIDUsersUserIDExplorationsExplorationID(o.context, o.GetSourcesIDUsersUserIDExplorationsExplorationIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/token"] = NewGetToken(o.context, o.GetTokenHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers[strings.ToUpper("PATCH")] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/sources/{id}"] = NewPatchSourcesID(o.context, o.PatchSourcesIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers[strings.ToUpper("PATCH")] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/sources/{id}/kapacitors/{kapa_id}"] = NewPatchSourcesIDKapacitorsKapaID(o.context, o.PatchSourcesIDKapacitorsKapaIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers[strings.ToUpper("PATCH")] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/sources/{id}/kapacitors/{kapa_id}/proxy"] = NewPatchSourcesIDKapacitorsKapaIDProxy(o.context, o.PatchSourcesIDKapacitorsKapaIDProxyHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers[strings.ToUpper("PATCH")] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/sources/{id}/roles/{role_id}"] = NewPatchSourcesIDRolesRoleID(o.context, o.PatchSourcesIDRolesRoleIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers[strings.ToUpper("PATCH")] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/sources/{id}/users/{user_id}"] = NewPatchSourcesIDUsersUserID(o.context, o.PatchSourcesIDUsersUserIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers[strings.ToUpper("PATCH")] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/sources/{id}/users/{user_id}/explorations/{exploration_id}"] = NewPatchSourcesIDUsersUserIDExplorationsExplorationID(o.context, o.PatchSourcesIDUsersUserIDExplorationsExplorationIDHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/layouts"] = NewPostLayouts(o.context, o.PostLayoutsHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sources"] = NewPostSources(o.context, o.PostSourcesHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sources/{id}/kapacitors"] = NewPostSourcesIDKapacitors(o.context, o.PostSourcesIDKapacitorsHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sources/{id}/kapacitors/{kapa_id}/proxy"] = NewPostSourcesIDKapacitorsKapaIDProxy(o.context, o.PostSourcesIDKapacitorsKapaIDProxyHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sources/{id}/proxy"] = NewPostSourcesIDProxy(o.context, o.PostSourcesIDProxyHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sources/{id}/roles"] = NewPostSourcesIDRoles(o.context, o.PostSourcesIDRolesHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sources/{id}/users"] = NewPostSourcesIDUsers(o.context, o.PostSourcesIDUsersHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sources/{id}/users/{user_id}/explorations"] = NewPostSourcesIDUsersUserIDExplorations(o.context, o.PostSourcesIDUsersUserIDExplorationsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/layouts/{id}"] = NewPutLayoutsID(o.context, o.PutLayoutsIDHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ChronografAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *ChronografAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
