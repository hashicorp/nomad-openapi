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
    instance = new nomad-client.TaskGroupScaleStatus();
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

  describe('TaskGroupScaleStatus', function() {
    it('should create an instance of TaskGroupScaleStatus', function() {
      // uncomment below and update the code to test TaskGroupScaleStatus
      //var instane = new nomad-client.TaskGroupScaleStatus();
      //expect(instance).to.be.a(nomad-client.TaskGroupScaleStatus);
    });

    it('should have the property desired (base name: "Desired")', function() {
      // uncomment below and update the code to test the property desired
      //var instance = new nomad-client.TaskGroupScaleStatus();
      //expect(instance).to.be();
    });

    it('should have the property events (base name: "Events")', function() {
      // uncomment below and update the code to test the property events
      //var instance = new nomad-client.TaskGroupScaleStatus();
      //expect(instance).to.be();
    });

    it('should have the property healthy (base name: "Healthy")', function() {
      // uncomment below and update the code to test the property healthy
      //var instance = new nomad-client.TaskGroupScaleStatus();
      //expect(instance).to.be();
    });

    it('should have the property placed (base name: "Placed")', function() {
      // uncomment below and update the code to test the property placed
      //var instance = new nomad-client.TaskGroupScaleStatus();
      //expect(instance).to.be();
    });

    it('should have the property running (base name: "Running")', function() {
      // uncomment below and update the code to test the property running
      //var instance = new nomad-client.TaskGroupScaleStatus();
      //expect(instance).to.be();
    });

    it('should have the property unhealthy (base name: "Unhealthy")', function() {
      // uncomment below and update the code to test the property unhealthy
      //var instance = new nomad-client.TaskGroupScaleStatus();
      //expect(instance).to.be();
    });

  });

}));
