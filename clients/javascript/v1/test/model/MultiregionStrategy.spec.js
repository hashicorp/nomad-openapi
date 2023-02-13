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
    instance = new nomad-client.MultiregionStrategy();
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

  describe('MultiregionStrategy', function() {
    it('should create an instance of MultiregionStrategy', function() {
      // uncomment below and update the code to test MultiregionStrategy
      //var instane = new nomad-client.MultiregionStrategy();
      //expect(instance).to.be.a(nomad-client.MultiregionStrategy);
    });

    it('should have the property maxParallel (base name: "MaxParallel")', function() {
      // uncomment below and update the code to test the property maxParallel
      //var instance = new nomad-client.MultiregionStrategy();
      //expect(instance).to.be();
    });

    it('should have the property onFailure (base name: "OnFailure")', function() {
      // uncomment below and update the code to test the property onFailure
      //var instance = new nomad-client.MultiregionStrategy();
      //expect(instance).to.be();
    });

  });

}));
