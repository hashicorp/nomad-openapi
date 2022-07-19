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
    /// DeploymentAllocHealthRequest
    /// </summary>
    [DataContract(Name = "DeploymentAllocHealthRequest")]
    public partial class DeploymentAllocHealthRequest : IEquatable<DeploymentAllocHealthRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="DeploymentAllocHealthRequest" /> class.
        /// </summary>
        /// <param name="deploymentID">deploymentID.</param>
        /// <param name="healthyAllocationIDs">healthyAllocationIDs.</param>
        /// <param name="_namespace">_namespace.</param>
        /// <param name="region">region.</param>
        /// <param name="secretID">secretID.</param>
        /// <param name="unhealthyAllocationIDs">unhealthyAllocationIDs.</param>
        public DeploymentAllocHealthRequest(string deploymentID = default(string), List<string> healthyAllocationIDs = default(List<string>), string _namespace = default(string), string region = default(string), string secretID = default(string), List<string> unhealthyAllocationIDs = default(List<string>))
        {
            this.DeploymentID = deploymentID;
            this.HealthyAllocationIDs = healthyAllocationIDs;
            this.Namespace = _namespace;
            this.Region = region;
            this.SecretID = secretID;
            this.UnhealthyAllocationIDs = unhealthyAllocationIDs;
        }

        /// <summary>
        /// Gets or Sets DeploymentID
        /// </summary>
        [DataMember(Name = "DeploymentID", EmitDefaultValue = false)]
        public string DeploymentID { get; set; }

        /// <summary>
        /// Gets or Sets HealthyAllocationIDs
        /// </summary>
        [DataMember(Name = "HealthyAllocationIDs", EmitDefaultValue = false)]
        public List<string> HealthyAllocationIDs { get; set; }

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
        /// Gets or Sets UnhealthyAllocationIDs
        /// </summary>
        [DataMember(Name = "UnhealthyAllocationIDs", EmitDefaultValue = false)]
        public List<string> UnhealthyAllocationIDs { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class DeploymentAllocHealthRequest {\n");
            sb.Append("  DeploymentID: ").Append(DeploymentID).Append("\n");
            sb.Append("  HealthyAllocationIDs: ").Append(HealthyAllocationIDs).Append("\n");
            sb.Append("  Namespace: ").Append(Namespace).Append("\n");
            sb.Append("  Region: ").Append(Region).Append("\n");
            sb.Append("  SecretID: ").Append(SecretID).Append("\n");
            sb.Append("  UnhealthyAllocationIDs: ").Append(UnhealthyAllocationIDs).Append("\n");
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
            return this.Equals(input as DeploymentAllocHealthRequest);
        }

        /// <summary>
        /// Returns true if DeploymentAllocHealthRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of DeploymentAllocHealthRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(DeploymentAllocHealthRequest input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.DeploymentID == input.DeploymentID ||
                    (this.DeploymentID != null &&
                    this.DeploymentID.Equals(input.DeploymentID))
                ) && 
                (
                    this.HealthyAllocationIDs == input.HealthyAllocationIDs ||
                    this.HealthyAllocationIDs != null &&
                    input.HealthyAllocationIDs != null &&
                    this.HealthyAllocationIDs.SequenceEqual(input.HealthyAllocationIDs)
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
                ) && 
                (
                    this.UnhealthyAllocationIDs == input.UnhealthyAllocationIDs ||
                    this.UnhealthyAllocationIDs != null &&
                    input.UnhealthyAllocationIDs != null &&
                    this.UnhealthyAllocationIDs.SequenceEqual(input.UnhealthyAllocationIDs)
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
                if (this.DeploymentID != null)
                    hashCode = hashCode * 59 + this.DeploymentID.GetHashCode();
                if (this.HealthyAllocationIDs != null)
                    hashCode = hashCode * 59 + this.HealthyAllocationIDs.GetHashCode();
                if (this.Namespace != null)
                    hashCode = hashCode * 59 + this.Namespace.GetHashCode();
                if (this.Region != null)
                    hashCode = hashCode * 59 + this.Region.GetHashCode();
                if (this.SecretID != null)
                    hashCode = hashCode * 59 + this.SecretID.GetHashCode();
                if (this.UnhealthyAllocationIDs != null)
                    hashCode = hashCode * 59 + this.UnhealthyAllocationIDs.GetHashCode();
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
