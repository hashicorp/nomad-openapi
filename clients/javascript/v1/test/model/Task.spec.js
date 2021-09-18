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
    instance = new nomad-client.Task();
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

  describe('Task', function() {
    it('should create an instance of Task', function() {
      // uncomment below and update the code to test Task
      //var instane = new nomad-client.Task();
      //expect(instance).to.be.a(nomad-client.Task);
    });

    it('should have the property affinities (base name: "Affinities")', function() {
      // uncomment below and update the code to test the property affinities
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property artifacts (base name: "Artifacts")', function() {
      // uncomment below and update the code to test the property artifacts
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property cSIPluginConfig (base name: "CSIPluginConfig")', function() {
      // uncomment below and update the code to test the property cSIPluginConfig
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "Config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property constraints (base name: "Constraints")', function() {
      // uncomment below and update the code to test the property constraints
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property dispatchPayload (base name: "DispatchPayload")', function() {
      // uncomment below and update the code to test the property dispatchPayload
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property driver (base name: "Driver")', function() {
      // uncomment below and update the code to test the property driver
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property env (base name: "Env")', function() {
      // uncomment below and update the code to test the property env
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property killSignal (base name: "KillSignal")', function() {
      // uncomment below and update the code to test the property killSignal
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property killTimeout (base name: "KillTimeout")', function() {
      // uncomment below and update the code to test the property killTimeout
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property kind (base name: "Kind")', function() {
      // uncomment below and update the code to test the property kind
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property leader (base name: "Leader")', function() {
      // uncomment below and update the code to test the property leader
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property lifecycle (base name: "Lifecycle")', function() {
      // uncomment below and update the code to test the property lifecycle
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property logConfig (base name: "LogConfig")', function() {
      // uncomment below and update the code to test the property logConfig
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property meta (base name: "Meta")', function() {
      // uncomment below and update the code to test the property meta
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "Name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property resources (base name: "Resources")', function() {
      // uncomment below and update the code to test the property resources
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property restartPolicy (base name: "RestartPolicy")', function() {
      // uncomment below and update the code to test the property restartPolicy
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property scalingPolicies (base name: "ScalingPolicies")', function() {
      // uncomment below and update the code to test the property scalingPolicies
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property services (base name: "Services")', function() {
      // uncomment below and update the code to test the property services
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property shutdownDelay (base name: "ShutdownDelay")', function() {
      // uncomment below and update the code to test the property shutdownDelay
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property templates (base name: "Templates")', function() {
      // uncomment below and update the code to test the property templates
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property user (base name: "User")', function() {
      // uncomment below and update the code to test the property user
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property vault (base name: "Vault")', function() {
      // uncomment below and update the code to test the property vault
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

    it('should have the property volumeMounts (base name: "VolumeMounts")', function() {
      // uncomment below and update the code to test the property volumeMounts
      //var instance = new nomad-client.Task();
      //expect(instance).to.be();
    });

  });

}));
