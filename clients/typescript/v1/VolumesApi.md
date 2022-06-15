# nomad-client.VolumesApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createVolume**](VolumesApi.md#createVolume) | **POST** /volume/csi/{volumeId}/{action} | 
[**deleteSnapshot**](VolumesApi.md#deleteSnapshot) | **DELETE** /volumes/snapshot | 
[**deleteVolumeRegistration**](VolumesApi.md#deleteVolumeRegistration) | **DELETE** /volume/csi/{volumeId} | 
[**detachOrDeleteVolume**](VolumesApi.md#detachOrDeleteVolume) | **DELETE** /volume/csi/{volumeId}/{action} | 
[**getExternalVolumes**](VolumesApi.md#getExternalVolumes) | **GET** /volumes/external | 
[**getSnapshots**](VolumesApi.md#getSnapshots) | **GET** /volumes/snapshot | 
[**getVolume**](VolumesApi.md#getVolume) | **GET** /volume/csi/{volumeId} | 
[**getVolumes**](VolumesApi.md#getVolumes) | **GET** /volumes | 
[**postSnapshot**](VolumesApi.md#postSnapshot) | **POST** /volumes/snapshot | 
[**postVolume**](VolumesApi.md#postVolume) | **POST** /volumes | 
[**postVolumeRegistration**](VolumesApi.md#postVolumeRegistration) | **POST** /volume/csi/{volumeId} | 


# **createVolume**
> void createVolume(cSIVolumeCreateRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiCreateVolumeRequest = {
  // string | Volume unique identifier.
  volumeId: "volumeId_example",
  // string | The action to perform on the Volume (create, detach, delete).
  action: "action_example",
  // CSIVolumeCreateRequest
  cSIVolumeCreateRequest: {
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
    volumes: [
      {
        accessMode: "accessMode_example",
        allocations: [
          {
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            jobID: "jobID_example",
            jobType: "jobType_example",
            jobVersion: 0,
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            taskGroup: "taskGroup_example",
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        ],
        attachmentMode: "attachmentMode_example",
        capacity: 1,
        cloneID: "cloneID_example",
        context: {
          "key": "key_example",
        },
        controllerRequired: true,
        controllersExpected: 1,
        controllersHealthy: 1,
        createIndex: 0,
        externalID: "externalID_example",
        ID: "ID_example",
        modifyIndex: 0,
        mountOptions: {
          fSType: "fSType_example",
          mountFlags: [
            "mountFlags_example",
          ],
        },
        name: "name_example",
        namespace: "namespace_example",
        nodesExpected: 1,
        nodesHealthy: 1,
        parameters: {
          "key": "key_example",
        },
        pluginID: "pluginID_example",
        provider: "provider_example",
        providerVersion: "providerVersion_example",
        readAllocs: {
          "key": {
            allocModifyIndex: 0,
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentID: "deploymentID_example",
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            desiredTransition: {
              migrate: true,
              reschedule: true,
            },
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            job: {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              allAtOnce: true,
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              consulNamespace: "consulNamespace_example",
              consulToken: "consulToken_example",
              createIndex: 0,
              datacenters: [
                "datacenters_example",
              ],
              dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
              dispatched: true,
              ID: "ID_example",
              jobModifyIndex: 0,
              meta: {
                "key": "key_example",
              },
              migrate: {
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
              },
              modifyIndex: 0,
              multiregion: {
                regions: [
                  {
                    count: 1,
                    datacenters: [
                      "datacenters_example",
                    ],
                    meta: {
                      "key": "key_example",
                    },
                    name: "name_example",
                  },
                ],
                strategy: {
                  maxParallel: 1,
                  onFailure: "onFailure_example",
                },
              },
              name: "name_example",
              namespace: "namespace_example",
              nomadTokenID: "nomadTokenID_example",
              parameterizedJob: {
                metaOptional: [
                  "metaOptional_example",
                ],
                metaRequired: [
                  "metaRequired_example",
                ],
                payload: "payload_example",
              },
              parentID: "parentID_example",
              payload: 'YQ==',
              periodic: {
                enabled: true,
                prohibitOverlap: true,
                spec: "spec_example",
                specType: "specType_example",
                timeZone: "timeZone_example",
              },
              priority: 1,
              region: "region_example",
              reschedule: {
                attempts: 1,
                delay: 1,
                delayFunction: "delayFunction_example",
                interval: 1,
                maxDelay: 1,
                unlimited: true,
              },
              spreads: [
                {
                  attribute: "attribute_example",
                  spreadTarget: [
                    {
                      percent: 0,
                      value: "value_example",
                    },
                  ],
                  weight: -128,
                },
              ],
              stable: true,
              status: "status_example",
              statusDescription: "statusDescription_example",
              stop: true,
              submitTime: 1,
              taskGroups: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  consul: {
                    namespace: "namespace_example",
                  },
                  count: 1,
                  ephemeralDisk: {
                    migrate: true,
                    sizeMB: 1,
                    sticky: true,
                  },
                  maxClientDisconnect: 1,
                  meta: {
                    "key": "key_example",
                  },
                  migrate: {
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                  },
                  name: "name_example",
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                  reschedulePolicy: {
                    attempts: 1,
                    delay: 1,
                    delayFunction: "delayFunction_example",
                    interval: 1,
                    maxDelay: 1,
                    unlimited: true,
                  },
                  restartPolicy: {
                    attempts: 1,
                    delay: 1,
                    interval: 1,
                    mode: "mode_example",
                  },
                  scaling: {
                    createIndex: 0,
                    enabled: true,
                    ID: "ID_example",
                    max: 1,
                    min: 1,
                    modifyIndex: 0,
                    namespace: "namespace_example",
                    policy: {
                      "key": null,
                    },
                    target: {
                      "key": "key_example",
                    },
                    type: "type_example",
                  },
                  services: [
                    {
                      address: "address_example",
                      addressMode: "addressMode_example",
                      canaryMeta: {
                        "key": "key_example",
                      },
                      canaryTags: [
                        "canaryTags_example",
                      ],
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      checks: [
                        {
                          addressMode: "addressMode_example",
                          advertise: "advertise_example",
                          args: [
                            "args_example",
                          ],
                          body: "body_example",
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          command: "command_example",
                          expose: true,
                          failuresBeforeCritical: 1,
                          gRPCService: "gRPCService_example",
                          gRPCUseTLS: true,
                          header: {
                            "key": [
                              "key_example",
                            ],
                          },
                          initialStatus: "initialStatus_example",
                          interval: 1,
                          method: "method_example",
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          path: "path_example",
                          portLabel: "portLabel_example",
                          protocol: "protocol_example",
                          successBeforePassing: 1,
                          tLSSkipVerify: true,
                          taskName: "taskName_example",
                          timeout: 1,
                          type: "type_example",
                        },
                      ],
                      connect: {
                        gateway: {
                          ingress: {
                            listeners: [
                              {
                                port: 1,
                                protocol: "protocol_example",
                                services: [
                                  {
                                    hosts: [
                                      "hosts_example",
                                    ],
                                    name: "name_example",
                                  },
                                ],
                              },
                            ],
                            TLS: {
                              enabled: true,
                            },
                          },
                          mesh: null,
                          proxy: {
                            config: {
                              "key": null,
                            },
                            connectTimeout: 1,
                            envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                            envoyGatewayBindAddresses: {
                              "key": {
                                address: "address_example",
                                name: "name_example",
                                port: 1,
                              },
                            },
                            envoyGatewayBindTaggedAddresses: true,
                            envoyGatewayNoDefaultBind: true,
                          },
                          terminating: {
                            services: [
                              {
                                cAFile: "cAFile_example",
                                certFile: "certFile_example",
                                keyFile: "keyFile_example",
                                name: "name_example",
                                SNI: "SNI_example",
                              },
                            ],
                          },
                        },
                        _native: true,
                        sidecarService: {
                          disableDefaultTCPCheck: true,
                          port: "port_example",
                          proxy: {
                            config: {
                              "key": null,
                            },
                            exposeConfig: {
                              path: [
                                {
                                  listenerPort: "listenerPort_example",
                                  localPathPort: 1,
                                  path: "path_example",
                                  protocol: "protocol_example",
                                },
                              ],
                            },
                            localServiceAddress: "localServiceAddress_example",
                            localServicePort: 1,
                            upstreams: [
                              {
                                datacenter: "datacenter_example",
                                destinationName: "destinationName_example",
                                localBindAddress: "localBindAddress_example",
                                localBindPort: 1,
                                meshGateway: {
                                  mode: "mode_example",
                                },
                              },
                            ],
                          },
                          tags: [
                            "tags_example",
                          ],
                        },
                        sidecarTask: {
                          config: {
                            "key": null,
                          },
                          driver: "driver_example",
                          env: {
                            "key": "key_example",
                          },
                          killSignal: "killSignal_example",
                          killTimeout: 1,
                          logConfig: {
                            maxFileSizeMB: 1,
                            maxFiles: 1,
                          },
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          resources: {
                            CPU: 1,
                            cores: 1,
                            devices: [
                              {
                                affinities: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                    weight: -128,
                                  },
                                ],
                                constraints: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                  },
                                ],
                                count: 0,
                                name: "name_example",
                              },
                            ],
                            diskMB: 1,
                            IOPS: 1,
                            memoryMB: 1,
                            memoryMaxMB: 1,
                            networks: [
                              {
                                CIDR: "CIDR_example",
                                DNS: {
                                  options: [
                                    "options_example",
                                  ],
                                  searches: [
                                    "searches_example",
                                  ],
                                  servers: [
                                    "servers_example",
                                  ],
                                },
                                device: "device_example",
                                dynamicPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                                hostname: "hostname_example",
                                IP: "IP_example",
                                mBits: 1,
                                mode: "mode_example",
                                reservedPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                              },
                            ],
                          },
                          shutdownDelay: 1,
                          user: "user_example",
                        },
                      },
                      enableTagOverride: true,
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      portLabel: "portLabel_example",
                      provider: "provider_example",
                      tags: [
                        "tags_example",
                      ],
                      taskName: "taskName_example",
                    },
                  ],
                  shutdownDelay: 1,
                  spreads: [
                    {
                      attribute: "attribute_example",
                      spreadTarget: [
                        {
                          percent: 0,
                          value: "value_example",
                        },
                      ],
                      weight: -128,
                    },
                  ],
                  stopAfterClientDisconnect: 1,
                  tasks: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      artifacts: [
                        {
                          getterHeaders: {
                            "key": "key_example",
                          },
                          getterMode: "getterMode_example",
                          getterOptions: {
                            "key": "key_example",
                          },
                          getterSource: "getterSource_example",
                          relativeDest: "relativeDest_example",
                        },
                      ],
                      cSIPluginConfig: {
                        ID: "ID_example",
                        mountDir: "mountDir_example",
                        type: "type_example",
                      },
                      config: {
                        "key": null,
                      },
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      dispatchPayload: {
                        file: "file_example",
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      kind: "kind_example",
                      leader: true,
                      lifecycle: {
                        hook: "hook_example",
                        sidecar: true,
                      },
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
                        CPU: 1,
                        cores: 1,
                        devices: [
                          {
                            affinities: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                                weight: -128,
                              },
                            ],
                            constraints: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                              },
                            ],
                            count: 0,
                            name: "name_example",
                          },
                        ],
                        diskMB: 1,
                        IOPS: 1,
                        memoryMB: 1,
                        memoryMaxMB: 1,
                        networks: [
                          {
                            CIDR: "CIDR_example",
                            DNS: {
                              options: [
                                "options_example",
                              ],
                              searches: [
                                "searches_example",
                              ],
                              servers: [
                                "servers_example",
                              ],
                            },
                            device: "device_example",
                            dynamicPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                            hostname: "hostname_example",
                            IP: "IP_example",
                            mBits: 1,
                            mode: "mode_example",
                            reservedPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                          },
                        ],
                      },
                      restartPolicy: {
                        attempts: 1,
                        delay: 1,
                        interval: 1,
                        mode: "mode_example",
                      },
                      scalingPolicies: [
                        {
                          createIndex: 0,
                          enabled: true,
                          ID: "ID_example",
                          max: 1,
                          min: 1,
                          modifyIndex: 0,
                          namespace: "namespace_example",
                          policy: {
                            "key": null,
                          },
                          target: {
                            "key": "key_example",
                          },
                          type: "type_example",
                        },
                      ],
                      services: [
                        {
                          address: "address_example",
                          addressMode: "addressMode_example",
                          canaryMeta: {
                            "key": "key_example",
                          },
                          canaryTags: [
                            "canaryTags_example",
                          ],
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          checks: [
                            {
                              addressMode: "addressMode_example",
                              advertise: "advertise_example",
                              args: [
                                "args_example",
                              ],
                              body: "body_example",
                              checkRestart: {
                                grace: 1,
                                ignoreWarnings: true,
                                limit: 1,
                              },
                              command: "command_example",
                              expose: true,
                              failuresBeforeCritical: 1,
                              gRPCService: "gRPCService_example",
                              gRPCUseTLS: true,
                              header: {
                                "key": [
                                  "key_example",
                                ],
                              },
                              initialStatus: "initialStatus_example",
                              interval: 1,
                              method: "method_example",
                              name: "name_example",
                              onUpdate: "onUpdate_example",
                              path: "path_example",
                              portLabel: "portLabel_example",
                              protocol: "protocol_example",
                              successBeforePassing: 1,
                              tLSSkipVerify: true,
                              taskName: "taskName_example",
                              timeout: 1,
                              type: "type_example",
                            },
                          ],
                          connect: {
                            gateway: {
                              ingress: {
                                listeners: [
                                  {
                                    port: 1,
                                    protocol: "protocol_example",
                                    services: [
                                      {
                                        hosts: [
                                          "hosts_example",
                                        ],
                                        name: "name_example",
                                      },
                                    ],
                                  },
                                ],
                                TLS: {
                                  enabled: true,
                                },
                              },
                              mesh: null,
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                connectTimeout: 1,
                                envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                                envoyGatewayBindAddresses: {
                                  "key": {
                                    address: "address_example",
                                    name: "name_example",
                                    port: 1,
                                  },
                                },
                                envoyGatewayBindTaggedAddresses: true,
                                envoyGatewayNoDefaultBind: true,
                              },
                              terminating: {
                                services: [
                                  {
                                    cAFile: "cAFile_example",
                                    certFile: "certFile_example",
                                    keyFile: "keyFile_example",
                                    name: "name_example",
                                    SNI: "SNI_example",
                                  },
                                ],
                              },
                            },
                            _native: true,
                            sidecarService: {
                              disableDefaultTCPCheck: true,
                              port: "port_example",
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                exposeConfig: {
                                  path: [
                                    {
                                      listenerPort: "listenerPort_example",
                                      localPathPort: 1,
                                      path: "path_example",
                                      protocol: "protocol_example",
                                    },
                                  ],
                                },
                                localServiceAddress: "localServiceAddress_example",
                                localServicePort: 1,
                                upstreams: [
                                  {
                                    datacenter: "datacenter_example",
                                    destinationName: "destinationName_example",
                                    localBindAddress: "localBindAddress_example",
                                    localBindPort: 1,
                                    meshGateway: {
                                      mode: "mode_example",
                                    },
                                  },
                                ],
                              },
                              tags: [
                                "tags_example",
                              ],
                            },
                            sidecarTask: {
                              config: {
                                "key": null,
                              },
                              driver: "driver_example",
                              env: {
                                "key": "key_example",
                              },
                              killSignal: "killSignal_example",
                              killTimeout: 1,
                              logConfig: {
                                maxFileSizeMB: 1,
                                maxFiles: 1,
                              },
                              meta: {
                                "key": "key_example",
                              },
                              name: "name_example",
                              resources: {
                                CPU: 1,
                                cores: 1,
                                devices: [
                                  {
                                    affinities: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                        weight: -128,
                                      },
                                    ],
                                    constraints: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                      },
                                    ],
                                    count: 0,
                                    name: "name_example",
                                  },
                                ],
                                diskMB: 1,
                                IOPS: 1,
                                memoryMB: 1,
                                memoryMaxMB: 1,
                                networks: [
                                  {
                                    CIDR: "CIDR_example",
                                    DNS: {
                                      options: [
                                        "options_example",
                                      ],
                                      searches: [
                                        "searches_example",
                                      ],
                                      servers: [
                                        "servers_example",
                                      ],
                                    },
                                    device: "device_example",
                                    dynamicPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                    hostname: "hostname_example",
                                    IP: "IP_example",
                                    mBits: 1,
                                    mode: "mode_example",
                                    reservedPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                  },
                                ],
                              },
                              shutdownDelay: 1,
                              user: "user_example",
                            },
                          },
                          enableTagOverride: true,
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          portLabel: "portLabel_example",
                          provider: "provider_example",
                          tags: [
                            "tags_example",
                          ],
                          taskName: "taskName_example",
                        },
                      ],
                      shutdownDelay: 1,
                      templates: [
                        {
                          changeMode: "changeMode_example",
                          changeSignal: "changeSignal_example",
                          destPath: "destPath_example",
                          embeddedTmpl: "embeddedTmpl_example",
                          envvars: true,
                          leftDelim: "leftDelim_example",
                          perms: "perms_example",
                          rightDelim: "rightDelim_example",
                          sourcePath: "sourcePath_example",
                          splay: 1,
                          vaultGrace: 1,
                          wait: {
                            max: 1,
                            min: 1,
                          },
                        },
                      ],
                      user: "user_example",
                      vault: {
                        changeMode: "changeMode_example",
                        changeSignal: "changeSignal_example",
                        env: true,
                        namespace: "namespace_example",
                        policies: [
                          "policies_example",
                        ],
                      },
                      volumeMounts: [
                        {
                          destination: "destination_example",
                          propagationMode: "propagationMode_example",
                          readOnly: true,
                          volume: "volume_example",
                        },
                      ],
                    },
                  ],
                  update: {
                    autoPromote: true,
                    autoRevert: true,
                    canary: 1,
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                    progressDeadline: 1,
                    stagger: 1,
                  },
                  volumes: {
                    "key": {
                      accessMode: "accessMode_example",
                      attachmentMode: "attachmentMode_example",
                      mountOptions: {
                        fSType: "fSType_example",
                        mountFlags: [
                          "mountFlags_example",
                        ],
                      },
                      name: "name_example",
                      perAlloc: true,
                      readOnly: true,
                      source: "source_example",
                      type: "type_example",
                    },
                  },
                },
              ],
              type: "type_example",
              update: {
                autoPromote: true,
                autoRevert: true,
                canary: 1,
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
                progressDeadline: 1,
                stagger: 1,
              },
              vaultNamespace: "vaultNamespace_example",
              vaultToken: "vaultToken_example",
              version: 0,
            },
            jobID: "jobID_example",
            metrics: {
              allocationTime: 1,
              classExhausted: {
                "key": 1,
              },
              classFiltered: {
                "key": 1,
              },
              coalescedFailures: 1,
              constraintFiltered: {
                "key": 1,
              },
              dimensionExhausted: {
                "key": 1,
              },
              nodesAvailable: {
                "key": 1,
              },
              nodesEvaluated: 1,
              nodesExhausted: 1,
              nodesFiltered: 1,
              quotaExhausted: [
                "quotaExhausted_example",
              ],
              resourcesExhausted: {
                "key": {
                  CPU: 1,
                  cores: 1,
                  devices: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      count: 0,
                      name: "name_example",
                    },
                  ],
                  diskMB: 1,
                  IOPS: 1,
                  memoryMB: 1,
                  memoryMaxMB: 1,
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
              scoreMetaData: [
                {
                  nodeID: "nodeID_example",
                  normScore: 3.14,
                  scores: {
                    "key": 3.14,
                  },
                },
              ],
              scores: {
                "key": 3.14,
              },
            },
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nextAllocation: "nextAllocation_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            previousAllocation: "previousAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            resources: {
              CPU: 1,
              cores: 1,
              devices: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  count: 0,
                  name: "name_example",
                },
              ],
              diskMB: 1,
              IOPS: 1,
              memoryMB: 1,
              memoryMaxMB: 1,
              networks: [
                {
                  CIDR: "CIDR_example",
                  DNS: {
                    options: [
                      "options_example",
                    ],
                    searches: [
                      "searches_example",
                    ],
                    servers: [
                      "servers_example",
                    ],
                  },
                  device: "device_example",
                  dynamicPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                  hostname: "hostname_example",
                  IP: "IP_example",
                  mBits: 1,
                  mode: "mode_example",
                  reservedPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                },
              ],
            },
            services: {
              "key": "key_example",
            },
            taskGroup: "taskGroup_example",
            taskResources: {
              "key": {
                CPU: 1,
                cores: 1,
                devices: [
                  {
                    affinities: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                        weight: -128,
                      },
                    ],
                    constraints: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                      },
                    ],
                    count: 0,
                    name: "name_example",
                  },
                ],
                diskMB: 1,
                IOPS: 1,
                memoryMB: 1,
                memoryMaxMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
              },
            },
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        },
        requestedCapabilities: [
          {
            accessMode: "accessMode_example",
            attachmentMode: "attachmentMode_example",
          },
        ],
        requestedCapacityMax: 1,
        requestedCapacityMin: 1,
        requestedTopologies: {
          preferred: [
            {
              segments: {
                "key": "key_example",
              },
            },
          ],
          required: [
            {
              segments: {
                "key": "key_example",
              },
            },
          ],
        },
        resourceExhausted: new Date('1970-01-01T00:00:00.00Z'),
        schedulable: true,
        secrets: {
          "key": "key_example",
        },
        snapshotID: "snapshotID_example",
        topologies: [
          {
            segments: {
              "key": "key_example",
            },
          },
        ],
        writeAllocs: {
          "key": {
            allocModifyIndex: 0,
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentID: "deploymentID_example",
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            desiredTransition: {
              migrate: true,
              reschedule: true,
            },
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            job: {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              allAtOnce: true,
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              consulNamespace: "consulNamespace_example",
              consulToken: "consulToken_example",
              createIndex: 0,
              datacenters: [
                "datacenters_example",
              ],
              dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
              dispatched: true,
              ID: "ID_example",
              jobModifyIndex: 0,
              meta: {
                "key": "key_example",
              },
              migrate: {
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
              },
              modifyIndex: 0,
              multiregion: {
                regions: [
                  {
                    count: 1,
                    datacenters: [
                      "datacenters_example",
                    ],
                    meta: {
                      "key": "key_example",
                    },
                    name: "name_example",
                  },
                ],
                strategy: {
                  maxParallel: 1,
                  onFailure: "onFailure_example",
                },
              },
              name: "name_example",
              namespace: "namespace_example",
              nomadTokenID: "nomadTokenID_example",
              parameterizedJob: {
                metaOptional: [
                  "metaOptional_example",
                ],
                metaRequired: [
                  "metaRequired_example",
                ],
                payload: "payload_example",
              },
              parentID: "parentID_example",
              payload: 'YQ==',
              periodic: {
                enabled: true,
                prohibitOverlap: true,
                spec: "spec_example",
                specType: "specType_example",
                timeZone: "timeZone_example",
              },
              priority: 1,
              region: "region_example",
              reschedule: {
                attempts: 1,
                delay: 1,
                delayFunction: "delayFunction_example",
                interval: 1,
                maxDelay: 1,
                unlimited: true,
              },
              spreads: [
                {
                  attribute: "attribute_example",
                  spreadTarget: [
                    {
                      percent: 0,
                      value: "value_example",
                    },
                  ],
                  weight: -128,
                },
              ],
              stable: true,
              status: "status_example",
              statusDescription: "statusDescription_example",
              stop: true,
              submitTime: 1,
              taskGroups: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  consul: {
                    namespace: "namespace_example",
                  },
                  count: 1,
                  ephemeralDisk: {
                    migrate: true,
                    sizeMB: 1,
                    sticky: true,
                  },
                  maxClientDisconnect: 1,
                  meta: {
                    "key": "key_example",
                  },
                  migrate: {
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                  },
                  name: "name_example",
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                  reschedulePolicy: {
                    attempts: 1,
                    delay: 1,
                    delayFunction: "delayFunction_example",
                    interval: 1,
                    maxDelay: 1,
                    unlimited: true,
                  },
                  restartPolicy: {
                    attempts: 1,
                    delay: 1,
                    interval: 1,
                    mode: "mode_example",
                  },
                  scaling: {
                    createIndex: 0,
                    enabled: true,
                    ID: "ID_example",
                    max: 1,
                    min: 1,
                    modifyIndex: 0,
                    namespace: "namespace_example",
                    policy: {
                      "key": null,
                    },
                    target: {
                      "key": "key_example",
                    },
                    type: "type_example",
                  },
                  services: [
                    {
                      address: "address_example",
                      addressMode: "addressMode_example",
                      canaryMeta: {
                        "key": "key_example",
                      },
                      canaryTags: [
                        "canaryTags_example",
                      ],
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      checks: [
                        {
                          addressMode: "addressMode_example",
                          advertise: "advertise_example",
                          args: [
                            "args_example",
                          ],
                          body: "body_example",
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          command: "command_example",
                          expose: true,
                          failuresBeforeCritical: 1,
                          gRPCService: "gRPCService_example",
                          gRPCUseTLS: true,
                          header: {
                            "key": [
                              "key_example",
                            ],
                          },
                          initialStatus: "initialStatus_example",
                          interval: 1,
                          method: "method_example",
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          path: "path_example",
                          portLabel: "portLabel_example",
                          protocol: "protocol_example",
                          successBeforePassing: 1,
                          tLSSkipVerify: true,
                          taskName: "taskName_example",
                          timeout: 1,
                          type: "type_example",
                        },
                      ],
                      connect: {
                        gateway: {
                          ingress: {
                            listeners: [
                              {
                                port: 1,
                                protocol: "protocol_example",
                                services: [
                                  {
                                    hosts: [
                                      "hosts_example",
                                    ],
                                    name: "name_example",
                                  },
                                ],
                              },
                            ],
                            TLS: {
                              enabled: true,
                            },
                          },
                          mesh: null,
                          proxy: {
                            config: {
                              "key": null,
                            },
                            connectTimeout: 1,
                            envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                            envoyGatewayBindAddresses: {
                              "key": {
                                address: "address_example",
                                name: "name_example",
                                port: 1,
                              },
                            },
                            envoyGatewayBindTaggedAddresses: true,
                            envoyGatewayNoDefaultBind: true,
                          },
                          terminating: {
                            services: [
                              {
                                cAFile: "cAFile_example",
                                certFile: "certFile_example",
                                keyFile: "keyFile_example",
                                name: "name_example",
                                SNI: "SNI_example",
                              },
                            ],
                          },
                        },
                        _native: true,
                        sidecarService: {
                          disableDefaultTCPCheck: true,
                          port: "port_example",
                          proxy: {
                            config: {
                              "key": null,
                            },
                            exposeConfig: {
                              path: [
                                {
                                  listenerPort: "listenerPort_example",
                                  localPathPort: 1,
                                  path: "path_example",
                                  protocol: "protocol_example",
                                },
                              ],
                            },
                            localServiceAddress: "localServiceAddress_example",
                            localServicePort: 1,
                            upstreams: [
                              {
                                datacenter: "datacenter_example",
                                destinationName: "destinationName_example",
                                localBindAddress: "localBindAddress_example",
                                localBindPort: 1,
                                meshGateway: {
                                  mode: "mode_example",
                                },
                              },
                            ],
                          },
                          tags: [
                            "tags_example",
                          ],
                        },
                        sidecarTask: {
                          config: {
                            "key": null,
                          },
                          driver: "driver_example",
                          env: {
                            "key": "key_example",
                          },
                          killSignal: "killSignal_example",
                          killTimeout: 1,
                          logConfig: {
                            maxFileSizeMB: 1,
                            maxFiles: 1,
                          },
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          resources: {
                            CPU: 1,
                            cores: 1,
                            devices: [
                              {
                                affinities: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                    weight: -128,
                                  },
                                ],
                                constraints: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                  },
                                ],
                                count: 0,
                                name: "name_example",
                              },
                            ],
                            diskMB: 1,
                            IOPS: 1,
                            memoryMB: 1,
                            memoryMaxMB: 1,
                            networks: [
                              {
                                CIDR: "CIDR_example",
                                DNS: {
                                  options: [
                                    "options_example",
                                  ],
                                  searches: [
                                    "searches_example",
                                  ],
                                  servers: [
                                    "servers_example",
                                  ],
                                },
                                device: "device_example",
                                dynamicPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                                hostname: "hostname_example",
                                IP: "IP_example",
                                mBits: 1,
                                mode: "mode_example",
                                reservedPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                              },
                            ],
                          },
                          shutdownDelay: 1,
                          user: "user_example",
                        },
                      },
                      enableTagOverride: true,
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      portLabel: "portLabel_example",
                      provider: "provider_example",
                      tags: [
                        "tags_example",
                      ],
                      taskName: "taskName_example",
                    },
                  ],
                  shutdownDelay: 1,
                  spreads: [
                    {
                      attribute: "attribute_example",
                      spreadTarget: [
                        {
                          percent: 0,
                          value: "value_example",
                        },
                      ],
                      weight: -128,
                    },
                  ],
                  stopAfterClientDisconnect: 1,
                  tasks: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      artifacts: [
                        {
                          getterHeaders: {
                            "key": "key_example",
                          },
                          getterMode: "getterMode_example",
                          getterOptions: {
                            "key": "key_example",
                          },
                          getterSource: "getterSource_example",
                          relativeDest: "relativeDest_example",
                        },
                      ],
                      cSIPluginConfig: {
                        ID: "ID_example",
                        mountDir: "mountDir_example",
                        type: "type_example",
                      },
                      config: {
                        "key": null,
                      },
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      dispatchPayload: {
                        file: "file_example",
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      kind: "kind_example",
                      leader: true,
                      lifecycle: {
                        hook: "hook_example",
                        sidecar: true,
                      },
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
                        CPU: 1,
                        cores: 1,
                        devices: [
                          {
                            affinities: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                                weight: -128,
                              },
                            ],
                            constraints: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                              },
                            ],
                            count: 0,
                            name: "name_example",
                          },
                        ],
                        diskMB: 1,
                        IOPS: 1,
                        memoryMB: 1,
                        memoryMaxMB: 1,
                        networks: [
                          {
                            CIDR: "CIDR_example",
                            DNS: {
                              options: [
                                "options_example",
                              ],
                              searches: [
                                "searches_example",
                              ],
                              servers: [
                                "servers_example",
                              ],
                            },
                            device: "device_example",
                            dynamicPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                            hostname: "hostname_example",
                            IP: "IP_example",
                            mBits: 1,
                            mode: "mode_example",
                            reservedPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                          },
                        ],
                      },
                      restartPolicy: {
                        attempts: 1,
                        delay: 1,
                        interval: 1,
                        mode: "mode_example",
                      },
                      scalingPolicies: [
                        {
                          createIndex: 0,
                          enabled: true,
                          ID: "ID_example",
                          max: 1,
                          min: 1,
                          modifyIndex: 0,
                          namespace: "namespace_example",
                          policy: {
                            "key": null,
                          },
                          target: {
                            "key": "key_example",
                          },
                          type: "type_example",
                        },
                      ],
                      services: [
                        {
                          address: "address_example",
                          addressMode: "addressMode_example",
                          canaryMeta: {
                            "key": "key_example",
                          },
                          canaryTags: [
                            "canaryTags_example",
                          ],
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          checks: [
                            {
                              addressMode: "addressMode_example",
                              advertise: "advertise_example",
                              args: [
                                "args_example",
                              ],
                              body: "body_example",
                              checkRestart: {
                                grace: 1,
                                ignoreWarnings: true,
                                limit: 1,
                              },
                              command: "command_example",
                              expose: true,
                              failuresBeforeCritical: 1,
                              gRPCService: "gRPCService_example",
                              gRPCUseTLS: true,
                              header: {
                                "key": [
                                  "key_example",
                                ],
                              },
                              initialStatus: "initialStatus_example",
                              interval: 1,
                              method: "method_example",
                              name: "name_example",
                              onUpdate: "onUpdate_example",
                              path: "path_example",
                              portLabel: "portLabel_example",
                              protocol: "protocol_example",
                              successBeforePassing: 1,
                              tLSSkipVerify: true,
                              taskName: "taskName_example",
                              timeout: 1,
                              type: "type_example",
                            },
                          ],
                          connect: {
                            gateway: {
                              ingress: {
                                listeners: [
                                  {
                                    port: 1,
                                    protocol: "protocol_example",
                                    services: [
                                      {
                                        hosts: [
                                          "hosts_example",
                                        ],
                                        name: "name_example",
                                      },
                                    ],
                                  },
                                ],
                                TLS: {
                                  enabled: true,
                                },
                              },
                              mesh: null,
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                connectTimeout: 1,
                                envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                                envoyGatewayBindAddresses: {
                                  "key": {
                                    address: "address_example",
                                    name: "name_example",
                                    port: 1,
                                  },
                                },
                                envoyGatewayBindTaggedAddresses: true,
                                envoyGatewayNoDefaultBind: true,
                              },
                              terminating: {
                                services: [
                                  {
                                    cAFile: "cAFile_example",
                                    certFile: "certFile_example",
                                    keyFile: "keyFile_example",
                                    name: "name_example",
                                    SNI: "SNI_example",
                                  },
                                ],
                              },
                            },
                            _native: true,
                            sidecarService: {
                              disableDefaultTCPCheck: true,
                              port: "port_example",
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                exposeConfig: {
                                  path: [
                                    {
                                      listenerPort: "listenerPort_example",
                                      localPathPort: 1,
                                      path: "path_example",
                                      protocol: "protocol_example",
                                    },
                                  ],
                                },
                                localServiceAddress: "localServiceAddress_example",
                                localServicePort: 1,
                                upstreams: [
                                  {
                                    datacenter: "datacenter_example",
                                    destinationName: "destinationName_example",
                                    localBindAddress: "localBindAddress_example",
                                    localBindPort: 1,
                                    meshGateway: {
                                      mode: "mode_example",
                                    },
                                  },
                                ],
                              },
                              tags: [
                                "tags_example",
                              ],
                            },
                            sidecarTask: {
                              config: {
                                "key": null,
                              },
                              driver: "driver_example",
                              env: {
                                "key": "key_example",
                              },
                              killSignal: "killSignal_example",
                              killTimeout: 1,
                              logConfig: {
                                maxFileSizeMB: 1,
                                maxFiles: 1,
                              },
                              meta: {
                                "key": "key_example",
                              },
                              name: "name_example",
                              resources: {
                                CPU: 1,
                                cores: 1,
                                devices: [
                                  {
                                    affinities: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                        weight: -128,
                                      },
                                    ],
                                    constraints: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                      },
                                    ],
                                    count: 0,
                                    name: "name_example",
                                  },
                                ],
                                diskMB: 1,
                                IOPS: 1,
                                memoryMB: 1,
                                memoryMaxMB: 1,
                                networks: [
                                  {
                                    CIDR: "CIDR_example",
                                    DNS: {
                                      options: [
                                        "options_example",
                                      ],
                                      searches: [
                                        "searches_example",
                                      ],
                                      servers: [
                                        "servers_example",
                                      ],
                                    },
                                    device: "device_example",
                                    dynamicPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                    hostname: "hostname_example",
                                    IP: "IP_example",
                                    mBits: 1,
                                    mode: "mode_example",
                                    reservedPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                  },
                                ],
                              },
                              shutdownDelay: 1,
                              user: "user_example",
                            },
                          },
                          enableTagOverride: true,
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          portLabel: "portLabel_example",
                          provider: "provider_example",
                          tags: [
                            "tags_example",
                          ],
                          taskName: "taskName_example",
                        },
                      ],
                      shutdownDelay: 1,
                      templates: [
                        {
                          changeMode: "changeMode_example",
                          changeSignal: "changeSignal_example",
                          destPath: "destPath_example",
                          embeddedTmpl: "embeddedTmpl_example",
                          envvars: true,
                          leftDelim: "leftDelim_example",
                          perms: "perms_example",
                          rightDelim: "rightDelim_example",
                          sourcePath: "sourcePath_example",
                          splay: 1,
                          vaultGrace: 1,
                          wait: {
                            max: 1,
                            min: 1,
                          },
                        },
                      ],
                      user: "user_example",
                      vault: {
                        changeMode: "changeMode_example",
                        changeSignal: "changeSignal_example",
                        env: true,
                        namespace: "namespace_example",
                        policies: [
                          "policies_example",
                        ],
                      },
                      volumeMounts: [
                        {
                          destination: "destination_example",
                          propagationMode: "propagationMode_example",
                          readOnly: true,
                          volume: "volume_example",
                        },
                      ],
                    },
                  ],
                  update: {
                    autoPromote: true,
                    autoRevert: true,
                    canary: 1,
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                    progressDeadline: 1,
                    stagger: 1,
                  },
                  volumes: {
                    "key": {
                      accessMode: "accessMode_example",
                      attachmentMode: "attachmentMode_example",
                      mountOptions: {
                        fSType: "fSType_example",
                        mountFlags: [
                          "mountFlags_example",
                        ],
                      },
                      name: "name_example",
                      perAlloc: true,
                      readOnly: true,
                      source: "source_example",
                      type: "type_example",
                    },
                  },
                },
              ],
              type: "type_example",
              update: {
                autoPromote: true,
                autoRevert: true,
                canary: 1,
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
                progressDeadline: 1,
                stagger: 1,
              },
              vaultNamespace: "vaultNamespace_example",
              vaultToken: "vaultToken_example",
              version: 0,
            },
            jobID: "jobID_example",
            metrics: {
              allocationTime: 1,
              classExhausted: {
                "key": 1,
              },
              classFiltered: {
                "key": 1,
              },
              coalescedFailures: 1,
              constraintFiltered: {
                "key": 1,
              },
              dimensionExhausted: {
                "key": 1,
              },
              nodesAvailable: {
                "key": 1,
              },
              nodesEvaluated: 1,
              nodesExhausted: 1,
              nodesFiltered: 1,
              quotaExhausted: [
                "quotaExhausted_example",
              ],
              resourcesExhausted: {
                "key": {
                  CPU: 1,
                  cores: 1,
                  devices: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      count: 0,
                      name: "name_example",
                    },
                  ],
                  diskMB: 1,
                  IOPS: 1,
                  memoryMB: 1,
                  memoryMaxMB: 1,
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
              scoreMetaData: [
                {
                  nodeID: "nodeID_example",
                  normScore: 3.14,
                  scores: {
                    "key": 3.14,
                  },
                },
              ],
              scores: {
                "key": 3.14,
              },
            },
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nextAllocation: "nextAllocation_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            previousAllocation: "previousAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            resources: {
              CPU: 1,
              cores: 1,
              devices: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  count: 0,
                  name: "name_example",
                },
              ],
              diskMB: 1,
              IOPS: 1,
              memoryMB: 1,
              memoryMaxMB: 1,
              networks: [
                {
                  CIDR: "CIDR_example",
                  DNS: {
                    options: [
                      "options_example",
                    ],
                    searches: [
                      "searches_example",
                    ],
                    servers: [
                      "servers_example",
                    ],
                  },
                  device: "device_example",
                  dynamicPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                  hostname: "hostname_example",
                  IP: "IP_example",
                  mBits: 1,
                  mode: "mode_example",
                  reservedPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                },
              ],
            },
            services: {
              "key": "key_example",
            },
            taskGroup: "taskGroup_example",
            taskResources: {
              "key": {
                CPU: 1,
                cores: 1,
                devices: [
                  {
                    affinities: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                        weight: -128,
                      },
                    ],
                    constraints: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                      },
                    ],
                    count: 0,
                    name: "name_example",
                  },
                ],
                diskMB: 1,
                IOPS: 1,
                memoryMB: 1,
                memoryMaxMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
              },
            },
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        },
      },
    ],
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.createVolume(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSIVolumeCreateRequest** | **CSIVolumeCreateRequest**|  |
 **volumeId** | [**string**] | Volume unique identifier. | defaults to undefined
 **action** | [**string**] | The action to perform on the Volume (create, detach, delete). | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **deleteSnapshot**
> void deleteSnapshot()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiDeleteSnapshotRequest = {
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
  // string | Filters volume lists by plugin ID. (optional)
  pluginId: "plugin_id_example",
  // string | The ID of the snapshot to target. (optional)
  snapshotId: "snapshot_id_example",
};

apiInstance.deleteSnapshot(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined
 **pluginId** | [**string**] | Filters volume lists by plugin ID. | (optional) defaults to undefined
 **snapshotId** | [**string**] | The ID of the snapshot to target. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **deleteVolumeRegistration**
> void deleteVolumeRegistration()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiDeleteVolumeRegistrationRequest = {
  // string | Volume unique identifier.
  volumeId: "volumeId_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
  // string | Used to force the de-registration of a volume. (optional)
  force: "force_example",
};

apiInstance.deleteVolumeRegistration(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | [**string**] | Volume unique identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined
 **force** | [**string**] | Used to force the de-registration of a volume. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **detachOrDeleteVolume**
> void detachOrDeleteVolume()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiDetachOrDeleteVolumeRequest = {
  // string | Volume unique identifier.
  volumeId: "volumeId_example",
  // string | The action to perform on the Volume (create, detach, delete).
  action: "action_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
  // string | Specifies node to target volume operation for. (optional)
  node: "node_example",
};

apiInstance.detachOrDeleteVolume(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | [**string**] | Volume unique identifier. | defaults to undefined
 **action** | [**string**] | The action to perform on the Volume (create, detach, delete). | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined
 **node** | [**string**] | Specifies node to target volume operation for. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **getExternalVolumes**
> CSIVolumeListExternalResponse getExternalVolumes()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiGetExternalVolumesRequest = {
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
  // string | Filters volume lists by plugin ID. (optional)
  pluginId: "plugin_id_example",
};

apiInstance.getExternalVolumes(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined
 **pluginId** | [**string**] | Filters volume lists by plugin ID. | (optional) defaults to undefined


### Return type

**CSIVolumeListExternalResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **getSnapshots**
> CSISnapshotListResponse getSnapshots()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiGetSnapshotsRequest = {
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
  // string | Filters volume lists by plugin ID. (optional)
  pluginId: "plugin_id_example",
};

apiInstance.getSnapshots(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined
 **pluginId** | [**string**] | Filters volume lists by plugin ID. | (optional) defaults to undefined


### Return type

**CSISnapshotListResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **getVolume**
> CSIVolume getVolume()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiGetVolumeRequest = {
  // string | Volume unique identifier.
  volumeId: "volumeId_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
};

apiInstance.getVolume(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | [**string**] | Volume unique identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined


### Return type

**CSIVolume**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **getVolumes**
> Array<CSIVolumeListStub> getVolumes()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiGetVolumesRequest = {
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
  // string | Filters volume lists by node ID. (optional)
  nodeId: "node_id_example",
  // string | Filters volume lists by plugin ID. (optional)
  pluginId: "plugin_id_example",
  // string | Filters volume lists to a specific type. (optional)
  type: "type_example",
};

apiInstance.getVolumes(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined
 **nodeId** | [**string**] | Filters volume lists by node ID. | (optional) defaults to undefined
 **pluginId** | [**string**] | Filters volume lists by plugin ID. | (optional) defaults to undefined
 **type** | [**string**] | Filters volume lists to a specific type. | (optional) defaults to undefined


### Return type

**Array<CSIVolumeListStub>**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **postSnapshot**
> CSISnapshotCreateResponse postSnapshot(cSISnapshotCreateRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiPostSnapshotRequest = {
  // CSISnapshotCreateRequest
  cSISnapshotCreateRequest: {
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
    snapshots: [
      {
        createTime: 1,
        externalSourceVolumeID: "externalSourceVolumeID_example",
        ID: "ID_example",
        isReady: true,
        name: "name_example",
        parameters: {
          "key": "key_example",
        },
        pluginID: "pluginID_example",
        secrets: {
          "key": "key_example",
        },
        sizeBytes: 1,
        sourceVolumeID: "sourceVolumeID_example",
      },
    ],
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postSnapshot(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSISnapshotCreateRequest** | **CSISnapshotCreateRequest**|  |
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**CSISnapshotCreateResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **postVolume**
> void postVolume(cSIVolumeRegisterRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiPostVolumeRequest = {
  // CSIVolumeRegisterRequest
  cSIVolumeRegisterRequest: {
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
    volumes: [
      {
        accessMode: "accessMode_example",
        allocations: [
          {
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            jobID: "jobID_example",
            jobType: "jobType_example",
            jobVersion: 0,
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            taskGroup: "taskGroup_example",
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        ],
        attachmentMode: "attachmentMode_example",
        capacity: 1,
        cloneID: "cloneID_example",
        context: {
          "key": "key_example",
        },
        controllerRequired: true,
        controllersExpected: 1,
        controllersHealthy: 1,
        createIndex: 0,
        externalID: "externalID_example",
        ID: "ID_example",
        modifyIndex: 0,
        mountOptions: {
          fSType: "fSType_example",
          mountFlags: [
            "mountFlags_example",
          ],
        },
        name: "name_example",
        namespace: "namespace_example",
        nodesExpected: 1,
        nodesHealthy: 1,
        parameters: {
          "key": "key_example",
        },
        pluginID: "pluginID_example",
        provider: "provider_example",
        providerVersion: "providerVersion_example",
        readAllocs: {
          "key": {
            allocModifyIndex: 0,
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentID: "deploymentID_example",
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            desiredTransition: {
              migrate: true,
              reschedule: true,
            },
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            job: {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              allAtOnce: true,
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              consulNamespace: "consulNamespace_example",
              consulToken: "consulToken_example",
              createIndex: 0,
              datacenters: [
                "datacenters_example",
              ],
              dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
              dispatched: true,
              ID: "ID_example",
              jobModifyIndex: 0,
              meta: {
                "key": "key_example",
              },
              migrate: {
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
              },
              modifyIndex: 0,
              multiregion: {
                regions: [
                  {
                    count: 1,
                    datacenters: [
                      "datacenters_example",
                    ],
                    meta: {
                      "key": "key_example",
                    },
                    name: "name_example",
                  },
                ],
                strategy: {
                  maxParallel: 1,
                  onFailure: "onFailure_example",
                },
              },
              name: "name_example",
              namespace: "namespace_example",
              nomadTokenID: "nomadTokenID_example",
              parameterizedJob: {
                metaOptional: [
                  "metaOptional_example",
                ],
                metaRequired: [
                  "metaRequired_example",
                ],
                payload: "payload_example",
              },
              parentID: "parentID_example",
              payload: 'YQ==',
              periodic: {
                enabled: true,
                prohibitOverlap: true,
                spec: "spec_example",
                specType: "specType_example",
                timeZone: "timeZone_example",
              },
              priority: 1,
              region: "region_example",
              reschedule: {
                attempts: 1,
                delay: 1,
                delayFunction: "delayFunction_example",
                interval: 1,
                maxDelay: 1,
                unlimited: true,
              },
              spreads: [
                {
                  attribute: "attribute_example",
                  spreadTarget: [
                    {
                      percent: 0,
                      value: "value_example",
                    },
                  ],
                  weight: -128,
                },
              ],
              stable: true,
              status: "status_example",
              statusDescription: "statusDescription_example",
              stop: true,
              submitTime: 1,
              taskGroups: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  consul: {
                    namespace: "namespace_example",
                  },
                  count: 1,
                  ephemeralDisk: {
                    migrate: true,
                    sizeMB: 1,
                    sticky: true,
                  },
                  maxClientDisconnect: 1,
                  meta: {
                    "key": "key_example",
                  },
                  migrate: {
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                  },
                  name: "name_example",
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                  reschedulePolicy: {
                    attempts: 1,
                    delay: 1,
                    delayFunction: "delayFunction_example",
                    interval: 1,
                    maxDelay: 1,
                    unlimited: true,
                  },
                  restartPolicy: {
                    attempts: 1,
                    delay: 1,
                    interval: 1,
                    mode: "mode_example",
                  },
                  scaling: {
                    createIndex: 0,
                    enabled: true,
                    ID: "ID_example",
                    max: 1,
                    min: 1,
                    modifyIndex: 0,
                    namespace: "namespace_example",
                    policy: {
                      "key": null,
                    },
                    target: {
                      "key": "key_example",
                    },
                    type: "type_example",
                  },
                  services: [
                    {
                      address: "address_example",
                      addressMode: "addressMode_example",
                      canaryMeta: {
                        "key": "key_example",
                      },
                      canaryTags: [
                        "canaryTags_example",
                      ],
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      checks: [
                        {
                          addressMode: "addressMode_example",
                          advertise: "advertise_example",
                          args: [
                            "args_example",
                          ],
                          body: "body_example",
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          command: "command_example",
                          expose: true,
                          failuresBeforeCritical: 1,
                          gRPCService: "gRPCService_example",
                          gRPCUseTLS: true,
                          header: {
                            "key": [
                              "key_example",
                            ],
                          },
                          initialStatus: "initialStatus_example",
                          interval: 1,
                          method: "method_example",
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          path: "path_example",
                          portLabel: "portLabel_example",
                          protocol: "protocol_example",
                          successBeforePassing: 1,
                          tLSSkipVerify: true,
                          taskName: "taskName_example",
                          timeout: 1,
                          type: "type_example",
                        },
                      ],
                      connect: {
                        gateway: {
                          ingress: {
                            listeners: [
                              {
                                port: 1,
                                protocol: "protocol_example",
                                services: [
                                  {
                                    hosts: [
                                      "hosts_example",
                                    ],
                                    name: "name_example",
                                  },
                                ],
                              },
                            ],
                            TLS: {
                              enabled: true,
                            },
                          },
                          mesh: null,
                          proxy: {
                            config: {
                              "key": null,
                            },
                            connectTimeout: 1,
                            envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                            envoyGatewayBindAddresses: {
                              "key": {
                                address: "address_example",
                                name: "name_example",
                                port: 1,
                              },
                            },
                            envoyGatewayBindTaggedAddresses: true,
                            envoyGatewayNoDefaultBind: true,
                          },
                          terminating: {
                            services: [
                              {
                                cAFile: "cAFile_example",
                                certFile: "certFile_example",
                                keyFile: "keyFile_example",
                                name: "name_example",
                                SNI: "SNI_example",
                              },
                            ],
                          },
                        },
                        _native: true,
                        sidecarService: {
                          disableDefaultTCPCheck: true,
                          port: "port_example",
                          proxy: {
                            config: {
                              "key": null,
                            },
                            exposeConfig: {
                              path: [
                                {
                                  listenerPort: "listenerPort_example",
                                  localPathPort: 1,
                                  path: "path_example",
                                  protocol: "protocol_example",
                                },
                              ],
                            },
                            localServiceAddress: "localServiceAddress_example",
                            localServicePort: 1,
                            upstreams: [
                              {
                                datacenter: "datacenter_example",
                                destinationName: "destinationName_example",
                                localBindAddress: "localBindAddress_example",
                                localBindPort: 1,
                                meshGateway: {
                                  mode: "mode_example",
                                },
                              },
                            ],
                          },
                          tags: [
                            "tags_example",
                          ],
                        },
                        sidecarTask: {
                          config: {
                            "key": null,
                          },
                          driver: "driver_example",
                          env: {
                            "key": "key_example",
                          },
                          killSignal: "killSignal_example",
                          killTimeout: 1,
                          logConfig: {
                            maxFileSizeMB: 1,
                            maxFiles: 1,
                          },
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          resources: {
                            CPU: 1,
                            cores: 1,
                            devices: [
                              {
                                affinities: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                    weight: -128,
                                  },
                                ],
                                constraints: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                  },
                                ],
                                count: 0,
                                name: "name_example",
                              },
                            ],
                            diskMB: 1,
                            IOPS: 1,
                            memoryMB: 1,
                            memoryMaxMB: 1,
                            networks: [
                              {
                                CIDR: "CIDR_example",
                                DNS: {
                                  options: [
                                    "options_example",
                                  ],
                                  searches: [
                                    "searches_example",
                                  ],
                                  servers: [
                                    "servers_example",
                                  ],
                                },
                                device: "device_example",
                                dynamicPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                                hostname: "hostname_example",
                                IP: "IP_example",
                                mBits: 1,
                                mode: "mode_example",
                                reservedPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                              },
                            ],
                          },
                          shutdownDelay: 1,
                          user: "user_example",
                        },
                      },
                      enableTagOverride: true,
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      portLabel: "portLabel_example",
                      provider: "provider_example",
                      tags: [
                        "tags_example",
                      ],
                      taskName: "taskName_example",
                    },
                  ],
                  shutdownDelay: 1,
                  spreads: [
                    {
                      attribute: "attribute_example",
                      spreadTarget: [
                        {
                          percent: 0,
                          value: "value_example",
                        },
                      ],
                      weight: -128,
                    },
                  ],
                  stopAfterClientDisconnect: 1,
                  tasks: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      artifacts: [
                        {
                          getterHeaders: {
                            "key": "key_example",
                          },
                          getterMode: "getterMode_example",
                          getterOptions: {
                            "key": "key_example",
                          },
                          getterSource: "getterSource_example",
                          relativeDest: "relativeDest_example",
                        },
                      ],
                      cSIPluginConfig: {
                        ID: "ID_example",
                        mountDir: "mountDir_example",
                        type: "type_example",
                      },
                      config: {
                        "key": null,
                      },
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      dispatchPayload: {
                        file: "file_example",
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      kind: "kind_example",
                      leader: true,
                      lifecycle: {
                        hook: "hook_example",
                        sidecar: true,
                      },
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
                        CPU: 1,
                        cores: 1,
                        devices: [
                          {
                            affinities: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                                weight: -128,
                              },
                            ],
                            constraints: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                              },
                            ],
                            count: 0,
                            name: "name_example",
                          },
                        ],
                        diskMB: 1,
                        IOPS: 1,
                        memoryMB: 1,
                        memoryMaxMB: 1,
                        networks: [
                          {
                            CIDR: "CIDR_example",
                            DNS: {
                              options: [
                                "options_example",
                              ],
                              searches: [
                                "searches_example",
                              ],
                              servers: [
                                "servers_example",
                              ],
                            },
                            device: "device_example",
                            dynamicPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                            hostname: "hostname_example",
                            IP: "IP_example",
                            mBits: 1,
                            mode: "mode_example",
                            reservedPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                          },
                        ],
                      },
                      restartPolicy: {
                        attempts: 1,
                        delay: 1,
                        interval: 1,
                        mode: "mode_example",
                      },
                      scalingPolicies: [
                        {
                          createIndex: 0,
                          enabled: true,
                          ID: "ID_example",
                          max: 1,
                          min: 1,
                          modifyIndex: 0,
                          namespace: "namespace_example",
                          policy: {
                            "key": null,
                          },
                          target: {
                            "key": "key_example",
                          },
                          type: "type_example",
                        },
                      ],
                      services: [
                        {
                          address: "address_example",
                          addressMode: "addressMode_example",
                          canaryMeta: {
                            "key": "key_example",
                          },
                          canaryTags: [
                            "canaryTags_example",
                          ],
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          checks: [
                            {
                              addressMode: "addressMode_example",
                              advertise: "advertise_example",
                              args: [
                                "args_example",
                              ],
                              body: "body_example",
                              checkRestart: {
                                grace: 1,
                                ignoreWarnings: true,
                                limit: 1,
                              },
                              command: "command_example",
                              expose: true,
                              failuresBeforeCritical: 1,
                              gRPCService: "gRPCService_example",
                              gRPCUseTLS: true,
                              header: {
                                "key": [
                                  "key_example",
                                ],
                              },
                              initialStatus: "initialStatus_example",
                              interval: 1,
                              method: "method_example",
                              name: "name_example",
                              onUpdate: "onUpdate_example",
                              path: "path_example",
                              portLabel: "portLabel_example",
                              protocol: "protocol_example",
                              successBeforePassing: 1,
                              tLSSkipVerify: true,
                              taskName: "taskName_example",
                              timeout: 1,
                              type: "type_example",
                            },
                          ],
                          connect: {
                            gateway: {
                              ingress: {
                                listeners: [
                                  {
                                    port: 1,
                                    protocol: "protocol_example",
                                    services: [
                                      {
                                        hosts: [
                                          "hosts_example",
                                        ],
                                        name: "name_example",
                                      },
                                    ],
                                  },
                                ],
                                TLS: {
                                  enabled: true,
                                },
                              },
                              mesh: null,
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                connectTimeout: 1,
                                envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                                envoyGatewayBindAddresses: {
                                  "key": {
                                    address: "address_example",
                                    name: "name_example",
                                    port: 1,
                                  },
                                },
                                envoyGatewayBindTaggedAddresses: true,
                                envoyGatewayNoDefaultBind: true,
                              },
                              terminating: {
                                services: [
                                  {
                                    cAFile: "cAFile_example",
                                    certFile: "certFile_example",
                                    keyFile: "keyFile_example",
                                    name: "name_example",
                                    SNI: "SNI_example",
                                  },
                                ],
                              },
                            },
                            _native: true,
                            sidecarService: {
                              disableDefaultTCPCheck: true,
                              port: "port_example",
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                exposeConfig: {
                                  path: [
                                    {
                                      listenerPort: "listenerPort_example",
                                      localPathPort: 1,
                                      path: "path_example",
                                      protocol: "protocol_example",
                                    },
                                  ],
                                },
                                localServiceAddress: "localServiceAddress_example",
                                localServicePort: 1,
                                upstreams: [
                                  {
                                    datacenter: "datacenter_example",
                                    destinationName: "destinationName_example",
                                    localBindAddress: "localBindAddress_example",
                                    localBindPort: 1,
                                    meshGateway: {
                                      mode: "mode_example",
                                    },
                                  },
                                ],
                              },
                              tags: [
                                "tags_example",
                              ],
                            },
                            sidecarTask: {
                              config: {
                                "key": null,
                              },
                              driver: "driver_example",
                              env: {
                                "key": "key_example",
                              },
                              killSignal: "killSignal_example",
                              killTimeout: 1,
                              logConfig: {
                                maxFileSizeMB: 1,
                                maxFiles: 1,
                              },
                              meta: {
                                "key": "key_example",
                              },
                              name: "name_example",
                              resources: {
                                CPU: 1,
                                cores: 1,
                                devices: [
                                  {
                                    affinities: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                        weight: -128,
                                      },
                                    ],
                                    constraints: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                      },
                                    ],
                                    count: 0,
                                    name: "name_example",
                                  },
                                ],
                                diskMB: 1,
                                IOPS: 1,
                                memoryMB: 1,
                                memoryMaxMB: 1,
                                networks: [
                                  {
                                    CIDR: "CIDR_example",
                                    DNS: {
                                      options: [
                                        "options_example",
                                      ],
                                      searches: [
                                        "searches_example",
                                      ],
                                      servers: [
                                        "servers_example",
                                      ],
                                    },
                                    device: "device_example",
                                    dynamicPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                    hostname: "hostname_example",
                                    IP: "IP_example",
                                    mBits: 1,
                                    mode: "mode_example",
                                    reservedPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                  },
                                ],
                              },
                              shutdownDelay: 1,
                              user: "user_example",
                            },
                          },
                          enableTagOverride: true,
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          portLabel: "portLabel_example",
                          provider: "provider_example",
                          tags: [
                            "tags_example",
                          ],
                          taskName: "taskName_example",
                        },
                      ],
                      shutdownDelay: 1,
                      templates: [
                        {
                          changeMode: "changeMode_example",
                          changeSignal: "changeSignal_example",
                          destPath: "destPath_example",
                          embeddedTmpl: "embeddedTmpl_example",
                          envvars: true,
                          leftDelim: "leftDelim_example",
                          perms: "perms_example",
                          rightDelim: "rightDelim_example",
                          sourcePath: "sourcePath_example",
                          splay: 1,
                          vaultGrace: 1,
                          wait: {
                            max: 1,
                            min: 1,
                          },
                        },
                      ],
                      user: "user_example",
                      vault: {
                        changeMode: "changeMode_example",
                        changeSignal: "changeSignal_example",
                        env: true,
                        namespace: "namespace_example",
                        policies: [
                          "policies_example",
                        ],
                      },
                      volumeMounts: [
                        {
                          destination: "destination_example",
                          propagationMode: "propagationMode_example",
                          readOnly: true,
                          volume: "volume_example",
                        },
                      ],
                    },
                  ],
                  update: {
                    autoPromote: true,
                    autoRevert: true,
                    canary: 1,
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                    progressDeadline: 1,
                    stagger: 1,
                  },
                  volumes: {
                    "key": {
                      accessMode: "accessMode_example",
                      attachmentMode: "attachmentMode_example",
                      mountOptions: {
                        fSType: "fSType_example",
                        mountFlags: [
                          "mountFlags_example",
                        ],
                      },
                      name: "name_example",
                      perAlloc: true,
                      readOnly: true,
                      source: "source_example",
                      type: "type_example",
                    },
                  },
                },
              ],
              type: "type_example",
              update: {
                autoPromote: true,
                autoRevert: true,
                canary: 1,
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
                progressDeadline: 1,
                stagger: 1,
              },
              vaultNamespace: "vaultNamespace_example",
              vaultToken: "vaultToken_example",
              version: 0,
            },
            jobID: "jobID_example",
            metrics: {
              allocationTime: 1,
              classExhausted: {
                "key": 1,
              },
              classFiltered: {
                "key": 1,
              },
              coalescedFailures: 1,
              constraintFiltered: {
                "key": 1,
              },
              dimensionExhausted: {
                "key": 1,
              },
              nodesAvailable: {
                "key": 1,
              },
              nodesEvaluated: 1,
              nodesExhausted: 1,
              nodesFiltered: 1,
              quotaExhausted: [
                "quotaExhausted_example",
              ],
              resourcesExhausted: {
                "key": {
                  CPU: 1,
                  cores: 1,
                  devices: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      count: 0,
                      name: "name_example",
                    },
                  ],
                  diskMB: 1,
                  IOPS: 1,
                  memoryMB: 1,
                  memoryMaxMB: 1,
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
              scoreMetaData: [
                {
                  nodeID: "nodeID_example",
                  normScore: 3.14,
                  scores: {
                    "key": 3.14,
                  },
                },
              ],
              scores: {
                "key": 3.14,
              },
            },
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nextAllocation: "nextAllocation_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            previousAllocation: "previousAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            resources: {
              CPU: 1,
              cores: 1,
              devices: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  count: 0,
                  name: "name_example",
                },
              ],
              diskMB: 1,
              IOPS: 1,
              memoryMB: 1,
              memoryMaxMB: 1,
              networks: [
                {
                  CIDR: "CIDR_example",
                  DNS: {
                    options: [
                      "options_example",
                    ],
                    searches: [
                      "searches_example",
                    ],
                    servers: [
                      "servers_example",
                    ],
                  },
                  device: "device_example",
                  dynamicPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                  hostname: "hostname_example",
                  IP: "IP_example",
                  mBits: 1,
                  mode: "mode_example",
                  reservedPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                },
              ],
            },
            services: {
              "key": "key_example",
            },
            taskGroup: "taskGroup_example",
            taskResources: {
              "key": {
                CPU: 1,
                cores: 1,
                devices: [
                  {
                    affinities: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                        weight: -128,
                      },
                    ],
                    constraints: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                      },
                    ],
                    count: 0,
                    name: "name_example",
                  },
                ],
                diskMB: 1,
                IOPS: 1,
                memoryMB: 1,
                memoryMaxMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
              },
            },
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        },
        requestedCapabilities: [
          {
            accessMode: "accessMode_example",
            attachmentMode: "attachmentMode_example",
          },
        ],
        requestedCapacityMax: 1,
        requestedCapacityMin: 1,
        requestedTopologies: {
          preferred: [
            {
              segments: {
                "key": "key_example",
              },
            },
          ],
          required: [
            {
              segments: {
                "key": "key_example",
              },
            },
          ],
        },
        resourceExhausted: new Date('1970-01-01T00:00:00.00Z'),
        schedulable: true,
        secrets: {
          "key": "key_example",
        },
        snapshotID: "snapshotID_example",
        topologies: [
          {
            segments: {
              "key": "key_example",
            },
          },
        ],
        writeAllocs: {
          "key": {
            allocModifyIndex: 0,
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentID: "deploymentID_example",
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            desiredTransition: {
              migrate: true,
              reschedule: true,
            },
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            job: {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              allAtOnce: true,
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              consulNamespace: "consulNamespace_example",
              consulToken: "consulToken_example",
              createIndex: 0,
              datacenters: [
                "datacenters_example",
              ],
              dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
              dispatched: true,
              ID: "ID_example",
              jobModifyIndex: 0,
              meta: {
                "key": "key_example",
              },
              migrate: {
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
              },
              modifyIndex: 0,
              multiregion: {
                regions: [
                  {
                    count: 1,
                    datacenters: [
                      "datacenters_example",
                    ],
                    meta: {
                      "key": "key_example",
                    },
                    name: "name_example",
                  },
                ],
                strategy: {
                  maxParallel: 1,
                  onFailure: "onFailure_example",
                },
              },
              name: "name_example",
              namespace: "namespace_example",
              nomadTokenID: "nomadTokenID_example",
              parameterizedJob: {
                metaOptional: [
                  "metaOptional_example",
                ],
                metaRequired: [
                  "metaRequired_example",
                ],
                payload: "payload_example",
              },
              parentID: "parentID_example",
              payload: 'YQ==',
              periodic: {
                enabled: true,
                prohibitOverlap: true,
                spec: "spec_example",
                specType: "specType_example",
                timeZone: "timeZone_example",
              },
              priority: 1,
              region: "region_example",
              reschedule: {
                attempts: 1,
                delay: 1,
                delayFunction: "delayFunction_example",
                interval: 1,
                maxDelay: 1,
                unlimited: true,
              },
              spreads: [
                {
                  attribute: "attribute_example",
                  spreadTarget: [
                    {
                      percent: 0,
                      value: "value_example",
                    },
                  ],
                  weight: -128,
                },
              ],
              stable: true,
              status: "status_example",
              statusDescription: "statusDescription_example",
              stop: true,
              submitTime: 1,
              taskGroups: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  consul: {
                    namespace: "namespace_example",
                  },
                  count: 1,
                  ephemeralDisk: {
                    migrate: true,
                    sizeMB: 1,
                    sticky: true,
                  },
                  maxClientDisconnect: 1,
                  meta: {
                    "key": "key_example",
                  },
                  migrate: {
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                  },
                  name: "name_example",
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                  reschedulePolicy: {
                    attempts: 1,
                    delay: 1,
                    delayFunction: "delayFunction_example",
                    interval: 1,
                    maxDelay: 1,
                    unlimited: true,
                  },
                  restartPolicy: {
                    attempts: 1,
                    delay: 1,
                    interval: 1,
                    mode: "mode_example",
                  },
                  scaling: {
                    createIndex: 0,
                    enabled: true,
                    ID: "ID_example",
                    max: 1,
                    min: 1,
                    modifyIndex: 0,
                    namespace: "namespace_example",
                    policy: {
                      "key": null,
                    },
                    target: {
                      "key": "key_example",
                    },
                    type: "type_example",
                  },
                  services: [
                    {
                      address: "address_example",
                      addressMode: "addressMode_example",
                      canaryMeta: {
                        "key": "key_example",
                      },
                      canaryTags: [
                        "canaryTags_example",
                      ],
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      checks: [
                        {
                          addressMode: "addressMode_example",
                          advertise: "advertise_example",
                          args: [
                            "args_example",
                          ],
                          body: "body_example",
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          command: "command_example",
                          expose: true,
                          failuresBeforeCritical: 1,
                          gRPCService: "gRPCService_example",
                          gRPCUseTLS: true,
                          header: {
                            "key": [
                              "key_example",
                            ],
                          },
                          initialStatus: "initialStatus_example",
                          interval: 1,
                          method: "method_example",
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          path: "path_example",
                          portLabel: "portLabel_example",
                          protocol: "protocol_example",
                          successBeforePassing: 1,
                          tLSSkipVerify: true,
                          taskName: "taskName_example",
                          timeout: 1,
                          type: "type_example",
                        },
                      ],
                      connect: {
                        gateway: {
                          ingress: {
                            listeners: [
                              {
                                port: 1,
                                protocol: "protocol_example",
                                services: [
                                  {
                                    hosts: [
                                      "hosts_example",
                                    ],
                                    name: "name_example",
                                  },
                                ],
                              },
                            ],
                            TLS: {
                              enabled: true,
                            },
                          },
                          mesh: null,
                          proxy: {
                            config: {
                              "key": null,
                            },
                            connectTimeout: 1,
                            envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                            envoyGatewayBindAddresses: {
                              "key": {
                                address: "address_example",
                                name: "name_example",
                                port: 1,
                              },
                            },
                            envoyGatewayBindTaggedAddresses: true,
                            envoyGatewayNoDefaultBind: true,
                          },
                          terminating: {
                            services: [
                              {
                                cAFile: "cAFile_example",
                                certFile: "certFile_example",
                                keyFile: "keyFile_example",
                                name: "name_example",
                                SNI: "SNI_example",
                              },
                            ],
                          },
                        },
                        _native: true,
                        sidecarService: {
                          disableDefaultTCPCheck: true,
                          port: "port_example",
                          proxy: {
                            config: {
                              "key": null,
                            },
                            exposeConfig: {
                              path: [
                                {
                                  listenerPort: "listenerPort_example",
                                  localPathPort: 1,
                                  path: "path_example",
                                  protocol: "protocol_example",
                                },
                              ],
                            },
                            localServiceAddress: "localServiceAddress_example",
                            localServicePort: 1,
                            upstreams: [
                              {
                                datacenter: "datacenter_example",
                                destinationName: "destinationName_example",
                                localBindAddress: "localBindAddress_example",
                                localBindPort: 1,
                                meshGateway: {
                                  mode: "mode_example",
                                },
                              },
                            ],
                          },
                          tags: [
                            "tags_example",
                          ],
                        },
                        sidecarTask: {
                          config: {
                            "key": null,
                          },
                          driver: "driver_example",
                          env: {
                            "key": "key_example",
                          },
                          killSignal: "killSignal_example",
                          killTimeout: 1,
                          logConfig: {
                            maxFileSizeMB: 1,
                            maxFiles: 1,
                          },
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          resources: {
                            CPU: 1,
                            cores: 1,
                            devices: [
                              {
                                affinities: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                    weight: -128,
                                  },
                                ],
                                constraints: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                  },
                                ],
                                count: 0,
                                name: "name_example",
                              },
                            ],
                            diskMB: 1,
                            IOPS: 1,
                            memoryMB: 1,
                            memoryMaxMB: 1,
                            networks: [
                              {
                                CIDR: "CIDR_example",
                                DNS: {
                                  options: [
                                    "options_example",
                                  ],
                                  searches: [
                                    "searches_example",
                                  ],
                                  servers: [
                                    "servers_example",
                                  ],
                                },
                                device: "device_example",
                                dynamicPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                                hostname: "hostname_example",
                                IP: "IP_example",
                                mBits: 1,
                                mode: "mode_example",
                                reservedPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                              },
                            ],
                          },
                          shutdownDelay: 1,
                          user: "user_example",
                        },
                      },
                      enableTagOverride: true,
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      portLabel: "portLabel_example",
                      provider: "provider_example",
                      tags: [
                        "tags_example",
                      ],
                      taskName: "taskName_example",
                    },
                  ],
                  shutdownDelay: 1,
                  spreads: [
                    {
                      attribute: "attribute_example",
                      spreadTarget: [
                        {
                          percent: 0,
                          value: "value_example",
                        },
                      ],
                      weight: -128,
                    },
                  ],
                  stopAfterClientDisconnect: 1,
                  tasks: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      artifacts: [
                        {
                          getterHeaders: {
                            "key": "key_example",
                          },
                          getterMode: "getterMode_example",
                          getterOptions: {
                            "key": "key_example",
                          },
                          getterSource: "getterSource_example",
                          relativeDest: "relativeDest_example",
                        },
                      ],
                      cSIPluginConfig: {
                        ID: "ID_example",
                        mountDir: "mountDir_example",
                        type: "type_example",
                      },
                      config: {
                        "key": null,
                      },
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      dispatchPayload: {
                        file: "file_example",
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      kind: "kind_example",
                      leader: true,
                      lifecycle: {
                        hook: "hook_example",
                        sidecar: true,
                      },
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
                        CPU: 1,
                        cores: 1,
                        devices: [
                          {
                            affinities: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                                weight: -128,
                              },
                            ],
                            constraints: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                              },
                            ],
                            count: 0,
                            name: "name_example",
                          },
                        ],
                        diskMB: 1,
                        IOPS: 1,
                        memoryMB: 1,
                        memoryMaxMB: 1,
                        networks: [
                          {
                            CIDR: "CIDR_example",
                            DNS: {
                              options: [
                                "options_example",
                              ],
                              searches: [
                                "searches_example",
                              ],
                              servers: [
                                "servers_example",
                              ],
                            },
                            device: "device_example",
                            dynamicPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                            hostname: "hostname_example",
                            IP: "IP_example",
                            mBits: 1,
                            mode: "mode_example",
                            reservedPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                          },
                        ],
                      },
                      restartPolicy: {
                        attempts: 1,
                        delay: 1,
                        interval: 1,
                        mode: "mode_example",
                      },
                      scalingPolicies: [
                        {
                          createIndex: 0,
                          enabled: true,
                          ID: "ID_example",
                          max: 1,
                          min: 1,
                          modifyIndex: 0,
                          namespace: "namespace_example",
                          policy: {
                            "key": null,
                          },
                          target: {
                            "key": "key_example",
                          },
                          type: "type_example",
                        },
                      ],
                      services: [
                        {
                          address: "address_example",
                          addressMode: "addressMode_example",
                          canaryMeta: {
                            "key": "key_example",
                          },
                          canaryTags: [
                            "canaryTags_example",
                          ],
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          checks: [
                            {
                              addressMode: "addressMode_example",
                              advertise: "advertise_example",
                              args: [
                                "args_example",
                              ],
                              body: "body_example",
                              checkRestart: {
                                grace: 1,
                                ignoreWarnings: true,
                                limit: 1,
                              },
                              command: "command_example",
                              expose: true,
                              failuresBeforeCritical: 1,
                              gRPCService: "gRPCService_example",
                              gRPCUseTLS: true,
                              header: {
                                "key": [
                                  "key_example",
                                ],
                              },
                              initialStatus: "initialStatus_example",
                              interval: 1,
                              method: "method_example",
                              name: "name_example",
                              onUpdate: "onUpdate_example",
                              path: "path_example",
                              portLabel: "portLabel_example",
                              protocol: "protocol_example",
                              successBeforePassing: 1,
                              tLSSkipVerify: true,
                              taskName: "taskName_example",
                              timeout: 1,
                              type: "type_example",
                            },
                          ],
                          connect: {
                            gateway: {
                              ingress: {
                                listeners: [
                                  {
                                    port: 1,
                                    protocol: "protocol_example",
                                    services: [
                                      {
                                        hosts: [
                                          "hosts_example",
                                        ],
                                        name: "name_example",
                                      },
                                    ],
                                  },
                                ],
                                TLS: {
                                  enabled: true,
                                },
                              },
                              mesh: null,
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                connectTimeout: 1,
                                envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                                envoyGatewayBindAddresses: {
                                  "key": {
                                    address: "address_example",
                                    name: "name_example",
                                    port: 1,
                                  },
                                },
                                envoyGatewayBindTaggedAddresses: true,
                                envoyGatewayNoDefaultBind: true,
                              },
                              terminating: {
                                services: [
                                  {
                                    cAFile: "cAFile_example",
                                    certFile: "certFile_example",
                                    keyFile: "keyFile_example",
                                    name: "name_example",
                                    SNI: "SNI_example",
                                  },
                                ],
                              },
                            },
                            _native: true,
                            sidecarService: {
                              disableDefaultTCPCheck: true,
                              port: "port_example",
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                exposeConfig: {
                                  path: [
                                    {
                                      listenerPort: "listenerPort_example",
                                      localPathPort: 1,
                                      path: "path_example",
                                      protocol: "protocol_example",
                                    },
                                  ],
                                },
                                localServiceAddress: "localServiceAddress_example",
                                localServicePort: 1,
                                upstreams: [
                                  {
                                    datacenter: "datacenter_example",
                                    destinationName: "destinationName_example",
                                    localBindAddress: "localBindAddress_example",
                                    localBindPort: 1,
                                    meshGateway: {
                                      mode: "mode_example",
                                    },
                                  },
                                ],
                              },
                              tags: [
                                "tags_example",
                              ],
                            },
                            sidecarTask: {
                              config: {
                                "key": null,
                              },
                              driver: "driver_example",
                              env: {
                                "key": "key_example",
                              },
                              killSignal: "killSignal_example",
                              killTimeout: 1,
                              logConfig: {
                                maxFileSizeMB: 1,
                                maxFiles: 1,
                              },
                              meta: {
                                "key": "key_example",
                              },
                              name: "name_example",
                              resources: {
                                CPU: 1,
                                cores: 1,
                                devices: [
                                  {
                                    affinities: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                        weight: -128,
                                      },
                                    ],
                                    constraints: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                      },
                                    ],
                                    count: 0,
                                    name: "name_example",
                                  },
                                ],
                                diskMB: 1,
                                IOPS: 1,
                                memoryMB: 1,
                                memoryMaxMB: 1,
                                networks: [
                                  {
                                    CIDR: "CIDR_example",
                                    DNS: {
                                      options: [
                                        "options_example",
                                      ],
                                      searches: [
                                        "searches_example",
                                      ],
                                      servers: [
                                        "servers_example",
                                      ],
                                    },
                                    device: "device_example",
                                    dynamicPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                    hostname: "hostname_example",
                                    IP: "IP_example",
                                    mBits: 1,
                                    mode: "mode_example",
                                    reservedPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                  },
                                ],
                              },
                              shutdownDelay: 1,
                              user: "user_example",
                            },
                          },
                          enableTagOverride: true,
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          portLabel: "portLabel_example",
                          provider: "provider_example",
                          tags: [
                            "tags_example",
                          ],
                          taskName: "taskName_example",
                        },
                      ],
                      shutdownDelay: 1,
                      templates: [
                        {
                          changeMode: "changeMode_example",
                          changeSignal: "changeSignal_example",
                          destPath: "destPath_example",
                          embeddedTmpl: "embeddedTmpl_example",
                          envvars: true,
                          leftDelim: "leftDelim_example",
                          perms: "perms_example",
                          rightDelim: "rightDelim_example",
                          sourcePath: "sourcePath_example",
                          splay: 1,
                          vaultGrace: 1,
                          wait: {
                            max: 1,
                            min: 1,
                          },
                        },
                      ],
                      user: "user_example",
                      vault: {
                        changeMode: "changeMode_example",
                        changeSignal: "changeSignal_example",
                        env: true,
                        namespace: "namespace_example",
                        policies: [
                          "policies_example",
                        ],
                      },
                      volumeMounts: [
                        {
                          destination: "destination_example",
                          propagationMode: "propagationMode_example",
                          readOnly: true,
                          volume: "volume_example",
                        },
                      ],
                    },
                  ],
                  update: {
                    autoPromote: true,
                    autoRevert: true,
                    canary: 1,
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                    progressDeadline: 1,
                    stagger: 1,
                  },
                  volumes: {
                    "key": {
                      accessMode: "accessMode_example",
                      attachmentMode: "attachmentMode_example",
                      mountOptions: {
                        fSType: "fSType_example",
                        mountFlags: [
                          "mountFlags_example",
                        ],
                      },
                      name: "name_example",
                      perAlloc: true,
                      readOnly: true,
                      source: "source_example",
                      type: "type_example",
                    },
                  },
                },
              ],
              type: "type_example",
              update: {
                autoPromote: true,
                autoRevert: true,
                canary: 1,
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
                progressDeadline: 1,
                stagger: 1,
              },
              vaultNamespace: "vaultNamespace_example",
              vaultToken: "vaultToken_example",
              version: 0,
            },
            jobID: "jobID_example",
            metrics: {
              allocationTime: 1,
              classExhausted: {
                "key": 1,
              },
              classFiltered: {
                "key": 1,
              },
              coalescedFailures: 1,
              constraintFiltered: {
                "key": 1,
              },
              dimensionExhausted: {
                "key": 1,
              },
              nodesAvailable: {
                "key": 1,
              },
              nodesEvaluated: 1,
              nodesExhausted: 1,
              nodesFiltered: 1,
              quotaExhausted: [
                "quotaExhausted_example",
              ],
              resourcesExhausted: {
                "key": {
                  CPU: 1,
                  cores: 1,
                  devices: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      count: 0,
                      name: "name_example",
                    },
                  ],
                  diskMB: 1,
                  IOPS: 1,
                  memoryMB: 1,
                  memoryMaxMB: 1,
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
              scoreMetaData: [
                {
                  nodeID: "nodeID_example",
                  normScore: 3.14,
                  scores: {
                    "key": 3.14,
                  },
                },
              ],
              scores: {
                "key": 3.14,
              },
            },
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nextAllocation: "nextAllocation_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            previousAllocation: "previousAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            resources: {
              CPU: 1,
              cores: 1,
              devices: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  count: 0,
                  name: "name_example",
                },
              ],
              diskMB: 1,
              IOPS: 1,
              memoryMB: 1,
              memoryMaxMB: 1,
              networks: [
                {
                  CIDR: "CIDR_example",
                  DNS: {
                    options: [
                      "options_example",
                    ],
                    searches: [
                      "searches_example",
                    ],
                    servers: [
                      "servers_example",
                    ],
                  },
                  device: "device_example",
                  dynamicPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                  hostname: "hostname_example",
                  IP: "IP_example",
                  mBits: 1,
                  mode: "mode_example",
                  reservedPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                },
              ],
            },
            services: {
              "key": "key_example",
            },
            taskGroup: "taskGroup_example",
            taskResources: {
              "key": {
                CPU: 1,
                cores: 1,
                devices: [
                  {
                    affinities: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                        weight: -128,
                      },
                    ],
                    constraints: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                      },
                    ],
                    count: 0,
                    name: "name_example",
                  },
                ],
                diskMB: 1,
                IOPS: 1,
                memoryMB: 1,
                memoryMaxMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
              },
            },
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        },
      },
    ],
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postVolume(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSIVolumeRegisterRequest** | **CSIVolumeRegisterRequest**|  |
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **postVolumeRegistration**
> void postVolumeRegistration(cSIVolumeRegisterRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.VolumesApi(configuration);

let body:nomad-client.VolumesApiPostVolumeRegistrationRequest = {
  // string | Volume unique identifier.
  volumeId: "volumeId_example",
  // CSIVolumeRegisterRequest
  cSIVolumeRegisterRequest: {
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
    volumes: [
      {
        accessMode: "accessMode_example",
        allocations: [
          {
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            jobID: "jobID_example",
            jobType: "jobType_example",
            jobVersion: 0,
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            taskGroup: "taskGroup_example",
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        ],
        attachmentMode: "attachmentMode_example",
        capacity: 1,
        cloneID: "cloneID_example",
        context: {
          "key": "key_example",
        },
        controllerRequired: true,
        controllersExpected: 1,
        controllersHealthy: 1,
        createIndex: 0,
        externalID: "externalID_example",
        ID: "ID_example",
        modifyIndex: 0,
        mountOptions: {
          fSType: "fSType_example",
          mountFlags: [
            "mountFlags_example",
          ],
        },
        name: "name_example",
        namespace: "namespace_example",
        nodesExpected: 1,
        nodesHealthy: 1,
        parameters: {
          "key": "key_example",
        },
        pluginID: "pluginID_example",
        provider: "provider_example",
        providerVersion: "providerVersion_example",
        readAllocs: {
          "key": {
            allocModifyIndex: 0,
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentID: "deploymentID_example",
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            desiredTransition: {
              migrate: true,
              reschedule: true,
            },
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            job: {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              allAtOnce: true,
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              consulNamespace: "consulNamespace_example",
              consulToken: "consulToken_example",
              createIndex: 0,
              datacenters: [
                "datacenters_example",
              ],
              dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
              dispatched: true,
              ID: "ID_example",
              jobModifyIndex: 0,
              meta: {
                "key": "key_example",
              },
              migrate: {
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
              },
              modifyIndex: 0,
              multiregion: {
                regions: [
                  {
                    count: 1,
                    datacenters: [
                      "datacenters_example",
                    ],
                    meta: {
                      "key": "key_example",
                    },
                    name: "name_example",
                  },
                ],
                strategy: {
                  maxParallel: 1,
                  onFailure: "onFailure_example",
                },
              },
              name: "name_example",
              namespace: "namespace_example",
              nomadTokenID: "nomadTokenID_example",
              parameterizedJob: {
                metaOptional: [
                  "metaOptional_example",
                ],
                metaRequired: [
                  "metaRequired_example",
                ],
                payload: "payload_example",
              },
              parentID: "parentID_example",
              payload: 'YQ==',
              periodic: {
                enabled: true,
                prohibitOverlap: true,
                spec: "spec_example",
                specType: "specType_example",
                timeZone: "timeZone_example",
              },
              priority: 1,
              region: "region_example",
              reschedule: {
                attempts: 1,
                delay: 1,
                delayFunction: "delayFunction_example",
                interval: 1,
                maxDelay: 1,
                unlimited: true,
              },
              spreads: [
                {
                  attribute: "attribute_example",
                  spreadTarget: [
                    {
                      percent: 0,
                      value: "value_example",
                    },
                  ],
                  weight: -128,
                },
              ],
              stable: true,
              status: "status_example",
              statusDescription: "statusDescription_example",
              stop: true,
              submitTime: 1,
              taskGroups: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  consul: {
                    namespace: "namespace_example",
                  },
                  count: 1,
                  ephemeralDisk: {
                    migrate: true,
                    sizeMB: 1,
                    sticky: true,
                  },
                  maxClientDisconnect: 1,
                  meta: {
                    "key": "key_example",
                  },
                  migrate: {
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                  },
                  name: "name_example",
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                  reschedulePolicy: {
                    attempts: 1,
                    delay: 1,
                    delayFunction: "delayFunction_example",
                    interval: 1,
                    maxDelay: 1,
                    unlimited: true,
                  },
                  restartPolicy: {
                    attempts: 1,
                    delay: 1,
                    interval: 1,
                    mode: "mode_example",
                  },
                  scaling: {
                    createIndex: 0,
                    enabled: true,
                    ID: "ID_example",
                    max: 1,
                    min: 1,
                    modifyIndex: 0,
                    namespace: "namespace_example",
                    policy: {
                      "key": null,
                    },
                    target: {
                      "key": "key_example",
                    },
                    type: "type_example",
                  },
                  services: [
                    {
                      address: "address_example",
                      addressMode: "addressMode_example",
                      canaryMeta: {
                        "key": "key_example",
                      },
                      canaryTags: [
                        "canaryTags_example",
                      ],
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      checks: [
                        {
                          addressMode: "addressMode_example",
                          advertise: "advertise_example",
                          args: [
                            "args_example",
                          ],
                          body: "body_example",
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          command: "command_example",
                          expose: true,
                          failuresBeforeCritical: 1,
                          gRPCService: "gRPCService_example",
                          gRPCUseTLS: true,
                          header: {
                            "key": [
                              "key_example",
                            ],
                          },
                          initialStatus: "initialStatus_example",
                          interval: 1,
                          method: "method_example",
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          path: "path_example",
                          portLabel: "portLabel_example",
                          protocol: "protocol_example",
                          successBeforePassing: 1,
                          tLSSkipVerify: true,
                          taskName: "taskName_example",
                          timeout: 1,
                          type: "type_example",
                        },
                      ],
                      connect: {
                        gateway: {
                          ingress: {
                            listeners: [
                              {
                                port: 1,
                                protocol: "protocol_example",
                                services: [
                                  {
                                    hosts: [
                                      "hosts_example",
                                    ],
                                    name: "name_example",
                                  },
                                ],
                              },
                            ],
                            TLS: {
                              enabled: true,
                            },
                          },
                          mesh: null,
                          proxy: {
                            config: {
                              "key": null,
                            },
                            connectTimeout: 1,
                            envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                            envoyGatewayBindAddresses: {
                              "key": {
                                address: "address_example",
                                name: "name_example",
                                port: 1,
                              },
                            },
                            envoyGatewayBindTaggedAddresses: true,
                            envoyGatewayNoDefaultBind: true,
                          },
                          terminating: {
                            services: [
                              {
                                cAFile: "cAFile_example",
                                certFile: "certFile_example",
                                keyFile: "keyFile_example",
                                name: "name_example",
                                SNI: "SNI_example",
                              },
                            ],
                          },
                        },
                        _native: true,
                        sidecarService: {
                          disableDefaultTCPCheck: true,
                          port: "port_example",
                          proxy: {
                            config: {
                              "key": null,
                            },
                            exposeConfig: {
                              path: [
                                {
                                  listenerPort: "listenerPort_example",
                                  localPathPort: 1,
                                  path: "path_example",
                                  protocol: "protocol_example",
                                },
                              ],
                            },
                            localServiceAddress: "localServiceAddress_example",
                            localServicePort: 1,
                            upstreams: [
                              {
                                datacenter: "datacenter_example",
                                destinationName: "destinationName_example",
                                localBindAddress: "localBindAddress_example",
                                localBindPort: 1,
                                meshGateway: {
                                  mode: "mode_example",
                                },
                              },
                            ],
                          },
                          tags: [
                            "tags_example",
                          ],
                        },
                        sidecarTask: {
                          config: {
                            "key": null,
                          },
                          driver: "driver_example",
                          env: {
                            "key": "key_example",
                          },
                          killSignal: "killSignal_example",
                          killTimeout: 1,
                          logConfig: {
                            maxFileSizeMB: 1,
                            maxFiles: 1,
                          },
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          resources: {
                            CPU: 1,
                            cores: 1,
                            devices: [
                              {
                                affinities: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                    weight: -128,
                                  },
                                ],
                                constraints: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                  },
                                ],
                                count: 0,
                                name: "name_example",
                              },
                            ],
                            diskMB: 1,
                            IOPS: 1,
                            memoryMB: 1,
                            memoryMaxMB: 1,
                            networks: [
                              {
                                CIDR: "CIDR_example",
                                DNS: {
                                  options: [
                                    "options_example",
                                  ],
                                  searches: [
                                    "searches_example",
                                  ],
                                  servers: [
                                    "servers_example",
                                  ],
                                },
                                device: "device_example",
                                dynamicPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                                hostname: "hostname_example",
                                IP: "IP_example",
                                mBits: 1,
                                mode: "mode_example",
                                reservedPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                              },
                            ],
                          },
                          shutdownDelay: 1,
                          user: "user_example",
                        },
                      },
                      enableTagOverride: true,
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      portLabel: "portLabel_example",
                      provider: "provider_example",
                      tags: [
                        "tags_example",
                      ],
                      taskName: "taskName_example",
                    },
                  ],
                  shutdownDelay: 1,
                  spreads: [
                    {
                      attribute: "attribute_example",
                      spreadTarget: [
                        {
                          percent: 0,
                          value: "value_example",
                        },
                      ],
                      weight: -128,
                    },
                  ],
                  stopAfterClientDisconnect: 1,
                  tasks: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      artifacts: [
                        {
                          getterHeaders: {
                            "key": "key_example",
                          },
                          getterMode: "getterMode_example",
                          getterOptions: {
                            "key": "key_example",
                          },
                          getterSource: "getterSource_example",
                          relativeDest: "relativeDest_example",
                        },
                      ],
                      cSIPluginConfig: {
                        ID: "ID_example",
                        mountDir: "mountDir_example",
                        type: "type_example",
                      },
                      config: {
                        "key": null,
                      },
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      dispatchPayload: {
                        file: "file_example",
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      kind: "kind_example",
                      leader: true,
                      lifecycle: {
                        hook: "hook_example",
                        sidecar: true,
                      },
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
                        CPU: 1,
                        cores: 1,
                        devices: [
                          {
                            affinities: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                                weight: -128,
                              },
                            ],
                            constraints: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                              },
                            ],
                            count: 0,
                            name: "name_example",
                          },
                        ],
                        diskMB: 1,
                        IOPS: 1,
                        memoryMB: 1,
                        memoryMaxMB: 1,
                        networks: [
                          {
                            CIDR: "CIDR_example",
                            DNS: {
                              options: [
                                "options_example",
                              ],
                              searches: [
                                "searches_example",
                              ],
                              servers: [
                                "servers_example",
                              ],
                            },
                            device: "device_example",
                            dynamicPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                            hostname: "hostname_example",
                            IP: "IP_example",
                            mBits: 1,
                            mode: "mode_example",
                            reservedPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                          },
                        ],
                      },
                      restartPolicy: {
                        attempts: 1,
                        delay: 1,
                        interval: 1,
                        mode: "mode_example",
                      },
                      scalingPolicies: [
                        {
                          createIndex: 0,
                          enabled: true,
                          ID: "ID_example",
                          max: 1,
                          min: 1,
                          modifyIndex: 0,
                          namespace: "namespace_example",
                          policy: {
                            "key": null,
                          },
                          target: {
                            "key": "key_example",
                          },
                          type: "type_example",
                        },
                      ],
                      services: [
                        {
                          address: "address_example",
                          addressMode: "addressMode_example",
                          canaryMeta: {
                            "key": "key_example",
                          },
                          canaryTags: [
                            "canaryTags_example",
                          ],
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          checks: [
                            {
                              addressMode: "addressMode_example",
                              advertise: "advertise_example",
                              args: [
                                "args_example",
                              ],
                              body: "body_example",
                              checkRestart: {
                                grace: 1,
                                ignoreWarnings: true,
                                limit: 1,
                              },
                              command: "command_example",
                              expose: true,
                              failuresBeforeCritical: 1,
                              gRPCService: "gRPCService_example",
                              gRPCUseTLS: true,
                              header: {
                                "key": [
                                  "key_example",
                                ],
                              },
                              initialStatus: "initialStatus_example",
                              interval: 1,
                              method: "method_example",
                              name: "name_example",
                              onUpdate: "onUpdate_example",
                              path: "path_example",
                              portLabel: "portLabel_example",
                              protocol: "protocol_example",
                              successBeforePassing: 1,
                              tLSSkipVerify: true,
                              taskName: "taskName_example",
                              timeout: 1,
                              type: "type_example",
                            },
                          ],
                          connect: {
                            gateway: {
                              ingress: {
                                listeners: [
                                  {
                                    port: 1,
                                    protocol: "protocol_example",
                                    services: [
                                      {
                                        hosts: [
                                          "hosts_example",
                                        ],
                                        name: "name_example",
                                      },
                                    ],
                                  },
                                ],
                                TLS: {
                                  enabled: true,
                                },
                              },
                              mesh: null,
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                connectTimeout: 1,
                                envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                                envoyGatewayBindAddresses: {
                                  "key": {
                                    address: "address_example",
                                    name: "name_example",
                                    port: 1,
                                  },
                                },
                                envoyGatewayBindTaggedAddresses: true,
                                envoyGatewayNoDefaultBind: true,
                              },
                              terminating: {
                                services: [
                                  {
                                    cAFile: "cAFile_example",
                                    certFile: "certFile_example",
                                    keyFile: "keyFile_example",
                                    name: "name_example",
                                    SNI: "SNI_example",
                                  },
                                ],
                              },
                            },
                            _native: true,
                            sidecarService: {
                              disableDefaultTCPCheck: true,
                              port: "port_example",
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                exposeConfig: {
                                  path: [
                                    {
                                      listenerPort: "listenerPort_example",
                                      localPathPort: 1,
                                      path: "path_example",
                                      protocol: "protocol_example",
                                    },
                                  ],
                                },
                                localServiceAddress: "localServiceAddress_example",
                                localServicePort: 1,
                                upstreams: [
                                  {
                                    datacenter: "datacenter_example",
                                    destinationName: "destinationName_example",
                                    localBindAddress: "localBindAddress_example",
                                    localBindPort: 1,
                                    meshGateway: {
                                      mode: "mode_example",
                                    },
                                  },
                                ],
                              },
                              tags: [
                                "tags_example",
                              ],
                            },
                            sidecarTask: {
                              config: {
                                "key": null,
                              },
                              driver: "driver_example",
                              env: {
                                "key": "key_example",
                              },
                              killSignal: "killSignal_example",
                              killTimeout: 1,
                              logConfig: {
                                maxFileSizeMB: 1,
                                maxFiles: 1,
                              },
                              meta: {
                                "key": "key_example",
                              },
                              name: "name_example",
                              resources: {
                                CPU: 1,
                                cores: 1,
                                devices: [
                                  {
                                    affinities: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                        weight: -128,
                                      },
                                    ],
                                    constraints: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                      },
                                    ],
                                    count: 0,
                                    name: "name_example",
                                  },
                                ],
                                diskMB: 1,
                                IOPS: 1,
                                memoryMB: 1,
                                memoryMaxMB: 1,
                                networks: [
                                  {
                                    CIDR: "CIDR_example",
                                    DNS: {
                                      options: [
                                        "options_example",
                                      ],
                                      searches: [
                                        "searches_example",
                                      ],
                                      servers: [
                                        "servers_example",
                                      ],
                                    },
                                    device: "device_example",
                                    dynamicPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                    hostname: "hostname_example",
                                    IP: "IP_example",
                                    mBits: 1,
                                    mode: "mode_example",
                                    reservedPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                  },
                                ],
                              },
                              shutdownDelay: 1,
                              user: "user_example",
                            },
                          },
                          enableTagOverride: true,
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          portLabel: "portLabel_example",
                          provider: "provider_example",
                          tags: [
                            "tags_example",
                          ],
                          taskName: "taskName_example",
                        },
                      ],
                      shutdownDelay: 1,
                      templates: [
                        {
                          changeMode: "changeMode_example",
                          changeSignal: "changeSignal_example",
                          destPath: "destPath_example",
                          embeddedTmpl: "embeddedTmpl_example",
                          envvars: true,
                          leftDelim: "leftDelim_example",
                          perms: "perms_example",
                          rightDelim: "rightDelim_example",
                          sourcePath: "sourcePath_example",
                          splay: 1,
                          vaultGrace: 1,
                          wait: {
                            max: 1,
                            min: 1,
                          },
                        },
                      ],
                      user: "user_example",
                      vault: {
                        changeMode: "changeMode_example",
                        changeSignal: "changeSignal_example",
                        env: true,
                        namespace: "namespace_example",
                        policies: [
                          "policies_example",
                        ],
                      },
                      volumeMounts: [
                        {
                          destination: "destination_example",
                          propagationMode: "propagationMode_example",
                          readOnly: true,
                          volume: "volume_example",
                        },
                      ],
                    },
                  ],
                  update: {
                    autoPromote: true,
                    autoRevert: true,
                    canary: 1,
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                    progressDeadline: 1,
                    stagger: 1,
                  },
                  volumes: {
                    "key": {
                      accessMode: "accessMode_example",
                      attachmentMode: "attachmentMode_example",
                      mountOptions: {
                        fSType: "fSType_example",
                        mountFlags: [
                          "mountFlags_example",
                        ],
                      },
                      name: "name_example",
                      perAlloc: true,
                      readOnly: true,
                      source: "source_example",
                      type: "type_example",
                    },
                  },
                },
              ],
              type: "type_example",
              update: {
                autoPromote: true,
                autoRevert: true,
                canary: 1,
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
                progressDeadline: 1,
                stagger: 1,
              },
              vaultNamespace: "vaultNamespace_example",
              vaultToken: "vaultToken_example",
              version: 0,
            },
            jobID: "jobID_example",
            metrics: {
              allocationTime: 1,
              classExhausted: {
                "key": 1,
              },
              classFiltered: {
                "key": 1,
              },
              coalescedFailures: 1,
              constraintFiltered: {
                "key": 1,
              },
              dimensionExhausted: {
                "key": 1,
              },
              nodesAvailable: {
                "key": 1,
              },
              nodesEvaluated: 1,
              nodesExhausted: 1,
              nodesFiltered: 1,
              quotaExhausted: [
                "quotaExhausted_example",
              ],
              resourcesExhausted: {
                "key": {
                  CPU: 1,
                  cores: 1,
                  devices: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      count: 0,
                      name: "name_example",
                    },
                  ],
                  diskMB: 1,
                  IOPS: 1,
                  memoryMB: 1,
                  memoryMaxMB: 1,
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
              scoreMetaData: [
                {
                  nodeID: "nodeID_example",
                  normScore: 3.14,
                  scores: {
                    "key": 3.14,
                  },
                },
              ],
              scores: {
                "key": 3.14,
              },
            },
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nextAllocation: "nextAllocation_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            previousAllocation: "previousAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            resources: {
              CPU: 1,
              cores: 1,
              devices: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  count: 0,
                  name: "name_example",
                },
              ],
              diskMB: 1,
              IOPS: 1,
              memoryMB: 1,
              memoryMaxMB: 1,
              networks: [
                {
                  CIDR: "CIDR_example",
                  DNS: {
                    options: [
                      "options_example",
                    ],
                    searches: [
                      "searches_example",
                    ],
                    servers: [
                      "servers_example",
                    ],
                  },
                  device: "device_example",
                  dynamicPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                  hostname: "hostname_example",
                  IP: "IP_example",
                  mBits: 1,
                  mode: "mode_example",
                  reservedPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                },
              ],
            },
            services: {
              "key": "key_example",
            },
            taskGroup: "taskGroup_example",
            taskResources: {
              "key": {
                CPU: 1,
                cores: 1,
                devices: [
                  {
                    affinities: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                        weight: -128,
                      },
                    ],
                    constraints: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                      },
                    ],
                    count: 0,
                    name: "name_example",
                  },
                ],
                diskMB: 1,
                IOPS: 1,
                memoryMB: 1,
                memoryMaxMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
              },
            },
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        },
        requestedCapabilities: [
          {
            accessMode: "accessMode_example",
            attachmentMode: "attachmentMode_example",
          },
        ],
        requestedCapacityMax: 1,
        requestedCapacityMin: 1,
        requestedTopologies: {
          preferred: [
            {
              segments: {
                "key": "key_example",
              },
            },
          ],
          required: [
            {
              segments: {
                "key": "key_example",
              },
            },
          ],
        },
        resourceExhausted: new Date('1970-01-01T00:00:00.00Z'),
        schedulable: true,
        secrets: {
          "key": "key_example",
        },
        snapshotID: "snapshotID_example",
        topologies: [
          {
            segments: {
              "key": "key_example",
            },
          },
        ],
        writeAllocs: {
          "key": {
            allocModifyIndex: 0,
            allocatedResources: {
              shared: {
                diskMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
                ports: [
                  {
                    hostIP: "hostIP_example",
                    label: "label_example",
                    to: 1,
                    value: 1,
                  },
                ],
              },
              tasks: {
                "key": {
                  cpu: {
                    cpuShares: 1,
                  },
                  devices: [
                    {
                      deviceIDs: [
                        "deviceIDs_example",
                      ],
                      name: "name_example",
                      type: "type_example",
                      vendor: "vendor_example",
                    },
                  ],
                  memory: {
                    memoryMB: 1,
                    memoryMaxMB: 1,
                  },
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
            },
            clientDescription: "clientDescription_example",
            clientStatus: "clientStatus_example",
            createIndex: 0,
            createTime: 1,
            deploymentID: "deploymentID_example",
            deploymentStatus: {
              canary: true,
              healthy: true,
              modifyIndex: 0,
              timestamp: new Date('1970-01-01T00:00:00.00Z'),
            },
            desiredDescription: "desiredDescription_example",
            desiredStatus: "desiredStatus_example",
            desiredTransition: {
              migrate: true,
              reschedule: true,
            },
            evalID: "evalID_example",
            followupEvalID: "followupEvalID_example",
            ID: "ID_example",
            job: {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              allAtOnce: true,
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              consulNamespace: "consulNamespace_example",
              consulToken: "consulToken_example",
              createIndex: 0,
              datacenters: [
                "datacenters_example",
              ],
              dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
              dispatched: true,
              ID: "ID_example",
              jobModifyIndex: 0,
              meta: {
                "key": "key_example",
              },
              migrate: {
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
              },
              modifyIndex: 0,
              multiregion: {
                regions: [
                  {
                    count: 1,
                    datacenters: [
                      "datacenters_example",
                    ],
                    meta: {
                      "key": "key_example",
                    },
                    name: "name_example",
                  },
                ],
                strategy: {
                  maxParallel: 1,
                  onFailure: "onFailure_example",
                },
              },
              name: "name_example",
              namespace: "namespace_example",
              nomadTokenID: "nomadTokenID_example",
              parameterizedJob: {
                metaOptional: [
                  "metaOptional_example",
                ],
                metaRequired: [
                  "metaRequired_example",
                ],
                payload: "payload_example",
              },
              parentID: "parentID_example",
              payload: 'YQ==',
              periodic: {
                enabled: true,
                prohibitOverlap: true,
                spec: "spec_example",
                specType: "specType_example",
                timeZone: "timeZone_example",
              },
              priority: 1,
              region: "region_example",
              reschedule: {
                attempts: 1,
                delay: 1,
                delayFunction: "delayFunction_example",
                interval: 1,
                maxDelay: 1,
                unlimited: true,
              },
              spreads: [
                {
                  attribute: "attribute_example",
                  spreadTarget: [
                    {
                      percent: 0,
                      value: "value_example",
                    },
                  ],
                  weight: -128,
                },
              ],
              stable: true,
              status: "status_example",
              statusDescription: "statusDescription_example",
              stop: true,
              submitTime: 1,
              taskGroups: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  consul: {
                    namespace: "namespace_example",
                  },
                  count: 1,
                  ephemeralDisk: {
                    migrate: true,
                    sizeMB: 1,
                    sticky: true,
                  },
                  maxClientDisconnect: 1,
                  meta: {
                    "key": "key_example",
                  },
                  migrate: {
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                  },
                  name: "name_example",
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                  reschedulePolicy: {
                    attempts: 1,
                    delay: 1,
                    delayFunction: "delayFunction_example",
                    interval: 1,
                    maxDelay: 1,
                    unlimited: true,
                  },
                  restartPolicy: {
                    attempts: 1,
                    delay: 1,
                    interval: 1,
                    mode: "mode_example",
                  },
                  scaling: {
                    createIndex: 0,
                    enabled: true,
                    ID: "ID_example",
                    max: 1,
                    min: 1,
                    modifyIndex: 0,
                    namespace: "namespace_example",
                    policy: {
                      "key": null,
                    },
                    target: {
                      "key": "key_example",
                    },
                    type: "type_example",
                  },
                  services: [
                    {
                      address: "address_example",
                      addressMode: "addressMode_example",
                      canaryMeta: {
                        "key": "key_example",
                      },
                      canaryTags: [
                        "canaryTags_example",
                      ],
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      checks: [
                        {
                          addressMode: "addressMode_example",
                          advertise: "advertise_example",
                          args: [
                            "args_example",
                          ],
                          body: "body_example",
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          command: "command_example",
                          expose: true,
                          failuresBeforeCritical: 1,
                          gRPCService: "gRPCService_example",
                          gRPCUseTLS: true,
                          header: {
                            "key": [
                              "key_example",
                            ],
                          },
                          initialStatus: "initialStatus_example",
                          interval: 1,
                          method: "method_example",
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          path: "path_example",
                          portLabel: "portLabel_example",
                          protocol: "protocol_example",
                          successBeforePassing: 1,
                          tLSSkipVerify: true,
                          taskName: "taskName_example",
                          timeout: 1,
                          type: "type_example",
                        },
                      ],
                      connect: {
                        gateway: {
                          ingress: {
                            listeners: [
                              {
                                port: 1,
                                protocol: "protocol_example",
                                services: [
                                  {
                                    hosts: [
                                      "hosts_example",
                                    ],
                                    name: "name_example",
                                  },
                                ],
                              },
                            ],
                            TLS: {
                              enabled: true,
                            },
                          },
                          mesh: null,
                          proxy: {
                            config: {
                              "key": null,
                            },
                            connectTimeout: 1,
                            envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                            envoyGatewayBindAddresses: {
                              "key": {
                                address: "address_example",
                                name: "name_example",
                                port: 1,
                              },
                            },
                            envoyGatewayBindTaggedAddresses: true,
                            envoyGatewayNoDefaultBind: true,
                          },
                          terminating: {
                            services: [
                              {
                                cAFile: "cAFile_example",
                                certFile: "certFile_example",
                                keyFile: "keyFile_example",
                                name: "name_example",
                                SNI: "SNI_example",
                              },
                            ],
                          },
                        },
                        _native: true,
                        sidecarService: {
                          disableDefaultTCPCheck: true,
                          port: "port_example",
                          proxy: {
                            config: {
                              "key": null,
                            },
                            exposeConfig: {
                              path: [
                                {
                                  listenerPort: "listenerPort_example",
                                  localPathPort: 1,
                                  path: "path_example",
                                  protocol: "protocol_example",
                                },
                              ],
                            },
                            localServiceAddress: "localServiceAddress_example",
                            localServicePort: 1,
                            upstreams: [
                              {
                                datacenter: "datacenter_example",
                                destinationName: "destinationName_example",
                                localBindAddress: "localBindAddress_example",
                                localBindPort: 1,
                                meshGateway: {
                                  mode: "mode_example",
                                },
                              },
                            ],
                          },
                          tags: [
                            "tags_example",
                          ],
                        },
                        sidecarTask: {
                          config: {
                            "key": null,
                          },
                          driver: "driver_example",
                          env: {
                            "key": "key_example",
                          },
                          killSignal: "killSignal_example",
                          killTimeout: 1,
                          logConfig: {
                            maxFileSizeMB: 1,
                            maxFiles: 1,
                          },
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          resources: {
                            CPU: 1,
                            cores: 1,
                            devices: [
                              {
                                affinities: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                    weight: -128,
                                  },
                                ],
                                constraints: [
                                  {
                                    lTarget: "lTarget_example",
                                    operand: "operand_example",
                                    rTarget: "rTarget_example",
                                  },
                                ],
                                count: 0,
                                name: "name_example",
                              },
                            ],
                            diskMB: 1,
                            IOPS: 1,
                            memoryMB: 1,
                            memoryMaxMB: 1,
                            networks: [
                              {
                                CIDR: "CIDR_example",
                                DNS: {
                                  options: [
                                    "options_example",
                                  ],
                                  searches: [
                                    "searches_example",
                                  ],
                                  servers: [
                                    "servers_example",
                                  ],
                                },
                                device: "device_example",
                                dynamicPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                                hostname: "hostname_example",
                                IP: "IP_example",
                                mBits: 1,
                                mode: "mode_example",
                                reservedPorts: [
                                  {
                                    hostNetwork: "hostNetwork_example",
                                    label: "label_example",
                                    to: 1,
                                    value: 1,
                                  },
                                ],
                              },
                            ],
                          },
                          shutdownDelay: 1,
                          user: "user_example",
                        },
                      },
                      enableTagOverride: true,
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      portLabel: "portLabel_example",
                      provider: "provider_example",
                      tags: [
                        "tags_example",
                      ],
                      taskName: "taskName_example",
                    },
                  ],
                  shutdownDelay: 1,
                  spreads: [
                    {
                      attribute: "attribute_example",
                      spreadTarget: [
                        {
                          percent: 0,
                          value: "value_example",
                        },
                      ],
                      weight: -128,
                    },
                  ],
                  stopAfterClientDisconnect: 1,
                  tasks: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      artifacts: [
                        {
                          getterHeaders: {
                            "key": "key_example",
                          },
                          getterMode: "getterMode_example",
                          getterOptions: {
                            "key": "key_example",
                          },
                          getterSource: "getterSource_example",
                          relativeDest: "relativeDest_example",
                        },
                      ],
                      cSIPluginConfig: {
                        ID: "ID_example",
                        mountDir: "mountDir_example",
                        type: "type_example",
                      },
                      config: {
                        "key": null,
                      },
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      dispatchPayload: {
                        file: "file_example",
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      kind: "kind_example",
                      leader: true,
                      lifecycle: {
                        hook: "hook_example",
                        sidecar: true,
                      },
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
                        CPU: 1,
                        cores: 1,
                        devices: [
                          {
                            affinities: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                                weight: -128,
                              },
                            ],
                            constraints: [
                              {
                                lTarget: "lTarget_example",
                                operand: "operand_example",
                                rTarget: "rTarget_example",
                              },
                            ],
                            count: 0,
                            name: "name_example",
                          },
                        ],
                        diskMB: 1,
                        IOPS: 1,
                        memoryMB: 1,
                        memoryMaxMB: 1,
                        networks: [
                          {
                            CIDR: "CIDR_example",
                            DNS: {
                              options: [
                                "options_example",
                              ],
                              searches: [
                                "searches_example",
                              ],
                              servers: [
                                "servers_example",
                              ],
                            },
                            device: "device_example",
                            dynamicPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                            hostname: "hostname_example",
                            IP: "IP_example",
                            mBits: 1,
                            mode: "mode_example",
                            reservedPorts: [
                              {
                                hostNetwork: "hostNetwork_example",
                                label: "label_example",
                                to: 1,
                                value: 1,
                              },
                            ],
                          },
                        ],
                      },
                      restartPolicy: {
                        attempts: 1,
                        delay: 1,
                        interval: 1,
                        mode: "mode_example",
                      },
                      scalingPolicies: [
                        {
                          createIndex: 0,
                          enabled: true,
                          ID: "ID_example",
                          max: 1,
                          min: 1,
                          modifyIndex: 0,
                          namespace: "namespace_example",
                          policy: {
                            "key": null,
                          },
                          target: {
                            "key": "key_example",
                          },
                          type: "type_example",
                        },
                      ],
                      services: [
                        {
                          address: "address_example",
                          addressMode: "addressMode_example",
                          canaryMeta: {
                            "key": "key_example",
                          },
                          canaryTags: [
                            "canaryTags_example",
                          ],
                          checkRestart: {
                            grace: 1,
                            ignoreWarnings: true,
                            limit: 1,
                          },
                          checks: [
                            {
                              addressMode: "addressMode_example",
                              advertise: "advertise_example",
                              args: [
                                "args_example",
                              ],
                              body: "body_example",
                              checkRestart: {
                                grace: 1,
                                ignoreWarnings: true,
                                limit: 1,
                              },
                              command: "command_example",
                              expose: true,
                              failuresBeforeCritical: 1,
                              gRPCService: "gRPCService_example",
                              gRPCUseTLS: true,
                              header: {
                                "key": [
                                  "key_example",
                                ],
                              },
                              initialStatus: "initialStatus_example",
                              interval: 1,
                              method: "method_example",
                              name: "name_example",
                              onUpdate: "onUpdate_example",
                              path: "path_example",
                              portLabel: "portLabel_example",
                              protocol: "protocol_example",
                              successBeforePassing: 1,
                              tLSSkipVerify: true,
                              taskName: "taskName_example",
                              timeout: 1,
                              type: "type_example",
                            },
                          ],
                          connect: {
                            gateway: {
                              ingress: {
                                listeners: [
                                  {
                                    port: 1,
                                    protocol: "protocol_example",
                                    services: [
                                      {
                                        hosts: [
                                          "hosts_example",
                                        ],
                                        name: "name_example",
                                      },
                                    ],
                                  },
                                ],
                                TLS: {
                                  enabled: true,
                                },
                              },
                              mesh: null,
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                connectTimeout: 1,
                                envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                                envoyGatewayBindAddresses: {
                                  "key": {
                                    address: "address_example",
                                    name: "name_example",
                                    port: 1,
                                  },
                                },
                                envoyGatewayBindTaggedAddresses: true,
                                envoyGatewayNoDefaultBind: true,
                              },
                              terminating: {
                                services: [
                                  {
                                    cAFile: "cAFile_example",
                                    certFile: "certFile_example",
                                    keyFile: "keyFile_example",
                                    name: "name_example",
                                    SNI: "SNI_example",
                                  },
                                ],
                              },
                            },
                            _native: true,
                            sidecarService: {
                              disableDefaultTCPCheck: true,
                              port: "port_example",
                              proxy: {
                                config: {
                                  "key": null,
                                },
                                exposeConfig: {
                                  path: [
                                    {
                                      listenerPort: "listenerPort_example",
                                      localPathPort: 1,
                                      path: "path_example",
                                      protocol: "protocol_example",
                                    },
                                  ],
                                },
                                localServiceAddress: "localServiceAddress_example",
                                localServicePort: 1,
                                upstreams: [
                                  {
                                    datacenter: "datacenter_example",
                                    destinationName: "destinationName_example",
                                    localBindAddress: "localBindAddress_example",
                                    localBindPort: 1,
                                    meshGateway: {
                                      mode: "mode_example",
                                    },
                                  },
                                ],
                              },
                              tags: [
                                "tags_example",
                              ],
                            },
                            sidecarTask: {
                              config: {
                                "key": null,
                              },
                              driver: "driver_example",
                              env: {
                                "key": "key_example",
                              },
                              killSignal: "killSignal_example",
                              killTimeout: 1,
                              logConfig: {
                                maxFileSizeMB: 1,
                                maxFiles: 1,
                              },
                              meta: {
                                "key": "key_example",
                              },
                              name: "name_example",
                              resources: {
                                CPU: 1,
                                cores: 1,
                                devices: [
                                  {
                                    affinities: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                        weight: -128,
                                      },
                                    ],
                                    constraints: [
                                      {
                                        lTarget: "lTarget_example",
                                        operand: "operand_example",
                                        rTarget: "rTarget_example",
                                      },
                                    ],
                                    count: 0,
                                    name: "name_example",
                                  },
                                ],
                                diskMB: 1,
                                IOPS: 1,
                                memoryMB: 1,
                                memoryMaxMB: 1,
                                networks: [
                                  {
                                    CIDR: "CIDR_example",
                                    DNS: {
                                      options: [
                                        "options_example",
                                      ],
                                      searches: [
                                        "searches_example",
                                      ],
                                      servers: [
                                        "servers_example",
                                      ],
                                    },
                                    device: "device_example",
                                    dynamicPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                    hostname: "hostname_example",
                                    IP: "IP_example",
                                    mBits: 1,
                                    mode: "mode_example",
                                    reservedPorts: [
                                      {
                                        hostNetwork: "hostNetwork_example",
                                        label: "label_example",
                                        to: 1,
                                        value: 1,
                                      },
                                    ],
                                  },
                                ],
                              },
                              shutdownDelay: 1,
                              user: "user_example",
                            },
                          },
                          enableTagOverride: true,
                          meta: {
                            "key": "key_example",
                          },
                          name: "name_example",
                          onUpdate: "onUpdate_example",
                          portLabel: "portLabel_example",
                          provider: "provider_example",
                          tags: [
                            "tags_example",
                          ],
                          taskName: "taskName_example",
                        },
                      ],
                      shutdownDelay: 1,
                      templates: [
                        {
                          changeMode: "changeMode_example",
                          changeSignal: "changeSignal_example",
                          destPath: "destPath_example",
                          embeddedTmpl: "embeddedTmpl_example",
                          envvars: true,
                          leftDelim: "leftDelim_example",
                          perms: "perms_example",
                          rightDelim: "rightDelim_example",
                          sourcePath: "sourcePath_example",
                          splay: 1,
                          vaultGrace: 1,
                          wait: {
                            max: 1,
                            min: 1,
                          },
                        },
                      ],
                      user: "user_example",
                      vault: {
                        changeMode: "changeMode_example",
                        changeSignal: "changeSignal_example",
                        env: true,
                        namespace: "namespace_example",
                        policies: [
                          "policies_example",
                        ],
                      },
                      volumeMounts: [
                        {
                          destination: "destination_example",
                          propagationMode: "propagationMode_example",
                          readOnly: true,
                          volume: "volume_example",
                        },
                      ],
                    },
                  ],
                  update: {
                    autoPromote: true,
                    autoRevert: true,
                    canary: 1,
                    healthCheck: "healthCheck_example",
                    healthyDeadline: 1,
                    maxParallel: 1,
                    minHealthyTime: 1,
                    progressDeadline: 1,
                    stagger: 1,
                  },
                  volumes: {
                    "key": {
                      accessMode: "accessMode_example",
                      attachmentMode: "attachmentMode_example",
                      mountOptions: {
                        fSType: "fSType_example",
                        mountFlags: [
                          "mountFlags_example",
                        ],
                      },
                      name: "name_example",
                      perAlloc: true,
                      readOnly: true,
                      source: "source_example",
                      type: "type_example",
                    },
                  },
                },
              ],
              type: "type_example",
              update: {
                autoPromote: true,
                autoRevert: true,
                canary: 1,
                healthCheck: "healthCheck_example",
                healthyDeadline: 1,
                maxParallel: 1,
                minHealthyTime: 1,
                progressDeadline: 1,
                stagger: 1,
              },
              vaultNamespace: "vaultNamespace_example",
              vaultToken: "vaultToken_example",
              version: 0,
            },
            jobID: "jobID_example",
            metrics: {
              allocationTime: 1,
              classExhausted: {
                "key": 1,
              },
              classFiltered: {
                "key": 1,
              },
              coalescedFailures: 1,
              constraintFiltered: {
                "key": 1,
              },
              dimensionExhausted: {
                "key": 1,
              },
              nodesAvailable: {
                "key": 1,
              },
              nodesEvaluated: 1,
              nodesExhausted: 1,
              nodesFiltered: 1,
              quotaExhausted: [
                "quotaExhausted_example",
              ],
              resourcesExhausted: {
                "key": {
                  CPU: 1,
                  cores: 1,
                  devices: [
                    {
                      affinities: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                          weight: -128,
                        },
                      ],
                      constraints: [
                        {
                          lTarget: "lTarget_example",
                          operand: "operand_example",
                          rTarget: "rTarget_example",
                        },
                      ],
                      count: 0,
                      name: "name_example",
                    },
                  ],
                  diskMB: 1,
                  IOPS: 1,
                  memoryMB: 1,
                  memoryMaxMB: 1,
                  networks: [
                    {
                      CIDR: "CIDR_example",
                      DNS: {
                        options: [
                          "options_example",
                        ],
                        searches: [
                          "searches_example",
                        ],
                        servers: [
                          "servers_example",
                        ],
                      },
                      device: "device_example",
                      dynamicPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                      hostname: "hostname_example",
                      IP: "IP_example",
                      mBits: 1,
                      mode: "mode_example",
                      reservedPorts: [
                        {
                          hostNetwork: "hostNetwork_example",
                          label: "label_example",
                          to: 1,
                          value: 1,
                        },
                      ],
                    },
                  ],
                },
              },
              scoreMetaData: [
                {
                  nodeID: "nodeID_example",
                  normScore: 3.14,
                  scores: {
                    "key": 3.14,
                  },
                },
              ],
              scores: {
                "key": 3.14,
              },
            },
            modifyIndex: 0,
            modifyTime: 1,
            name: "name_example",
            namespace: "namespace_example",
            nextAllocation: "nextAllocation_example",
            nodeID: "nodeID_example",
            nodeName: "nodeName_example",
            preemptedAllocations: [
              "preemptedAllocations_example",
            ],
            preemptedByAllocation: "preemptedByAllocation_example",
            previousAllocation: "previousAllocation_example",
            rescheduleTracker: {
              events: [
                {
                  prevAllocID: "prevAllocID_example",
                  prevNodeID: "prevNodeID_example",
                  rescheduleTime: 1,
                },
              ],
            },
            resources: {
              CPU: 1,
              cores: 1,
              devices: [
                {
                  affinities: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                      weight: -128,
                    },
                  ],
                  constraints: [
                    {
                      lTarget: "lTarget_example",
                      operand: "operand_example",
                      rTarget: "rTarget_example",
                    },
                  ],
                  count: 0,
                  name: "name_example",
                },
              ],
              diskMB: 1,
              IOPS: 1,
              memoryMB: 1,
              memoryMaxMB: 1,
              networks: [
                {
                  CIDR: "CIDR_example",
                  DNS: {
                    options: [
                      "options_example",
                    ],
                    searches: [
                      "searches_example",
                    ],
                    servers: [
                      "servers_example",
                    ],
                  },
                  device: "device_example",
                  dynamicPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                  hostname: "hostname_example",
                  IP: "IP_example",
                  mBits: 1,
                  mode: "mode_example",
                  reservedPorts: [
                    {
                      hostNetwork: "hostNetwork_example",
                      label: "label_example",
                      to: 1,
                      value: 1,
                    },
                  ],
                },
              ],
            },
            services: {
              "key": "key_example",
            },
            taskGroup: "taskGroup_example",
            taskResources: {
              "key": {
                CPU: 1,
                cores: 1,
                devices: [
                  {
                    affinities: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                        weight: -128,
                      },
                    ],
                    constraints: [
                      {
                        lTarget: "lTarget_example",
                        operand: "operand_example",
                        rTarget: "rTarget_example",
                      },
                    ],
                    count: 0,
                    name: "name_example",
                  },
                ],
                diskMB: 1,
                IOPS: 1,
                memoryMB: 1,
                memoryMaxMB: 1,
                networks: [
                  {
                    CIDR: "CIDR_example",
                    DNS: {
                      options: [
                        "options_example",
                      ],
                      searches: [
                        "searches_example",
                      ],
                      servers: [
                        "servers_example",
                      ],
                    },
                    device: "device_example",
                    dynamicPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                    hostname: "hostname_example",
                    IP: "IP_example",
                    mBits: 1,
                    mode: "mode_example",
                    reservedPorts: [
                      {
                        hostNetwork: "hostNetwork_example",
                        label: "label_example",
                        to: 1,
                        value: 1,
                      },
                    ],
                  },
                ],
              },
            },
            taskStates: {
              "key": {
                events: [
                  {
                    details: {
                      "key": "key_example",
                    },
                    diskLimit: 1,
                    diskSize: 1,
                    displayMessage: "displayMessage_example",
                    downloadError: "downloadError_example",
                    driverError: "driverError_example",
                    driverMessage: "driverMessage_example",
                    exitCode: 1,
                    failedSibling: "failedSibling_example",
                    failsTask: true,
                    genericSource: "genericSource_example",
                    killError: "killError_example",
                    killReason: "killReason_example",
                    killTimeout: 1,
                    message: "message_example",
                    restartReason: "restartReason_example",
                    setupError: "setupError_example",
                    signal: 1,
                    startDelay: 1,
                    taskSignal: "taskSignal_example",
                    taskSignalReason: "taskSignalReason_example",
                    time: 1,
                    type: "type_example",
                    validationError: "validationError_example",
                    vaultError: "vaultError_example",
                  },
                ],
                failed: true,
                finishedAt: new Date('1970-01-01T00:00:00.00Z'),
                lastRestart: new Date('1970-01-01T00:00:00.00Z'),
                restarts: 0,
                startedAt: new Date('1970-01-01T00:00:00.00Z'),
                state: "state_example",
                taskHandle: {
                  driverState: 'YQ==',
                  version: 1,
                },
              },
            },
          },
        },
      },
    ],
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postVolumeRegistration(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSIVolumeRegisterRequest** | **CSIVolumeRegisterRequest**|  |
 **volumeId** | [**string**] | Volume unique identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)


