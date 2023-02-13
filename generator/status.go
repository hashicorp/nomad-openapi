// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"net/http"
)

func (v *v1api) getStatusPaths() []*apiPath {
	tags := []string{"Status"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/status/leader", s.wrap(s.StatusLeaderRequest))
		{
			Template: "/status/leader",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.StatusLeaderRequest,
					tags,
					"GetStatusLeader",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						objectSchema,
						"string",
						nil,
						"GetStatusLeaderResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/status/peers", s.wrap(s.StatusPeersRequest))
		{
			Template: "/status/peers",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.StatusPeersRequest,
					tags,
					"GetStatusPeers",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						"string",
						nil,
						"GetStatusPeersResponse",
					),
				),
			},
		},
	}
}
