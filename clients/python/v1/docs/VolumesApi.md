# nomad_client.VolumesApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_volume**](VolumesApi.md#create_volume) | **POST** /volume/csi/{volumeId}/{action} | 
[**delete_snapshot**](VolumesApi.md#delete_snapshot) | **DELETE** /volumes/snapshot | 
[**delete_volume_registration**](VolumesApi.md#delete_volume_registration) | **DELETE** /volume/csi/{volumeId} | 
[**detach_or_delete_volume**](VolumesApi.md#detach_or_delete_volume) | **DELETE** /volume/csi/{volumeId}/{action} | 
[**get_external_volumes**](VolumesApi.md#get_external_volumes) | **GET** /volumes/external | 
[**get_snapshots**](VolumesApi.md#get_snapshots) | **GET** /volumes/snapshot | 
[**get_volume**](VolumesApi.md#get_volume) | **GET** /volume/csi/{volumeId} | 
[**get_volumes**](VolumesApi.md#get_volumes) | **GET** /volumes | 
[**post_snapshot**](VolumesApi.md#post_snapshot) | **POST** /volumes/snapshot | 
[**post_volume**](VolumesApi.md#post_volume) | **POST** /volumes | 
[**post_volume_registration**](VolumesApi.md#post_volume_registration) | **POST** /volume/csi/{volumeId} | 


# **create_volume**
> create_volume(volume_id, action, csi_volume_create_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_volume_create_request import CSIVolumeCreateRequest
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    volume_id = "volumeId_example" # str | Volume unique identifier.
    action = "action_example" # str | The action to perform on the Volume (create, detach, delete).
    csi_volume_create_request = CSIVolumeCreateRequest(
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
        volumes=[
            CSIVolume(
                access_mode="access_mode_example",
                allocations=[
                    AllocationListStub(
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job_id="job_id_example",
                        job_type="job_type_example",
                        job_version=0,
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        task_group="task_group_example",
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                ],
                attachment_mode="attachment_mode_example",
                capacity=1,
                clone_id="clone_id_example",
                context={
                    "key": "key_example",
                },
                controller_required=True,
                controllers_expected=1,
                controllers_healthy=1,
                create_index=0,
                external_id="external_id_example",
                id="id_example",
                modify_index=0,
                mount_options=CSIMountOptions(
                    fs_type="fs_type_example",
                    mount_flags=[
                        "mount_flags_example",
                    ],
                ),
                name="name_example",
                namespace="namespace_example",
                nodes_expected=1,
                nodes_healthy=1,
                parameters={
                    "key": "key_example",
                },
                plugin_id="plugin_id_example",
                provider="provider_example",
                provider_version="provider_version_example",
                read_allocs={
                    "key": Allocation(
                        alloc_modify_index=0,
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_id="deployment_id_example",
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        desired_transition=DesiredTransition(
                            migrate=True,
                            reschedule=True,
                        ),
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job=Job(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            all_at_once=True,
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            consul_namespace="consul_namespace_example",
                            consul_token="consul_token_example",
                            create_index=0,
                            datacenters=[
                                "datacenters_example",
                            ],
                            dispatch_idempotency_token="dispatch_idempotency_token_example",
                            dispatched=True,
                            id="id_example",
                            job_modify_index=0,
                            meta={
                                "key": "key_example",
                            },
                            migrate=MigrateStrategy(
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                            ),
                            modify_index=0,
                            multiregion=Multiregion(
                                regions=[
                                    MultiregionRegion(
                                        count=1,
                                        datacenters=[
                                            "datacenters_example",
                                        ],
                                        meta={
                                            "key": "key_example",
                                        },
                                        name="name_example",
                                    ),
                                ],
                                strategy=MultiregionStrategy(
                                    max_parallel=1,
                                    on_failure="on_failure_example",
                                ),
                            ),
                            name="name_example",
                            namespace="namespace_example",
                            nomad_token_id="nomad_token_id_example",
                            parameterized_job=ParameterizedJobConfig(
                                meta_optional=[
                                    "meta_optional_example",
                                ],
                                meta_required=[
                                    "meta_required_example",
                                ],
                                payload="payload_example",
                            ),
                            parent_id="parent_id_example",
                            payload='YQ==',
                            periodic=PeriodicConfig(
                                enabled=True,
                                prohibit_overlap=True,
                                spec="spec_example",
                                spec_type="spec_type_example",
                                time_zone="time_zone_example",
                            ),
                            priority=1,
                            region="region_example",
                            reschedule=ReschedulePolicy(
                                attempts=1,
                                delay=1,
                                delay_function="delay_function_example",
                                interval=1,
                                max_delay=1,
                                unlimited=True,
                            ),
                            spreads=[
                                Spread(
                                    attribute="attribute_example",
                                    spread_target=[
                                        SpreadTarget(
                                            percent=0,
                                            value="value_example",
                                        ),
                                    ],
                                    weight=-128,
                                ),
                            ],
                            stable=True,
                            status="status_example",
                            status_description="status_description_example",
                            stop=True,
                            submit_time=1,
                            task_groups=[
                                TaskGroup(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    consul=Consul(
                                        namespace="namespace_example",
                                    ),
                                    count=1,
                                    ephemeral_disk=EphemeralDisk(
                                        migrate=True,
                                        size_mb=1,
                                        sticky=True,
                                    ),
                                    max_client_disconnect=1,
                                    meta={
                                        "key": "key_example",
                                    },
                                    migrate=MigrateStrategy(
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                    ),
                                    name="name_example",
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                    reschedule_policy=ReschedulePolicy(
                                        attempts=1,
                                        delay=1,
                                        delay_function="delay_function_example",
                                        interval=1,
                                        max_delay=1,
                                        unlimited=True,
                                    ),
                                    restart_policy=RestartPolicy(
                                        attempts=1,
                                        delay=1,
                                        interval=1,
                                        mode="mode_example",
                                    ),
                                    scaling=ScalingPolicy(
                                        create_index=0,
                                        enabled=True,
                                        id="id_example",
                                        max=1,
                                        min=1,
                                        modify_index=0,
                                        namespace="namespace_example",
                                        policy={
                                            "key": None,
                                        },
                                        target={
                                            "key": "key_example",
                                        },
                                        type="type_example",
                                    ),
                                    services=[
                                        Service(
                                            address_mode="address_mode_example",
                                            canary_meta={
                                                "key": "key_example",
                                            },
                                            canary_tags=[
                                                "canary_tags_example",
                                            ],
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            checks=[
                                                ServiceCheck(
                                                    address_mode="address_mode_example",
                                                    args=[
                                                        "args_example",
                                                    ],
                                                    body="body_example",
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    command="command_example",
                                                    expose=True,
                                                    failures_before_critical=1,
                                                    grpc_service="grpc_service_example",
                                                    grpc_use_tls=True,
                                                    header={
                                                        "key": [
                                                            "key_example",
                                                        ],
                                                    },
                                                    id="id_example",
                                                    initial_status="initial_status_example",
                                                    interval=1,
                                                    method="method_example",
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    path="path_example",
                                                    port_label="port_label_example",
                                                    protocol="protocol_example",
                                                    success_before_passing=1,
                                                    tls_skip_verify=True,
                                                    task_name="task_name_example",
                                                    timeout=1,
                                                    type="type_example",
                                                ),
                                            ],
                                            connect=ConsulConnect(
                                                gateway=ConsulGateway(
                                                    ingress=ConsulIngressConfigEntry(
                                                        listeners=[
                                                            ConsulIngressListener(
                                                                port=1,
                                                                protocol="protocol_example",
                                                                services=[
                                                                    ConsulIngressService(
                                                                        hosts=[
                                                                            "hosts_example",
                                                                        ],
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                        tls=ConsulGatewayTLSConfig(
                                                            enabled=True,
                                                        ),
                                                    ),
                                                    mesh=None,
                                                    proxy=ConsulGatewayProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        connect_timeout=1,
                                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                        envoy_gateway_bind_addresses={
                                                            "key": ConsulGatewayBindAddress(
                                                                address="address_example",
                                                                name="name_example",
                                                                port=1,
                                                            ),
                                                        },
                                                        envoy_gateway_bind_tagged_addresses=True,
                                                        envoy_gateway_no_default_bind=True,
                                                    ),
                                                    terminating=ConsulTerminatingConfigEntry(
                                                        services=[
                                                            ConsulLinkedService(
                                                                ca_file="ca_file_example",
                                                                cert_file="cert_file_example",
                                                                key_file="key_file_example",
                                                                name="name_example",
                                                                sni="sni_example",
                                                            ),
                                                        ],
                                                    ),
                                                ),
                                                native=True,
                                                sidecar_service=ConsulSidecarService(
                                                    disable_default_tcp_check=True,
                                                    port="port_example",
                                                    proxy=ConsulProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        expose_config=ConsulExposeConfig(
                                                            path=[
                                                                ConsulExposePath(
                                                                    listener_port="listener_port_example",
                                                                    local_path_port=1,
                                                                    path="path_example",
                                                                    protocol="protocol_example",
                                                                ),
                                                            ],
                                                        ),
                                                        local_service_address="local_service_address_example",
                                                        local_service_port=1,
                                                        upstreams=[
                                                            ConsulUpstream(
                                                                datacenter="datacenter_example",
                                                                destination_name="destination_name_example",
                                                                local_bind_address="local_bind_address_example",
                                                                local_bind_port=1,
                                                                mesh_gateway=ConsulMeshGateway(
                                                                    mode="mode_example",
                                                                ),
                                                            ),
                                                        ],
                                                    ),
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                ),
                                                sidecar_task=SidecarTask(
                                                    config={
                                                        "key": None,
                                                    },
                                                    driver="driver_example",
                                                    env={
                                                        "key": "key_example",
                                                    },
                                                    kill_signal="kill_signal_example",
                                                    kill_timeout=1,
                                                    log_config=LogConfig(
                                                        max_file_size_mb=1,
                                                        max_files=1,
                                                    ),
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    resources=Resources(
                                                        cpu=1,
                                                        cores=1,
                                                        devices=[
                                                            RequestedDevice(
                                                                affinities=[
                                                                    Affinity(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                        weight=-128,
                                                                    ),
                                                                ],
                                                                constraints=[
                                                                    Constraint(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                    ),
                                                                ],
                                                                count=0,
                                                                name="name_example",
                                                            ),
                                                        ],
                                                        disk_mb=1,
                                                        iops=1,
                                                        memory_mb=1,
                                                        memory_max_mb=1,
                                                        networks=[
                                                            NetworkResource(
                                                                cidr="cidr_example",
                                                                dns=DNSConfig(
                                                                    options=[
                                                                        "options_example",
                                                                    ],
                                                                    searches=[
                                                                        "searches_example",
                                                                    ],
                                                                    servers=[
                                                                        "servers_example",
                                                                    ],
                                                                ),
                                                                device="device_example",
                                                                dynamic_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                                hostname="hostname_example",
                                                                ip="ip_example",
                                                                m_bits=1,
                                                                mode="mode_example",
                                                                reserved_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                    ),
                                                    shutdown_delay=1,
                                                    user="user_example",
                                                ),
                                            ),
                                            enable_tag_override=True,
                                            id="id_example",
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            on_update="on_update_example",
                                            port_label="port_label_example",
                                            provider="provider_example",
                                            tags=[
                                                "tags_example",
                                            ],
                                            task_name="task_name_example",
                                        ),
                                    ],
                                    shutdown_delay=1,
                                    spreads=[
                                        Spread(
                                            attribute="attribute_example",
                                            spread_target=[
                                                SpreadTarget(
                                                    percent=0,
                                                    value="value_example",
                                                ),
                                            ],
                                            weight=-128,
                                        ),
                                    ],
                                    stop_after_client_disconnect=1,
                                    tasks=[
                                        Task(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            artifacts=[
                                                TaskArtifact(
                                                    getter_headers={
                                                        "key": "key_example",
                                                    },
                                                    getter_mode="getter_mode_example",
                                                    getter_options={
                                                        "key": "key_example",
                                                    },
                                                    getter_source="getter_source_example",
                                                    relative_dest="relative_dest_example",
                                                ),
                                            ],
                                            csi_plugin_config=TaskCSIPluginConfig(
                                                id="id_example",
                                                mount_dir="mount_dir_example",
                                                type="type_example",
                                            ),
                                            config={
                                                "key": None,
                                            },
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            dispatch_payload=DispatchPayloadConfig(
                                                file="file_example",
                                            ),
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            kind="kind_example",
                                            leader=True,
                                            lifecycle=TaskLifecycle(
                                                hook="hook_example",
                                                sidecar=True,
                                            ),
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            restart_policy=RestartPolicy(
                                                attempts=1,
                                                delay=1,
                                                interval=1,
                                                mode="mode_example",
                                            ),
                                            scaling_policies=[
                                                ScalingPolicy(
                                                    create_index=0,
                                                    enabled=True,
                                                    id="id_example",
                                                    max=1,
                                                    min=1,
                                                    modify_index=0,
                                                    namespace="namespace_example",
                                                    policy={
                                                        "key": None,
                                                    },
                                                    target={
                                                        "key": "key_example",
                                                    },
                                                    type="type_example",
                                                ),
                                            ],
                                            services=[
                                                Service(
                                                    address_mode="address_mode_example",
                                                    canary_meta={
                                                        "key": "key_example",
                                                    },
                                                    canary_tags=[
                                                        "canary_tags_example",
                                                    ],
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    checks=[
                                                        ServiceCheck(
                                                            address_mode="address_mode_example",
                                                            args=[
                                                                "args_example",
                                                            ],
                                                            body="body_example",
                                                            check_restart=CheckRestart(
                                                                grace=1,
                                                                ignore_warnings=True,
                                                                limit=1,
                                                            ),
                                                            command="command_example",
                                                            expose=True,
                                                            failures_before_critical=1,
                                                            grpc_service="grpc_service_example",
                                                            grpc_use_tls=True,
                                                            header={
                                                                "key": [
                                                                    "key_example",
                                                                ],
                                                            },
                                                            id="id_example",
                                                            initial_status="initial_status_example",
                                                            interval=1,
                                                            method="method_example",
                                                            name="name_example",
                                                            on_update="on_update_example",
                                                            path="path_example",
                                                            port_label="port_label_example",
                                                            protocol="protocol_example",
                                                            success_before_passing=1,
                                                            tls_skip_verify=True,
                                                            task_name="task_name_example",
                                                            timeout=1,
                                                            type="type_example",
                                                        ),
                                                    ],
                                                    connect=ConsulConnect(
                                                        gateway=ConsulGateway(
                                                            ingress=ConsulIngressConfigEntry(
                                                                listeners=[
                                                                    ConsulIngressListener(
                                                                        port=1,
                                                                        protocol="protocol_example",
                                                                        services=[
                                                                            ConsulIngressService(
                                                                                hosts=[
                                                                                    "hosts_example",
                                                                                ],
                                                                                name="name_example",
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                                tls=ConsulGatewayTLSConfig(
                                                                    enabled=True,
                                                                ),
                                                            ),
                                                            mesh=None,
                                                            proxy=ConsulGatewayProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                connect_timeout=1,
                                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                                envoy_gateway_bind_addresses={
                                                                    "key": ConsulGatewayBindAddress(
                                                                        address="address_example",
                                                                        name="name_example",
                                                                        port=1,
                                                                    ),
                                                                },
                                                                envoy_gateway_bind_tagged_addresses=True,
                                                                envoy_gateway_no_default_bind=True,
                                                            ),
                                                            terminating=ConsulTerminatingConfigEntry(
                                                                services=[
                                                                    ConsulLinkedService(
                                                                        ca_file="ca_file_example",
                                                                        cert_file="cert_file_example",
                                                                        key_file="key_file_example",
                                                                        name="name_example",
                                                                        sni="sni_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ),
                                                        native=True,
                                                        sidecar_service=ConsulSidecarService(
                                                            disable_default_tcp_check=True,
                                                            port="port_example",
                                                            proxy=ConsulProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                expose_config=ConsulExposeConfig(
                                                                    path=[
                                                                        ConsulExposePath(
                                                                            listener_port="listener_port_example",
                                                                            local_path_port=1,
                                                                            path="path_example",
                                                                            protocol="protocol_example",
                                                                        ),
                                                                    ],
                                                                ),
                                                                local_service_address="local_service_address_example",
                                                                local_service_port=1,
                                                                upstreams=[
                                                                    ConsulUpstream(
                                                                        datacenter="datacenter_example",
                                                                        destination_name="destination_name_example",
                                                                        local_bind_address="local_bind_address_example",
                                                                        local_bind_port=1,
                                                                        mesh_gateway=ConsulMeshGateway(
                                                                            mode="mode_example",
                                                                        ),
                                                                    ),
                                                                ],
                                                            ),
                                                            tags=[
                                                                "tags_example",
                                                            ],
                                                        ),
                                                        sidecar_task=SidecarTask(
                                                            config={
                                                                "key": None,
                                                            },
                                                            driver="driver_example",
                                                            env={
                                                                "key": "key_example",
                                                            },
                                                            kill_signal="kill_signal_example",
                                                            kill_timeout=1,
                                                            log_config=LogConfig(
                                                                max_file_size_mb=1,
                                                                max_files=1,
                                                            ),
                                                            meta={
                                                                "key": "key_example",
                                                            },
                                                            name="name_example",
                                                            resources=Resources(
                                                                cpu=1,
                                                                cores=1,
                                                                devices=[
                                                                    RequestedDevice(
                                                                        affinities=[
                                                                            Affinity(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                                weight=-128,
                                                                            ),
                                                                        ],
                                                                        constraints=[
                                                                            Constraint(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                            ),
                                                                        ],
                                                                        count=0,
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                                disk_mb=1,
                                                                iops=1,
                                                                memory_mb=1,
                                                                memory_max_mb=1,
                                                                networks=[
                                                                    NetworkResource(
                                                                        cidr="cidr_example",
                                                                        dns=DNSConfig(
                                                                            options=[
                                                                                "options_example",
                                                                            ],
                                                                            searches=[
                                                                                "searches_example",
                                                                            ],
                                                                            servers=[
                                                                                "servers_example",
                                                                            ],
                                                                        ),
                                                                        device="device_example",
                                                                        dynamic_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                        hostname="hostname_example",
                                                                        ip="ip_example",
                                                                        m_bits=1,
                                                                        mode="mode_example",
                                                                        reserved_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                            ),
                                                            shutdown_delay=1,
                                                            user="user_example",
                                                        ),
                                                    ),
                                                    enable_tag_override=True,
                                                    id="id_example",
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    port_label="port_label_example",
                                                    provider="provider_example",
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                    task_name="task_name_example",
                                                ),
                                            ],
                                            shutdown_delay=1,
                                            templates=[
                                                Template(
                                                    change_mode="change_mode_example",
                                                    change_signal="change_signal_example",
                                                    dest_path="dest_path_example",
                                                    embedded_tmpl="embedded_tmpl_example",
                                                    envvars=True,
                                                    left_delim="left_delim_example",
                                                    perms="perms_example",
                                                    right_delim="right_delim_example",
                                                    source_path="source_path_example",
                                                    splay=1,
                                                    vault_grace=1,
                                                    wait=WaitConfig(
                                                        max=1,
                                                        min=1,
                                                    ),
                                                ),
                                            ],
                                            user="user_example",
                                            vault=Vault(
                                                change_mode="change_mode_example",
                                                change_signal="change_signal_example",
                                                entity_alias="entity_alias_example",
                                                env=True,
                                                namespace="namespace_example",
                                                policies=[
                                                    "policies_example",
                                                ],
                                            ),
                                            volume_mounts=[
                                                VolumeMount(
                                                    destination="destination_example",
                                                    propagation_mode="propagation_mode_example",
                                                    read_only=True,
                                                    volume="volume_example",
                                                ),
                                            ],
                                        ),
                                    ],
                                    update=UpdateStrategy(
                                        auto_promote=True,
                                        auto_revert=True,
                                        canary=1,
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                        progress_deadline=1,
                                        stagger=1,
                                    ),
                                    volumes={
                                        "key": VolumeRequest(
                                            access_mode="access_mode_example",
                                            attachment_mode="attachment_mode_example",
                                            mount_options=CSIMountOptions(
                                                fs_type="fs_type_example",
                                                mount_flags=[
                                                    "mount_flags_example",
                                                ],
                                            ),
                                            name="name_example",
                                            per_alloc=True,
                                            read_only=True,
                                            source="source_example",
                                            type="type_example",
                                        ),
                                    },
                                ),
                            ],
                            type="type_example",
                            update=UpdateStrategy(
                                auto_promote=True,
                                auto_revert=True,
                                canary=1,
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                                progress_deadline=1,
                                stagger=1,
                            ),
                            vault_namespace="vault_namespace_example",
                            vault_token="vault_token_example",
                            version=0,
                        ),
                        job_id="job_id_example",
                        metrics=AllocationMetric(
                            allocation_time=1,
                            class_exhausted={
                                "key": 1,
                            },
                            class_filtered={
                                "key": 1,
                            },
                            coalesced_failures=1,
                            constraint_filtered={
                                "key": 1,
                            },
                            dimension_exhausted={
                                "key": 1,
                            },
                            nodes_available={
                                "key": 1,
                            },
                            nodes_evaluated=1,
                            nodes_exhausted=1,
                            nodes_filtered=1,
                            quota_exhausted=[
                                "quota_exhausted_example",
                            ],
                            resources_exhausted={
                                "key": Resources(
                                    cpu=1,
                                    cores=1,
                                    devices=[
                                        RequestedDevice(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            count=0,
                                            name="name_example",
                                        ),
                                    ],
                                    disk_mb=1,
                                    iops=1,
                                    memory_mb=1,
                                    memory_max_mb=1,
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                            score_meta_data=[
                                NodeScoreMeta(
                                    node_id="node_id_example",
                                    norm_score=3.14,
                                    scores={
                                        "key": 3.14,
                                    },
                                ),
                            ],
                            scores={
                                "key": 3.14,
                            },
                        ),
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        next_allocation="next_allocation_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        previous_allocation="previous_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        resources=Resources(
                            cpu=1,
                            cores=1,
                            devices=[
                                RequestedDevice(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    count=0,
                                    name="name_example",
                                ),
                            ],
                            disk_mb=1,
                            iops=1,
                            memory_mb=1,
                            memory_max_mb=1,
                            networks=[
                                NetworkResource(
                                    cidr="cidr_example",
                                    dns=DNSConfig(
                                        options=[
                                            "options_example",
                                        ],
                                        searches=[
                                            "searches_example",
                                        ],
                                        servers=[
                                            "servers_example",
                                        ],
                                    ),
                                    device="device_example",
                                    dynamic_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                    hostname="hostname_example",
                                    ip="ip_example",
                                    m_bits=1,
                                    mode="mode_example",
                                    reserved_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                ),
                            ],
                        ),
                        services={
                            "key": "key_example",
                        },
                        task_group="task_group_example",
                        task_resources={
                            "key": Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                        },
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                },
                requested_capabilities=[
                    CSIVolumeCapability(
                        access_mode="access_mode_example",
                        attachment_mode="attachment_mode_example",
                    ),
                ],
                requested_capacity_max=1,
                requested_capacity_min=1,
                requested_topologies=CSITopologyRequest(
                    preferred=[
                        CSITopology(
                            segments={
                                "key": "key_example",
                            },
                        ),
                    ],
                    required=[
                        CSITopology(
                            segments={
                                "key": "key_example",
                            },
                        ),
                    ],
                ),
                resource_exhausted=dateutil_parser('1970-01-01T00:00:00.00Z'),
                schedulable=True,
                secrets=CSISecrets(
                    key="key_example",
                ),
                snapshot_id="snapshot_id_example",
                topologies=[
                    CSITopology(
                        segments={
                            "key": "key_example",
                        },
                    ),
                ],
                write_allocs={
                    "key": Allocation(
                        alloc_modify_index=0,
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_id="deployment_id_example",
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        desired_transition=DesiredTransition(
                            migrate=True,
                            reschedule=True,
                        ),
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job=Job(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            all_at_once=True,
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            consul_namespace="consul_namespace_example",
                            consul_token="consul_token_example",
                            create_index=0,
                            datacenters=[
                                "datacenters_example",
                            ],
                            dispatch_idempotency_token="dispatch_idempotency_token_example",
                            dispatched=True,
                            id="id_example",
                            job_modify_index=0,
                            meta={
                                "key": "key_example",
                            },
                            migrate=MigrateStrategy(
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                            ),
                            modify_index=0,
                            multiregion=Multiregion(
                                regions=[
                                    MultiregionRegion(
                                        count=1,
                                        datacenters=[
                                            "datacenters_example",
                                        ],
                                        meta={
                                            "key": "key_example",
                                        },
                                        name="name_example",
                                    ),
                                ],
                                strategy=MultiregionStrategy(
                                    max_parallel=1,
                                    on_failure="on_failure_example",
                                ),
                            ),
                            name="name_example",
                            namespace="namespace_example",
                            nomad_token_id="nomad_token_id_example",
                            parameterized_job=ParameterizedJobConfig(
                                meta_optional=[
                                    "meta_optional_example",
                                ],
                                meta_required=[
                                    "meta_required_example",
                                ],
                                payload="payload_example",
                            ),
                            parent_id="parent_id_example",
                            payload='YQ==',
                            periodic=PeriodicConfig(
                                enabled=True,
                                prohibit_overlap=True,
                                spec="spec_example",
                                spec_type="spec_type_example",
                                time_zone="time_zone_example",
                            ),
                            priority=1,
                            region="region_example",
                            reschedule=ReschedulePolicy(
                                attempts=1,
                                delay=1,
                                delay_function="delay_function_example",
                                interval=1,
                                max_delay=1,
                                unlimited=True,
                            ),
                            spreads=[
                                Spread(
                                    attribute="attribute_example",
                                    spread_target=[
                                        SpreadTarget(
                                            percent=0,
                                            value="value_example",
                                        ),
                                    ],
                                    weight=-128,
                                ),
                            ],
                            stable=True,
                            status="status_example",
                            status_description="status_description_example",
                            stop=True,
                            submit_time=1,
                            task_groups=[
                                TaskGroup(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    consul=Consul(
                                        namespace="namespace_example",
                                    ),
                                    count=1,
                                    ephemeral_disk=EphemeralDisk(
                                        migrate=True,
                                        size_mb=1,
                                        sticky=True,
                                    ),
                                    max_client_disconnect=1,
                                    meta={
                                        "key": "key_example",
                                    },
                                    migrate=MigrateStrategy(
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                    ),
                                    name="name_example",
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                    reschedule_policy=ReschedulePolicy(
                                        attempts=1,
                                        delay=1,
                                        delay_function="delay_function_example",
                                        interval=1,
                                        max_delay=1,
                                        unlimited=True,
                                    ),
                                    restart_policy=RestartPolicy(
                                        attempts=1,
                                        delay=1,
                                        interval=1,
                                        mode="mode_example",
                                    ),
                                    scaling=ScalingPolicy(
                                        create_index=0,
                                        enabled=True,
                                        id="id_example",
                                        max=1,
                                        min=1,
                                        modify_index=0,
                                        namespace="namespace_example",
                                        policy={
                                            "key": None,
                                        },
                                        target={
                                            "key": "key_example",
                                        },
                                        type="type_example",
                                    ),
                                    services=[
                                        Service(
                                            address_mode="address_mode_example",
                                            canary_meta={
                                                "key": "key_example",
                                            },
                                            canary_tags=[
                                                "canary_tags_example",
                                            ],
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            checks=[
                                                ServiceCheck(
                                                    address_mode="address_mode_example",
                                                    args=[
                                                        "args_example",
                                                    ],
                                                    body="body_example",
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    command="command_example",
                                                    expose=True,
                                                    failures_before_critical=1,
                                                    grpc_service="grpc_service_example",
                                                    grpc_use_tls=True,
                                                    header={
                                                        "key": [
                                                            "key_example",
                                                        ],
                                                    },
                                                    id="id_example",
                                                    initial_status="initial_status_example",
                                                    interval=1,
                                                    method="method_example",
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    path="path_example",
                                                    port_label="port_label_example",
                                                    protocol="protocol_example",
                                                    success_before_passing=1,
                                                    tls_skip_verify=True,
                                                    task_name="task_name_example",
                                                    timeout=1,
                                                    type="type_example",
                                                ),
                                            ],
                                            connect=ConsulConnect(
                                                gateway=ConsulGateway(
                                                    ingress=ConsulIngressConfigEntry(
                                                        listeners=[
                                                            ConsulIngressListener(
                                                                port=1,
                                                                protocol="protocol_example",
                                                                services=[
                                                                    ConsulIngressService(
                                                                        hosts=[
                                                                            "hosts_example",
                                                                        ],
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                        tls=ConsulGatewayTLSConfig(
                                                            enabled=True,
                                                        ),
                                                    ),
                                                    mesh=None,
                                                    proxy=ConsulGatewayProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        connect_timeout=1,
                                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                        envoy_gateway_bind_addresses={
                                                            "key": ConsulGatewayBindAddress(
                                                                address="address_example",
                                                                name="name_example",
                                                                port=1,
                                                            ),
                                                        },
                                                        envoy_gateway_bind_tagged_addresses=True,
                                                        envoy_gateway_no_default_bind=True,
                                                    ),
                                                    terminating=ConsulTerminatingConfigEntry(
                                                        services=[
                                                            ConsulLinkedService(
                                                                ca_file="ca_file_example",
                                                                cert_file="cert_file_example",
                                                                key_file="key_file_example",
                                                                name="name_example",
                                                                sni="sni_example",
                                                            ),
                                                        ],
                                                    ),
                                                ),
                                                native=True,
                                                sidecar_service=ConsulSidecarService(
                                                    disable_default_tcp_check=True,
                                                    port="port_example",
                                                    proxy=ConsulProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        expose_config=ConsulExposeConfig(
                                                            path=[
                                                                ConsulExposePath(
                                                                    listener_port="listener_port_example",
                                                                    local_path_port=1,
                                                                    path="path_example",
                                                                    protocol="protocol_example",
                                                                ),
                                                            ],
                                                        ),
                                                        local_service_address="local_service_address_example",
                                                        local_service_port=1,
                                                        upstreams=[
                                                            ConsulUpstream(
                                                                datacenter="datacenter_example",
                                                                destination_name="destination_name_example",
                                                                local_bind_address="local_bind_address_example",
                                                                local_bind_port=1,
                                                                mesh_gateway=ConsulMeshGateway(
                                                                    mode="mode_example",
                                                                ),
                                                            ),
                                                        ],
                                                    ),
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                ),
                                                sidecar_task=SidecarTask(
                                                    config={
                                                        "key": None,
                                                    },
                                                    driver="driver_example",
                                                    env={
                                                        "key": "key_example",
                                                    },
                                                    kill_signal="kill_signal_example",
                                                    kill_timeout=1,
                                                    log_config=LogConfig(
                                                        max_file_size_mb=1,
                                                        max_files=1,
                                                    ),
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    resources=Resources(
                                                        cpu=1,
                                                        cores=1,
                                                        devices=[
                                                            RequestedDevice(
                                                                affinities=[
                                                                    Affinity(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                        weight=-128,
                                                                    ),
                                                                ],
                                                                constraints=[
                                                                    Constraint(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                    ),
                                                                ],
                                                                count=0,
                                                                name="name_example",
                                                            ),
                                                        ],
                                                        disk_mb=1,
                                                        iops=1,
                                                        memory_mb=1,
                                                        memory_max_mb=1,
                                                        networks=[
                                                            NetworkResource(
                                                                cidr="cidr_example",
                                                                dns=DNSConfig(
                                                                    options=[
                                                                        "options_example",
                                                                    ],
                                                                    searches=[
                                                                        "searches_example",
                                                                    ],
                                                                    servers=[
                                                                        "servers_example",
                                                                    ],
                                                                ),
                                                                device="device_example",
                                                                dynamic_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                                hostname="hostname_example",
                                                                ip="ip_example",
                                                                m_bits=1,
                                                                mode="mode_example",
                                                                reserved_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                    ),
                                                    shutdown_delay=1,
                                                    user="user_example",
                                                ),
                                            ),
                                            enable_tag_override=True,
                                            id="id_example",
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            on_update="on_update_example",
                                            port_label="port_label_example",
                                            provider="provider_example",
                                            tags=[
                                                "tags_example",
                                            ],
                                            task_name="task_name_example",
                                        ),
                                    ],
                                    shutdown_delay=1,
                                    spreads=[
                                        Spread(
                                            attribute="attribute_example",
                                            spread_target=[
                                                SpreadTarget(
                                                    percent=0,
                                                    value="value_example",
                                                ),
                                            ],
                                            weight=-128,
                                        ),
                                    ],
                                    stop_after_client_disconnect=1,
                                    tasks=[
                                        Task(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            artifacts=[
                                                TaskArtifact(
                                                    getter_headers={
                                                        "key": "key_example",
                                                    },
                                                    getter_mode="getter_mode_example",
                                                    getter_options={
                                                        "key": "key_example",
                                                    },
                                                    getter_source="getter_source_example",
                                                    relative_dest="relative_dest_example",
                                                ),
                                            ],
                                            csi_plugin_config=TaskCSIPluginConfig(
                                                id="id_example",
                                                mount_dir="mount_dir_example",
                                                type="type_example",
                                            ),
                                            config={
                                                "key": None,
                                            },
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            dispatch_payload=DispatchPayloadConfig(
                                                file="file_example",
                                            ),
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            kind="kind_example",
                                            leader=True,
                                            lifecycle=TaskLifecycle(
                                                hook="hook_example",
                                                sidecar=True,
                                            ),
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            restart_policy=RestartPolicy(
                                                attempts=1,
                                                delay=1,
                                                interval=1,
                                                mode="mode_example",
                                            ),
                                            scaling_policies=[
                                                ScalingPolicy(
                                                    create_index=0,
                                                    enabled=True,
                                                    id="id_example",
                                                    max=1,
                                                    min=1,
                                                    modify_index=0,
                                                    namespace="namespace_example",
                                                    policy={
                                                        "key": None,
                                                    },
                                                    target={
                                                        "key": "key_example",
                                                    },
                                                    type="type_example",
                                                ),
                                            ],
                                            services=[
                                                Service(
                                                    address_mode="address_mode_example",
                                                    canary_meta={
                                                        "key": "key_example",
                                                    },
                                                    canary_tags=[
                                                        "canary_tags_example",
                                                    ],
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    checks=[
                                                        ServiceCheck(
                                                            address_mode="address_mode_example",
                                                            args=[
                                                                "args_example",
                                                            ],
                                                            body="body_example",
                                                            check_restart=CheckRestart(
                                                                grace=1,
                                                                ignore_warnings=True,
                                                                limit=1,
                                                            ),
                                                            command="command_example",
                                                            expose=True,
                                                            failures_before_critical=1,
                                                            grpc_service="grpc_service_example",
                                                            grpc_use_tls=True,
                                                            header={
                                                                "key": [
                                                                    "key_example",
                                                                ],
                                                            },
                                                            id="id_example",
                                                            initial_status="initial_status_example",
                                                            interval=1,
                                                            method="method_example",
                                                            name="name_example",
                                                            on_update="on_update_example",
                                                            path="path_example",
                                                            port_label="port_label_example",
                                                            protocol="protocol_example",
                                                            success_before_passing=1,
                                                            tls_skip_verify=True,
                                                            task_name="task_name_example",
                                                            timeout=1,
                                                            type="type_example",
                                                        ),
                                                    ],
                                                    connect=ConsulConnect(
                                                        gateway=ConsulGateway(
                                                            ingress=ConsulIngressConfigEntry(
                                                                listeners=[
                                                                    ConsulIngressListener(
                                                                        port=1,
                                                                        protocol="protocol_example",
                                                                        services=[
                                                                            ConsulIngressService(
                                                                                hosts=[
                                                                                    "hosts_example",
                                                                                ],
                                                                                name="name_example",
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                                tls=ConsulGatewayTLSConfig(
                                                                    enabled=True,
                                                                ),
                                                            ),
                                                            mesh=None,
                                                            proxy=ConsulGatewayProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                connect_timeout=1,
                                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                                envoy_gateway_bind_addresses={
                                                                    "key": ConsulGatewayBindAddress(
                                                                        address="address_example",
                                                                        name="name_example",
                                                                        port=1,
                                                                    ),
                                                                },
                                                                envoy_gateway_bind_tagged_addresses=True,
                                                                envoy_gateway_no_default_bind=True,
                                                            ),
                                                            terminating=ConsulTerminatingConfigEntry(
                                                                services=[
                                                                    ConsulLinkedService(
                                                                        ca_file="ca_file_example",
                                                                        cert_file="cert_file_example",
                                                                        key_file="key_file_example",
                                                                        name="name_example",
                                                                        sni="sni_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ),
                                                        native=True,
                                                        sidecar_service=ConsulSidecarService(
                                                            disable_default_tcp_check=True,
                                                            port="port_example",
                                                            proxy=ConsulProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                expose_config=ConsulExposeConfig(
                                                                    path=[
                                                                        ConsulExposePath(
                                                                            listener_port="listener_port_example",
                                                                            local_path_port=1,
                                                                            path="path_example",
                                                                            protocol="protocol_example",
                                                                        ),
                                                                    ],
                                                                ),
                                                                local_service_address="local_service_address_example",
                                                                local_service_port=1,
                                                                upstreams=[
                                                                    ConsulUpstream(
                                                                        datacenter="datacenter_example",
                                                                        destination_name="destination_name_example",
                                                                        local_bind_address="local_bind_address_example",
                                                                        local_bind_port=1,
                                                                        mesh_gateway=ConsulMeshGateway(
                                                                            mode="mode_example",
                                                                        ),
                                                                    ),
                                                                ],
                                                            ),
                                                            tags=[
                                                                "tags_example",
                                                            ],
                                                        ),
                                                        sidecar_task=SidecarTask(
                                                            config={
                                                                "key": None,
                                                            },
                                                            driver="driver_example",
                                                            env={
                                                                "key": "key_example",
                                                            },
                                                            kill_signal="kill_signal_example",
                                                            kill_timeout=1,
                                                            log_config=LogConfig(
                                                                max_file_size_mb=1,
                                                                max_files=1,
                                                            ),
                                                            meta={
                                                                "key": "key_example",
                                                            },
                                                            name="name_example",
                                                            resources=Resources(
                                                                cpu=1,
                                                                cores=1,
                                                                devices=[
                                                                    RequestedDevice(
                                                                        affinities=[
                                                                            Affinity(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                                weight=-128,
                                                                            ),
                                                                        ],
                                                                        constraints=[
                                                                            Constraint(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                            ),
                                                                        ],
                                                                        count=0,
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                                disk_mb=1,
                                                                iops=1,
                                                                memory_mb=1,
                                                                memory_max_mb=1,
                                                                networks=[
                                                                    NetworkResource(
                                                                        cidr="cidr_example",
                                                                        dns=DNSConfig(
                                                                            options=[
                                                                                "options_example",
                                                                            ],
                                                                            searches=[
                                                                                "searches_example",
                                                                            ],
                                                                            servers=[
                                                                                "servers_example",
                                                                            ],
                                                                        ),
                                                                        device="device_example",
                                                                        dynamic_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                        hostname="hostname_example",
                                                                        ip="ip_example",
                                                                        m_bits=1,
                                                                        mode="mode_example",
                                                                        reserved_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                            ),
                                                            shutdown_delay=1,
                                                            user="user_example",
                                                        ),
                                                    ),
                                                    enable_tag_override=True,
                                                    id="id_example",
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    port_label="port_label_example",
                                                    provider="provider_example",
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                    task_name="task_name_example",
                                                ),
                                            ],
                                            shutdown_delay=1,
                                            templates=[
                                                Template(
                                                    change_mode="change_mode_example",
                                                    change_signal="change_signal_example",
                                                    dest_path="dest_path_example",
                                                    embedded_tmpl="embedded_tmpl_example",
                                                    envvars=True,
                                                    left_delim="left_delim_example",
                                                    perms="perms_example",
                                                    right_delim="right_delim_example",
                                                    source_path="source_path_example",
                                                    splay=1,
                                                    vault_grace=1,
                                                    wait=WaitConfig(
                                                        max=1,
                                                        min=1,
                                                    ),
                                                ),
                                            ],
                                            user="user_example",
                                            vault=Vault(
                                                change_mode="change_mode_example",
                                                change_signal="change_signal_example",
                                                entity_alias="entity_alias_example",
                                                env=True,
                                                namespace="namespace_example",
                                                policies=[
                                                    "policies_example",
                                                ],
                                            ),
                                            volume_mounts=[
                                                VolumeMount(
                                                    destination="destination_example",
                                                    propagation_mode="propagation_mode_example",
                                                    read_only=True,
                                                    volume="volume_example",
                                                ),
                                            ],
                                        ),
                                    ],
                                    update=UpdateStrategy(
                                        auto_promote=True,
                                        auto_revert=True,
                                        canary=1,
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                        progress_deadline=1,
                                        stagger=1,
                                    ),
                                    volumes={
                                        "key": VolumeRequest(
                                            access_mode="access_mode_example",
                                            attachment_mode="attachment_mode_example",
                                            mount_options=CSIMountOptions(
                                                fs_type="fs_type_example",
                                                mount_flags=[
                                                    "mount_flags_example",
                                                ],
                                            ),
                                            name="name_example",
                                            per_alloc=True,
                                            read_only=True,
                                            source="source_example",
                                            type="type_example",
                                        ),
                                    },
                                ),
                            ],
                            type="type_example",
                            update=UpdateStrategy(
                                auto_promote=True,
                                auto_revert=True,
                                canary=1,
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                                progress_deadline=1,
                                stagger=1,
                            ),
                            vault_namespace="vault_namespace_example",
                            vault_token="vault_token_example",
                            version=0,
                        ),
                        job_id="job_id_example",
                        metrics=AllocationMetric(
                            allocation_time=1,
                            class_exhausted={
                                "key": 1,
                            },
                            class_filtered={
                                "key": 1,
                            },
                            coalesced_failures=1,
                            constraint_filtered={
                                "key": 1,
                            },
                            dimension_exhausted={
                                "key": 1,
                            },
                            nodes_available={
                                "key": 1,
                            },
                            nodes_evaluated=1,
                            nodes_exhausted=1,
                            nodes_filtered=1,
                            quota_exhausted=[
                                "quota_exhausted_example",
                            ],
                            resources_exhausted={
                                "key": Resources(
                                    cpu=1,
                                    cores=1,
                                    devices=[
                                        RequestedDevice(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            count=0,
                                            name="name_example",
                                        ),
                                    ],
                                    disk_mb=1,
                                    iops=1,
                                    memory_mb=1,
                                    memory_max_mb=1,
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                            score_meta_data=[
                                NodeScoreMeta(
                                    node_id="node_id_example",
                                    norm_score=3.14,
                                    scores={
                                        "key": 3.14,
                                    },
                                ),
                            ],
                            scores={
                                "key": 3.14,
                            },
                        ),
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        next_allocation="next_allocation_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        previous_allocation="previous_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        resources=Resources(
                            cpu=1,
                            cores=1,
                            devices=[
                                RequestedDevice(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    count=0,
                                    name="name_example",
                                ),
                            ],
                            disk_mb=1,
                            iops=1,
                            memory_mb=1,
                            memory_max_mb=1,
                            networks=[
                                NetworkResource(
                                    cidr="cidr_example",
                                    dns=DNSConfig(
                                        options=[
                                            "options_example",
                                        ],
                                        searches=[
                                            "searches_example",
                                        ],
                                        servers=[
                                            "servers_example",
                                        ],
                                    ),
                                    device="device_example",
                                    dynamic_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                    hostname="hostname_example",
                                    ip="ip_example",
                                    m_bits=1,
                                    mode="mode_example",
                                    reserved_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                ),
                            ],
                        ),
                        services={
                            "key": "key_example",
                        },
                        task_group="task_group_example",
                        task_resources={
                            "key": Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                        },
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                },
            ),
        ],
    ) # CSIVolumeCreateRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.create_volume(volume_id, action, csi_volume_create_request)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->create_volume: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.create_volume(volume_id, action, csi_volume_create_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->create_volume: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume unique identifier. |
 **action** | **str**| The action to perform on the Volume (create, detach, delete). |
 **csi_volume_create_request** | [**CSIVolumeCreateRequest**](CSIVolumeCreateRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_snapshot**
> delete_snapshot()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)
    plugin_id = "plugin_id_example" # str | Filters volume lists by plugin ID. (optional)
    snapshot_id = "snapshot_id_example" # str | The ID of the snapshot to target. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.delete_snapshot(region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token, plugin_id=plugin_id, snapshot_id=snapshot_id)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->delete_snapshot: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]
 **plugin_id** | **str**| Filters volume lists by plugin ID. | [optional]
 **snapshot_id** | **str**| The ID of the snapshot to target. | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_volume_registration**
> delete_volume_registration(volume_id)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    volume_id = "volumeId_example" # str | Volume unique identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)
    force = "force_example" # str | Used to force the de-registration of a volume. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.delete_volume_registration(volume_id)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->delete_volume_registration: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.delete_volume_registration(volume_id, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token, force=force)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->delete_volume_registration: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume unique identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]
 **force** | **str**| Used to force the de-registration of a volume. | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **detach_or_delete_volume**
> detach_or_delete_volume(volume_id, action)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    volume_id = "volumeId_example" # str | Volume unique identifier.
    action = "action_example" # str | The action to perform on the Volume (create, detach, delete).
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)
    node = "node_example" # str | Specifies node to target volume operation for. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.detach_or_delete_volume(volume_id, action)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->detach_or_delete_volume: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.detach_or_delete_volume(volume_id, action, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token, node=node)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->detach_or_delete_volume: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume unique identifier. |
 **action** | **str**| The action to perform on the Volume (create, detach, delete). |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]
 **node** | **str**| Specifies node to target volume operation for. | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_external_volumes**
> CSIVolumeListExternalResponse get_external_volumes()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_volume_list_external_response import CSIVolumeListExternalResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    plugin_id = "plugin_id_example" # str | Filters volume lists by plugin ID. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_external_volumes(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, plugin_id=plugin_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->get_external_volumes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]
 **plugin_id** | **str**| Filters volume lists by plugin ID. | [optional]

### Return type

[**CSIVolumeListExternalResponse**](CSIVolumeListExternalResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_snapshots**
> CSISnapshotListResponse get_snapshots()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_snapshot_list_response import CSISnapshotListResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    plugin_id = "plugin_id_example" # str | Filters volume lists by plugin ID. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_snapshots(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, plugin_id=plugin_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->get_snapshots: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]
 **plugin_id** | **str**| Filters volume lists by plugin ID. | [optional]

### Return type

[**CSISnapshotListResponse**](CSISnapshotListResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_volume**
> CSIVolume get_volume(volume_id)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_volume import CSIVolume
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    volume_id = "volumeId_example" # str | Volume unique identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_volume(volume_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->get_volume: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_volume(volume_id, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->get_volume: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume unique identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]

### Return type

[**CSIVolume**](CSIVolume.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_volumes**
> [CSIVolumeListStub] get_volumes()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_volume_list_stub import CSIVolumeListStub
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    node_id = "node_id_example" # str | Filters volume lists by node ID. (optional)
    plugin_id = "plugin_id_example" # str | Filters volume lists by plugin ID. (optional)
    type = "type_example" # str | Filters volume lists to a specific type. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_volumes(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, node_id=node_id, plugin_id=plugin_id, type=type)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->get_volumes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]
 **node_id** | **str**| Filters volume lists by node ID. | [optional]
 **plugin_id** | **str**| Filters volume lists by plugin ID. | [optional]
 **type** | **str**| Filters volume lists to a specific type. | [optional]

### Return type

[**[CSIVolumeListStub]**](CSIVolumeListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_snapshot**
> CSISnapshotCreateResponse post_snapshot(csi_snapshot_create_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_snapshot_create_request import CSISnapshotCreateRequest
from nomad_client.model.csi_snapshot_create_response import CSISnapshotCreateResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    csi_snapshot_create_request = CSISnapshotCreateRequest(
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
        snapshots=[
            CSISnapshot(
                create_time=1,
                external_source_volume_id="external_source_volume_id_example",
                id="id_example",
                is_ready=True,
                name="name_example",
                parameters={
                    "key": "key_example",
                },
                plugin_id="plugin_id_example",
                secrets=CSISecrets(
                    key="key_example",
                ),
                size_bytes=1,
                source_volume_id="source_volume_id_example",
            ),
        ],
    ) # CSISnapshotCreateRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_snapshot(csi_snapshot_create_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->post_snapshot: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_snapshot(csi_snapshot_create_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->post_snapshot: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **csi_snapshot_create_request** | [**CSISnapshotCreateRequest**](CSISnapshotCreateRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**CSISnapshotCreateResponse**](CSISnapshotCreateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_volume**
> post_volume(csi_volume_register_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_volume_register_request import CSIVolumeRegisterRequest
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    csi_volume_register_request = CSIVolumeRegisterRequest(
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
        volumes=[
            CSIVolume(
                access_mode="access_mode_example",
                allocations=[
                    AllocationListStub(
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job_id="job_id_example",
                        job_type="job_type_example",
                        job_version=0,
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        task_group="task_group_example",
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                ],
                attachment_mode="attachment_mode_example",
                capacity=1,
                clone_id="clone_id_example",
                context={
                    "key": "key_example",
                },
                controller_required=True,
                controllers_expected=1,
                controllers_healthy=1,
                create_index=0,
                external_id="external_id_example",
                id="id_example",
                modify_index=0,
                mount_options=CSIMountOptions(
                    fs_type="fs_type_example",
                    mount_flags=[
                        "mount_flags_example",
                    ],
                ),
                name="name_example",
                namespace="namespace_example",
                nodes_expected=1,
                nodes_healthy=1,
                parameters={
                    "key": "key_example",
                },
                plugin_id="plugin_id_example",
                provider="provider_example",
                provider_version="provider_version_example",
                read_allocs={
                    "key": Allocation(
                        alloc_modify_index=0,
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_id="deployment_id_example",
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        desired_transition=DesiredTransition(
                            migrate=True,
                            reschedule=True,
                        ),
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job=Job(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            all_at_once=True,
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            consul_namespace="consul_namespace_example",
                            consul_token="consul_token_example",
                            create_index=0,
                            datacenters=[
                                "datacenters_example",
                            ],
                            dispatch_idempotency_token="dispatch_idempotency_token_example",
                            dispatched=True,
                            id="id_example",
                            job_modify_index=0,
                            meta={
                                "key": "key_example",
                            },
                            migrate=MigrateStrategy(
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                            ),
                            modify_index=0,
                            multiregion=Multiregion(
                                regions=[
                                    MultiregionRegion(
                                        count=1,
                                        datacenters=[
                                            "datacenters_example",
                                        ],
                                        meta={
                                            "key": "key_example",
                                        },
                                        name="name_example",
                                    ),
                                ],
                                strategy=MultiregionStrategy(
                                    max_parallel=1,
                                    on_failure="on_failure_example",
                                ),
                            ),
                            name="name_example",
                            namespace="namespace_example",
                            nomad_token_id="nomad_token_id_example",
                            parameterized_job=ParameterizedJobConfig(
                                meta_optional=[
                                    "meta_optional_example",
                                ],
                                meta_required=[
                                    "meta_required_example",
                                ],
                                payload="payload_example",
                            ),
                            parent_id="parent_id_example",
                            payload='YQ==',
                            periodic=PeriodicConfig(
                                enabled=True,
                                prohibit_overlap=True,
                                spec="spec_example",
                                spec_type="spec_type_example",
                                time_zone="time_zone_example",
                            ),
                            priority=1,
                            region="region_example",
                            reschedule=ReschedulePolicy(
                                attempts=1,
                                delay=1,
                                delay_function="delay_function_example",
                                interval=1,
                                max_delay=1,
                                unlimited=True,
                            ),
                            spreads=[
                                Spread(
                                    attribute="attribute_example",
                                    spread_target=[
                                        SpreadTarget(
                                            percent=0,
                                            value="value_example",
                                        ),
                                    ],
                                    weight=-128,
                                ),
                            ],
                            stable=True,
                            status="status_example",
                            status_description="status_description_example",
                            stop=True,
                            submit_time=1,
                            task_groups=[
                                TaskGroup(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    consul=Consul(
                                        namespace="namespace_example",
                                    ),
                                    count=1,
                                    ephemeral_disk=EphemeralDisk(
                                        migrate=True,
                                        size_mb=1,
                                        sticky=True,
                                    ),
                                    max_client_disconnect=1,
                                    meta={
                                        "key": "key_example",
                                    },
                                    migrate=MigrateStrategy(
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                    ),
                                    name="name_example",
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                    reschedule_policy=ReschedulePolicy(
                                        attempts=1,
                                        delay=1,
                                        delay_function="delay_function_example",
                                        interval=1,
                                        max_delay=1,
                                        unlimited=True,
                                    ),
                                    restart_policy=RestartPolicy(
                                        attempts=1,
                                        delay=1,
                                        interval=1,
                                        mode="mode_example",
                                    ),
                                    scaling=ScalingPolicy(
                                        create_index=0,
                                        enabled=True,
                                        id="id_example",
                                        max=1,
                                        min=1,
                                        modify_index=0,
                                        namespace="namespace_example",
                                        policy={
                                            "key": None,
                                        },
                                        target={
                                            "key": "key_example",
                                        },
                                        type="type_example",
                                    ),
                                    services=[
                                        Service(
                                            address_mode="address_mode_example",
                                            canary_meta={
                                                "key": "key_example",
                                            },
                                            canary_tags=[
                                                "canary_tags_example",
                                            ],
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            checks=[
                                                ServiceCheck(
                                                    address_mode="address_mode_example",
                                                    args=[
                                                        "args_example",
                                                    ],
                                                    body="body_example",
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    command="command_example",
                                                    expose=True,
                                                    failures_before_critical=1,
                                                    grpc_service="grpc_service_example",
                                                    grpc_use_tls=True,
                                                    header={
                                                        "key": [
                                                            "key_example",
                                                        ],
                                                    },
                                                    id="id_example",
                                                    initial_status="initial_status_example",
                                                    interval=1,
                                                    method="method_example",
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    path="path_example",
                                                    port_label="port_label_example",
                                                    protocol="protocol_example",
                                                    success_before_passing=1,
                                                    tls_skip_verify=True,
                                                    task_name="task_name_example",
                                                    timeout=1,
                                                    type="type_example",
                                                ),
                                            ],
                                            connect=ConsulConnect(
                                                gateway=ConsulGateway(
                                                    ingress=ConsulIngressConfigEntry(
                                                        listeners=[
                                                            ConsulIngressListener(
                                                                port=1,
                                                                protocol="protocol_example",
                                                                services=[
                                                                    ConsulIngressService(
                                                                        hosts=[
                                                                            "hosts_example",
                                                                        ],
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                        tls=ConsulGatewayTLSConfig(
                                                            enabled=True,
                                                        ),
                                                    ),
                                                    mesh=None,
                                                    proxy=ConsulGatewayProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        connect_timeout=1,
                                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                        envoy_gateway_bind_addresses={
                                                            "key": ConsulGatewayBindAddress(
                                                                address="address_example",
                                                                name="name_example",
                                                                port=1,
                                                            ),
                                                        },
                                                        envoy_gateway_bind_tagged_addresses=True,
                                                        envoy_gateway_no_default_bind=True,
                                                    ),
                                                    terminating=ConsulTerminatingConfigEntry(
                                                        services=[
                                                            ConsulLinkedService(
                                                                ca_file="ca_file_example",
                                                                cert_file="cert_file_example",
                                                                key_file="key_file_example",
                                                                name="name_example",
                                                                sni="sni_example",
                                                            ),
                                                        ],
                                                    ),
                                                ),
                                                native=True,
                                                sidecar_service=ConsulSidecarService(
                                                    disable_default_tcp_check=True,
                                                    port="port_example",
                                                    proxy=ConsulProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        expose_config=ConsulExposeConfig(
                                                            path=[
                                                                ConsulExposePath(
                                                                    listener_port="listener_port_example",
                                                                    local_path_port=1,
                                                                    path="path_example",
                                                                    protocol="protocol_example",
                                                                ),
                                                            ],
                                                        ),
                                                        local_service_address="local_service_address_example",
                                                        local_service_port=1,
                                                        upstreams=[
                                                            ConsulUpstream(
                                                                datacenter="datacenter_example",
                                                                destination_name="destination_name_example",
                                                                local_bind_address="local_bind_address_example",
                                                                local_bind_port=1,
                                                                mesh_gateway=ConsulMeshGateway(
                                                                    mode="mode_example",
                                                                ),
                                                            ),
                                                        ],
                                                    ),
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                ),
                                                sidecar_task=SidecarTask(
                                                    config={
                                                        "key": None,
                                                    },
                                                    driver="driver_example",
                                                    env={
                                                        "key": "key_example",
                                                    },
                                                    kill_signal="kill_signal_example",
                                                    kill_timeout=1,
                                                    log_config=LogConfig(
                                                        max_file_size_mb=1,
                                                        max_files=1,
                                                    ),
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    resources=Resources(
                                                        cpu=1,
                                                        cores=1,
                                                        devices=[
                                                            RequestedDevice(
                                                                affinities=[
                                                                    Affinity(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                        weight=-128,
                                                                    ),
                                                                ],
                                                                constraints=[
                                                                    Constraint(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                    ),
                                                                ],
                                                                count=0,
                                                                name="name_example",
                                                            ),
                                                        ],
                                                        disk_mb=1,
                                                        iops=1,
                                                        memory_mb=1,
                                                        memory_max_mb=1,
                                                        networks=[
                                                            NetworkResource(
                                                                cidr="cidr_example",
                                                                dns=DNSConfig(
                                                                    options=[
                                                                        "options_example",
                                                                    ],
                                                                    searches=[
                                                                        "searches_example",
                                                                    ],
                                                                    servers=[
                                                                        "servers_example",
                                                                    ],
                                                                ),
                                                                device="device_example",
                                                                dynamic_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                                hostname="hostname_example",
                                                                ip="ip_example",
                                                                m_bits=1,
                                                                mode="mode_example",
                                                                reserved_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                    ),
                                                    shutdown_delay=1,
                                                    user="user_example",
                                                ),
                                            ),
                                            enable_tag_override=True,
                                            id="id_example",
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            on_update="on_update_example",
                                            port_label="port_label_example",
                                            provider="provider_example",
                                            tags=[
                                                "tags_example",
                                            ],
                                            task_name="task_name_example",
                                        ),
                                    ],
                                    shutdown_delay=1,
                                    spreads=[
                                        Spread(
                                            attribute="attribute_example",
                                            spread_target=[
                                                SpreadTarget(
                                                    percent=0,
                                                    value="value_example",
                                                ),
                                            ],
                                            weight=-128,
                                        ),
                                    ],
                                    stop_after_client_disconnect=1,
                                    tasks=[
                                        Task(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            artifacts=[
                                                TaskArtifact(
                                                    getter_headers={
                                                        "key": "key_example",
                                                    },
                                                    getter_mode="getter_mode_example",
                                                    getter_options={
                                                        "key": "key_example",
                                                    },
                                                    getter_source="getter_source_example",
                                                    relative_dest="relative_dest_example",
                                                ),
                                            ],
                                            csi_plugin_config=TaskCSIPluginConfig(
                                                id="id_example",
                                                mount_dir="mount_dir_example",
                                                type="type_example",
                                            ),
                                            config={
                                                "key": None,
                                            },
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            dispatch_payload=DispatchPayloadConfig(
                                                file="file_example",
                                            ),
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            kind="kind_example",
                                            leader=True,
                                            lifecycle=TaskLifecycle(
                                                hook="hook_example",
                                                sidecar=True,
                                            ),
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            restart_policy=RestartPolicy(
                                                attempts=1,
                                                delay=1,
                                                interval=1,
                                                mode="mode_example",
                                            ),
                                            scaling_policies=[
                                                ScalingPolicy(
                                                    create_index=0,
                                                    enabled=True,
                                                    id="id_example",
                                                    max=1,
                                                    min=1,
                                                    modify_index=0,
                                                    namespace="namespace_example",
                                                    policy={
                                                        "key": None,
                                                    },
                                                    target={
                                                        "key": "key_example",
                                                    },
                                                    type="type_example",
                                                ),
                                            ],
                                            services=[
                                                Service(
                                                    address_mode="address_mode_example",
                                                    canary_meta={
                                                        "key": "key_example",
                                                    },
                                                    canary_tags=[
                                                        "canary_tags_example",
                                                    ],
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    checks=[
                                                        ServiceCheck(
                                                            address_mode="address_mode_example",
                                                            args=[
                                                                "args_example",
                                                            ],
                                                            body="body_example",
                                                            check_restart=CheckRestart(
                                                                grace=1,
                                                                ignore_warnings=True,
                                                                limit=1,
                                                            ),
                                                            command="command_example",
                                                            expose=True,
                                                            failures_before_critical=1,
                                                            grpc_service="grpc_service_example",
                                                            grpc_use_tls=True,
                                                            header={
                                                                "key": [
                                                                    "key_example",
                                                                ],
                                                            },
                                                            id="id_example",
                                                            initial_status="initial_status_example",
                                                            interval=1,
                                                            method="method_example",
                                                            name="name_example",
                                                            on_update="on_update_example",
                                                            path="path_example",
                                                            port_label="port_label_example",
                                                            protocol="protocol_example",
                                                            success_before_passing=1,
                                                            tls_skip_verify=True,
                                                            task_name="task_name_example",
                                                            timeout=1,
                                                            type="type_example",
                                                        ),
                                                    ],
                                                    connect=ConsulConnect(
                                                        gateway=ConsulGateway(
                                                            ingress=ConsulIngressConfigEntry(
                                                                listeners=[
                                                                    ConsulIngressListener(
                                                                        port=1,
                                                                        protocol="protocol_example",
                                                                        services=[
                                                                            ConsulIngressService(
                                                                                hosts=[
                                                                                    "hosts_example",
                                                                                ],
                                                                                name="name_example",
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                                tls=ConsulGatewayTLSConfig(
                                                                    enabled=True,
                                                                ),
                                                            ),
                                                            mesh=None,
                                                            proxy=ConsulGatewayProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                connect_timeout=1,
                                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                                envoy_gateway_bind_addresses={
                                                                    "key": ConsulGatewayBindAddress(
                                                                        address="address_example",
                                                                        name="name_example",
                                                                        port=1,
                                                                    ),
                                                                },
                                                                envoy_gateway_bind_tagged_addresses=True,
                                                                envoy_gateway_no_default_bind=True,
                                                            ),
                                                            terminating=ConsulTerminatingConfigEntry(
                                                                services=[
                                                                    ConsulLinkedService(
                                                                        ca_file="ca_file_example",
                                                                        cert_file="cert_file_example",
                                                                        key_file="key_file_example",
                                                                        name="name_example",
                                                                        sni="sni_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ),
                                                        native=True,
                                                        sidecar_service=ConsulSidecarService(
                                                            disable_default_tcp_check=True,
                                                            port="port_example",
                                                            proxy=ConsulProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                expose_config=ConsulExposeConfig(
                                                                    path=[
                                                                        ConsulExposePath(
                                                                            listener_port="listener_port_example",
                                                                            local_path_port=1,
                                                                            path="path_example",
                                                                            protocol="protocol_example",
                                                                        ),
                                                                    ],
                                                                ),
                                                                local_service_address="local_service_address_example",
                                                                local_service_port=1,
                                                                upstreams=[
                                                                    ConsulUpstream(
                                                                        datacenter="datacenter_example",
                                                                        destination_name="destination_name_example",
                                                                        local_bind_address="local_bind_address_example",
                                                                        local_bind_port=1,
                                                                        mesh_gateway=ConsulMeshGateway(
                                                                            mode="mode_example",
                                                                        ),
                                                                    ),
                                                                ],
                                                            ),
                                                            tags=[
                                                                "tags_example",
                                                            ],
                                                        ),
                                                        sidecar_task=SidecarTask(
                                                            config={
                                                                "key": None,
                                                            },
                                                            driver="driver_example",
                                                            env={
                                                                "key": "key_example",
                                                            },
                                                            kill_signal="kill_signal_example",
                                                            kill_timeout=1,
                                                            log_config=LogConfig(
                                                                max_file_size_mb=1,
                                                                max_files=1,
                                                            ),
                                                            meta={
                                                                "key": "key_example",
                                                            },
                                                            name="name_example",
                                                            resources=Resources(
                                                                cpu=1,
                                                                cores=1,
                                                                devices=[
                                                                    RequestedDevice(
                                                                        affinities=[
                                                                            Affinity(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                                weight=-128,
                                                                            ),
                                                                        ],
                                                                        constraints=[
                                                                            Constraint(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                            ),
                                                                        ],
                                                                        count=0,
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                                disk_mb=1,
                                                                iops=1,
                                                                memory_mb=1,
                                                                memory_max_mb=1,
                                                                networks=[
                                                                    NetworkResource(
                                                                        cidr="cidr_example",
                                                                        dns=DNSConfig(
                                                                            options=[
                                                                                "options_example",
                                                                            ],
                                                                            searches=[
                                                                                "searches_example",
                                                                            ],
                                                                            servers=[
                                                                                "servers_example",
                                                                            ],
                                                                        ),
                                                                        device="device_example",
                                                                        dynamic_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                        hostname="hostname_example",
                                                                        ip="ip_example",
                                                                        m_bits=1,
                                                                        mode="mode_example",
                                                                        reserved_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                            ),
                                                            shutdown_delay=1,
                                                            user="user_example",
                                                        ),
                                                    ),
                                                    enable_tag_override=True,
                                                    id="id_example",
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    port_label="port_label_example",
                                                    provider="provider_example",
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                    task_name="task_name_example",
                                                ),
                                            ],
                                            shutdown_delay=1,
                                            templates=[
                                                Template(
                                                    change_mode="change_mode_example",
                                                    change_signal="change_signal_example",
                                                    dest_path="dest_path_example",
                                                    embedded_tmpl="embedded_tmpl_example",
                                                    envvars=True,
                                                    left_delim="left_delim_example",
                                                    perms="perms_example",
                                                    right_delim="right_delim_example",
                                                    source_path="source_path_example",
                                                    splay=1,
                                                    vault_grace=1,
                                                    wait=WaitConfig(
                                                        max=1,
                                                        min=1,
                                                    ),
                                                ),
                                            ],
                                            user="user_example",
                                            vault=Vault(
                                                change_mode="change_mode_example",
                                                change_signal="change_signal_example",
                                                entity_alias="entity_alias_example",
                                                env=True,
                                                namespace="namespace_example",
                                                policies=[
                                                    "policies_example",
                                                ],
                                            ),
                                            volume_mounts=[
                                                VolumeMount(
                                                    destination="destination_example",
                                                    propagation_mode="propagation_mode_example",
                                                    read_only=True,
                                                    volume="volume_example",
                                                ),
                                            ],
                                        ),
                                    ],
                                    update=UpdateStrategy(
                                        auto_promote=True,
                                        auto_revert=True,
                                        canary=1,
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                        progress_deadline=1,
                                        stagger=1,
                                    ),
                                    volumes={
                                        "key": VolumeRequest(
                                            access_mode="access_mode_example",
                                            attachment_mode="attachment_mode_example",
                                            mount_options=CSIMountOptions(
                                                fs_type="fs_type_example",
                                                mount_flags=[
                                                    "mount_flags_example",
                                                ],
                                            ),
                                            name="name_example",
                                            per_alloc=True,
                                            read_only=True,
                                            source="source_example",
                                            type="type_example",
                                        ),
                                    },
                                ),
                            ],
                            type="type_example",
                            update=UpdateStrategy(
                                auto_promote=True,
                                auto_revert=True,
                                canary=1,
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                                progress_deadline=1,
                                stagger=1,
                            ),
                            vault_namespace="vault_namespace_example",
                            vault_token="vault_token_example",
                            version=0,
                        ),
                        job_id="job_id_example",
                        metrics=AllocationMetric(
                            allocation_time=1,
                            class_exhausted={
                                "key": 1,
                            },
                            class_filtered={
                                "key": 1,
                            },
                            coalesced_failures=1,
                            constraint_filtered={
                                "key": 1,
                            },
                            dimension_exhausted={
                                "key": 1,
                            },
                            nodes_available={
                                "key": 1,
                            },
                            nodes_evaluated=1,
                            nodes_exhausted=1,
                            nodes_filtered=1,
                            quota_exhausted=[
                                "quota_exhausted_example",
                            ],
                            resources_exhausted={
                                "key": Resources(
                                    cpu=1,
                                    cores=1,
                                    devices=[
                                        RequestedDevice(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            count=0,
                                            name="name_example",
                                        ),
                                    ],
                                    disk_mb=1,
                                    iops=1,
                                    memory_mb=1,
                                    memory_max_mb=1,
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                            score_meta_data=[
                                NodeScoreMeta(
                                    node_id="node_id_example",
                                    norm_score=3.14,
                                    scores={
                                        "key": 3.14,
                                    },
                                ),
                            ],
                            scores={
                                "key": 3.14,
                            },
                        ),
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        next_allocation="next_allocation_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        previous_allocation="previous_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        resources=Resources(
                            cpu=1,
                            cores=1,
                            devices=[
                                RequestedDevice(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    count=0,
                                    name="name_example",
                                ),
                            ],
                            disk_mb=1,
                            iops=1,
                            memory_mb=1,
                            memory_max_mb=1,
                            networks=[
                                NetworkResource(
                                    cidr="cidr_example",
                                    dns=DNSConfig(
                                        options=[
                                            "options_example",
                                        ],
                                        searches=[
                                            "searches_example",
                                        ],
                                        servers=[
                                            "servers_example",
                                        ],
                                    ),
                                    device="device_example",
                                    dynamic_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                    hostname="hostname_example",
                                    ip="ip_example",
                                    m_bits=1,
                                    mode="mode_example",
                                    reserved_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                ),
                            ],
                        ),
                        services={
                            "key": "key_example",
                        },
                        task_group="task_group_example",
                        task_resources={
                            "key": Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                        },
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                },
                requested_capabilities=[
                    CSIVolumeCapability(
                        access_mode="access_mode_example",
                        attachment_mode="attachment_mode_example",
                    ),
                ],
                requested_capacity_max=1,
                requested_capacity_min=1,
                requested_topologies=CSITopologyRequest(
                    preferred=[
                        CSITopology(
                            segments={
                                "key": "key_example",
                            },
                        ),
                    ],
                    required=[
                        CSITopology(
                            segments={
                                "key": "key_example",
                            },
                        ),
                    ],
                ),
                resource_exhausted=dateutil_parser('1970-01-01T00:00:00.00Z'),
                schedulable=True,
                secrets=CSISecrets(
                    key="key_example",
                ),
                snapshot_id="snapshot_id_example",
                topologies=[
                    CSITopology(
                        segments={
                            "key": "key_example",
                        },
                    ),
                ],
                write_allocs={
                    "key": Allocation(
                        alloc_modify_index=0,
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_id="deployment_id_example",
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        desired_transition=DesiredTransition(
                            migrate=True,
                            reschedule=True,
                        ),
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job=Job(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            all_at_once=True,
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            consul_namespace="consul_namespace_example",
                            consul_token="consul_token_example",
                            create_index=0,
                            datacenters=[
                                "datacenters_example",
                            ],
                            dispatch_idempotency_token="dispatch_idempotency_token_example",
                            dispatched=True,
                            id="id_example",
                            job_modify_index=0,
                            meta={
                                "key": "key_example",
                            },
                            migrate=MigrateStrategy(
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                            ),
                            modify_index=0,
                            multiregion=Multiregion(
                                regions=[
                                    MultiregionRegion(
                                        count=1,
                                        datacenters=[
                                            "datacenters_example",
                                        ],
                                        meta={
                                            "key": "key_example",
                                        },
                                        name="name_example",
                                    ),
                                ],
                                strategy=MultiregionStrategy(
                                    max_parallel=1,
                                    on_failure="on_failure_example",
                                ),
                            ),
                            name="name_example",
                            namespace="namespace_example",
                            nomad_token_id="nomad_token_id_example",
                            parameterized_job=ParameterizedJobConfig(
                                meta_optional=[
                                    "meta_optional_example",
                                ],
                                meta_required=[
                                    "meta_required_example",
                                ],
                                payload="payload_example",
                            ),
                            parent_id="parent_id_example",
                            payload='YQ==',
                            periodic=PeriodicConfig(
                                enabled=True,
                                prohibit_overlap=True,
                                spec="spec_example",
                                spec_type="spec_type_example",
                                time_zone="time_zone_example",
                            ),
                            priority=1,
                            region="region_example",
                            reschedule=ReschedulePolicy(
                                attempts=1,
                                delay=1,
                                delay_function="delay_function_example",
                                interval=1,
                                max_delay=1,
                                unlimited=True,
                            ),
                            spreads=[
                                Spread(
                                    attribute="attribute_example",
                                    spread_target=[
                                        SpreadTarget(
                                            percent=0,
                                            value="value_example",
                                        ),
                                    ],
                                    weight=-128,
                                ),
                            ],
                            stable=True,
                            status="status_example",
                            status_description="status_description_example",
                            stop=True,
                            submit_time=1,
                            task_groups=[
                                TaskGroup(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    consul=Consul(
                                        namespace="namespace_example",
                                    ),
                                    count=1,
                                    ephemeral_disk=EphemeralDisk(
                                        migrate=True,
                                        size_mb=1,
                                        sticky=True,
                                    ),
                                    max_client_disconnect=1,
                                    meta={
                                        "key": "key_example",
                                    },
                                    migrate=MigrateStrategy(
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                    ),
                                    name="name_example",
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                    reschedule_policy=ReschedulePolicy(
                                        attempts=1,
                                        delay=1,
                                        delay_function="delay_function_example",
                                        interval=1,
                                        max_delay=1,
                                        unlimited=True,
                                    ),
                                    restart_policy=RestartPolicy(
                                        attempts=1,
                                        delay=1,
                                        interval=1,
                                        mode="mode_example",
                                    ),
                                    scaling=ScalingPolicy(
                                        create_index=0,
                                        enabled=True,
                                        id="id_example",
                                        max=1,
                                        min=1,
                                        modify_index=0,
                                        namespace="namespace_example",
                                        policy={
                                            "key": None,
                                        },
                                        target={
                                            "key": "key_example",
                                        },
                                        type="type_example",
                                    ),
                                    services=[
                                        Service(
                                            address_mode="address_mode_example",
                                            canary_meta={
                                                "key": "key_example",
                                            },
                                            canary_tags=[
                                                "canary_tags_example",
                                            ],
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            checks=[
                                                ServiceCheck(
                                                    address_mode="address_mode_example",
                                                    args=[
                                                        "args_example",
                                                    ],
                                                    body="body_example",
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    command="command_example",
                                                    expose=True,
                                                    failures_before_critical=1,
                                                    grpc_service="grpc_service_example",
                                                    grpc_use_tls=True,
                                                    header={
                                                        "key": [
                                                            "key_example",
                                                        ],
                                                    },
                                                    id="id_example",
                                                    initial_status="initial_status_example",
                                                    interval=1,
                                                    method="method_example",
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    path="path_example",
                                                    port_label="port_label_example",
                                                    protocol="protocol_example",
                                                    success_before_passing=1,
                                                    tls_skip_verify=True,
                                                    task_name="task_name_example",
                                                    timeout=1,
                                                    type="type_example",
                                                ),
                                            ],
                                            connect=ConsulConnect(
                                                gateway=ConsulGateway(
                                                    ingress=ConsulIngressConfigEntry(
                                                        listeners=[
                                                            ConsulIngressListener(
                                                                port=1,
                                                                protocol="protocol_example",
                                                                services=[
                                                                    ConsulIngressService(
                                                                        hosts=[
                                                                            "hosts_example",
                                                                        ],
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                        tls=ConsulGatewayTLSConfig(
                                                            enabled=True,
                                                        ),
                                                    ),
                                                    mesh=None,
                                                    proxy=ConsulGatewayProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        connect_timeout=1,
                                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                        envoy_gateway_bind_addresses={
                                                            "key": ConsulGatewayBindAddress(
                                                                address="address_example",
                                                                name="name_example",
                                                                port=1,
                                                            ),
                                                        },
                                                        envoy_gateway_bind_tagged_addresses=True,
                                                        envoy_gateway_no_default_bind=True,
                                                    ),
                                                    terminating=ConsulTerminatingConfigEntry(
                                                        services=[
                                                            ConsulLinkedService(
                                                                ca_file="ca_file_example",
                                                                cert_file="cert_file_example",
                                                                key_file="key_file_example",
                                                                name="name_example",
                                                                sni="sni_example",
                                                            ),
                                                        ],
                                                    ),
                                                ),
                                                native=True,
                                                sidecar_service=ConsulSidecarService(
                                                    disable_default_tcp_check=True,
                                                    port="port_example",
                                                    proxy=ConsulProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        expose_config=ConsulExposeConfig(
                                                            path=[
                                                                ConsulExposePath(
                                                                    listener_port="listener_port_example",
                                                                    local_path_port=1,
                                                                    path="path_example",
                                                                    protocol="protocol_example",
                                                                ),
                                                            ],
                                                        ),
                                                        local_service_address="local_service_address_example",
                                                        local_service_port=1,
                                                        upstreams=[
                                                            ConsulUpstream(
                                                                datacenter="datacenter_example",
                                                                destination_name="destination_name_example",
                                                                local_bind_address="local_bind_address_example",
                                                                local_bind_port=1,
                                                                mesh_gateway=ConsulMeshGateway(
                                                                    mode="mode_example",
                                                                ),
                                                            ),
                                                        ],
                                                    ),
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                ),
                                                sidecar_task=SidecarTask(
                                                    config={
                                                        "key": None,
                                                    },
                                                    driver="driver_example",
                                                    env={
                                                        "key": "key_example",
                                                    },
                                                    kill_signal="kill_signal_example",
                                                    kill_timeout=1,
                                                    log_config=LogConfig(
                                                        max_file_size_mb=1,
                                                        max_files=1,
                                                    ),
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    resources=Resources(
                                                        cpu=1,
                                                        cores=1,
                                                        devices=[
                                                            RequestedDevice(
                                                                affinities=[
                                                                    Affinity(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                        weight=-128,
                                                                    ),
                                                                ],
                                                                constraints=[
                                                                    Constraint(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                    ),
                                                                ],
                                                                count=0,
                                                                name="name_example",
                                                            ),
                                                        ],
                                                        disk_mb=1,
                                                        iops=1,
                                                        memory_mb=1,
                                                        memory_max_mb=1,
                                                        networks=[
                                                            NetworkResource(
                                                                cidr="cidr_example",
                                                                dns=DNSConfig(
                                                                    options=[
                                                                        "options_example",
                                                                    ],
                                                                    searches=[
                                                                        "searches_example",
                                                                    ],
                                                                    servers=[
                                                                        "servers_example",
                                                                    ],
                                                                ),
                                                                device="device_example",
                                                                dynamic_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                                hostname="hostname_example",
                                                                ip="ip_example",
                                                                m_bits=1,
                                                                mode="mode_example",
                                                                reserved_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                    ),
                                                    shutdown_delay=1,
                                                    user="user_example",
                                                ),
                                            ),
                                            enable_tag_override=True,
                                            id="id_example",
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            on_update="on_update_example",
                                            port_label="port_label_example",
                                            provider="provider_example",
                                            tags=[
                                                "tags_example",
                                            ],
                                            task_name="task_name_example",
                                        ),
                                    ],
                                    shutdown_delay=1,
                                    spreads=[
                                        Spread(
                                            attribute="attribute_example",
                                            spread_target=[
                                                SpreadTarget(
                                                    percent=0,
                                                    value="value_example",
                                                ),
                                            ],
                                            weight=-128,
                                        ),
                                    ],
                                    stop_after_client_disconnect=1,
                                    tasks=[
                                        Task(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            artifacts=[
                                                TaskArtifact(
                                                    getter_headers={
                                                        "key": "key_example",
                                                    },
                                                    getter_mode="getter_mode_example",
                                                    getter_options={
                                                        "key": "key_example",
                                                    },
                                                    getter_source="getter_source_example",
                                                    relative_dest="relative_dest_example",
                                                ),
                                            ],
                                            csi_plugin_config=TaskCSIPluginConfig(
                                                id="id_example",
                                                mount_dir="mount_dir_example",
                                                type="type_example",
                                            ),
                                            config={
                                                "key": None,
                                            },
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            dispatch_payload=DispatchPayloadConfig(
                                                file="file_example",
                                            ),
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            kind="kind_example",
                                            leader=True,
                                            lifecycle=TaskLifecycle(
                                                hook="hook_example",
                                                sidecar=True,
                                            ),
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            restart_policy=RestartPolicy(
                                                attempts=1,
                                                delay=1,
                                                interval=1,
                                                mode="mode_example",
                                            ),
                                            scaling_policies=[
                                                ScalingPolicy(
                                                    create_index=0,
                                                    enabled=True,
                                                    id="id_example",
                                                    max=1,
                                                    min=1,
                                                    modify_index=0,
                                                    namespace="namespace_example",
                                                    policy={
                                                        "key": None,
                                                    },
                                                    target={
                                                        "key": "key_example",
                                                    },
                                                    type="type_example",
                                                ),
                                            ],
                                            services=[
                                                Service(
                                                    address_mode="address_mode_example",
                                                    canary_meta={
                                                        "key": "key_example",
                                                    },
                                                    canary_tags=[
                                                        "canary_tags_example",
                                                    ],
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    checks=[
                                                        ServiceCheck(
                                                            address_mode="address_mode_example",
                                                            args=[
                                                                "args_example",
                                                            ],
                                                            body="body_example",
                                                            check_restart=CheckRestart(
                                                                grace=1,
                                                                ignore_warnings=True,
                                                                limit=1,
                                                            ),
                                                            command="command_example",
                                                            expose=True,
                                                            failures_before_critical=1,
                                                            grpc_service="grpc_service_example",
                                                            grpc_use_tls=True,
                                                            header={
                                                                "key": [
                                                                    "key_example",
                                                                ],
                                                            },
                                                            id="id_example",
                                                            initial_status="initial_status_example",
                                                            interval=1,
                                                            method="method_example",
                                                            name="name_example",
                                                            on_update="on_update_example",
                                                            path="path_example",
                                                            port_label="port_label_example",
                                                            protocol="protocol_example",
                                                            success_before_passing=1,
                                                            tls_skip_verify=True,
                                                            task_name="task_name_example",
                                                            timeout=1,
                                                            type="type_example",
                                                        ),
                                                    ],
                                                    connect=ConsulConnect(
                                                        gateway=ConsulGateway(
                                                            ingress=ConsulIngressConfigEntry(
                                                                listeners=[
                                                                    ConsulIngressListener(
                                                                        port=1,
                                                                        protocol="protocol_example",
                                                                        services=[
                                                                            ConsulIngressService(
                                                                                hosts=[
                                                                                    "hosts_example",
                                                                                ],
                                                                                name="name_example",
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                                tls=ConsulGatewayTLSConfig(
                                                                    enabled=True,
                                                                ),
                                                            ),
                                                            mesh=None,
                                                            proxy=ConsulGatewayProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                connect_timeout=1,
                                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                                envoy_gateway_bind_addresses={
                                                                    "key": ConsulGatewayBindAddress(
                                                                        address="address_example",
                                                                        name="name_example",
                                                                        port=1,
                                                                    ),
                                                                },
                                                                envoy_gateway_bind_tagged_addresses=True,
                                                                envoy_gateway_no_default_bind=True,
                                                            ),
                                                            terminating=ConsulTerminatingConfigEntry(
                                                                services=[
                                                                    ConsulLinkedService(
                                                                        ca_file="ca_file_example",
                                                                        cert_file="cert_file_example",
                                                                        key_file="key_file_example",
                                                                        name="name_example",
                                                                        sni="sni_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ),
                                                        native=True,
                                                        sidecar_service=ConsulSidecarService(
                                                            disable_default_tcp_check=True,
                                                            port="port_example",
                                                            proxy=ConsulProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                expose_config=ConsulExposeConfig(
                                                                    path=[
                                                                        ConsulExposePath(
                                                                            listener_port="listener_port_example",
                                                                            local_path_port=1,
                                                                            path="path_example",
                                                                            protocol="protocol_example",
                                                                        ),
                                                                    ],
                                                                ),
                                                                local_service_address="local_service_address_example",
                                                                local_service_port=1,
                                                                upstreams=[
                                                                    ConsulUpstream(
                                                                        datacenter="datacenter_example",
                                                                        destination_name="destination_name_example",
                                                                        local_bind_address="local_bind_address_example",
                                                                        local_bind_port=1,
                                                                        mesh_gateway=ConsulMeshGateway(
                                                                            mode="mode_example",
                                                                        ),
                                                                    ),
                                                                ],
                                                            ),
                                                            tags=[
                                                                "tags_example",
                                                            ],
                                                        ),
                                                        sidecar_task=SidecarTask(
                                                            config={
                                                                "key": None,
                                                            },
                                                            driver="driver_example",
                                                            env={
                                                                "key": "key_example",
                                                            },
                                                            kill_signal="kill_signal_example",
                                                            kill_timeout=1,
                                                            log_config=LogConfig(
                                                                max_file_size_mb=1,
                                                                max_files=1,
                                                            ),
                                                            meta={
                                                                "key": "key_example",
                                                            },
                                                            name="name_example",
                                                            resources=Resources(
                                                                cpu=1,
                                                                cores=1,
                                                                devices=[
                                                                    RequestedDevice(
                                                                        affinities=[
                                                                            Affinity(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                                weight=-128,
                                                                            ),
                                                                        ],
                                                                        constraints=[
                                                                            Constraint(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                            ),
                                                                        ],
                                                                        count=0,
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                                disk_mb=1,
                                                                iops=1,
                                                                memory_mb=1,
                                                                memory_max_mb=1,
                                                                networks=[
                                                                    NetworkResource(
                                                                        cidr="cidr_example",
                                                                        dns=DNSConfig(
                                                                            options=[
                                                                                "options_example",
                                                                            ],
                                                                            searches=[
                                                                                "searches_example",
                                                                            ],
                                                                            servers=[
                                                                                "servers_example",
                                                                            ],
                                                                        ),
                                                                        device="device_example",
                                                                        dynamic_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                        hostname="hostname_example",
                                                                        ip="ip_example",
                                                                        m_bits=1,
                                                                        mode="mode_example",
                                                                        reserved_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                            ),
                                                            shutdown_delay=1,
                                                            user="user_example",
                                                        ),
                                                    ),
                                                    enable_tag_override=True,
                                                    id="id_example",
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    port_label="port_label_example",
                                                    provider="provider_example",
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                    task_name="task_name_example",
                                                ),
                                            ],
                                            shutdown_delay=1,
                                            templates=[
                                                Template(
                                                    change_mode="change_mode_example",
                                                    change_signal="change_signal_example",
                                                    dest_path="dest_path_example",
                                                    embedded_tmpl="embedded_tmpl_example",
                                                    envvars=True,
                                                    left_delim="left_delim_example",
                                                    perms="perms_example",
                                                    right_delim="right_delim_example",
                                                    source_path="source_path_example",
                                                    splay=1,
                                                    vault_grace=1,
                                                    wait=WaitConfig(
                                                        max=1,
                                                        min=1,
                                                    ),
                                                ),
                                            ],
                                            user="user_example",
                                            vault=Vault(
                                                change_mode="change_mode_example",
                                                change_signal="change_signal_example",
                                                entity_alias="entity_alias_example",
                                                env=True,
                                                namespace="namespace_example",
                                                policies=[
                                                    "policies_example",
                                                ],
                                            ),
                                            volume_mounts=[
                                                VolumeMount(
                                                    destination="destination_example",
                                                    propagation_mode="propagation_mode_example",
                                                    read_only=True,
                                                    volume="volume_example",
                                                ),
                                            ],
                                        ),
                                    ],
                                    update=UpdateStrategy(
                                        auto_promote=True,
                                        auto_revert=True,
                                        canary=1,
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                        progress_deadline=1,
                                        stagger=1,
                                    ),
                                    volumes={
                                        "key": VolumeRequest(
                                            access_mode="access_mode_example",
                                            attachment_mode="attachment_mode_example",
                                            mount_options=CSIMountOptions(
                                                fs_type="fs_type_example",
                                                mount_flags=[
                                                    "mount_flags_example",
                                                ],
                                            ),
                                            name="name_example",
                                            per_alloc=True,
                                            read_only=True,
                                            source="source_example",
                                            type="type_example",
                                        ),
                                    },
                                ),
                            ],
                            type="type_example",
                            update=UpdateStrategy(
                                auto_promote=True,
                                auto_revert=True,
                                canary=1,
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                                progress_deadline=1,
                                stagger=1,
                            ),
                            vault_namespace="vault_namespace_example",
                            vault_token="vault_token_example",
                            version=0,
                        ),
                        job_id="job_id_example",
                        metrics=AllocationMetric(
                            allocation_time=1,
                            class_exhausted={
                                "key": 1,
                            },
                            class_filtered={
                                "key": 1,
                            },
                            coalesced_failures=1,
                            constraint_filtered={
                                "key": 1,
                            },
                            dimension_exhausted={
                                "key": 1,
                            },
                            nodes_available={
                                "key": 1,
                            },
                            nodes_evaluated=1,
                            nodes_exhausted=1,
                            nodes_filtered=1,
                            quota_exhausted=[
                                "quota_exhausted_example",
                            ],
                            resources_exhausted={
                                "key": Resources(
                                    cpu=1,
                                    cores=1,
                                    devices=[
                                        RequestedDevice(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            count=0,
                                            name="name_example",
                                        ),
                                    ],
                                    disk_mb=1,
                                    iops=1,
                                    memory_mb=1,
                                    memory_max_mb=1,
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                            score_meta_data=[
                                NodeScoreMeta(
                                    node_id="node_id_example",
                                    norm_score=3.14,
                                    scores={
                                        "key": 3.14,
                                    },
                                ),
                            ],
                            scores={
                                "key": 3.14,
                            },
                        ),
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        next_allocation="next_allocation_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        previous_allocation="previous_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        resources=Resources(
                            cpu=1,
                            cores=1,
                            devices=[
                                RequestedDevice(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    count=0,
                                    name="name_example",
                                ),
                            ],
                            disk_mb=1,
                            iops=1,
                            memory_mb=1,
                            memory_max_mb=1,
                            networks=[
                                NetworkResource(
                                    cidr="cidr_example",
                                    dns=DNSConfig(
                                        options=[
                                            "options_example",
                                        ],
                                        searches=[
                                            "searches_example",
                                        ],
                                        servers=[
                                            "servers_example",
                                        ],
                                    ),
                                    device="device_example",
                                    dynamic_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                    hostname="hostname_example",
                                    ip="ip_example",
                                    m_bits=1,
                                    mode="mode_example",
                                    reserved_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                ),
                            ],
                        ),
                        services={
                            "key": "key_example",
                        },
                        task_group="task_group_example",
                        task_resources={
                            "key": Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                        },
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                },
            ),
        ],
    ) # CSIVolumeRegisterRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.post_volume(csi_volume_register_request)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->post_volume: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.post_volume(csi_volume_register_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->post_volume: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **csi_volume_register_request** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_volume_registration**
> post_volume_registration(volume_id, csi_volume_register_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import volumes_api
from nomad_client.model.csi_volume_register_request import CSIVolumeRegisterRequest
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = volumes_api.VolumesApi(api_client)
    volume_id = "volumeId_example" # str | Volume unique identifier.
    csi_volume_register_request = CSIVolumeRegisterRequest(
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
        volumes=[
            CSIVolume(
                access_mode="access_mode_example",
                allocations=[
                    AllocationListStub(
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job_id="job_id_example",
                        job_type="job_type_example",
                        job_version=0,
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        task_group="task_group_example",
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                ],
                attachment_mode="attachment_mode_example",
                capacity=1,
                clone_id="clone_id_example",
                context={
                    "key": "key_example",
                },
                controller_required=True,
                controllers_expected=1,
                controllers_healthy=1,
                create_index=0,
                external_id="external_id_example",
                id="id_example",
                modify_index=0,
                mount_options=CSIMountOptions(
                    fs_type="fs_type_example",
                    mount_flags=[
                        "mount_flags_example",
                    ],
                ),
                name="name_example",
                namespace="namespace_example",
                nodes_expected=1,
                nodes_healthy=1,
                parameters={
                    "key": "key_example",
                },
                plugin_id="plugin_id_example",
                provider="provider_example",
                provider_version="provider_version_example",
                read_allocs={
                    "key": Allocation(
                        alloc_modify_index=0,
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_id="deployment_id_example",
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        desired_transition=DesiredTransition(
                            migrate=True,
                            reschedule=True,
                        ),
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job=Job(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            all_at_once=True,
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            consul_namespace="consul_namespace_example",
                            consul_token="consul_token_example",
                            create_index=0,
                            datacenters=[
                                "datacenters_example",
                            ],
                            dispatch_idempotency_token="dispatch_idempotency_token_example",
                            dispatched=True,
                            id="id_example",
                            job_modify_index=0,
                            meta={
                                "key": "key_example",
                            },
                            migrate=MigrateStrategy(
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                            ),
                            modify_index=0,
                            multiregion=Multiregion(
                                regions=[
                                    MultiregionRegion(
                                        count=1,
                                        datacenters=[
                                            "datacenters_example",
                                        ],
                                        meta={
                                            "key": "key_example",
                                        },
                                        name="name_example",
                                    ),
                                ],
                                strategy=MultiregionStrategy(
                                    max_parallel=1,
                                    on_failure="on_failure_example",
                                ),
                            ),
                            name="name_example",
                            namespace="namespace_example",
                            nomad_token_id="nomad_token_id_example",
                            parameterized_job=ParameterizedJobConfig(
                                meta_optional=[
                                    "meta_optional_example",
                                ],
                                meta_required=[
                                    "meta_required_example",
                                ],
                                payload="payload_example",
                            ),
                            parent_id="parent_id_example",
                            payload='YQ==',
                            periodic=PeriodicConfig(
                                enabled=True,
                                prohibit_overlap=True,
                                spec="spec_example",
                                spec_type="spec_type_example",
                                time_zone="time_zone_example",
                            ),
                            priority=1,
                            region="region_example",
                            reschedule=ReschedulePolicy(
                                attempts=1,
                                delay=1,
                                delay_function="delay_function_example",
                                interval=1,
                                max_delay=1,
                                unlimited=True,
                            ),
                            spreads=[
                                Spread(
                                    attribute="attribute_example",
                                    spread_target=[
                                        SpreadTarget(
                                            percent=0,
                                            value="value_example",
                                        ),
                                    ],
                                    weight=-128,
                                ),
                            ],
                            stable=True,
                            status="status_example",
                            status_description="status_description_example",
                            stop=True,
                            submit_time=1,
                            task_groups=[
                                TaskGroup(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    consul=Consul(
                                        namespace="namespace_example",
                                    ),
                                    count=1,
                                    ephemeral_disk=EphemeralDisk(
                                        migrate=True,
                                        size_mb=1,
                                        sticky=True,
                                    ),
                                    max_client_disconnect=1,
                                    meta={
                                        "key": "key_example",
                                    },
                                    migrate=MigrateStrategy(
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                    ),
                                    name="name_example",
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                    reschedule_policy=ReschedulePolicy(
                                        attempts=1,
                                        delay=1,
                                        delay_function="delay_function_example",
                                        interval=1,
                                        max_delay=1,
                                        unlimited=True,
                                    ),
                                    restart_policy=RestartPolicy(
                                        attempts=1,
                                        delay=1,
                                        interval=1,
                                        mode="mode_example",
                                    ),
                                    scaling=ScalingPolicy(
                                        create_index=0,
                                        enabled=True,
                                        id="id_example",
                                        max=1,
                                        min=1,
                                        modify_index=0,
                                        namespace="namespace_example",
                                        policy={
                                            "key": None,
                                        },
                                        target={
                                            "key": "key_example",
                                        },
                                        type="type_example",
                                    ),
                                    services=[
                                        Service(
                                            address_mode="address_mode_example",
                                            canary_meta={
                                                "key": "key_example",
                                            },
                                            canary_tags=[
                                                "canary_tags_example",
                                            ],
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            checks=[
                                                ServiceCheck(
                                                    address_mode="address_mode_example",
                                                    args=[
                                                        "args_example",
                                                    ],
                                                    body="body_example",
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    command="command_example",
                                                    expose=True,
                                                    failures_before_critical=1,
                                                    grpc_service="grpc_service_example",
                                                    grpc_use_tls=True,
                                                    header={
                                                        "key": [
                                                            "key_example",
                                                        ],
                                                    },
                                                    id="id_example",
                                                    initial_status="initial_status_example",
                                                    interval=1,
                                                    method="method_example",
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    path="path_example",
                                                    port_label="port_label_example",
                                                    protocol="protocol_example",
                                                    success_before_passing=1,
                                                    tls_skip_verify=True,
                                                    task_name="task_name_example",
                                                    timeout=1,
                                                    type="type_example",
                                                ),
                                            ],
                                            connect=ConsulConnect(
                                                gateway=ConsulGateway(
                                                    ingress=ConsulIngressConfigEntry(
                                                        listeners=[
                                                            ConsulIngressListener(
                                                                port=1,
                                                                protocol="protocol_example",
                                                                services=[
                                                                    ConsulIngressService(
                                                                        hosts=[
                                                                            "hosts_example",
                                                                        ],
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                        tls=ConsulGatewayTLSConfig(
                                                            enabled=True,
                                                        ),
                                                    ),
                                                    mesh=None,
                                                    proxy=ConsulGatewayProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        connect_timeout=1,
                                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                        envoy_gateway_bind_addresses={
                                                            "key": ConsulGatewayBindAddress(
                                                                address="address_example",
                                                                name="name_example",
                                                                port=1,
                                                            ),
                                                        },
                                                        envoy_gateway_bind_tagged_addresses=True,
                                                        envoy_gateway_no_default_bind=True,
                                                    ),
                                                    terminating=ConsulTerminatingConfigEntry(
                                                        services=[
                                                            ConsulLinkedService(
                                                                ca_file="ca_file_example",
                                                                cert_file="cert_file_example",
                                                                key_file="key_file_example",
                                                                name="name_example",
                                                                sni="sni_example",
                                                            ),
                                                        ],
                                                    ),
                                                ),
                                                native=True,
                                                sidecar_service=ConsulSidecarService(
                                                    disable_default_tcp_check=True,
                                                    port="port_example",
                                                    proxy=ConsulProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        expose_config=ConsulExposeConfig(
                                                            path=[
                                                                ConsulExposePath(
                                                                    listener_port="listener_port_example",
                                                                    local_path_port=1,
                                                                    path="path_example",
                                                                    protocol="protocol_example",
                                                                ),
                                                            ],
                                                        ),
                                                        local_service_address="local_service_address_example",
                                                        local_service_port=1,
                                                        upstreams=[
                                                            ConsulUpstream(
                                                                datacenter="datacenter_example",
                                                                destination_name="destination_name_example",
                                                                local_bind_address="local_bind_address_example",
                                                                local_bind_port=1,
                                                                mesh_gateway=ConsulMeshGateway(
                                                                    mode="mode_example",
                                                                ),
                                                            ),
                                                        ],
                                                    ),
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                ),
                                                sidecar_task=SidecarTask(
                                                    config={
                                                        "key": None,
                                                    },
                                                    driver="driver_example",
                                                    env={
                                                        "key": "key_example",
                                                    },
                                                    kill_signal="kill_signal_example",
                                                    kill_timeout=1,
                                                    log_config=LogConfig(
                                                        max_file_size_mb=1,
                                                        max_files=1,
                                                    ),
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    resources=Resources(
                                                        cpu=1,
                                                        cores=1,
                                                        devices=[
                                                            RequestedDevice(
                                                                affinities=[
                                                                    Affinity(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                        weight=-128,
                                                                    ),
                                                                ],
                                                                constraints=[
                                                                    Constraint(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                    ),
                                                                ],
                                                                count=0,
                                                                name="name_example",
                                                            ),
                                                        ],
                                                        disk_mb=1,
                                                        iops=1,
                                                        memory_mb=1,
                                                        memory_max_mb=1,
                                                        networks=[
                                                            NetworkResource(
                                                                cidr="cidr_example",
                                                                dns=DNSConfig(
                                                                    options=[
                                                                        "options_example",
                                                                    ],
                                                                    searches=[
                                                                        "searches_example",
                                                                    ],
                                                                    servers=[
                                                                        "servers_example",
                                                                    ],
                                                                ),
                                                                device="device_example",
                                                                dynamic_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                                hostname="hostname_example",
                                                                ip="ip_example",
                                                                m_bits=1,
                                                                mode="mode_example",
                                                                reserved_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                    ),
                                                    shutdown_delay=1,
                                                    user="user_example",
                                                ),
                                            ),
                                            enable_tag_override=True,
                                            id="id_example",
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            on_update="on_update_example",
                                            port_label="port_label_example",
                                            provider="provider_example",
                                            tags=[
                                                "tags_example",
                                            ],
                                            task_name="task_name_example",
                                        ),
                                    ],
                                    shutdown_delay=1,
                                    spreads=[
                                        Spread(
                                            attribute="attribute_example",
                                            spread_target=[
                                                SpreadTarget(
                                                    percent=0,
                                                    value="value_example",
                                                ),
                                            ],
                                            weight=-128,
                                        ),
                                    ],
                                    stop_after_client_disconnect=1,
                                    tasks=[
                                        Task(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            artifacts=[
                                                TaskArtifact(
                                                    getter_headers={
                                                        "key": "key_example",
                                                    },
                                                    getter_mode="getter_mode_example",
                                                    getter_options={
                                                        "key": "key_example",
                                                    },
                                                    getter_source="getter_source_example",
                                                    relative_dest="relative_dest_example",
                                                ),
                                            ],
                                            csi_plugin_config=TaskCSIPluginConfig(
                                                id="id_example",
                                                mount_dir="mount_dir_example",
                                                type="type_example",
                                            ),
                                            config={
                                                "key": None,
                                            },
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            dispatch_payload=DispatchPayloadConfig(
                                                file="file_example",
                                            ),
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            kind="kind_example",
                                            leader=True,
                                            lifecycle=TaskLifecycle(
                                                hook="hook_example",
                                                sidecar=True,
                                            ),
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            restart_policy=RestartPolicy(
                                                attempts=1,
                                                delay=1,
                                                interval=1,
                                                mode="mode_example",
                                            ),
                                            scaling_policies=[
                                                ScalingPolicy(
                                                    create_index=0,
                                                    enabled=True,
                                                    id="id_example",
                                                    max=1,
                                                    min=1,
                                                    modify_index=0,
                                                    namespace="namespace_example",
                                                    policy={
                                                        "key": None,
                                                    },
                                                    target={
                                                        "key": "key_example",
                                                    },
                                                    type="type_example",
                                                ),
                                            ],
                                            services=[
                                                Service(
                                                    address_mode="address_mode_example",
                                                    canary_meta={
                                                        "key": "key_example",
                                                    },
                                                    canary_tags=[
                                                        "canary_tags_example",
                                                    ],
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    checks=[
                                                        ServiceCheck(
                                                            address_mode="address_mode_example",
                                                            args=[
                                                                "args_example",
                                                            ],
                                                            body="body_example",
                                                            check_restart=CheckRestart(
                                                                grace=1,
                                                                ignore_warnings=True,
                                                                limit=1,
                                                            ),
                                                            command="command_example",
                                                            expose=True,
                                                            failures_before_critical=1,
                                                            grpc_service="grpc_service_example",
                                                            grpc_use_tls=True,
                                                            header={
                                                                "key": [
                                                                    "key_example",
                                                                ],
                                                            },
                                                            id="id_example",
                                                            initial_status="initial_status_example",
                                                            interval=1,
                                                            method="method_example",
                                                            name="name_example",
                                                            on_update="on_update_example",
                                                            path="path_example",
                                                            port_label="port_label_example",
                                                            protocol="protocol_example",
                                                            success_before_passing=1,
                                                            tls_skip_verify=True,
                                                            task_name="task_name_example",
                                                            timeout=1,
                                                            type="type_example",
                                                        ),
                                                    ],
                                                    connect=ConsulConnect(
                                                        gateway=ConsulGateway(
                                                            ingress=ConsulIngressConfigEntry(
                                                                listeners=[
                                                                    ConsulIngressListener(
                                                                        port=1,
                                                                        protocol="protocol_example",
                                                                        services=[
                                                                            ConsulIngressService(
                                                                                hosts=[
                                                                                    "hosts_example",
                                                                                ],
                                                                                name="name_example",
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                                tls=ConsulGatewayTLSConfig(
                                                                    enabled=True,
                                                                ),
                                                            ),
                                                            mesh=None,
                                                            proxy=ConsulGatewayProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                connect_timeout=1,
                                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                                envoy_gateway_bind_addresses={
                                                                    "key": ConsulGatewayBindAddress(
                                                                        address="address_example",
                                                                        name="name_example",
                                                                        port=1,
                                                                    ),
                                                                },
                                                                envoy_gateway_bind_tagged_addresses=True,
                                                                envoy_gateway_no_default_bind=True,
                                                            ),
                                                            terminating=ConsulTerminatingConfigEntry(
                                                                services=[
                                                                    ConsulLinkedService(
                                                                        ca_file="ca_file_example",
                                                                        cert_file="cert_file_example",
                                                                        key_file="key_file_example",
                                                                        name="name_example",
                                                                        sni="sni_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ),
                                                        native=True,
                                                        sidecar_service=ConsulSidecarService(
                                                            disable_default_tcp_check=True,
                                                            port="port_example",
                                                            proxy=ConsulProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                expose_config=ConsulExposeConfig(
                                                                    path=[
                                                                        ConsulExposePath(
                                                                            listener_port="listener_port_example",
                                                                            local_path_port=1,
                                                                            path="path_example",
                                                                            protocol="protocol_example",
                                                                        ),
                                                                    ],
                                                                ),
                                                                local_service_address="local_service_address_example",
                                                                local_service_port=1,
                                                                upstreams=[
                                                                    ConsulUpstream(
                                                                        datacenter="datacenter_example",
                                                                        destination_name="destination_name_example",
                                                                        local_bind_address="local_bind_address_example",
                                                                        local_bind_port=1,
                                                                        mesh_gateway=ConsulMeshGateway(
                                                                            mode="mode_example",
                                                                        ),
                                                                    ),
                                                                ],
                                                            ),
                                                            tags=[
                                                                "tags_example",
                                                            ],
                                                        ),
                                                        sidecar_task=SidecarTask(
                                                            config={
                                                                "key": None,
                                                            },
                                                            driver="driver_example",
                                                            env={
                                                                "key": "key_example",
                                                            },
                                                            kill_signal="kill_signal_example",
                                                            kill_timeout=1,
                                                            log_config=LogConfig(
                                                                max_file_size_mb=1,
                                                                max_files=1,
                                                            ),
                                                            meta={
                                                                "key": "key_example",
                                                            },
                                                            name="name_example",
                                                            resources=Resources(
                                                                cpu=1,
                                                                cores=1,
                                                                devices=[
                                                                    RequestedDevice(
                                                                        affinities=[
                                                                            Affinity(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                                weight=-128,
                                                                            ),
                                                                        ],
                                                                        constraints=[
                                                                            Constraint(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                            ),
                                                                        ],
                                                                        count=0,
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                                disk_mb=1,
                                                                iops=1,
                                                                memory_mb=1,
                                                                memory_max_mb=1,
                                                                networks=[
                                                                    NetworkResource(
                                                                        cidr="cidr_example",
                                                                        dns=DNSConfig(
                                                                            options=[
                                                                                "options_example",
                                                                            ],
                                                                            searches=[
                                                                                "searches_example",
                                                                            ],
                                                                            servers=[
                                                                                "servers_example",
                                                                            ],
                                                                        ),
                                                                        device="device_example",
                                                                        dynamic_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                        hostname="hostname_example",
                                                                        ip="ip_example",
                                                                        m_bits=1,
                                                                        mode="mode_example",
                                                                        reserved_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                            ),
                                                            shutdown_delay=1,
                                                            user="user_example",
                                                        ),
                                                    ),
                                                    enable_tag_override=True,
                                                    id="id_example",
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    port_label="port_label_example",
                                                    provider="provider_example",
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                    task_name="task_name_example",
                                                ),
                                            ],
                                            shutdown_delay=1,
                                            templates=[
                                                Template(
                                                    change_mode="change_mode_example",
                                                    change_signal="change_signal_example",
                                                    dest_path="dest_path_example",
                                                    embedded_tmpl="embedded_tmpl_example",
                                                    envvars=True,
                                                    left_delim="left_delim_example",
                                                    perms="perms_example",
                                                    right_delim="right_delim_example",
                                                    source_path="source_path_example",
                                                    splay=1,
                                                    vault_grace=1,
                                                    wait=WaitConfig(
                                                        max=1,
                                                        min=1,
                                                    ),
                                                ),
                                            ],
                                            user="user_example",
                                            vault=Vault(
                                                change_mode="change_mode_example",
                                                change_signal="change_signal_example",
                                                entity_alias="entity_alias_example",
                                                env=True,
                                                namespace="namespace_example",
                                                policies=[
                                                    "policies_example",
                                                ],
                                            ),
                                            volume_mounts=[
                                                VolumeMount(
                                                    destination="destination_example",
                                                    propagation_mode="propagation_mode_example",
                                                    read_only=True,
                                                    volume="volume_example",
                                                ),
                                            ],
                                        ),
                                    ],
                                    update=UpdateStrategy(
                                        auto_promote=True,
                                        auto_revert=True,
                                        canary=1,
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                        progress_deadline=1,
                                        stagger=1,
                                    ),
                                    volumes={
                                        "key": VolumeRequest(
                                            access_mode="access_mode_example",
                                            attachment_mode="attachment_mode_example",
                                            mount_options=CSIMountOptions(
                                                fs_type="fs_type_example",
                                                mount_flags=[
                                                    "mount_flags_example",
                                                ],
                                            ),
                                            name="name_example",
                                            per_alloc=True,
                                            read_only=True,
                                            source="source_example",
                                            type="type_example",
                                        ),
                                    },
                                ),
                            ],
                            type="type_example",
                            update=UpdateStrategy(
                                auto_promote=True,
                                auto_revert=True,
                                canary=1,
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                                progress_deadline=1,
                                stagger=1,
                            ),
                            vault_namespace="vault_namespace_example",
                            vault_token="vault_token_example",
                            version=0,
                        ),
                        job_id="job_id_example",
                        metrics=AllocationMetric(
                            allocation_time=1,
                            class_exhausted={
                                "key": 1,
                            },
                            class_filtered={
                                "key": 1,
                            },
                            coalesced_failures=1,
                            constraint_filtered={
                                "key": 1,
                            },
                            dimension_exhausted={
                                "key": 1,
                            },
                            nodes_available={
                                "key": 1,
                            },
                            nodes_evaluated=1,
                            nodes_exhausted=1,
                            nodes_filtered=1,
                            quota_exhausted=[
                                "quota_exhausted_example",
                            ],
                            resources_exhausted={
                                "key": Resources(
                                    cpu=1,
                                    cores=1,
                                    devices=[
                                        RequestedDevice(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            count=0,
                                            name="name_example",
                                        ),
                                    ],
                                    disk_mb=1,
                                    iops=1,
                                    memory_mb=1,
                                    memory_max_mb=1,
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                            score_meta_data=[
                                NodeScoreMeta(
                                    node_id="node_id_example",
                                    norm_score=3.14,
                                    scores={
                                        "key": 3.14,
                                    },
                                ),
                            ],
                            scores={
                                "key": 3.14,
                            },
                        ),
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        next_allocation="next_allocation_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        previous_allocation="previous_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        resources=Resources(
                            cpu=1,
                            cores=1,
                            devices=[
                                RequestedDevice(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    count=0,
                                    name="name_example",
                                ),
                            ],
                            disk_mb=1,
                            iops=1,
                            memory_mb=1,
                            memory_max_mb=1,
                            networks=[
                                NetworkResource(
                                    cidr="cidr_example",
                                    dns=DNSConfig(
                                        options=[
                                            "options_example",
                                        ],
                                        searches=[
                                            "searches_example",
                                        ],
                                        servers=[
                                            "servers_example",
                                        ],
                                    ),
                                    device="device_example",
                                    dynamic_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                    hostname="hostname_example",
                                    ip="ip_example",
                                    m_bits=1,
                                    mode="mode_example",
                                    reserved_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                ),
                            ],
                        ),
                        services={
                            "key": "key_example",
                        },
                        task_group="task_group_example",
                        task_resources={
                            "key": Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                        },
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                },
                requested_capabilities=[
                    CSIVolumeCapability(
                        access_mode="access_mode_example",
                        attachment_mode="attachment_mode_example",
                    ),
                ],
                requested_capacity_max=1,
                requested_capacity_min=1,
                requested_topologies=CSITopologyRequest(
                    preferred=[
                        CSITopology(
                            segments={
                                "key": "key_example",
                            },
                        ),
                    ],
                    required=[
                        CSITopology(
                            segments={
                                "key": "key_example",
                            },
                        ),
                    ],
                ),
                resource_exhausted=dateutil_parser('1970-01-01T00:00:00.00Z'),
                schedulable=True,
                secrets=CSISecrets(
                    key="key_example",
                ),
                snapshot_id="snapshot_id_example",
                topologies=[
                    CSITopology(
                        segments={
                            "key": "key_example",
                        },
                    ),
                ],
                write_allocs={
                    "key": Allocation(
                        alloc_modify_index=0,
                        allocated_resources=AllocatedResources(
                            shared=AllocatedSharedResources(
                                disk_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                                ports=[
                                    PortMapping(
                                        host_ip="host_ip_example",
                                        label="label_example",
                                        to=1,
                                        value=1,
                                    ),
                                ],
                            ),
                            tasks={
                                "key": AllocatedTaskResources(
                                    cpu=AllocatedCpuResources(
                                        cpu_shares=1,
                                    ),
                                    devices=[
                                        AllocatedDeviceResource(
                                            device_ids=[
                                                "device_ids_example",
                                            ],
                                            name="name_example",
                                            type="type_example",
                                            vendor="vendor_example",
                                        ),
                                    ],
                                    memory=AllocatedMemoryResources(
                                        memory_mb=1,
                                        memory_max_mb=1,
                                    ),
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                        ),
                        client_description="client_description_example",
                        client_status="client_status_example",
                        create_index=0,
                        create_time=1,
                        deployment_id="deployment_id_example",
                        deployment_status=AllocDeploymentStatus(
                            canary=True,
                            healthy=True,
                            modify_index=0,
                            timestamp=dateutil_parser('1970-01-01T00:00:00.00Z'),
                        ),
                        desired_description="desired_description_example",
                        desired_status="desired_status_example",
                        desired_transition=DesiredTransition(
                            migrate=True,
                            reschedule=True,
                        ),
                        eval_id="eval_id_example",
                        followup_eval_id="followup_eval_id_example",
                        id="id_example",
                        job=Job(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            all_at_once=True,
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            consul_namespace="consul_namespace_example",
                            consul_token="consul_token_example",
                            create_index=0,
                            datacenters=[
                                "datacenters_example",
                            ],
                            dispatch_idempotency_token="dispatch_idempotency_token_example",
                            dispatched=True,
                            id="id_example",
                            job_modify_index=0,
                            meta={
                                "key": "key_example",
                            },
                            migrate=MigrateStrategy(
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                            ),
                            modify_index=0,
                            multiregion=Multiregion(
                                regions=[
                                    MultiregionRegion(
                                        count=1,
                                        datacenters=[
                                            "datacenters_example",
                                        ],
                                        meta={
                                            "key": "key_example",
                                        },
                                        name="name_example",
                                    ),
                                ],
                                strategy=MultiregionStrategy(
                                    max_parallel=1,
                                    on_failure="on_failure_example",
                                ),
                            ),
                            name="name_example",
                            namespace="namespace_example",
                            nomad_token_id="nomad_token_id_example",
                            parameterized_job=ParameterizedJobConfig(
                                meta_optional=[
                                    "meta_optional_example",
                                ],
                                meta_required=[
                                    "meta_required_example",
                                ],
                                payload="payload_example",
                            ),
                            parent_id="parent_id_example",
                            payload='YQ==',
                            periodic=PeriodicConfig(
                                enabled=True,
                                prohibit_overlap=True,
                                spec="spec_example",
                                spec_type="spec_type_example",
                                time_zone="time_zone_example",
                            ),
                            priority=1,
                            region="region_example",
                            reschedule=ReschedulePolicy(
                                attempts=1,
                                delay=1,
                                delay_function="delay_function_example",
                                interval=1,
                                max_delay=1,
                                unlimited=True,
                            ),
                            spreads=[
                                Spread(
                                    attribute="attribute_example",
                                    spread_target=[
                                        SpreadTarget(
                                            percent=0,
                                            value="value_example",
                                        ),
                                    ],
                                    weight=-128,
                                ),
                            ],
                            stable=True,
                            status="status_example",
                            status_description="status_description_example",
                            stop=True,
                            submit_time=1,
                            task_groups=[
                                TaskGroup(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    consul=Consul(
                                        namespace="namespace_example",
                                    ),
                                    count=1,
                                    ephemeral_disk=EphemeralDisk(
                                        migrate=True,
                                        size_mb=1,
                                        sticky=True,
                                    ),
                                    max_client_disconnect=1,
                                    meta={
                                        "key": "key_example",
                                    },
                                    migrate=MigrateStrategy(
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                    ),
                                    name="name_example",
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                    reschedule_policy=ReschedulePolicy(
                                        attempts=1,
                                        delay=1,
                                        delay_function="delay_function_example",
                                        interval=1,
                                        max_delay=1,
                                        unlimited=True,
                                    ),
                                    restart_policy=RestartPolicy(
                                        attempts=1,
                                        delay=1,
                                        interval=1,
                                        mode="mode_example",
                                    ),
                                    scaling=ScalingPolicy(
                                        create_index=0,
                                        enabled=True,
                                        id="id_example",
                                        max=1,
                                        min=1,
                                        modify_index=0,
                                        namespace="namespace_example",
                                        policy={
                                            "key": None,
                                        },
                                        target={
                                            "key": "key_example",
                                        },
                                        type="type_example",
                                    ),
                                    services=[
                                        Service(
                                            address_mode="address_mode_example",
                                            canary_meta={
                                                "key": "key_example",
                                            },
                                            canary_tags=[
                                                "canary_tags_example",
                                            ],
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            checks=[
                                                ServiceCheck(
                                                    address_mode="address_mode_example",
                                                    args=[
                                                        "args_example",
                                                    ],
                                                    body="body_example",
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    command="command_example",
                                                    expose=True,
                                                    failures_before_critical=1,
                                                    grpc_service="grpc_service_example",
                                                    grpc_use_tls=True,
                                                    header={
                                                        "key": [
                                                            "key_example",
                                                        ],
                                                    },
                                                    id="id_example",
                                                    initial_status="initial_status_example",
                                                    interval=1,
                                                    method="method_example",
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    path="path_example",
                                                    port_label="port_label_example",
                                                    protocol="protocol_example",
                                                    success_before_passing=1,
                                                    tls_skip_verify=True,
                                                    task_name="task_name_example",
                                                    timeout=1,
                                                    type="type_example",
                                                ),
                                            ],
                                            connect=ConsulConnect(
                                                gateway=ConsulGateway(
                                                    ingress=ConsulIngressConfigEntry(
                                                        listeners=[
                                                            ConsulIngressListener(
                                                                port=1,
                                                                protocol="protocol_example",
                                                                services=[
                                                                    ConsulIngressService(
                                                                        hosts=[
                                                                            "hosts_example",
                                                                        ],
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                        tls=ConsulGatewayTLSConfig(
                                                            enabled=True,
                                                        ),
                                                    ),
                                                    mesh=None,
                                                    proxy=ConsulGatewayProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        connect_timeout=1,
                                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                        envoy_gateway_bind_addresses={
                                                            "key": ConsulGatewayBindAddress(
                                                                address="address_example",
                                                                name="name_example",
                                                                port=1,
                                                            ),
                                                        },
                                                        envoy_gateway_bind_tagged_addresses=True,
                                                        envoy_gateway_no_default_bind=True,
                                                    ),
                                                    terminating=ConsulTerminatingConfigEntry(
                                                        services=[
                                                            ConsulLinkedService(
                                                                ca_file="ca_file_example",
                                                                cert_file="cert_file_example",
                                                                key_file="key_file_example",
                                                                name="name_example",
                                                                sni="sni_example",
                                                            ),
                                                        ],
                                                    ),
                                                ),
                                                native=True,
                                                sidecar_service=ConsulSidecarService(
                                                    disable_default_tcp_check=True,
                                                    port="port_example",
                                                    proxy=ConsulProxy(
                                                        config={
                                                            "key": None,
                                                        },
                                                        expose_config=ConsulExposeConfig(
                                                            path=[
                                                                ConsulExposePath(
                                                                    listener_port="listener_port_example",
                                                                    local_path_port=1,
                                                                    path="path_example",
                                                                    protocol="protocol_example",
                                                                ),
                                                            ],
                                                        ),
                                                        local_service_address="local_service_address_example",
                                                        local_service_port=1,
                                                        upstreams=[
                                                            ConsulUpstream(
                                                                datacenter="datacenter_example",
                                                                destination_name="destination_name_example",
                                                                local_bind_address="local_bind_address_example",
                                                                local_bind_port=1,
                                                                mesh_gateway=ConsulMeshGateway(
                                                                    mode="mode_example",
                                                                ),
                                                            ),
                                                        ],
                                                    ),
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                ),
                                                sidecar_task=SidecarTask(
                                                    config={
                                                        "key": None,
                                                    },
                                                    driver="driver_example",
                                                    env={
                                                        "key": "key_example",
                                                    },
                                                    kill_signal="kill_signal_example",
                                                    kill_timeout=1,
                                                    log_config=LogConfig(
                                                        max_file_size_mb=1,
                                                        max_files=1,
                                                    ),
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    resources=Resources(
                                                        cpu=1,
                                                        cores=1,
                                                        devices=[
                                                            RequestedDevice(
                                                                affinities=[
                                                                    Affinity(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                        weight=-128,
                                                                    ),
                                                                ],
                                                                constraints=[
                                                                    Constraint(
                                                                        l_target="l_target_example",
                                                                        operand="operand_example",
                                                                        r_target="r_target_example",
                                                                    ),
                                                                ],
                                                                count=0,
                                                                name="name_example",
                                                            ),
                                                        ],
                                                        disk_mb=1,
                                                        iops=1,
                                                        memory_mb=1,
                                                        memory_max_mb=1,
                                                        networks=[
                                                            NetworkResource(
                                                                cidr="cidr_example",
                                                                dns=DNSConfig(
                                                                    options=[
                                                                        "options_example",
                                                                    ],
                                                                    searches=[
                                                                        "searches_example",
                                                                    ],
                                                                    servers=[
                                                                        "servers_example",
                                                                    ],
                                                                ),
                                                                device="device_example",
                                                                dynamic_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                                hostname="hostname_example",
                                                                ip="ip_example",
                                                                m_bits=1,
                                                                mode="mode_example",
                                                                reserved_ports=[
                                                                    Port(
                                                                        host_network="host_network_example",
                                                                        label="label_example",
                                                                        to=1,
                                                                        value=1,
                                                                    ),
                                                                ],
                                                            ),
                                                        ],
                                                    ),
                                                    shutdown_delay=1,
                                                    user="user_example",
                                                ),
                                            ),
                                            enable_tag_override=True,
                                            id="id_example",
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            on_update="on_update_example",
                                            port_label="port_label_example",
                                            provider="provider_example",
                                            tags=[
                                                "tags_example",
                                            ],
                                            task_name="task_name_example",
                                        ),
                                    ],
                                    shutdown_delay=1,
                                    spreads=[
                                        Spread(
                                            attribute="attribute_example",
                                            spread_target=[
                                                SpreadTarget(
                                                    percent=0,
                                                    value="value_example",
                                                ),
                                            ],
                                            weight=-128,
                                        ),
                                    ],
                                    stop_after_client_disconnect=1,
                                    tasks=[
                                        Task(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            artifacts=[
                                                TaskArtifact(
                                                    getter_headers={
                                                        "key": "key_example",
                                                    },
                                                    getter_mode="getter_mode_example",
                                                    getter_options={
                                                        "key": "key_example",
                                                    },
                                                    getter_source="getter_source_example",
                                                    relative_dest="relative_dest_example",
                                                ),
                                            ],
                                            csi_plugin_config=TaskCSIPluginConfig(
                                                id="id_example",
                                                mount_dir="mount_dir_example",
                                                type="type_example",
                                            ),
                                            config={
                                                "key": None,
                                            },
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            dispatch_payload=DispatchPayloadConfig(
                                                file="file_example",
                                            ),
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            kind="kind_example",
                                            leader=True,
                                            lifecycle=TaskLifecycle(
                                                hook="hook_example",
                                                sidecar=True,
                                            ),
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            restart_policy=RestartPolicy(
                                                attempts=1,
                                                delay=1,
                                                interval=1,
                                                mode="mode_example",
                                            ),
                                            scaling_policies=[
                                                ScalingPolicy(
                                                    create_index=0,
                                                    enabled=True,
                                                    id="id_example",
                                                    max=1,
                                                    min=1,
                                                    modify_index=0,
                                                    namespace="namespace_example",
                                                    policy={
                                                        "key": None,
                                                    },
                                                    target={
                                                        "key": "key_example",
                                                    },
                                                    type="type_example",
                                                ),
                                            ],
                                            services=[
                                                Service(
                                                    address_mode="address_mode_example",
                                                    canary_meta={
                                                        "key": "key_example",
                                                    },
                                                    canary_tags=[
                                                        "canary_tags_example",
                                                    ],
                                                    check_restart=CheckRestart(
                                                        grace=1,
                                                        ignore_warnings=True,
                                                        limit=1,
                                                    ),
                                                    checks=[
                                                        ServiceCheck(
                                                            address_mode="address_mode_example",
                                                            args=[
                                                                "args_example",
                                                            ],
                                                            body="body_example",
                                                            check_restart=CheckRestart(
                                                                grace=1,
                                                                ignore_warnings=True,
                                                                limit=1,
                                                            ),
                                                            command="command_example",
                                                            expose=True,
                                                            failures_before_critical=1,
                                                            grpc_service="grpc_service_example",
                                                            grpc_use_tls=True,
                                                            header={
                                                                "key": [
                                                                    "key_example",
                                                                ],
                                                            },
                                                            id="id_example",
                                                            initial_status="initial_status_example",
                                                            interval=1,
                                                            method="method_example",
                                                            name="name_example",
                                                            on_update="on_update_example",
                                                            path="path_example",
                                                            port_label="port_label_example",
                                                            protocol="protocol_example",
                                                            success_before_passing=1,
                                                            tls_skip_verify=True,
                                                            task_name="task_name_example",
                                                            timeout=1,
                                                            type="type_example",
                                                        ),
                                                    ],
                                                    connect=ConsulConnect(
                                                        gateway=ConsulGateway(
                                                            ingress=ConsulIngressConfigEntry(
                                                                listeners=[
                                                                    ConsulIngressListener(
                                                                        port=1,
                                                                        protocol="protocol_example",
                                                                        services=[
                                                                            ConsulIngressService(
                                                                                hosts=[
                                                                                    "hosts_example",
                                                                                ],
                                                                                name="name_example",
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                                tls=ConsulGatewayTLSConfig(
                                                                    enabled=True,
                                                                ),
                                                            ),
                                                            mesh=None,
                                                            proxy=ConsulGatewayProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                connect_timeout=1,
                                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                                envoy_gateway_bind_addresses={
                                                                    "key": ConsulGatewayBindAddress(
                                                                        address="address_example",
                                                                        name="name_example",
                                                                        port=1,
                                                                    ),
                                                                },
                                                                envoy_gateway_bind_tagged_addresses=True,
                                                                envoy_gateway_no_default_bind=True,
                                                            ),
                                                            terminating=ConsulTerminatingConfigEntry(
                                                                services=[
                                                                    ConsulLinkedService(
                                                                        ca_file="ca_file_example",
                                                                        cert_file="cert_file_example",
                                                                        key_file="key_file_example",
                                                                        name="name_example",
                                                                        sni="sni_example",
                                                                    ),
                                                                ],
                                                            ),
                                                        ),
                                                        native=True,
                                                        sidecar_service=ConsulSidecarService(
                                                            disable_default_tcp_check=True,
                                                            port="port_example",
                                                            proxy=ConsulProxy(
                                                                config={
                                                                    "key": None,
                                                                },
                                                                expose_config=ConsulExposeConfig(
                                                                    path=[
                                                                        ConsulExposePath(
                                                                            listener_port="listener_port_example",
                                                                            local_path_port=1,
                                                                            path="path_example",
                                                                            protocol="protocol_example",
                                                                        ),
                                                                    ],
                                                                ),
                                                                local_service_address="local_service_address_example",
                                                                local_service_port=1,
                                                                upstreams=[
                                                                    ConsulUpstream(
                                                                        datacenter="datacenter_example",
                                                                        destination_name="destination_name_example",
                                                                        local_bind_address="local_bind_address_example",
                                                                        local_bind_port=1,
                                                                        mesh_gateway=ConsulMeshGateway(
                                                                            mode="mode_example",
                                                                        ),
                                                                    ),
                                                                ],
                                                            ),
                                                            tags=[
                                                                "tags_example",
                                                            ],
                                                        ),
                                                        sidecar_task=SidecarTask(
                                                            config={
                                                                "key": None,
                                                            },
                                                            driver="driver_example",
                                                            env={
                                                                "key": "key_example",
                                                            },
                                                            kill_signal="kill_signal_example",
                                                            kill_timeout=1,
                                                            log_config=LogConfig(
                                                                max_file_size_mb=1,
                                                                max_files=1,
                                                            ),
                                                            meta={
                                                                "key": "key_example",
                                                            },
                                                            name="name_example",
                                                            resources=Resources(
                                                                cpu=1,
                                                                cores=1,
                                                                devices=[
                                                                    RequestedDevice(
                                                                        affinities=[
                                                                            Affinity(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                                weight=-128,
                                                                            ),
                                                                        ],
                                                                        constraints=[
                                                                            Constraint(
                                                                                l_target="l_target_example",
                                                                                operand="operand_example",
                                                                                r_target="r_target_example",
                                                                            ),
                                                                        ],
                                                                        count=0,
                                                                        name="name_example",
                                                                    ),
                                                                ],
                                                                disk_mb=1,
                                                                iops=1,
                                                                memory_mb=1,
                                                                memory_max_mb=1,
                                                                networks=[
                                                                    NetworkResource(
                                                                        cidr="cidr_example",
                                                                        dns=DNSConfig(
                                                                            options=[
                                                                                "options_example",
                                                                            ],
                                                                            searches=[
                                                                                "searches_example",
                                                                            ],
                                                                            servers=[
                                                                                "servers_example",
                                                                            ],
                                                                        ),
                                                                        device="device_example",
                                                                        dynamic_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                        hostname="hostname_example",
                                                                        ip="ip_example",
                                                                        m_bits=1,
                                                                        mode="mode_example",
                                                                        reserved_ports=[
                                                                            Port(
                                                                                host_network="host_network_example",
                                                                                label="label_example",
                                                                                to=1,
                                                                                value=1,
                                                                            ),
                                                                        ],
                                                                    ),
                                                                ],
                                                            ),
                                                            shutdown_delay=1,
                                                            user="user_example",
                                                        ),
                                                    ),
                                                    enable_tag_override=True,
                                                    id="id_example",
                                                    meta={
                                                        "key": "key_example",
                                                    },
                                                    name="name_example",
                                                    on_update="on_update_example",
                                                    port_label="port_label_example",
                                                    provider="provider_example",
                                                    tags=[
                                                        "tags_example",
                                                    ],
                                                    task_name="task_name_example",
                                                ),
                                            ],
                                            shutdown_delay=1,
                                            templates=[
                                                Template(
                                                    change_mode="change_mode_example",
                                                    change_signal="change_signal_example",
                                                    dest_path="dest_path_example",
                                                    embedded_tmpl="embedded_tmpl_example",
                                                    envvars=True,
                                                    left_delim="left_delim_example",
                                                    perms="perms_example",
                                                    right_delim="right_delim_example",
                                                    source_path="source_path_example",
                                                    splay=1,
                                                    vault_grace=1,
                                                    wait=WaitConfig(
                                                        max=1,
                                                        min=1,
                                                    ),
                                                ),
                                            ],
                                            user="user_example",
                                            vault=Vault(
                                                change_mode="change_mode_example",
                                                change_signal="change_signal_example",
                                                entity_alias="entity_alias_example",
                                                env=True,
                                                namespace="namespace_example",
                                                policies=[
                                                    "policies_example",
                                                ],
                                            ),
                                            volume_mounts=[
                                                VolumeMount(
                                                    destination="destination_example",
                                                    propagation_mode="propagation_mode_example",
                                                    read_only=True,
                                                    volume="volume_example",
                                                ),
                                            ],
                                        ),
                                    ],
                                    update=UpdateStrategy(
                                        auto_promote=True,
                                        auto_revert=True,
                                        canary=1,
                                        health_check="health_check_example",
                                        healthy_deadline=1,
                                        max_parallel=1,
                                        min_healthy_time=1,
                                        progress_deadline=1,
                                        stagger=1,
                                    ),
                                    volumes={
                                        "key": VolumeRequest(
                                            access_mode="access_mode_example",
                                            attachment_mode="attachment_mode_example",
                                            mount_options=CSIMountOptions(
                                                fs_type="fs_type_example",
                                                mount_flags=[
                                                    "mount_flags_example",
                                                ],
                                            ),
                                            name="name_example",
                                            per_alloc=True,
                                            read_only=True,
                                            source="source_example",
                                            type="type_example",
                                        ),
                                    },
                                ),
                            ],
                            type="type_example",
                            update=UpdateStrategy(
                                auto_promote=True,
                                auto_revert=True,
                                canary=1,
                                health_check="health_check_example",
                                healthy_deadline=1,
                                max_parallel=1,
                                min_healthy_time=1,
                                progress_deadline=1,
                                stagger=1,
                            ),
                            vault_namespace="vault_namespace_example",
                            vault_token="vault_token_example",
                            version=0,
                        ),
                        job_id="job_id_example",
                        metrics=AllocationMetric(
                            allocation_time=1,
                            class_exhausted={
                                "key": 1,
                            },
                            class_filtered={
                                "key": 1,
                            },
                            coalesced_failures=1,
                            constraint_filtered={
                                "key": 1,
                            },
                            dimension_exhausted={
                                "key": 1,
                            },
                            nodes_available={
                                "key": 1,
                            },
                            nodes_evaluated=1,
                            nodes_exhausted=1,
                            nodes_filtered=1,
                            quota_exhausted=[
                                "quota_exhausted_example",
                            ],
                            resources_exhausted={
                                "key": Resources(
                                    cpu=1,
                                    cores=1,
                                    devices=[
                                        RequestedDevice(
                                            affinities=[
                                                Affinity(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                    weight=-128,
                                                ),
                                            ],
                                            constraints=[
                                                Constraint(
                                                    l_target="l_target_example",
                                                    operand="operand_example",
                                                    r_target="r_target_example",
                                                ),
                                            ],
                                            count=0,
                                            name="name_example",
                                        ),
                                    ],
                                    disk_mb=1,
                                    iops=1,
                                    memory_mb=1,
                                    memory_max_mb=1,
                                    networks=[
                                        NetworkResource(
                                            cidr="cidr_example",
                                            dns=DNSConfig(
                                                options=[
                                                    "options_example",
                                                ],
                                                searches=[
                                                    "searches_example",
                                                ],
                                                servers=[
                                                    "servers_example",
                                                ],
                                            ),
                                            device="device_example",
                                            dynamic_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                            hostname="hostname_example",
                                            ip="ip_example",
                                            m_bits=1,
                                            mode="mode_example",
                                            reserved_ports=[
                                                Port(
                                                    host_network="host_network_example",
                                                    label="label_example",
                                                    to=1,
                                                    value=1,
                                                ),
                                            ],
                                        ),
                                    ],
                                ),
                            },
                            score_meta_data=[
                                NodeScoreMeta(
                                    node_id="node_id_example",
                                    norm_score=3.14,
                                    scores={
                                        "key": 3.14,
                                    },
                                ),
                            ],
                            scores={
                                "key": 3.14,
                            },
                        ),
                        modify_index=0,
                        modify_time=1,
                        name="name_example",
                        namespace="namespace_example",
                        next_allocation="next_allocation_example",
                        node_id="node_id_example",
                        node_name="node_name_example",
                        preempted_allocations=[
                            "preempted_allocations_example",
                        ],
                        preempted_by_allocation="preempted_by_allocation_example",
                        previous_allocation="previous_allocation_example",
                        reschedule_tracker=RescheduleTracker(
                            events=[
                                RescheduleEvent(
                                    prev_alloc_id="prev_alloc_id_example",
                                    prev_node_id="prev_node_id_example",
                                    reschedule_time=1,
                                ),
                            ],
                        ),
                        resources=Resources(
                            cpu=1,
                            cores=1,
                            devices=[
                                RequestedDevice(
                                    affinities=[
                                        Affinity(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                            weight=-128,
                                        ),
                                    ],
                                    constraints=[
                                        Constraint(
                                            l_target="l_target_example",
                                            operand="operand_example",
                                            r_target="r_target_example",
                                        ),
                                    ],
                                    count=0,
                                    name="name_example",
                                ),
                            ],
                            disk_mb=1,
                            iops=1,
                            memory_mb=1,
                            memory_max_mb=1,
                            networks=[
                                NetworkResource(
                                    cidr="cidr_example",
                                    dns=DNSConfig(
                                        options=[
                                            "options_example",
                                        ],
                                        searches=[
                                            "searches_example",
                                        ],
                                        servers=[
                                            "servers_example",
                                        ],
                                    ),
                                    device="device_example",
                                    dynamic_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                    hostname="hostname_example",
                                    ip="ip_example",
                                    m_bits=1,
                                    mode="mode_example",
                                    reserved_ports=[
                                        Port(
                                            host_network="host_network_example",
                                            label="label_example",
                                            to=1,
                                            value=1,
                                        ),
                                    ],
                                ),
                            ],
                        ),
                        services={
                            "key": "key_example",
                        },
                        task_group="task_group_example",
                        task_resources={
                            "key": Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                        },
                        task_states={
                            "key": TaskState(
                                events=[
                                    TaskEvent(
                                        details={
                                            "key": "key_example",
                                        },
                                        disk_limit=1,
                                        disk_size=1,
                                        display_message="display_message_example",
                                        download_error="download_error_example",
                                        driver_error="driver_error_example",
                                        driver_message="driver_message_example",
                                        exit_code=1,
                                        failed_sibling="failed_sibling_example",
                                        fails_task=True,
                                        generic_source="generic_source_example",
                                        kill_error="kill_error_example",
                                        kill_reason="kill_reason_example",
                                        kill_timeout=1,
                                        message="message_example",
                                        restart_reason="restart_reason_example",
                                        setup_error="setup_error_example",
                                        signal=1,
                                        start_delay=1,
                                        task_signal="task_signal_example",
                                        task_signal_reason="task_signal_reason_example",
                                        time=1,
                                        type="type_example",
                                        validation_error="validation_error_example",
                                        vault_error="vault_error_example",
                                    ),
                                ],
                                failed=True,
                                finished_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                last_restart=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                restarts=0,
                                started_at=dateutil_parser('1970-01-01T00:00:00.00Z'),
                                state="state_example",
                                task_handle=TaskHandle(
                                    driver_state='YQ==',
                                    version=1,
                                ),
                            ),
                        },
                    ),
                },
            ),
        ],
    ) # CSIVolumeRegisterRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.post_volume_registration(volume_id, csi_volume_register_request)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->post_volume_registration: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.post_volume_registration(volume_id, csi_volume_register_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except nomad_client.ApiException as e:
        print("Exception when calling VolumesApi->post_volume_registration: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume unique identifier. |
 **csi_volume_register_request** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

