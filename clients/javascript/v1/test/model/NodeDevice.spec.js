/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
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
    instance = new nomad-client.NodeDevice();
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

  describe('NodeDevice', function() {
    it('should create an instance of NodeDevice', function() {
      // uncomment below and update the code to test NodeDevice
      //var instane = new nomad-client.NodeDevice();
      //expect(instance).to.be.a(nomad-client.NodeDevice);
    });

    it('should have the property healthDescription (base name: "HealthDescription")', function() {
      // uncomment below and update the code to test the property healthDescription
      //var instance = new nomad-client.NodeDevice();
      //expect(instance).to.be();
    });

    it('should have the property healthy (base name: "Healthy")', function() {
      // uncomment below and update the code to test the property healthy
      //var instance = new nomad-client.NodeDevice();
      //expect(instance).to.be();
    });

    it('should have the property ID (base name: "ID")', function() {
      // uncomment below and update the code to test the property ID
      //var instance = new nomad-client.NodeDevice();
      //expect(instance).to.be();
    });

    it('should have the property locality (base name: "Locality")', function() {
      // uncomment below and update the code to test the property locality
      //var instance = new nomad-client.NodeDevice();
      //expect(instance).to.be();
    });

  });

}));
