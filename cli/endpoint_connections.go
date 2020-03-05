package cli

import (
	"github.com/digitalrebar/provision/v4/models"
	"github.com/spf13/cobra"
)

func init() {
	addRegistrar(registerEndpointConnections)
}

func registerEndpointConnections(app *cobra.Command) {
	op := &ops{
		name:       "endpoint_connections",
		singleName: "endpoint_connection",
		example:    func() models.Model { return &models.EndpointConnection{} },
	}
	op.command(app)
}
