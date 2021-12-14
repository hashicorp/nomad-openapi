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
  class AutopilotConfiguration
    attr_accessor :cleanup_dead_servers

    attr_accessor :create_index

    attr_accessor :disable_upgrade_migration

    attr_accessor :enable_custom_upgrades

    attr_accessor :enable_redundancy_zones

    attr_accessor :last_contact_threshold

    attr_accessor :max_trailing_logs

    attr_accessor :min_quorum

    attr_accessor :modify_index

    attr_accessor :server_stabilization_time

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'cleanup_dead_servers' => :'CleanupDeadServers',
        :'create_index' => :'CreateIndex',
        :'disable_upgrade_migration' => :'DisableUpgradeMigration',
        :'enable_custom_upgrades' => :'EnableCustomUpgrades',
        :'enable_redundancy_zones' => :'EnableRedundancyZones',
        :'last_contact_threshold' => :'LastContactThreshold',
        :'max_trailing_logs' => :'MaxTrailingLogs',
        :'min_quorum' => :'MinQuorum',
        :'modify_index' => :'ModifyIndex',
        :'server_stabilization_time' => :'ServerStabilizationTime'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'cleanup_dead_servers' => :'Boolean',
        :'create_index' => :'Integer',
        :'disable_upgrade_migration' => :'Boolean',
        :'enable_custom_upgrades' => :'Boolean',
        :'enable_redundancy_zones' => :'Boolean',
        :'last_contact_threshold' => :'Integer',
        :'max_trailing_logs' => :'Integer',
        :'min_quorum' => :'Integer',
        :'modify_index' => :'Integer',
        :'server_stabilization_time' => :'Integer'
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
        fail ArgumentError, "The input argument (attributes) must be a hash in `NomadClient::AutopilotConfiguration` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `NomadClient::AutopilotConfiguration`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'cleanup_dead_servers')
        self.cleanup_dead_servers = attributes[:'cleanup_dead_servers']
      end

      if attributes.key?(:'create_index')
        self.create_index = attributes[:'create_index']
      end

      if attributes.key?(:'disable_upgrade_migration')
        self.disable_upgrade_migration = attributes[:'disable_upgrade_migration']
      end

      if attributes.key?(:'enable_custom_upgrades')
        self.enable_custom_upgrades = attributes[:'enable_custom_upgrades']
      end

      if attributes.key?(:'enable_redundancy_zones')
        self.enable_redundancy_zones = attributes[:'enable_redundancy_zones']
      end

      if attributes.key?(:'last_contact_threshold')
        self.last_contact_threshold = attributes[:'last_contact_threshold']
      end

      if attributes.key?(:'max_trailing_logs')
        self.max_trailing_logs = attributes[:'max_trailing_logs']
      end

      if attributes.key?(:'min_quorum')
        self.min_quorum = attributes[:'min_quorum']
      end

      if attributes.key?(:'modify_index')
        self.modify_index = attributes[:'modify_index']
      end

      if attributes.key?(:'server_stabilization_time')
        self.server_stabilization_time = attributes[:'server_stabilization_time']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if !@create_index.nil? && @create_index > 384
        invalid_properties.push('invalid value for "create_index", must be smaller than or equal to 384.')
      end

      if !@create_index.nil? && @create_index < 0
        invalid_properties.push('invalid value for "create_index", must be greater than or equal to 0.')
      end

      if !@max_trailing_logs.nil? && @max_trailing_logs > 384
        invalid_properties.push('invalid value for "max_trailing_logs", must be smaller than or equal to 384.')
      end

      if !@max_trailing_logs.nil? && @max_trailing_logs < 0
        invalid_properties.push('invalid value for "max_trailing_logs", must be greater than or equal to 0.')
      end

      if !@min_quorum.nil? && @min_quorum < 0
        invalid_properties.push('invalid value for "min_quorum", must be greater than or equal to 0.')
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
      return false if !@create_index.nil? && @create_index > 384
      return false if !@create_index.nil? && @create_index < 0
      return false if !@max_trailing_logs.nil? && @max_trailing_logs > 384
      return false if !@max_trailing_logs.nil? && @max_trailing_logs < 0
      return false if !@min_quorum.nil? && @min_quorum < 0
      return false if !@modify_index.nil? && @modify_index > 384
      return false if !@modify_index.nil? && @modify_index < 0
      true
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
    # @param [Object] max_trailing_logs Value to be assigned
    def max_trailing_logs=(max_trailing_logs)
      if !max_trailing_logs.nil? && max_trailing_logs > 384
        fail ArgumentError, 'invalid value for "max_trailing_logs", must be smaller than or equal to 384.'
      end

      if !max_trailing_logs.nil? && max_trailing_logs < 0
        fail ArgumentError, 'invalid value for "max_trailing_logs", must be greater than or equal to 0.'
      end

      @max_trailing_logs = max_trailing_logs
    end

    # Custom attribute writer method with validation
    # @param [Object] min_quorum Value to be assigned
    def min_quorum=(min_quorum)
      if !min_quorum.nil? && min_quorum < 0
        fail ArgumentError, 'invalid value for "min_quorum", must be greater than or equal to 0.'
      end

      @min_quorum = min_quorum
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
          cleanup_dead_servers == o.cleanup_dead_servers &&
          create_index == o.create_index &&
          disable_upgrade_migration == o.disable_upgrade_migration &&
          enable_custom_upgrades == o.enable_custom_upgrades &&
          enable_redundancy_zones == o.enable_redundancy_zones &&
          last_contact_threshold == o.last_contact_threshold &&
          max_trailing_logs == o.max_trailing_logs &&
          min_quorum == o.min_quorum &&
          modify_index == o.modify_index &&
          server_stabilization_time == o.server_stabilization_time
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [cleanup_dead_servers, create_index, disable_upgrade_migration, enable_custom_upgrades, enable_redundancy_zones, last_contact_threshold, max_trailing_logs, min_quorum, modify_index, server_stabilization_time].hash
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
