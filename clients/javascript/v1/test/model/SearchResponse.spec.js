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
    instance = new nomad-client.SearchResponse();
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

  describe('SearchResponse', function() {
    it('should create an instance of SearchResponse', function() {
      // uncomment below and update the code to test SearchResponse
      //var instane = new nomad-client.SearchResponse();
      //expect(instance).to.be.a(nomad-client.SearchResponse);
    });

    it('should have the property knownLeader (base name: "KnownLeader")', function() {
      // uncomment below and update the code to test the property knownLeader
      //var instance = new nomad-client.SearchResponse();
      //expect(instance).to.be();
    });

    it('should have the property lastContact (base name: "LastContact")', function() {
      // uncomment below and update the code to test the property lastContact
      //var instance = new nomad-client.SearchResponse();
      //expect(instance).to.be();
    });

    it('should have the property lastIndex (base name: "LastIndex")', function() {
      // uncomment below and update the code to test the property lastIndex
      //var instance = new nomad-client.SearchResponse();
      //expect(instance).to.be();
    });

    it('should have the property matches (base name: "Matches")', function() {
      // uncomment below and update the code to test the property matches
      //var instance = new nomad-client.SearchResponse();
      //expect(instance).to.be();
    });

    it('should have the property requestTime (base name: "RequestTime")', function() {
      // uncomment below and update the code to test the property requestTime
      //var instance = new nomad-client.SearchResponse();
      //expect(instance).to.be();
    });

    it('should have the property truncations (base name: "Truncations")', function() {
      // uncomment below and update the code to test the property truncations
      //var instance = new nomad-client.SearchResponse();
      //expect(instance).to.be();
    });

  });

}));
