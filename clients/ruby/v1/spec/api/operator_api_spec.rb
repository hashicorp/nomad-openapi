=begin
#Nomad

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.1.4
Contact: support@hashicorp.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.2.0

=end

require 'spec_helper'
require 'json'

# Unit tests for NomadClient::OperatorApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'OperatorApi' do
  before do
    # run before each test
    @api_instance = NomadClient::OperatorApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of OperatorApi' do
    it 'should create an instance of OperatorApi' do
      expect(@api_instance).to be_instance_of(NomadClient::OperatorApi)
    end
  end

  # unit tests for delete_operator_raft_peer
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
  # @return [nil]
  describe 'delete_operator_raft_peer test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_operator_autopilot_configuration
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
  # @option opts [String] :wait Provided with IndexParam to wait for change.
  # @option opts [String] :stale If present, results will include stale reads.
  # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [Integer] :per_page Maximum number of results to return.
  # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
  # @return [AutopilotConfiguration]
  describe 'get_operator_autopilot_configuration test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_operator_autopilot_health
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
  # @option opts [String] :wait Provided with IndexParam to wait for change.
  # @option opts [String] :stale If present, results will include stale reads.
  # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [Integer] :per_page Maximum number of results to return.
  # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
  # @return [OperatorHealthReply]
  describe 'get_operator_autopilot_health test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_operator_raft_configuration
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
  # @option opts [String] :wait Provided with IndexParam to wait for change.
  # @option opts [String] :stale If present, results will include stale reads.
  # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [Integer] :per_page Maximum number of results to return.
  # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
  # @return [Array<RaftConfigurationResponse>]
  describe 'get_operator_raft_configuration test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_operator_scheduler_configuration
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [Integer] :index If set, wait until query exceeds given index. Must be provided with WaitParam.
  # @option opts [String] :wait Provided with IndexParam to wait for change.
  # @option opts [String] :stale If present, results will include stale reads.
  # @option opts [String] :prefix Constrains results to jobs that start with the defined prefix
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [Integer] :per_page Maximum number of results to return.
  # @option opts [String] :next_token Indicates where to start paging for queries that support pagination.
  # @return [SchedulerConfigurationResponse]
  describe 'get_operator_scheduler_configuration test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for post_operator_scheduler_configuration
  # @param scheduler_configuration 
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
  # @return [SchedulerSetConfigurationResponse]
  describe 'post_operator_scheduler_configuration test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for put_operator_autopilot_configuration
  # @param autopilot_configuration 
  # @param [Hash] opts the optional parameters
  # @option opts [String] :region Filters results based on the specified region.
  # @option opts [String] :namespace Filters results based on the specified namespace.
  # @option opts [String] :x_nomad_token A Nomad ACL token.
  # @option opts [String] :idempotency_token Can be used to ensure operations are only run once.
  # @return [Boolean]
  describe 'put_operator_autopilot_configuration test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
