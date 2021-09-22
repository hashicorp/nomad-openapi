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
    instance = new nomad-client.ScalingEvent();
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

  describe('ScalingEvent', function() {
    it('should create an instance of ScalingEvent', function() {
      // uncomment below and update the code to test ScalingEvent
      //var instane = new nomad-client.ScalingEvent();
      //expect(instance).to.be.a(nomad-client.ScalingEvent);
    });

    it('should have the property count (base name: "Count")', function() {
      // uncomment below and update the code to test the property count
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

    it('should have the property createIndex (base name: "CreateIndex")', function() {
      // uncomment below and update the code to test the property createIndex
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

    it('should have the property error (base name: "Error")', function() {
      // uncomment below and update the code to test the property error
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

    it('should have the property evalID (base name: "EvalID")', function() {
      // uncomment below and update the code to test the property evalID
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

    it('should have the property message (base name: "Message")', function() {
      // uncomment below and update the code to test the property message
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

    it('should have the property meta (base name: "Meta")', function() {
      // uncomment below and update the code to test the property meta
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

    it('should have the property previousCount (base name: "PreviousCount")', function() {
      // uncomment below and update the code to test the property previousCount
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

    it('should have the property time (base name: "Time")', function() {
      // uncomment below and update the code to test the property time
      //var instance = new nomad-client.ScalingEvent();
      //expect(instance).to.be();
    });

  });

}));