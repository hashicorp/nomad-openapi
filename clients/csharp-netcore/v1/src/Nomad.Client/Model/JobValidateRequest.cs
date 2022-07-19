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
    /// JobValidateRequest
    /// </summary>
    [DataContract(Name = "JobValidateRequest")]
    public partial class JobValidateRequest : IEquatable<JobValidateRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="JobValidateRequest" /> class.
        /// </summary>
        /// <param name="job">job.</param>
        /// <param name="_namespace">_namespace.</param>
        /// <param name="region">region.</param>
        /// <param name="secretID">secretID.</param>
        public JobValidateRequest(Job job = default(Job), string _namespace = default(string), string region = default(string), string secretID = default(string))
        {
            this.Job = job;
            this.Namespace = _namespace;
            this.Region = region;
            this.SecretID = secretID;
        }

        /// <summary>
        /// Gets or Sets Job
        /// </summary>
        [DataMember(Name = "Job", EmitDefaultValue = false)]
        public Job Job { get; set; }

        /// <summary>
        /// Gets or Sets Namespace
        /// </summary>
        [DataMember(Name = "Namespace", EmitDefaultValue = false)]
        public string Namespace { get; set; }

        /// <summary>
        /// Gets or Sets Region
        /// </summary>
        [DataMember(Name = "Region", EmitDefaultValue = false)]
        public string Region { get; set; }

        /// <summary>
        /// Gets or Sets SecretID
        /// </summary>
        [DataMember(Name = "SecretID", EmitDefaultValue = false)]
        public string SecretID { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class JobValidateRequest {\n");
            sb.Append("  Job: ").Append(Job).Append("\n");
            sb.Append("  Namespace: ").Append(Namespace).Append("\n");
            sb.Append("  Region: ").Append(Region).Append("\n");
            sb.Append("  SecretID: ").Append(SecretID).Append("\n");
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
            return this.Equals(input as JobValidateRequest);
        }

        /// <summary>
        /// Returns true if JobValidateRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of JobValidateRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(JobValidateRequest input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Job == input.Job ||
                    (this.Job != null &&
                    this.Job.Equals(input.Job))
                ) && 
                (
                    this.Namespace == input.Namespace ||
                    (this.Namespace != null &&
                    this.Namespace.Equals(input.Namespace))
                ) && 
                (
                    this.Region == input.Region ||
                    (this.Region != null &&
                    this.Region.Equals(input.Region))
                ) && 
                (
                    this.SecretID == input.SecretID ||
                    (this.SecretID != null &&
                    this.SecretID.Equals(input.SecretID))
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
                if (this.Job != null)
                    hashCode = hashCode * 59 + this.Job.GetHashCode();
                if (this.Namespace != null)
                    hashCode = hashCode * 59 + this.Namespace.GetHashCode();
                if (this.Region != null)
                    hashCode = hashCode * 59 + this.Region.GetHashCode();
                if (this.SecretID != null)
                    hashCode = hashCode * 59 + this.SecretID.GetHashCode();
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
