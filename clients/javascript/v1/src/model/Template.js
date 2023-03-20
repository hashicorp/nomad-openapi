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

import ApiClient from '../ApiClient';
import ChangeScript from './ChangeScript';
import WaitConfig from './WaitConfig';

/**
 * The Template model module.
 * @module model/Template
 * @version 1.1.4
 */
class Template {
    /**
     * Constructs a new <code>Template</code>.
     * @alias module:model/Template
     */
    constructor() { 
        
        Template.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Template</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Template} obj Optional instance to populate.
     * @return {module:model/Template} The populated <code>Template</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Template();

            if (data.hasOwnProperty('ChangeMode')) {
                obj['ChangeMode'] = ApiClient.convertToType(data['ChangeMode'], 'String');
            }
            if (data.hasOwnProperty('ChangeScript')) {
                obj['ChangeScript'] = ChangeScript.constructFromObject(data['ChangeScript']);
            }
            if (data.hasOwnProperty('ChangeSignal')) {
                obj['ChangeSignal'] = ApiClient.convertToType(data['ChangeSignal'], 'String');
            }
            if (data.hasOwnProperty('DestPath')) {
                obj['DestPath'] = ApiClient.convertToType(data['DestPath'], 'String');
            }
            if (data.hasOwnProperty('EmbeddedTmpl')) {
                obj['EmbeddedTmpl'] = ApiClient.convertToType(data['EmbeddedTmpl'], 'String');
            }
            if (data.hasOwnProperty('Envvars')) {
                obj['Envvars'] = ApiClient.convertToType(data['Envvars'], 'Boolean');
            }
            if (data.hasOwnProperty('ErrMissingKey')) {
                obj['ErrMissingKey'] = ApiClient.convertToType(data['ErrMissingKey'], 'Boolean');
            }
            if (data.hasOwnProperty('Gid')) {
                obj['Gid'] = ApiClient.convertToType(data['Gid'], 'Number');
            }
            if (data.hasOwnProperty('LeftDelim')) {
                obj['LeftDelim'] = ApiClient.convertToType(data['LeftDelim'], 'String');
            }
            if (data.hasOwnProperty('Perms')) {
                obj['Perms'] = ApiClient.convertToType(data['Perms'], 'String');
            }
            if (data.hasOwnProperty('RightDelim')) {
                obj['RightDelim'] = ApiClient.convertToType(data['RightDelim'], 'String');
            }
            if (data.hasOwnProperty('SourcePath')) {
                obj['SourcePath'] = ApiClient.convertToType(data['SourcePath'], 'String');
            }
            if (data.hasOwnProperty('Splay')) {
                obj['Splay'] = ApiClient.convertToType(data['Splay'], 'Number');
            }
            if (data.hasOwnProperty('Uid')) {
                obj['Uid'] = ApiClient.convertToType(data['Uid'], 'Number');
            }
            if (data.hasOwnProperty('VaultGrace')) {
                obj['VaultGrace'] = ApiClient.convertToType(data['VaultGrace'], 'Number');
            }
            if (data.hasOwnProperty('Wait')) {
                obj['Wait'] = WaitConfig.constructFromObject(data['Wait']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} ChangeMode
 */
Template.prototype['ChangeMode'] = undefined;

/**
 * @member {module:model/ChangeScript} ChangeScript
 */
Template.prototype['ChangeScript'] = undefined;

/**
 * @member {String} ChangeSignal
 */
Template.prototype['ChangeSignal'] = undefined;

/**
 * @member {String} DestPath
 */
Template.prototype['DestPath'] = undefined;

/**
 * @member {String} EmbeddedTmpl
 */
Template.prototype['EmbeddedTmpl'] = undefined;

/**
 * @member {Boolean} Envvars
 */
Template.prototype['Envvars'] = undefined;

/**
 * @member {Boolean} ErrMissingKey
 */
Template.prototype['ErrMissingKey'] = undefined;

/**
 * @member {Number} Gid
 */
Template.prototype['Gid'] = undefined;

/**
 * @member {String} LeftDelim
 */
Template.prototype['LeftDelim'] = undefined;

/**
 * @member {String} Perms
 */
Template.prototype['Perms'] = undefined;

/**
 * @member {String} RightDelim
 */
Template.prototype['RightDelim'] = undefined;

/**
 * @member {String} SourcePath
 */
Template.prototype['SourcePath'] = undefined;

/**
 * @member {Number} Splay
 */
Template.prototype['Splay'] = undefined;

/**
 * @member {Number} Uid
 */
Template.prototype['Uid'] = undefined;

/**
 * @member {Number} VaultGrace
 */
Template.prototype['VaultGrace'] = undefined;

/**
 * @member {module:model/WaitConfig} Wait
 */
Template.prototype['Wait'] = undefined;






export default Template;

