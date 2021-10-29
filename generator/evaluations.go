package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getEvaluationsPaths() []*apiPath {
	tags := []string{"Evaluations"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/evaluations", s.wrap(s.EvalsRequest))
		{
			Template: "/evaluations",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.EvalsRequest,
					tags,
					"GetEvaluations",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.Evaluation{},
						defaultQueryMeta,
						"GetEvaluationsResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/evaluation/", s.wrap(s.EvalSpecificRequest))
		{
			Template: "/evaluation/{evalID}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.EvalSpecificRequest,
					tags,
					"GetEvaluation",
					nil,
					appendParams(defaultQueryOpts, &evalIDParam),
					newResponseConfig(200,
						objectSchema,
						api.Evaluation{},
						defaultQueryMeta,
						"GetEvaluationResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/evaluation/{evalID}/allocations", s.wrap(s.EvalSpecificRequest))
		{
			Template: "/evaluation/{evalID}/allocations",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.EvalSpecificRequest,
					tags,
					"GetEvaluationAllocations",
					nil,
					appendParams(defaultQueryOpts, &evalIDParam),
					newResponseConfig(200,
						arraySchema,
						api.AllocationListStub{},
						defaultQueryMeta,
						"GetEvaluationAllocationsResponse",
					),
				),
			},
		},
	}
}
