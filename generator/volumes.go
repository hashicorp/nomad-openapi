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
					appendParams(queryOptions, &volumeNodeIDParam, &volumePluginIDParam, &volumeTypeParam),
					newResponseConfig(200, arraySchema, api.CSIVolumeListStub{}, queryMeta, "GetVolumesResponse"),
				),
				newOperation(http.MethodPost, "csiVolumeRegister", tags, "PostVolume",
					newRequestBody(objectSchema, api.CSIVolumeRegisterRequest{}),
					writeOptions,
					newResponseConfig(200, nilSchema, nil, queryMeta, "PostVolumeResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/volumes/external", s.wrap(s.CSIExternalVolumesRequest))
		{
			Template: "/volumes/external",
			Operations: []*operation{
				newOperation(http.MethodGet, "CSIExternalVolumesRequest", tags, "GetExternalVolumes",
					nil,
					appendParams(queryOptions, &volumePluginIDParam),
					newResponseConfig(200, objectSchema, api.CSIVolumeListExternalResponse{}, queryMeta, "GetExternalVolumesResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/volumes/snapshot", s.wrap(s.CSISnapshotsRequest))
		{
			Template: "/volumes/snapshot",
			Operations: []*operation{
				newOperation(http.MethodGet, "csiSnapshotList", tags, "GetSnapshots",
					nil,
					appendParams(queryOptions, &volumePluginIDParam),
					newResponseConfig(200, objectSchema, api.CSISnapshotListResponse{}, queryMeta,
						"GetSnapshotsResponse"),
				),
				// TODO: See if there is a way to override mismatch between the naming convention and the struct name.
				newOperation(http.MethodPost, "csiSnapshotCreate", tags, "PostSnapshot",
					newRequestBody(objectSchema, api.CSISnapshotCreateRequest{}),
					writeOptions,
					newResponseConfig(200, objectSchema, api.CSISnapshotCreateResponse{}, writeMeta,
						"PostSnapshotResponse"),
				),
				newOperation(http.MethodDelete, "csiSnapshotDelete", tags, "DeleteSnapshot",
					nil,
					appendParams(writeOptions, &volumePluginIDParam, &snapshotIDParam),
					newResponseConfig(200, nilSchema, nil, writeMeta,
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
					appendParams(queryOptions, &volumeIDParam),
					newResponseConfig(200, objectSchema, api.CSIVolume{}, queryMeta,
						"GetVolumeResponse"),
				),
				// TODO: See if there is a way to override mismatch between the naming convention and the struct name.
				newOperation(http.MethodPost, "csiVolumeRegister", tags, "PostVolumeRegistration",
					newRequestBody(objectSchema, api.CSIVolumeRegisterRequest{}),
					appendParams(writeOptions, &volumeIDParam),
					newResponseConfig(200, nilSchema, nil, writeMeta,
						"PostVolumeRegistrationResponse"),
				),
				newOperation(http.MethodDelete, "csiVolumeDeregister", tags, "DeleteVolumeRegistration",
					nil,
					appendParams(writeOptions, &volumeIDParam, &volumeForceParam),
					newResponseConfig(200, nilSchema, nil, writeMeta,
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
					appendParams(writeOptions, &volumeIDParam, &volumeActionParam),
					newResponseConfig(200, objectSchema, api.CSIVolumeCreateResponse{}, queryMeta,
						"PostVolumeResponse"),
				),
				newOperation(http.MethodDelete, "csiVolumeDetach:csiVolumeDelete", tags, "DetachOrDeleteVolume",
					nil,
					appendParams(writeOptions, &volumeIDParam, &volumeNodeParam, &volumeActionParam),
					newResponseConfig(200, nilSchema, nil, queryMeta,
						"DetachOrDeleteVolumeResponse"),
				),
			},
		},
	}
}
