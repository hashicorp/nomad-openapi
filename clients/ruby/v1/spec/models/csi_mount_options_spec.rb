=begin
#Nomad

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.1.4
Contact: support@hashicorp.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.4.0

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for NomadClient::CSIMountOptions
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe NomadClient::CSIMountOptions do
  let(:instance) { NomadClient::CSIMountOptions.new }

  describe 'test an instance of CSIMountOptions' do
    it 'should create an instance of CSIMountOptions' do
      expect(instance).to be_instance_of(NomadClient::CSIMountOptions)
    end
  end
  describe 'test attribute "fs_type"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "mount_flags"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
