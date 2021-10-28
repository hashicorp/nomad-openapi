package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getPluginsPaths() []*apiPath {
	tags := []string{"Plugins"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/plugins", s.wrap(s.CSIPluginsRequest))
		{
			Template: "/plugins",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.CSIPluginsRequest,
					tags,
					"GetPlugins",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.CSIPluginListStub{},
						nil,
						"GetPluginsResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/plugin/csi/", s.wrap(s.CSIPluginSpecificRequest))
		{
			Template: "/plugin/csi/{pluginID}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.CSIPluginSpecificRequest,
					tags,
					"GetPluginCSI",
					nil,
					appendParams(defaultQueryOpts, &pluginIDParam),
					newResponseConfig(200,
						arraySchema,
						api.CSIPlugin{},
						defaultQueryMeta,
						"GetPluginCSIResponse",
					),
				),
			},
		},
	}
}
