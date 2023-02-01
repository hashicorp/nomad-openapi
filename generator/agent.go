// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

func (v *v1api) getAgentPaths() []*apiPath {
	// tags := []string{"Agent"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/agent/self", s.wrap(s.AgentSelfRequest))
		//s.mux.HandleFunc("/v1/agent/join", s.wrap(s.AgentJoinRequest))
		//s.mux.HandleFunc("/v1/agent/members", s.wrap(s.AgentMembersRequest))
		//s.mux.HandleFunc("/v1/agent/force-leave", s.wrap(s.AgentForceLeaveRequest))
		//s.mux.HandleFunc("/v1/agent/servers", s.wrap(s.AgentServersRequest))
		//s.mux.HandleFunc("/v1/agent/keyring/", s.wrap(s.KeyringOperationRequest))
		//s.mux.HandleFunc("/v1/agent/health", s.wrap(s.HealthRequest))
		//s.mux.HandleFunc("/v1/agent/host", s.wrap(s.AgentHostRequest))
		//s.mux.HandleFunc("/v1/agent/monitor", s.wrap(s.AgentMonitor))
		//s.mux.HandleFunc("/v1/agent/pprof/", s.wrapNonJSON(s.AgentPprofRequest))
	}
}
