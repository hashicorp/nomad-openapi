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
    instance = new nomad-client.ConsulSidecarService();
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

  describe('ConsulSidecarService', function() {
    it('should create an instance of ConsulSidecarService', function() {
      // uncomment below and update the code to test ConsulSidecarService
      //var instane = new nomad-client.ConsulSidecarService();
      //expect(instance).to.be.a(nomad-client.ConsulSidecarService);
    });

    it('should have the property disableDefaultTCPCheck (base name: "DisableDefaultTCPCheck")', function() {
      // uncomment below and update the code to test the property disableDefaultTCPCheck
      //var instance = new nomad-client.ConsulSidecarService();
      //expect(instance).to.be();
    });

    it('should have the property port (base name: "Port")', function() {
      // uncomment below and update the code to test the property port
      //var instance = new nomad-client.ConsulSidecarService();
      //expect(instance).to.be();
    });

    it('should have the property proxy (base name: "Proxy")', function() {
      // uncomment below and update the code to test the property proxy
      //var instance = new nomad-client.ConsulSidecarService();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "Tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new nomad-client.ConsulSidecarService();
      //expect(instance).to.be();
    });

  });

}));
