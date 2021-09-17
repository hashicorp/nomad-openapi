package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
	_ "github.com/hashicorp/nomad/client/allocrunner/taskrunner/getter"
	_ "github.com/hashicorp/nomad/client/allocrunner/taskrunner/template"
	_ "github.com/hashicorp/nomad/client/fingerprint"
	"github.com/hashicorp/nomad/command/agent"
)

func (v *v1api) getAgentPaths() []*apiPath {
	tags := []string{"Agent"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/agent/self", s.wrap(s.AgentSelfRequest))
		{
			Template: "/agent/self",
			Operations: []*operation{
				newOperation(http.MethodGet, "AgentSelfRequest", tags,
					"GetAgentSelf",
					nil,
					[]*parameter{&nomadTokenHeader},
					newResponseConfig(200,
						objectSchema,
						AgentSelf{},
						nil,
						"GetAgentSelfResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/join", s.wrap(s.AgentJoinRequest))
		{
			Template: "/agent/join",
			Operations: []*operation{
				newOperation(http.MethodPost,
					"AgentJoinRequest",
					tags,
					"PostAgentJoin",
					nil,
					appendParams(writeOptions, &agentAddressParam),
					newResponseConfig(200,
						objectSchema,
						JoinResponse{},
						nil,
						"PostAgentJoinResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/members", s.wrap(s.AgentMembersRequest))
		{
			Template: "/agent/members",
			Operations: []*operation{
				newOperation(http.MethodGet,
					"AgentMembersRequest",
					tags,
					"GetAgentMembers",
					nil,
					appendParams(queryOptions),
					newResponseConfig(200,
						objectSchema,
						api.ServerMembers{},
						nil,
						"GetAgentMembersResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/force-leave", s.wrap(s.AgentForceLeaveRequest))
		{
			Template: "/agent/force-leave",
			Operations: []*operation{
				newOperation(http.MethodPost,
					"TODO",
					tags,
					"PostAgentForceLeave",
					nil,
					[]*parameter{&nomadTokenHeader, &nodeNameParam},
					newResponseConfig(200,
						nilSchema,
						nil,
						nil,
						"PostAgentForceLeaveResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/servers", s.wrap(s.AgentServersRequest))
		{
			Template: "/agent/servers",
			Operations: []*operation{
				newOperation(http.MethodGet,
					"listServers",
					tags,
					"GetAgentServers",
					nil,
					[]*parameter{&nomadTokenHeader},
					newResponseConfig(200,
						arraySchema,
						"servernames",
						nil,
						"GetAgentServersResponse",
					),
				),
				newOperation(http.MethodPost,
					"updateServers",
					tags,
					"PostAgentServers",
					nil,
					[]*parameter{&nomadTokenHeader, &agentAddressParam},
					newResponseConfig(200,
						nilSchema,
						nil,
						nil,
						"PostAgentServersResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/keyring/", s.wrap(s.KeyringOperationRequest))
		{
			Template: "/agent/keyring/{action}",
			Operations: []*operation{
				newOperation(http.MethodPost,
					"KeyringOperationRequest",
					tags,
					"PostAgentKeyringAction",
					newRequestBody(objectSchema, api.KeyringRequest{}),
					[]*parameter{&nomadTokenHeader, &keyringActionParam},
					newResponseConfig(200,
						objectSchema,
						api.KeyringResponse{},
						queryMeta,
						"PostAgentKeyringActionResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/health", s.wrap(s.HealthRequest))
		{
			Template: "/agent/health",
			Operations: []*operation{
				newOperation(http.MethodGet,
					"HealthRequest",
					tags,
					"GetAgentHealth",
					nil,
					[]*parameter{&agentTypeParam},
					newResponseConfig(200,
						objectSchema,
						api.AgentHealthResponse{},
						nil,
						"GetAgentHealthResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/host", s.wrap(s.AgentHostRequest))
		{
			Template: "/agent/host",
			Operations: []*operation{
				newOperation(http.MethodGet,
					"AgentHostRequest",
					tags,
					"GetAgentHostData", //TODO: Bad naming
					nil,
					appendParams(queryOptions, &serverIDParam, &nodeIDParam),
					newResponseConfig(200,
						objectSchema,
						api.HostDataResponse{},
						nil,
						"GetAgentHostDataResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/agent/monitor", s.wrap(s.AgentMonitor))
		// TODO: This seems to be a streaming request, so punting for now.
		//s.mux.HandleFunc("/v1/agent/pprof/", s.wrapNonJSON(s.AgentPprofRequest))
		{
			Template: "/agent/pprof/{type}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					"agentPprof",
					tags,
					"GetAgentPprof",
					nil,
					appendParams(queryOptions, &nodeIDParam, &serverIDParam, &pprofProfileParam, &pprofTypeParam),
					newResponseConfig(200,
						arraySchema,
						[]byte{}, // TODO: Test. This may take tweaking.
						nil,      // TODO: Returns a variety of content headers. We'll see if we need them. May need to split into different endpoints by action to support this.
						"GetAgentPprofResponse",
					),
				),
			},
		},
	}
}

// JoinResponse is used to decode the response we get while
// sending a member join request. Helper struct exposes internal type
type JoinResponse struct {
	NumJoined int    `json:"num_joined"`
	Error     string `json:"error"`
}

// AgentSelf is used to report information about an agent and its configuration.
// Helper struct exposes internal type that is returned by the endpoint.
type AgentSelf struct {
	Config *agent.Config                `json:"config"`
	Member agent.Member                 `json:"member,omitempty"`
	Stats  map[string]map[string]string `json:"stats"`
}
