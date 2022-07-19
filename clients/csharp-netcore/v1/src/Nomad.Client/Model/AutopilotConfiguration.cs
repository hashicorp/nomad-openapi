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
    /// AutopilotConfiguration
    /// </summary>
    [DataContract(Name = "AutopilotConfiguration")]
    public partial class AutopilotConfiguration : IEquatable<AutopilotConfiguration>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="AutopilotConfiguration" /> class.
        /// </summary>
        /// <param name="cleanupDeadServers">cleanupDeadServers.</param>
        /// <param name="createIndex">createIndex.</param>
        /// <param name="disableUpgradeMigration">disableUpgradeMigration.</param>
        /// <param name="enableCustomUpgrades">enableCustomUpgrades.</param>
        /// <param name="enableRedundancyZones">enableRedundancyZones.</param>
        /// <param name="lastContactThreshold">lastContactThreshold.</param>
        /// <param name="maxTrailingLogs">maxTrailingLogs.</param>
        /// <param name="minQuorum">minQuorum.</param>
        /// <param name="modifyIndex">modifyIndex.</param>
        /// <param name="serverStabilizationTime">serverStabilizationTime.</param>
        public AutopilotConfiguration(bool cleanupDeadServers = default(bool), int createIndex = default(int), bool disableUpgradeMigration = default(bool), bool enableCustomUpgrades = default(bool), bool enableRedundancyZones = default(bool), string lastContactThreshold = default(string), int maxTrailingLogs = default(int), int minQuorum = default(int), int modifyIndex = default(int), string serverStabilizationTime = default(string))
        {
            this.CleanupDeadServers = cleanupDeadServers;
            this.CreateIndex = createIndex;
            this.DisableUpgradeMigration = disableUpgradeMigration;
            this.EnableCustomUpgrades = enableCustomUpgrades;
            this.EnableRedundancyZones = enableRedundancyZones;
            this.LastContactThreshold = lastContactThreshold;
            this.MaxTrailingLogs = maxTrailingLogs;
            this.MinQuorum = minQuorum;
            this.ModifyIndex = modifyIndex;
            this.ServerStabilizationTime = serverStabilizationTime;
        }

        /// <summary>
        /// Gets or Sets CleanupDeadServers
        /// </summary>
        [DataMember(Name = "CleanupDeadServers", EmitDefaultValue = true)]
        public bool CleanupDeadServers { get; set; }

        /// <summary>
        /// Gets or Sets CreateIndex
        /// </summary>
        [DataMember(Name = "CreateIndex", EmitDefaultValue = false)]
        public int CreateIndex { get; set; }

        /// <summary>
        /// Gets or Sets DisableUpgradeMigration
        /// </summary>
        [DataMember(Name = "DisableUpgradeMigration", EmitDefaultValue = true)]
        public bool DisableUpgradeMigration { get; set; }

        /// <summary>
        /// Gets or Sets EnableCustomUpgrades
        /// </summary>
        [DataMember(Name = "EnableCustomUpgrades", EmitDefaultValue = true)]
        public bool EnableCustomUpgrades { get; set; }

        /// <summary>
        /// Gets or Sets EnableRedundancyZones
        /// </summary>
        [DataMember(Name = "EnableRedundancyZones", EmitDefaultValue = true)]
        public bool EnableRedundancyZones { get; set; }

        /// <summary>
        /// Gets or Sets LastContactThreshold
        /// </summary>
        [DataMember(Name = "LastContactThreshold", EmitDefaultValue = false)]
        public string LastContactThreshold { get; set; }

        /// <summary>
        /// Gets or Sets MaxTrailingLogs
        /// </summary>
        [DataMember(Name = "MaxTrailingLogs", EmitDefaultValue = false)]
        public int MaxTrailingLogs { get; set; }

        /// <summary>
        /// Gets or Sets MinQuorum
        /// </summary>
        [DataMember(Name = "MinQuorum", EmitDefaultValue = false)]
        public int MinQuorum { get; set; }

        /// <summary>
        /// Gets or Sets ModifyIndex
        /// </summary>
        [DataMember(Name = "ModifyIndex", EmitDefaultValue = false)]
        public int ModifyIndex { get; set; }

        /// <summary>
        /// Gets or Sets ServerStabilizationTime
        /// </summary>
        [DataMember(Name = "ServerStabilizationTime", EmitDefaultValue = false)]
        public string ServerStabilizationTime { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class AutopilotConfiguration {\n");
            sb.Append("  CleanupDeadServers: ").Append(CleanupDeadServers).Append("\n");
            sb.Append("  CreateIndex: ").Append(CreateIndex).Append("\n");
            sb.Append("  DisableUpgradeMigration: ").Append(DisableUpgradeMigration).Append("\n");
            sb.Append("  EnableCustomUpgrades: ").Append(EnableCustomUpgrades).Append("\n");
            sb.Append("  EnableRedundancyZones: ").Append(EnableRedundancyZones).Append("\n");
            sb.Append("  LastContactThreshold: ").Append(LastContactThreshold).Append("\n");
            sb.Append("  MaxTrailingLogs: ").Append(MaxTrailingLogs).Append("\n");
            sb.Append("  MinQuorum: ").Append(MinQuorum).Append("\n");
            sb.Append("  ModifyIndex: ").Append(ModifyIndex).Append("\n");
            sb.Append("  ServerStabilizationTime: ").Append(ServerStabilizationTime).Append("\n");
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
            return this.Equals(input as AutopilotConfiguration);
        }

        /// <summary>
        /// Returns true if AutopilotConfiguration instances are equal
        /// </summary>
        /// <param name="input">Instance of AutopilotConfiguration to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(AutopilotConfiguration input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.CleanupDeadServers == input.CleanupDeadServers ||
                    this.CleanupDeadServers.Equals(input.CleanupDeadServers)
                ) && 
                (
                    this.CreateIndex == input.CreateIndex ||
                    this.CreateIndex.Equals(input.CreateIndex)
                ) && 
                (
                    this.DisableUpgradeMigration == input.DisableUpgradeMigration ||
                    this.DisableUpgradeMigration.Equals(input.DisableUpgradeMigration)
                ) && 
                (
                    this.EnableCustomUpgrades == input.EnableCustomUpgrades ||
                    this.EnableCustomUpgrades.Equals(input.EnableCustomUpgrades)
                ) && 
                (
                    this.EnableRedundancyZones == input.EnableRedundancyZones ||
                    this.EnableRedundancyZones.Equals(input.EnableRedundancyZones)
                ) && 
                (
                    this.LastContactThreshold == input.LastContactThreshold ||
                    (this.LastContactThreshold != null &&
                    this.LastContactThreshold.Equals(input.LastContactThreshold))
                ) && 
                (
                    this.MaxTrailingLogs == input.MaxTrailingLogs ||
                    this.MaxTrailingLogs.Equals(input.MaxTrailingLogs)
                ) && 
                (
                    this.MinQuorum == input.MinQuorum ||
                    this.MinQuorum.Equals(input.MinQuorum)
                ) && 
                (
                    this.ModifyIndex == input.ModifyIndex ||
                    this.ModifyIndex.Equals(input.ModifyIndex)
                ) && 
                (
                    this.ServerStabilizationTime == input.ServerStabilizationTime ||
                    (this.ServerStabilizationTime != null &&
                    this.ServerStabilizationTime.Equals(input.ServerStabilizationTime))
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
                hashCode = hashCode * 59 + this.CleanupDeadServers.GetHashCode();
                hashCode = hashCode * 59 + this.CreateIndex.GetHashCode();
                hashCode = hashCode * 59 + this.DisableUpgradeMigration.GetHashCode();
                hashCode = hashCode * 59 + this.EnableCustomUpgrades.GetHashCode();
                hashCode = hashCode * 59 + this.EnableRedundancyZones.GetHashCode();
                if (this.LastContactThreshold != null)
                    hashCode = hashCode * 59 + this.LastContactThreshold.GetHashCode();
                hashCode = hashCode * 59 + this.MaxTrailingLogs.GetHashCode();
                hashCode = hashCode * 59 + this.MinQuorum.GetHashCode();
                hashCode = hashCode * 59 + this.ModifyIndex.GetHashCode();
                if (this.ServerStabilizationTime != null)
                    hashCode = hashCode * 59 + this.ServerStabilizationTime.GetHashCode();
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
            // CreateIndex (int) maximum
            if(this.CreateIndex > (int)384)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for CreateIndex, must be a value less than or equal to 384.", new [] { "CreateIndex" });
            }

            // CreateIndex (int) minimum
            if(this.CreateIndex < (int)0)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for CreateIndex, must be a value greater than or equal to 0.", new [] { "CreateIndex" });
            }

            // MaxTrailingLogs (int) maximum
            if(this.MaxTrailingLogs > (int)384)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for MaxTrailingLogs, must be a value less than or equal to 384.", new [] { "MaxTrailingLogs" });
            }

            // MaxTrailingLogs (int) minimum
            if(this.MaxTrailingLogs < (int)0)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for MaxTrailingLogs, must be a value greater than or equal to 0.", new [] { "MaxTrailingLogs" });
            }

            // MinQuorum (int) minimum
            if(this.MinQuorum < (int)0)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for MinQuorum, must be a value greater than or equal to 0.", new [] { "MinQuorum" });
            }

            // ModifyIndex (int) maximum
            if(this.ModifyIndex > (int)384)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for ModifyIndex, must be a value less than or equal to 384.", new [] { "ModifyIndex" });
            }

            // ModifyIndex (int) minimum
            if(this.ModifyIndex < (int)0)
            {
                yield return new System.ComponentModel.DataAnnotations.ValidationResult("Invalid value for ModifyIndex, must be a value greater than or equal to 0.", new [] { "ModifyIndex" });
            }

            yield break;
        }
    }

}
