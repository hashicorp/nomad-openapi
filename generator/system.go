package main

import "net/http"

func (v *v1api) getSystemPaths() []*apiPath {
	tags := []string{"System"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/system/gc", s.wrap(s.GarbageCollectRequest))
		{
			Template: "/system/gc",
			Operations: []*operation{
				// ATTENTION: nomad system_endpoint only recognizes PUT and not POST
				newOperation(http.MethodPost,
					httpServer.GarbageCollectRequest,
					tags,
					"PostSystemGC",
					nil,
					defaultWriteOpts,
					newResponseConfig(200,
						nilSchema,
						nil,
						defaultWriteMeta,
						"PostSystemGCResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/system/reconcile/summaries", s.wrap(s.ReconcileJobSummaries))
		{
			Template: "/system/reconcile/summaries",
			Operations: []*operation{
				// ATTENTION: nomad system_endpoint only recognizes PUT and not POST
				newOperation(http.MethodPost,
					httpServer.ReconcileJobSummaries,
					tags,
					"PostSystemReconcileSummaries",
					nil,
					defaultWriteOpts,
					newResponseConfig(200,
						nilSchema,
						nil,
						defaultWriteMeta,
						"PostSystemReconcileSummariesResponse",
					),
				),
			},
		},
	}
}
