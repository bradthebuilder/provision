package models

/*
 * EndpointConnection defines what this manager should connect to.
 */

// EndpointConnection represents a connection to Endpoint
// swagger:model
type EndpointConnection struct {
	Validation
	Access
	Meta
	Owned
	Bundled

	Id string

	Description   string `json:"Description,omitempty"`
	Documentation string `json:"Documentation,omitempty"`

	// Holds the access parameters.
	Params map[string]interface{} `json:"Params,omitempty"`

	ConnectionStatus string `json:"ConnectionStatus,omitempty"`
}

// GetMeta get the meta data from the model
func (e *EndpointConnection) GetMeta() Meta {
	return e.Meta
}

// SetMeta set the meta data on the model
func (e *EndpointConnection) SetMeta(d Meta) {
	e.Meta = d
}

// Validate validates the object
func (e *EndpointConnection) Validate() {
	e.AddError(ValidName("Invalid Id", e.Id))
}

// Prefix returns the type of object
func (e *EndpointConnection) Prefix() string {
	return "endpoint_connections"
}

// Key returns the key for this object
func (e *EndpointConnection) Key() string {
	return e.Id
}

// KeyName returns the name of the field that is the key for this object
func (e *EndpointConnection) KeyName() string {
	return "Id"
}

// GetDescription returns the models Description
func (e *EndpointConnection) GetDescription() string {
	return e.Description
}

// Fill initials an Endpoint
func (e *EndpointConnection) Fill() {
	if e.Meta == nil {
		e.Meta = Meta{}
	}
	e.Validation.fill()
	if e.Params == nil {
		e.Params = map[string]interface{}{}
	}
}

// AuthKey returns the value of the key for auth purposes
func (e *EndpointConnection) AuthKey() string {
	return e.Key()
}

// SliceOf returns a slice of the model
func (e *EndpointConnection) SliceOf() interface{} {
	s := []*EndpointConnection{}
	return &s
}

// ToModels converts a slice of Endpoints into a slice of Model
func (e *EndpointConnection) ToModels(obj interface{}) []Model {
	items := obj.(*[]*EndpointConnection)
	res := make([]Model, len(*items))
	for i, item := range *items {
		res[i] = Model(item)
	}
	return res
}

// GetParams returns the parameters on the Endpoint
// The returned map is a shallow copy.
func (e *EndpointConnection) GetParams() map[string]interface{} {
	return copyMap(e.Params)
}

// SetParams replaces the current parameters with a shallow
// copy of the input map.
func (e *EndpointConnection) SetParams(p map[string]interface{}) {
	e.Params = copyMap(p)
}

// CanHaveActions indicates that the model can have actions
func (e *EndpointConnection) CanHaveActions() bool {
	return true
}

// SetName sets the name. In this case, it sets Id.
func (e *EndpointConnection) SetName(name string) {
	e.Id = name
}
