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
    /// CSITopologyRequest
    /// </summary>
    [DataContract(Name = "CSITopologyRequest")]
    public partial class CSITopologyRequest : IEquatable<CSITopologyRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="CSITopologyRequest" /> class.
        /// </summary>
        /// <param name="preferred">preferred.</param>
        /// <param name="required">required.</param>
        public CSITopologyRequest(List<CSITopology> preferred = default(List<CSITopology>), List<CSITopology> required = default(List<CSITopology>))
        {
            this.Preferred = preferred;
            this.Required = required;
        }

        /// <summary>
        /// Gets or Sets Preferred
        /// </summary>
        [DataMember(Name = "Preferred", EmitDefaultValue = false)]
        public List<CSITopology> Preferred { get; set; }

        /// <summary>
        /// Gets or Sets Required
        /// </summary>
        [DataMember(Name = "Required", EmitDefaultValue = false)]
        public List<CSITopology> Required { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class CSITopologyRequest {\n");
            sb.Append("  Preferred: ").Append(Preferred).Append("\n");
            sb.Append("  Required: ").Append(Required).Append("\n");
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
            return this.Equals(input as CSITopologyRequest);
        }

        /// <summary>
        /// Returns true if CSITopologyRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of CSITopologyRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(CSITopologyRequest input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Preferred == input.Preferred ||
                    this.Preferred != null &&
                    input.Preferred != null &&
                    this.Preferred.SequenceEqual(input.Preferred)
                ) && 
                (
                    this.Required == input.Required ||
                    this.Required != null &&
                    input.Required != null &&
                    this.Required.SequenceEqual(input.Required)
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
                if (this.Preferred != null)
                    hashCode = hashCode * 59 + this.Preferred.GetHashCode();
                if (this.Required != null)
                    hashCode = hashCode * 59 + this.Required.GetHashCode();
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
            yield break;
        }
    }

}
