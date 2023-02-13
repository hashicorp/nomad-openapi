/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.3
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.nomad-client);
  }
}(this, function(expect, nomad-client) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new nomad-client.QuotaSpec();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('QuotaSpec', function() {
    it('should create an instance of QuotaSpec', function() {
      // uncomment below and update the code to test QuotaSpec
      //var instane = new nomad-client.QuotaSpec();
      //expect(instance).to.be.a(nomad-client.QuotaSpec);
    });

    it('should have the property createIndex (base name: "CreateIndex")', function() {
      // uncomment below and update the code to test the property createIndex
      //var instance = new nomad-client.QuotaSpec();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "Description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new nomad-client.QuotaSpec();
      //expect(instance).to.be();
    });

    it('should have the property limits (base name: "Limits")', function() {
      // uncomment below and update the code to test the property limits
      //var instance = new nomad-client.QuotaSpec();
      //expect(instance).to.be();
    });

    it('should have the property modifyIndex (base name: "ModifyIndex")', function() {
      // uncomment below and update the code to test the property modifyIndex
      //var instance = new nomad-client.QuotaSpec();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "Name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new nomad-client.QuotaSpec();
      //expect(instance).to.be();
    });

  });

}));
