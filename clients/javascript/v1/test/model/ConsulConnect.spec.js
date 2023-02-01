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
    instance = new nomad-client.ConsulConnect();
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

  describe('ConsulConnect', function() {
    it('should create an instance of ConsulConnect', function() {
      // uncomment below and update the code to test ConsulConnect
      //var instane = new nomad-client.ConsulConnect();
      //expect(instance).to.be.a(nomad-client.ConsulConnect);
    });

    it('should have the property gateway (base name: "Gateway")', function() {
      // uncomment below and update the code to test the property gateway
      //var instance = new nomad-client.ConsulConnect();
      //expect(instance).to.be();
    });

    it('should have the property _native (base name: "Native")', function() {
      // uncomment below and update the code to test the property _native
      //var instance = new nomad-client.ConsulConnect();
      //expect(instance).to.be();
    });

    it('should have the property sidecarService (base name: "SidecarService")', function() {
      // uncomment below and update the code to test the property sidecarService
      //var instance = new nomad-client.ConsulConnect();
      //expect(instance).to.be();
    });

    it('should have the property sidecarTask (base name: "SidecarTask")', function() {
      // uncomment below and update the code to test the property sidecarTask
      //var instance = new nomad-client.ConsulConnect();
      //expect(instance).to.be();
    });

  });

}));
