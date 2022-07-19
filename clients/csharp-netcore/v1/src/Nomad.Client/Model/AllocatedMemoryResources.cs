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
    /// AllocatedMemoryResources
    /// </summary>
    [DataContract(Name = "AllocatedMemoryResources")]
    public partial class AllocatedMemoryResources : IEquatable<AllocatedMemoryResources>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="AllocatedMemoryResources" /> class.
        /// </summary>
        /// <param name="memoryMB">memoryMB.</param>
        /// <param name="memoryMaxMB">memoryMaxMB.</param>
        public AllocatedMemoryResources(long memoryMB = default(long), long memoryMaxMB = default(long))
        {
            this.MemoryMB = memoryMB;
            this.MemoryMaxMB = memoryMaxMB;
        }

        /// <summary>
        /// Gets or Sets MemoryMB
        /// </summary>
        [DataMember(Name = "MemoryMB", EmitDefaultValue = false)]
        public long MemoryMB { get; set; }

        /// <summary>
        /// Gets or Sets MemoryMaxMB
        /// </summary>
        [DataMember(Name = "MemoryMaxMB", EmitDefaultValue = false)]
        public long MemoryMaxMB { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class AllocatedMemoryResources {\n");
            sb.Append("  MemoryMB: ").Append(MemoryMB).Append("\n");
            sb.Append("  MemoryMaxMB: ").Append(MemoryMaxMB).Append("\n");
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
            return this.Equals(input as AllocatedMemoryResources);
        }

        /// <summary>
        /// Returns true if AllocatedMemoryResources instances are equal
        /// </summary>
        /// <param name="input">Instance of AllocatedMemoryResources to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(AllocatedMemoryResources input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.MemoryMB == input.MemoryMB ||
                    this.MemoryMB.Equals(input.MemoryMB)
                ) && 
                (
                    this.MemoryMaxMB == input.MemoryMaxMB ||
                    this.MemoryMaxMB.Equals(input.MemoryMaxMB)
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
                hashCode = hashCode * 59 + this.MemoryMB.GetHashCode();
                hashCode = hashCode * 59 + this.MemoryMaxMB.GetHashCode();
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
