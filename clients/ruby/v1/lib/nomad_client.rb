=begin
#Nomad

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.1.4
Contact: support@hashicorp.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.2.0

=end

# Common files
require 'nomad_client/api_client'
require 'nomad_client/api_error'
require 'nomad_client/version'
require 'nomad_client/configuration'

# Models
require 'nomad_client/models/acl_policy'
require 'nomad_client/models/acl_policy_list_stub'
require 'nomad_client/models/acl_token'
require 'nomad_client/models/acl_token_list_stub'
require 'nomad_client/models/affinity'
require 'nomad_client/models/alloc_deployment_status'
require 'nomad_client/models/allocated_cpu_resources'
require 'nomad_client/models/allocated_device_resource'
require 'nomad_client/models/allocated_memory_resources'
require 'nomad_client/models/allocated_resources'
require 'nomad_client/models/allocated_shared_resources'
require 'nomad_client/models/allocated_task_resources'
require 'nomad_client/models/allocation'
require 'nomad_client/models/allocation_list_stub'
require 'nomad_client/models/allocation_metric'
require 'nomad_client/models/attribute'
require 'nomad_client/models/autopilot_configuration'
require 'nomad_client/models/csi_controller_info'
require 'nomad_client/models/csi_info'
require 'nomad_client/models/csi_mount_options'
require 'nomad_client/models/csi_node_info'
require 'nomad_client/models/csi_plugin'
require 'nomad_client/models/csi_plugin_list_stub'
require 'nomad_client/models/csi_snapshot'
require 'nomad_client/models/csi_snapshot_create_request'
require 'nomad_client/models/csi_snapshot_create_response'
require 'nomad_client/models/csi_snapshot_list_response'
require 'nomad_client/models/csi_topology'
require 'nomad_client/models/csi_volume'
require 'nomad_client/models/csi_volume_capability'
require 'nomad_client/models/csi_volume_create_request'
require 'nomad_client/models/csi_volume_external_stub'
require 'nomad_client/models/csi_volume_list_external_response'
require 'nomad_client/models/csi_volume_list_stub'
require 'nomad_client/models/csi_volume_register_request'
require 'nomad_client/models/check_restart'
require 'nomad_client/models/constraint'
require 'nomad_client/models/consul'
require 'nomad_client/models/consul_connect'
require 'nomad_client/models/consul_expose_config'
require 'nomad_client/models/consul_expose_path'
require 'nomad_client/models/consul_gateway'
require 'nomad_client/models/consul_gateway_bind_address'
require 'nomad_client/models/consul_gateway_proxy'
require 'nomad_client/models/consul_gateway_tls_config'
require 'nomad_client/models/consul_ingress_config_entry'
require 'nomad_client/models/consul_ingress_listener'
require 'nomad_client/models/consul_ingress_service'
require 'nomad_client/models/consul_linked_service'
require 'nomad_client/models/consul_mesh_gateway'
require 'nomad_client/models/consul_proxy'
require 'nomad_client/models/consul_sidecar_service'
require 'nomad_client/models/consul_terminating_config_entry'
require 'nomad_client/models/consul_upstream'
require 'nomad_client/models/dns_config'
require 'nomad_client/models/deployment'
require 'nomad_client/models/deployment_alloc_health_request'
require 'nomad_client/models/deployment_pause_request'
require 'nomad_client/models/deployment_promote_request'
require 'nomad_client/models/deployment_state'
require 'nomad_client/models/deployment_unblock_request'
require 'nomad_client/models/deployment_update_response'
require 'nomad_client/models/desired_transition'
require 'nomad_client/models/desired_updates'
require 'nomad_client/models/dispatch_payload_config'
require 'nomad_client/models/drain_metadata'
require 'nomad_client/models/drain_spec'
require 'nomad_client/models/drain_strategy'
require 'nomad_client/models/driver_info'
require 'nomad_client/models/ephemeral_disk'
require 'nomad_client/models/eval_options'
require 'nomad_client/models/evaluation'
require 'nomad_client/models/field_diff'
require 'nomad_client/models/fuzzy_match'
require 'nomad_client/models/fuzzy_search_request'
require 'nomad_client/models/fuzzy_search_response'
require 'nomad_client/models/gauge_value'
require 'nomad_client/models/host_volume_info'
require 'nomad_client/models/job'
require 'nomad_client/models/job_children_summary'
require 'nomad_client/models/job_deregister_response'
require 'nomad_client/models/job_diff'
require 'nomad_client/models/job_dispatch_request'
require 'nomad_client/models/job_dispatch_response'
require 'nomad_client/models/job_evaluate_request'
require 'nomad_client/models/job_list_stub'
require 'nomad_client/models/job_plan_request'
require 'nomad_client/models/job_plan_response'
require 'nomad_client/models/job_register_request'
require 'nomad_client/models/job_register_response'
require 'nomad_client/models/job_revert_request'
require 'nomad_client/models/job_scale_status_response'
require 'nomad_client/models/job_stability_request'
require 'nomad_client/models/job_stability_response'
require 'nomad_client/models/job_summary'
require 'nomad_client/models/job_validate_request'
require 'nomad_client/models/job_validate_response'
require 'nomad_client/models/job_versions_response'
require 'nomad_client/models/jobs_parse_request'
require 'nomad_client/models/log_config'
require 'nomad_client/models/metrics_summary'
require 'nomad_client/models/migrate_strategy'
require 'nomad_client/models/multiregion'
require 'nomad_client/models/multiregion_region'
require 'nomad_client/models/multiregion_strategy'
require 'nomad_client/models/namespace'
require 'nomad_client/models/network_resource'
require 'nomad_client/models/node'
require 'nomad_client/models/node_cpu_resources'
require 'nomad_client/models/node_device'
require 'nomad_client/models/node_device_locality'
require 'nomad_client/models/node_device_resource'
require 'nomad_client/models/node_disk_resources'
require 'nomad_client/models/node_drain_update_response'
require 'nomad_client/models/node_eligibility_update_response'
require 'nomad_client/models/node_event'
require 'nomad_client/models/node_list_stub'
require 'nomad_client/models/node_memory_resources'
require 'nomad_client/models/node_purge_response'
require 'nomad_client/models/node_reserved_cpu_resources'
require 'nomad_client/models/node_reserved_disk_resources'
require 'nomad_client/models/node_reserved_memory_resources'
require 'nomad_client/models/node_reserved_network_resources'
require 'nomad_client/models/node_reserved_resources'
require 'nomad_client/models/node_resources'
require 'nomad_client/models/node_score_meta'
require 'nomad_client/models/node_update_drain_request'
require 'nomad_client/models/node_update_eligibility_request'
require 'nomad_client/models/object_diff'
require 'nomad_client/models/one_time_token'
require 'nomad_client/models/one_time_token_exchange_request'
require 'nomad_client/models/operator_health_reply'
require 'nomad_client/models/parameterized_job_config'
require 'nomad_client/models/periodic_config'
require 'nomad_client/models/periodic_force_response'
require 'nomad_client/models/plan_annotations'
require 'nomad_client/models/point_value'
require 'nomad_client/models/port'
require 'nomad_client/models/port_mapping'
require 'nomad_client/models/preemption_config'
require 'nomad_client/models/quota_limit'
require 'nomad_client/models/quota_spec'
require 'nomad_client/models/raft_server'
require 'nomad_client/models/requested_device'
require 'nomad_client/models/reschedule_event'
require 'nomad_client/models/reschedule_policy'
require 'nomad_client/models/reschedule_tracker'
require 'nomad_client/models/resources'
require 'nomad_client/models/restart_policy'
require 'nomad_client/models/sampled_value'
require 'nomad_client/models/scaling_event'
require 'nomad_client/models/scaling_policy'
require 'nomad_client/models/scaling_policy_list_stub'
require 'nomad_client/models/scaling_request'
require 'nomad_client/models/scheduler_configuration'
require 'nomad_client/models/scheduler_configuration_response'
require 'nomad_client/models/scheduler_set_configuration_response'
require 'nomad_client/models/search_request'
require 'nomad_client/models/search_response'
require 'nomad_client/models/server_health'
require 'nomad_client/models/service'
require 'nomad_client/models/service_check'
require 'nomad_client/models/sidecar_task'
require 'nomad_client/models/spread'
require 'nomad_client/models/spread_target'
require 'nomad_client/models/task'
require 'nomad_client/models/task_artifact'
require 'nomad_client/models/task_csi_plugin_config'
require 'nomad_client/models/task_diff'
require 'nomad_client/models/task_event'
require 'nomad_client/models/task_group'
require 'nomad_client/models/task_group_diff'
require 'nomad_client/models/task_group_scale_status'
require 'nomad_client/models/task_group_summary'
require 'nomad_client/models/task_handle'
require 'nomad_client/models/task_lifecycle'
require 'nomad_client/models/task_state'
require 'nomad_client/models/template'
require 'nomad_client/models/update_strategy'
require 'nomad_client/models/vault'
require 'nomad_client/models/volume_mount'
require 'nomad_client/models/volume_request'

# APIs
require 'nomad_client/api/acl_api'
require 'nomad_client/api/allocations_api'
require 'nomad_client/api/deployments_api'
require 'nomad_client/api/enterprise_api'
require 'nomad_client/api/evaluations_api'
require 'nomad_client/api/jobs_api'
require 'nomad_client/api/metrics_api'
require 'nomad_client/api/namespaces_api'
require 'nomad_client/api/nodes_api'
require 'nomad_client/api/operator_api'
require 'nomad_client/api/plugins_api'
require 'nomad_client/api/regions_api'
require 'nomad_client/api/scaling_api'
require 'nomad_client/api/search_api'
require 'nomad_client/api/system_api'
require 'nomad_client/api/volumes_api'

module NomadClient
  class << self
    # Customize default settings for the SDK using block.
    #   NomadClient.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end
