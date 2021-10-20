package main

import (
	"net/http"
)

func (v *v1api) getSystemPaths() []*apiPath {
	tags := []string{"System"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/system/gc", s.wrap(s.GarbageCollectRequest))
		{
			Template: "/system/gc",
			Operations: []*operation{
				newOperation(http.MethodPut,
					httpServer.GarbageCollectRequest,
					tags,
					"PutSystemGC",
					nil,
					defaultWriteOpts,
					newResponseConfig(200,
						nilSchema,
						nil,
						nil,
						"PutSystemGCResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/system/reconcile/summaries", s.wrap(s.ReconcileJobSummaries))
		{
			Template: "/system/reconcile/summaries",
			Operations: []*operation{
				newOperation(http.MethodPut,
					httpServer.ReconcileJobSummaries,
					tags,
					"PutSystemReconcileSummaries",
					nil,
					defaultWriteOpts,
					newResponseConfig(200,
						nilSchema,
						nil,
						nil,
						"PutSystemReconcileSummariesResponse",
					),
				),
			},
		},
	}
}
