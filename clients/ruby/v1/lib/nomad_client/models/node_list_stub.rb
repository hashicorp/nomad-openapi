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
  class NodeListStub
    attr_accessor :address

    attr_accessor :attributes

    attr_accessor :create_index

    attr_accessor :datacenter

    attr_accessor :drain

    attr_accessor :drivers

    attr_accessor :id

    attr_accessor :last_drain

    attr_accessor :modify_index

    attr_accessor :name

    attr_accessor :node_class

    attr_accessor :node_resources

    attr_accessor :reserved_resources

    attr_accessor :scheduling_eligibility

    attr_accessor :status

    attr_accessor :status_description

    attr_accessor :version

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'address' => :'Address',
        :'attributes' => :'Attributes',
        :'create_index' => :'CreateIndex',
        :'datacenter' => :'Datacenter',
        :'drain' => :'Drain',
        :'drivers' => :'Drivers',
        :'id' => :'ID',
        :'last_drain' => :'LastDrain',
        :'modify_index' => :'ModifyIndex',
        :'name' => :'Name',
        :'node_class' => :'NodeClass',
        :'node_resources' => :'NodeResources',
        :'reserved_resources' => :'ReservedResources',
        :'scheduling_eligibility' => :'SchedulingEligibility',
        :'status' => :'Status',
        :'status_description' => :'StatusDescription',
        :'version' => :'Version'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'address' => :'String',
        :'attributes' => :'Hash<String, String>',
        :'create_index' => :'Integer',
        :'datacenter' => :'String',
        :'drain' => :'Boolean',
        :'drivers' => :'Hash<String, DriverInfo>',
        :'id' => :'String',
        :'last_drain' => :'DrainMetadata',
        :'modify_index' => :'Integer',
        :'name' => :'String',
        :'node_class' => :'String',
        :'node_resources' => :'NodeResources',
        :'reserved_resources' => :'NodeReservedResources',
        :'scheduling_eligibility' => :'String',
        :'status' => :'String',
        :'status_description' => :'String',
        :'version' => :'String'
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
        fail ArgumentError, "The input argument (attributes) must be a hash in `NomadClient::NodeListStub` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `NomadClient::NodeListStub`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'address')
        self.address = attributes[:'address']
      end

      if attributes.key?(:'attributes')
        if (value = attributes[:'attributes']).is_a?(Hash)
          self.attributes = value
        end
      end

      if attributes.key?(:'create_index')
        self.create_index = attributes[:'create_index']
      end

      if attributes.key?(:'datacenter')
        self.datacenter = attributes[:'datacenter']
      end

      if attributes.key?(:'drain')
        self.drain = attributes[:'drain']
      end

      if attributes.key?(:'drivers')
        if (value = attributes[:'drivers']).is_a?(Hash)
          self.drivers = value
        end
      end

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'last_drain')
        self.last_drain = attributes[:'last_drain']
      end

      if attributes.key?(:'modify_index')
        self.modify_index = attributes[:'modify_index']
      end

      if attributes.key?(:'name')
        self.name = attributes[:'name']
      end

      if attributes.key?(:'node_class')
        self.node_class = attributes[:'node_class']
      end

      if attributes.key?(:'node_resources')
        self.node_resources = attributes[:'node_resources']
      end

      if attributes.key?(:'reserved_resources')
        self.reserved_resources = attributes[:'reserved_resources']
      end

      if attributes.key?(:'scheduling_eligibility')
        self.scheduling_eligibility = attributes[:'scheduling_eligibility']
      end

      if attributes.key?(:'status')
        self.status = attributes[:'status']
      end

      if attributes.key?(:'status_description')
        self.status_description = attributes[:'status_description']
      end

      if attributes.key?(:'version')
        self.version = attributes[:'version']
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
          address == o.address &&
          attributes == o.attributes &&
          create_index == o.create_index &&
          datacenter == o.datacenter &&
          drain == o.drain &&
          drivers == o.drivers &&
          id == o.id &&
          last_drain == o.last_drain &&
          modify_index == o.modify_index &&
          name == o.name &&
          node_class == o.node_class &&
          node_resources == o.node_resources &&
          reserved_resources == o.reserved_resources &&
          scheduling_eligibility == o.scheduling_eligibility &&
          status == o.status &&
          status_description == o.status_description &&
          version == o.version
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [address, attributes, create_index, datacenter, drain, drivers, id, last_drain, modify_index, name, node_class, node_resources, reserved_resources, scheduling_eligibility, status, status_description, version].hash
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
