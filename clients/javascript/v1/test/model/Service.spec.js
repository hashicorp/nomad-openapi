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
    instance = new nomad-client.Service();
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

  describe('Service', function() {
    it('should create an instance of Service', function() {
      // uncomment below and update the code to test Service
      //var instane = new nomad-client.Service();
      //expect(instance).to.be.a(nomad-client.Service);
    });

    it('should have the property addressMode (base name: "AddressMode")', function() {
      // uncomment below and update the code to test the property addressMode
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property canaryMeta (base name: "CanaryMeta")', function() {
      // uncomment below and update the code to test the property canaryMeta
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property canaryTags (base name: "CanaryTags")', function() {
      // uncomment below and update the code to test the property canaryTags
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property checkRestart (base name: "CheckRestart")', function() {
      // uncomment below and update the code to test the property checkRestart
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property checks (base name: "Checks")', function() {
      // uncomment below and update the code to test the property checks
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property connect (base name: "Connect")', function() {
      // uncomment below and update the code to test the property connect
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property enableTagOverride (base name: "EnableTagOverride")', function() {
      // uncomment below and update the code to test the property enableTagOverride
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property id (base name: "Id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property meta (base name: "Meta")', function() {
      // uncomment below and update the code to test the property meta
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "Name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property onUpdate (base name: "OnUpdate")', function() {
      // uncomment below and update the code to test the property onUpdate
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property portLabel (base name: "PortLabel")', function() {
      // uncomment below and update the code to test the property portLabel
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "Tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

    it('should have the property taskName (base name: "TaskName")', function() {
      // uncomment below and update the code to test the property taskName
      //var instance = new nomad-client.Service();
      //expect(instance).to.be();
    });

  });

}));
