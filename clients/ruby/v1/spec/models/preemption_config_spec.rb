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

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for NomadClient::PreemptionConfig
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe NomadClient::PreemptionConfig do
  let(:instance) { NomadClient::PreemptionConfig.new }

  describe 'test an instance of PreemptionConfig' do
    it 'should create an instance of PreemptionConfig' do
      expect(instance).to be_instance_of(NomadClient::PreemptionConfig)
    end
  end
  describe 'test attribute "batch_scheduler_enabled"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "service_scheduler_enabled"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "sys_batch_scheduler_enabled"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "system_scheduler_enabled"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
