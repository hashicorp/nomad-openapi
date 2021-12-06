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
  class CSIPluginListStub
    attr_accessor :controller_required

    attr_accessor :controllers_expected

    attr_accessor :controllers_healthy

    attr_accessor :create_index

    attr_accessor :id

    attr_accessor :modify_index

    attr_accessor :nodes_expected

    attr_accessor :nodes_healthy

    attr_accessor :provider

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'controller_required' => :'ControllerRequired',
        :'controllers_expected' => :'ControllersExpected',
        :'controllers_healthy' => :'ControllersHealthy',
        :'create_index' => :'CreateIndex',
        :'id' => :'ID',
        :'modify_index' => :'ModifyIndex',
        :'nodes_expected' => :'NodesExpected',
        :'nodes_healthy' => :'NodesHealthy',
        :'provider' => :'Provider'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'controller_required' => :'Boolean',
        :'controllers_expected' => :'Integer',
        :'controllers_healthy' => :'Integer',
        :'create_index' => :'Integer',
        :'id' => :'String',
        :'modify_index' => :'Integer',
        :'nodes_expected' => :'Integer',
        :'nodes_healthy' => :'Integer',
        :'provider' => :'String'
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
        fail ArgumentError, "The input argument (attributes) must be a hash in `NomadClient::CSIPluginListStub` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `NomadClient::CSIPluginListStub`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'controller_required')
        self.controller_required = attributes[:'controller_required']
      end

      if attributes.key?(:'controllers_expected')
        self.controllers_expected = attributes[:'controllers_expected']
      end

      if attributes.key?(:'controllers_healthy')
        self.controllers_healthy = attributes[:'controllers_healthy']
      end

      if attributes.key?(:'create_index')
        self.create_index = attributes[:'create_index']
      end

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'modify_index')
        self.modify_index = attributes[:'modify_index']
      end

      if attributes.key?(:'nodes_expected')
        self.nodes_expected = attributes[:'nodes_expected']
      end

      if attributes.key?(:'nodes_healthy')
        self.nodes_healthy = attributes[:'nodes_healthy']
      end

      if attributes.key?(:'provider')
        self.provider = attributes[:'provider']
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
          controller_required == o.controller_required &&
          controllers_expected == o.controllers_expected &&
          controllers_healthy == o.controllers_healthy &&
          create_index == o.create_index &&
          id == o.id &&
          modify_index == o.modify_index &&
          nodes_expected == o.nodes_expected &&
          nodes_healthy == o.nodes_healthy &&
          provider == o.provider
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [controller_required, controllers_expected, controllers_healthy, create_index, id, modify_index, nodes_expected, nodes_healthy, provider].hash
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