package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getVolumesPaths() []*apiPath {
	tags := []string{"Volumes"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/volumes", s.wrap(s.CSIVolumesRequest))
		{
			Template: "/volumes",
			Operations: []*operation{
				newOperation(http.MethodGet, "CSIVolumesRequest", tags, "GetVolumes",
					nil,
					appendParams(defaultQueryOpts, &volumeNodeIDParam, &volumePluginIDParam, &volumeTypeParam),
					newResponseConfig(200, arraySchema, api.CSIVolumeListStub{}, defaultQueryMeta, "GetVolumesResponse"),
				),
				newOperation(http.MethodPost, "csiVolumeRegister", tags, "PostVolume",
					newRequestBody(objectSchema, api.CSIVolumeRegisterRequest{}),
					defaultWriteOpts,
					newResponseConfig(200, nilSchema, nil, defaultQueryMeta, "PostVolumeResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/volumes/external", s.wrap(s.CSIExternalVolumesRequest))
		{
			Template: "/volumes/external",
			Operations: []*operation{
				newOperation(http.MethodGet, "CSIExternalVolumesRequest", tags, "GetExternalVolumes",
					nil,
					appendParams(defaultQueryOpts, &volumePluginIDParam),
					newResponseConfig(200, objectSchema, api.CSIVolumeListExternalResponse{}, defaultQueryMeta, "GetExternalVolumesResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/volumes/snapshot", s.wrap(s.CSISnapshotsRequest))
		{
			Template: "/volumes/snapshot",
			Operations: []*operation{
				newOperation(http.MethodGet, "csiSnapshotList", tags, "GetSnapshots",
					nil,
					appendParams(defaultQueryOpts, &volumePluginIDParam),
					newResponseConfig(200, objectSchema, api.CSISnapshotListResponse{}, defaultQueryMeta,
						"GetSnapshotsResponse"),
				),
				// TODO: See if there is a way to override mismatch between the naming convention and the struct name.
				newOperation(http.MethodPost, "csiSnapshotCreate", tags, "PostSnapshot",
					newRequestBody(objectSchema, api.CSISnapshotCreateRequest{}),
					defaultWriteOpts,
					newResponseConfig(200, objectSchema, api.CSISnapshotCreateResponse{}, defaultWriteMeta,
						"PostSnapshotResponse"),
				),
				newOperation(http.MethodDelete, "csiSnapshotDelete", tags, "DeleteSnapshot",
					nil,
					appendParams(defaultWriteOpts, &volumePluginIDParam, &snapshotIDParam),
					newResponseConfig(200, nilSchema, nil, defaultWriteMeta,
						"DeleteSnapshotResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/volume/csi/", s.wrap(s.CSIVolumeSpecificRequest))
		{
			Template: "/volume/csi/{volumeId}",
			Operations: []*operation{
				newOperation(http.MethodGet, "csiVolumeGet", tags, "GetVolume",
					nil,
					appendParams(defaultQueryOpts, &volumeIDParam),
					newResponseConfig(200, objectSchema, api.CSIVolume{}, defaultQueryMeta,
						"GetVolumeResponse"),
				),
				// TODO: See if there is a way to override mismatch between the naming convention and the struct name.
				newOperation(http.MethodPost, "csiVolumeRegister", tags, "PostVolumeRegistration",
					newRequestBody(objectSchema, api.CSIVolumeRegisterRequest{}),
					appendParams(defaultWriteOpts, &volumeIDParam),
					newResponseConfig(200, nilSchema, nil, defaultWriteMeta,
						"PostVolumeRegistrationResponse"),
				),
				newOperation(http.MethodDelete, "csiVolumeDeregister", tags, "DeleteVolumeRegistration",
					nil,
					appendParams(defaultWriteOpts, &volumeIDParam, &volumeForceParam),
					newResponseConfig(200, nilSchema, nil, defaultWriteMeta,
						"DeleteVolumeRegistrationResponse"),
				),
			},
		},
		{
			Template: "/volume/csi/{volumeId}/{action}",
			Operations: []*operation{
				// TODO: See if there is a way to override mismatch between the naming convention and the struct name.
				newOperation(http.MethodPost, "csiVolumeCreate", tags, "CreateVolume",
					newRequestBody(objectSchema, api.CSIVolumeCreateRequest{}),
					appendParams(defaultWriteOpts, &volumeIDParam, &volumeActionParam),
					newResponseConfig(200, objectSchema, api.CSIVolumeCreateResponse{}, defaultQueryMeta,
						"PostVolumeResponse"),
				),
				newOperation(http.MethodDelete, "csiVolumeDetach:csiVolumeDelete", tags, "DetachOrDeleteVolume",
					nil,
					appendParams(defaultWriteOpts, &volumeIDParam, &volumeNodeParam, &volumeActionParam),
					newResponseConfig(200, nilSchema, nil, defaultQueryMeta,
						"DetachOrDeleteVolumeResponse"),
				),
			},
		},
	}
}
