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
    /// JobDiff
    /// </summary>
    [DataContract(Name = "JobDiff")]
    public partial class JobDiff : IEquatable<JobDiff>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="JobDiff" /> class.
        /// </summary>
        /// <param name="fields">fields.</param>
        /// <param name="iD">iD.</param>
        /// <param name="objects">objects.</param>
        /// <param name="taskGroups">taskGroups.</param>
        /// <param name="type">type.</param>
        public JobDiff(List<FieldDiff> fields = default(List<FieldDiff>), string iD = default(string), List<ObjectDiff> objects = default(List<ObjectDiff>), List<TaskGroupDiff> taskGroups = default(List<TaskGroupDiff>), string type = default(string))
        {
            this.Fields = fields;
            this.ID = iD;
            this.Objects = objects;
            this.TaskGroups = taskGroups;
            this.Type = type;
        }

        /// <summary>
        /// Gets or Sets Fields
        /// </summary>
        [DataMember(Name = "Fields", EmitDefaultValue = false)]
        public List<FieldDiff> Fields { get; set; }

        /// <summary>
        /// Gets or Sets ID
        /// </summary>
        [DataMember(Name = "ID", EmitDefaultValue = false)]
        public string ID { get; set; }

        /// <summary>
        /// Gets or Sets Objects
        /// </summary>
        [DataMember(Name = "Objects", EmitDefaultValue = false)]
        public List<ObjectDiff> Objects { get; set; }

        /// <summary>
        /// Gets or Sets TaskGroups
        /// </summary>
        [DataMember(Name = "TaskGroups", EmitDefaultValue = false)]
        public List<TaskGroupDiff> TaskGroups { get; set; }

        /// <summary>
        /// Gets or Sets Type
        /// </summary>
        [DataMember(Name = "Type", EmitDefaultValue = false)]
        public string Type { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class JobDiff {\n");
            sb.Append("  Fields: ").Append(Fields).Append("\n");
            sb.Append("  ID: ").Append(ID).Append("\n");
            sb.Append("  Objects: ").Append(Objects).Append("\n");
            sb.Append("  TaskGroups: ").Append(TaskGroups).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
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
            return this.Equals(input as JobDiff);
        }

        /// <summary>
        /// Returns true if JobDiff instances are equal
        /// </summary>
        /// <param name="input">Instance of JobDiff to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(JobDiff input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Fields == input.Fields ||
                    this.Fields != null &&
                    input.Fields != null &&
                    this.Fields.SequenceEqual(input.Fields)
                ) && 
                (
                    this.ID == input.ID ||
                    (this.ID != null &&
                    this.ID.Equals(input.ID))
                ) && 
                (
                    this.Objects == input.Objects ||
                    this.Objects != null &&
                    input.Objects != null &&
                    this.Objects.SequenceEqual(input.Objects)
                ) && 
                (
                    this.TaskGroups == input.TaskGroups ||
                    this.TaskGroups != null &&
                    input.TaskGroups != null &&
                    this.TaskGroups.SequenceEqual(input.TaskGroups)
                ) && 
                (
                    this.Type == input.Type ||
                    (this.Type != null &&
                    this.Type.Equals(input.Type))
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
                if (this.Fields != null)
                    hashCode = hashCode * 59 + this.Fields.GetHashCode();
                if (this.ID != null)
                    hashCode = hashCode * 59 + this.ID.GetHashCode();
                if (this.Objects != null)
                    hashCode = hashCode * 59 + this.Objects.GetHashCode();
                if (this.TaskGroups != null)
                    hashCode = hashCode * 59 + this.TaskGroups.GetHashCode();
                if (this.Type != null)
                    hashCode = hashCode * 59 + this.Type.GetHashCode();
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
