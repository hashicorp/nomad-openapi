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
    /// EvalOptions
    /// </summary>
    [DataContract(Name = "EvalOptions")]
    public partial class EvalOptions : IEquatable<EvalOptions>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="EvalOptions" /> class.
        /// </summary>
        /// <param name="forceReschedule">forceReschedule.</param>
        public EvalOptions(bool forceReschedule = default(bool))
        {
            this.ForceReschedule = forceReschedule;
        }

        /// <summary>
        /// Gets or Sets ForceReschedule
        /// </summary>
        [DataMember(Name = "ForceReschedule", EmitDefaultValue = true)]
        public bool ForceReschedule { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class EvalOptions {\n");
            sb.Append("  ForceReschedule: ").Append(ForceReschedule).Append("\n");
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
            return this.Equals(input as EvalOptions);
        }

        /// <summary>
        /// Returns true if EvalOptions instances are equal
        /// </summary>
        /// <param name="input">Instance of EvalOptions to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(EvalOptions input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.ForceReschedule == input.ForceReschedule ||
                    this.ForceReschedule.Equals(input.ForceReschedule)
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
                hashCode = hashCode * 59 + this.ForceReschedule.GetHashCode();
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
