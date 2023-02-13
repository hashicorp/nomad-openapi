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
    instance = new nomad-client.EnterpriseApi();
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

  describe('EnterpriseApi', function() {
    describe('createQuotaSpec', function() {
      it('should call createQuotaSpec successfully', function(done) {
        //uncomment below and update the code to test createQuotaSpec
        //instance.createQuotaSpec(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteQuotaSpec', function() {
      it('should call deleteQuotaSpec successfully', function(done) {
        //uncomment below and update the code to test deleteQuotaSpec
        //instance.deleteQuotaSpec(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getQuotaSpec', function() {
      it('should call getQuotaSpec successfully', function(done) {
        //uncomment below and update the code to test getQuotaSpec
        //instance.getQuotaSpec(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getQuotas', function() {
      it('should call getQuotas successfully', function(done) {
        //uncomment below and update the code to test getQuotas
        //instance.getQuotas(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('postQuotaSpec', function() {
      it('should call postQuotaSpec successfully', function(done) {
        //uncomment below and update the code to test postQuotaSpec
        //instance.postQuotaSpec(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
