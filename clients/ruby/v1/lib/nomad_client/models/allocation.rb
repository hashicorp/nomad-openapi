# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

=begin
#Nomad

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.1.4
Contact: support@hashicorp.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.2.0

=end

require 'date'
require 'time'

module NomadClient
  class Allocation
    attr_accessor :alloc_modify_index

    attr_accessor :allocated_resources

    attr_accessor :client_description

    attr_accessor :client_status

    attr_accessor :create_index

    attr_accessor :create_time

    attr_accessor :deployment_id

    attr_accessor :deployment_status

    attr_accessor :desired_description

    attr_accessor :desired_status

    attr_accessor :desired_transition

    attr_accessor :eval_id

    attr_accessor :followup_eval_id

    attr_accessor :id

    attr_accessor :job

    attr_accessor :job_id

    attr_accessor :metrics

    attr_accessor :modify_index

    attr_accessor :modify_time

    attr_accessor :name

    attr_accessor :namespace

    attr_accessor :next_allocation

    attr_accessor :node_id

    attr_accessor :node_name

    attr_accessor :preempted_allocations

    attr_accessor :preempted_by_allocation

    attr_accessor :previous_allocation

    attr_accessor :reschedule_tracker

    attr_accessor :resources

    attr_accessor :services

    attr_accessor :task_group

    attr_accessor :task_resources

    attr_accessor :task_states

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'alloc_modify_index' => :'AllocModifyIndex',
        :'allocated_resources' => :'AllocatedResources',
        :'client_description' => :'ClientDescription',
        :'client_status' => :'ClientStatus',
        :'create_index' => :'CreateIndex',
        :'create_time' => :'CreateTime',
        :'deployment_id' => :'DeploymentID',
        :'deployment_status' => :'DeploymentStatus',
        :'desired_description' => :'DesiredDescription',
        :'desired_status' => :'DesiredStatus',
        :'desired_transition' => :'DesiredTransition',
        :'eval_id' => :'EvalID',
        :'followup_eval_id' => :'FollowupEvalID',
        :'id' => :'ID',
        :'job' => :'Job',
        :'job_id' => :'JobID',
        :'metrics' => :'Metrics',
        :'modify_index' => :'ModifyIndex',
        :'modify_time' => :'ModifyTime',
        :'name' => :'Name',
        :'namespace' => :'Namespace',
        :'next_allocation' => :'NextAllocation',
        :'node_id' => :'NodeID',
        :'node_name' => :'NodeName',
        :'preempted_allocations' => :'PreemptedAllocations',
        :'preempted_by_allocation' => :'PreemptedByAllocation',
        :'previous_allocation' => :'PreviousAllocation',
        :'reschedule_tracker' => :'RescheduleTracker',
        :'resources' => :'Resources',
        :'services' => :'Services',
        :'task_group' => :'TaskGroup',
        :'task_resources' => :'TaskResources',
        :'task_states' => :'TaskStates'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'alloc_modify_index' => :'Integer',
        :'allocated_resources' => :'AllocatedResources',
        :'client_description' => :'String',
        :'client_status' => :'String',
        :'create_index' => :'Integer',
        :'create_time' => :'Integer',
        :'deployment_id' => :'String',
        :'deployment_status' => :'AllocDeploymentStatus',
        :'desired_description' => :'String',
        :'desired_status' => :'String',
        :'desired_transition' => :'DesiredTransition',
        :'eval_id' => :'String',
        :'followup_eval_id' => :'String',
        :'id' => :'String',
        :'job' => :'Job',
        :'job_id' => :'String',
        :'metrics' => :'AllocationMetric',
        :'modify_index' => :'Integer',
        :'modify_time' => :'Integer',
        :'name' => :'String',
        :'namespace' => :'String',
        :'next_allocation' => :'String',
        :'node_id' => :'String',
        :'node_name' => :'String',
        :'preempted_allocations' => :'Array<String>',
        :'preempted_by_allocation' => :'String',
        :'previous_allocation' => :'String',
        :'reschedule_tracker' => :'RescheduleTracker',
        :'resources' => :'Resources',
        :'services' => :'Hash<String, String>',
        :'task_group' => :'String',
        :'task_resources' => :'Hash<String, Resources>',
        :'task_states' => :'Hash<String, TaskState>'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `NomadClient::Allocation` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `NomadClient::Allocation`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'alloc_modify_index')
        self.alloc_modify_index = attributes[:'alloc_modify_index']
      end

      if attributes.key?(:'allocated_resources')
        self.allocated_resources = attributes[:'allocated_resources']
      end

      if attributes.key?(:'client_description')
        self.client_description = attributes[:'client_description']
      end

      if attributes.key?(:'client_status')
        self.client_status = attributes[:'client_status']
      end

      if attributes.key?(:'create_index')
        self.create_index = attributes[:'create_index']
      end

      if attributes.key?(:'create_time')
        self.create_time = attributes[:'create_time']
      end

      if attributes.key?(:'deployment_id')
        self.deployment_id = attributes[:'deployment_id']
      end

      if attributes.key?(:'deployment_status')
        self.deployment_status = attributes[:'deployment_status']
      end

      if attributes.key?(:'desired_description')
        self.desired_description = attributes[:'desired_description']
      end

      if attributes.key?(:'desired_status')
        self.desired_status = attributes[:'desired_status']
      end

      if attributes.key?(:'desired_transition')
        self.desired_transition = attributes[:'desired_transition']
      end

      if attributes.key?(:'eval_id')
        self.eval_id = attributes[:'eval_id']
      end

      if attributes.key?(:'followup_eval_id')
        self.followup_eval_id = attributes[:'followup_eval_id']
      end

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'job')
        self.job = attributes[:'job']
      end

      if attributes.key?(:'job_id')
        self.job_id = attributes[:'job_id']
      end

      if attributes.key?(:'metrics')
        self.metrics = attributes[:'metrics']
      end

      if attributes.key?(:'modify_index')
        self.modify_index = attributes[:'modify_index']
      end

      if attributes.key?(:'modify_time')
        self.modify_time = attributes[:'modify_time']
      end

      if attributes.key?(:'name')
        self.name = attributes[:'name']
      end

      if attributes.key?(:'namespace')
        self.namespace = attributes[:'namespace']
      end

      if attributes.key?(:'next_allocation')
        self.next_allocation = attributes[:'next_allocation']
      end

      if attributes.key?(:'node_id')
        self.node_id = attributes[:'node_id']
      end

      if attributes.key?(:'node_name')
        self.node_name = attributes[:'node_name']
      end

      if attributes.key?(:'preempted_allocations')
        if (value = attributes[:'preempted_allocations']).is_a?(Array)
          self.preempted_allocations = value
        end
      end

      if attributes.key?(:'preempted_by_allocation')
        self.preempted_by_allocation = attributes[:'preempted_by_allocation']
      end

      if attributes.key?(:'previous_allocation')
        self.previous_allocation = attributes[:'previous_allocation']
      end

      if attributes.key?(:'reschedule_tracker')
        self.reschedule_tracker = attributes[:'reschedule_tracker']
      end

      if attributes.key?(:'resources')
        self.resources = attributes[:'resources']
      end

      if attributes.key?(:'services')
        if (value = attributes[:'services']).is_a?(Hash)
          self.services = value
        end
      end

      if attributes.key?(:'task_group')
        self.task_group = attributes[:'task_group']
      end

      if attributes.key?(:'task_resources')
        if (value = attributes[:'task_resources']).is_a?(Hash)
          self.task_resources = value
        end
      end

      if attributes.key?(:'task_states')
        if (value = attributes[:'task_states']).is_a?(Hash)
          self.task_states = value
        end
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if !@alloc_modify_index.nil? && @alloc_modify_index > 384
        invalid_properties.push('invalid value for "alloc_modify_index", must be smaller than or equal to 384.')
      end

      if !@alloc_modify_index.nil? && @alloc_modify_index < 0
        invalid_properties.push('invalid value for "alloc_modify_index", must be greater than or equal to 0.')
      end

      if !@create_index.nil? && @create_index > 384
        invalid_properties.push('invalid value for "create_index", must be smaller than or equal to 384.')
      end

      if !@create_index.nil? && @create_index < 0
        invalid_properties.push('invalid value for "create_index", must be greater than or equal to 0.')
      end

      if !@modify_index.nil? && @modify_index > 384
        invalid_properties.push('invalid value for "modify_index", must be smaller than or equal to 384.')
      end

      if !@modify_index.nil? && @modify_index < 0
        invalid_properties.push('invalid value for "modify_index", must be greater than or equal to 0.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if !@alloc_modify_index.nil? && @alloc_modify_index > 384
      return false if !@alloc_modify_index.nil? && @alloc_modify_index < 0
      return false if !@create_index.nil? && @create_index > 384
      return false if !@create_index.nil? && @create_index < 0
      return false if !@modify_index.nil? && @modify_index > 384
      return false if !@modify_index.nil? && @modify_index < 0
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] alloc_modify_index Value to be assigned
    def alloc_modify_index=(alloc_modify_index)
      if !alloc_modify_index.nil? && alloc_modify_index > 384
        fail ArgumentError, 'invalid value for "alloc_modify_index", must be smaller than or equal to 384.'
      end

      if !alloc_modify_index.nil? && alloc_modify_index < 0
        fail ArgumentError, 'invalid value for "alloc_modify_index", must be greater than or equal to 0.'
      end

      @alloc_modify_index = alloc_modify_index
    end

    # Custom attribute writer method with validation
    # @param [Object] create_index Value to be assigned
    def create_index=(create_index)
      if !create_index.nil? && create_index > 384
        fail ArgumentError, 'invalid value for "create_index", must be smaller than or equal to 384.'
      end

      if !create_index.nil? && create_index < 0
        fail ArgumentError, 'invalid value for "create_index", must be greater than or equal to 0.'
      end

      @create_index = create_index
    end

    # Custom attribute writer method with validation
    # @param [Object] modify_index Value to be assigned
    def modify_index=(modify_index)
      if !modify_index.nil? && modify_index > 384
        fail ArgumentError, 'invalid value for "modify_index", must be smaller than or equal to 384.'
      end

      if !modify_index.nil? && modify_index < 0
        fail ArgumentError, 'invalid value for "modify_index", must be greater than or equal to 0.'
      end

      @modify_index = modify_index
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          alloc_modify_index == o.alloc_modify_index &&
          allocated_resources == o.allocated_resources &&
          client_description == o.client_description &&
          client_status == o.client_status &&
          create_index == o.create_index &&
          create_time == o.create_time &&
          deployment_id == o.deployment_id &&
          deployment_status == o.deployment_status &&
          desired_description == o.desired_description &&
          desired_status == o.desired_status &&
          desired_transition == o.desired_transition &&
          eval_id == o.eval_id &&
          followup_eval_id == o.followup_eval_id &&
          id == o.id &&
          job == o.job &&
          job_id == o.job_id &&
          metrics == o.metrics &&
          modify_index == o.modify_index &&
          modify_time == o.modify_time &&
          name == o.name &&
          namespace == o.namespace &&
          next_allocation == o.next_allocation &&
          node_id == o.node_id &&
          node_name == o.node_name &&
          preempted_allocations == o.preempted_allocations &&
          preempted_by_allocation == o.preempted_by_allocation &&
          previous_allocation == o.previous_allocation &&
          reschedule_tracker == o.reschedule_tracker &&
          resources == o.resources &&
          services == o.services &&
          task_group == o.task_group &&
          task_resources == o.task_resources &&
          task_states == o.task_states
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [alloc_modify_index, allocated_resources, client_description, client_status, create_index, create_time, deployment_id, deployment_status, desired_description, desired_status, desired_transition, eval_id, followup_eval_id, id, job, job_id, metrics, modify_index, modify_time, name, namespace, next_allocation, node_id, node_name, preempted_allocations, preempted_by_allocation, previous_allocation, reschedule_tracker, resources, services, task_group, task_resources, task_states].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      new.build_from_hash(attributes)
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      self.class.openapi_types.each_pair do |key, type|
        if attributes[self.class.attribute_map[key]].nil? && self.class.openapi_nullable.include?(key)
          self.send("#{key}=", nil)
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[self.class.attribute_map[key]].is_a?(Array)
            self.send("#{key}=", attributes[self.class.attribute_map[key]].map { |v| _deserialize($1, v) })
          end
        elsif !attributes[self.class.attribute_map[key]].nil?
          self.send("#{key}=", _deserialize(type, attributes[self.class.attribute_map[key]]))
        end
      end

      self
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def _deserialize(type, value)
      case type.to_sym
      when :Time
        Time.parse(value)
      when :Date
        Date.parse(value)
      when :String
        value.to_s
      when :Integer
        value.to_i
      when :Float
        value.to_f
      when :Boolean
        if value.to_s =~ /\A(true|t|yes|y|1)\z/i
          true
        else
          false
        end
      when :Object
        # generic object (usually a Hash), return directly
        value
      when /\AArray<(?<inner_type>.+)>\z/
        inner_type = Regexp.last_match[:inner_type]
        value.map { |v| _deserialize(inner_type, v) }
      when /\AHash<(?<k_type>.+?), (?<v_type>.+)>\z/
        k_type = Regexp.last_match[:k_type]
        v_type = Regexp.last_match[:v_type]
        {}.tap do |hash|
          value.each do |k, v|
            hash[_deserialize(k_type, k)] = _deserialize(v_type, v)
          end
        end
      else # model
        # models (e.g. Pet) or oneOf
        klass = NomadClient.const_get(type)
        klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
      end
    end

    # Returns the string representation of the object
    # @return [String] String presentation of the object
    def to_s
      to_hash.to_s
    end

    # to_body is an alias to to_hash (backward compatibility)
    # @return [Hash] Returns the object in the form of hash
    def to_body
      to_hash
    end

    # Returns the object in the form of hash
    # @return [Hash] Returns the object in the form of hash
    def to_hash
      hash = {}
      self.class.attribute_map.each_pair do |attr, param|
        value = self.send(attr)
        if value.nil?
          is_nullable = self.class.openapi_nullable.include?(attr)
          next if !is_nullable || (is_nullable && !instance_variable_defined?(:"@#{attr}"))
        end

        hash[param] = _to_hash(value)
      end
      hash
    end

    # Outputs non-array value in the form of hash
    # For object, use to_hash. Otherwise, just return the value
    # @param [Object] value Any valid value
    # @return [Hash] Returns the value in the form of hash
    def _to_hash(value)
      if value.is_a?(Array)
        value.compact.map { |v| _to_hash(v) }
      elsif value.is_a?(Hash)
        {}.tap do |hash|
          value.each { |k, v| hash[k] = _to_hash(v) }
        end
      elsif value.respond_to? :to_hash
        value.to_hash
      else
        value
      end
    end

  end

end
