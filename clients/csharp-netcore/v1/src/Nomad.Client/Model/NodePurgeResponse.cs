/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Nomad.Client.Client.OpenAPIDateConverter;

namespace Nomad.Client.Model
{
    /// <summary>
    /// NodePurgeResponse
    /// </summary>
    [DataContract(Name = "NodePurgeResponse")]
    public partial class NodePurgeResponse : IEquatable<NodePurgeResponse>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="NodePurgeResponse" /> class.
        /// </summary>
        /// <param name="evalCreateIndex">evalCreateIndex.</param>
        /// <param name="evalIDs">evalIDs.</param>
        /// <param name="nodeModifyIndex">nodeModifyIndex.</param>
        public NodePurgeResponse(int evalCreateIndex = default(int), List<string> evalIDs = default(List<string>), int nodeModifyIndex = default(int))
        {
            this.EvalCreateIndex = evalCreateIndex;
            this.EvalIDs = evalIDs;
            this.NodeModifyIndex = nodeModifyIndex;
        }

        /// <summary>
        /// Gets or Sets EvalCreateIndex
        /// </summary>
        [DataMember(Name = "EvalCreateIndex", EmitDefaultValue = false)]
        public int EvalCreateIndex { get; set; }

        /// <summary>
        /// Gets or Sets EvalIDs
        /// </summary>
        [DataMember(Name = "EvalIDs", EmitDefaultValue = false)]
        public List<string> EvalIDs { get; set; }

        /// <summary>
        /// Gets or Sets NodeModifyIndex
        /// </summary>
        [DataMember(Name = "NodeModifyIndex", EmitDefaultValue = false)]
        public int NodeModifyIndex { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class NodePurgeResponse {\n");
            sb.Append("  EvalCreateIndex: ").Append(EvalCreateIndex).Append("\n");
            sb.Append("  EvalIDs: ").Append(EvalIDs).Append("\n");
            sb.Append("  NodeModifyIndex: ").Append(NodeModifyIndex).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as NodePurgeResponse);
        }

        /// <summary>
        /// Returns true if NodePurgeResponse instances are equal
        /// </summary>
        /// <param name="input">Instance of NodePurgeResponse to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(NodePurgeResponse input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.EvalCreateIndex == input.EvalCreateIndex ||
                    this.EvalCreateIndex.Equals(input.EvalCreateIndex)
                ) && 
                (
                    this.EvalIDs == input.EvalIDs ||
                    this.EvalIDs != null &&
                    input.EvalIDs != null &&
                    this.EvalIDs.SequenceEqual(input.EvalIDs)
                ) && 
                (
                    this.NodeModifyIndex == input.NodeModifyIndex ||
                    this.NodeModifyIndex.Equals(input.NodeModifyIndex)
                );
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                hashCode = hashCode * 59 + this.EvalCreateIndex.GetHashCode();
                if (this.EvalIDs != null)
                    hashCode = hashCode * 59 + this.EvalIDs.GetHashCode();
                hashCode = hashCode * 59 + this.NodeModifyIndex.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            // EvalCreateIndex (int) maximum
            if(this.EvalCreateIndex > (int)384)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for EvalCreateIndex, must be a value less than or equal to 384.", new [] { "EvalCreateIndex" });
            }

            // EvalCreateIndex (int) minimum
            if(this.EvalCreateIndex < (int)0)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for EvalCreateIndex, must be a value greater than or equal to 0.", new [] { "EvalCreateIndex" });
            }

            // NodeModifyIndex (int) maximum
            if(this.NodeModifyIndex > (int)384)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for NodeModifyIndex, must be a value less than or equal to 384.", new [] { "NodeModifyIndex" });
            }

            // NodeModifyIndex (int) minimum
            if(this.NodeModifyIndex < (int)0)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for NodeModifyIndex, must be a value greater than or equal to 0.", new [] { "NodeModifyIndex" });
            }

            yield break;
        }
    }

}
