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
    /// AllocationMetric
    /// </summary>
    [DataContract(Name = "AllocationMetric")]
    public partial class AllocationMetric : IEquatable<AllocationMetric>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="AllocationMetric" /> class.
        /// </summary>
        /// <param name="allocationTime">allocationTime.</param>
        /// <param name="classExhausted">classExhausted.</param>
        /// <param name="classFiltered">classFiltered.</param>
        /// <param name="coalescedFailures">coalescedFailures.</param>
        /// <param name="constraintFiltered">constraintFiltered.</param>
        /// <param name="dimensionExhausted">dimensionExhausted.</param>
        /// <param name="nodesAvailable">nodesAvailable.</param>
        /// <param name="nodesEvaluated">nodesEvaluated.</param>
        /// <param name="nodesExhausted">nodesExhausted.</param>
        /// <param name="nodesFiltered">nodesFiltered.</param>
        /// <param name="quotaExhausted">quotaExhausted.</param>
        /// <param name="resourcesExhausted">resourcesExhausted.</param>
        /// <param name="scoreMetaData">scoreMetaData.</param>
        /// <param name="scores">scores.</param>
        public AllocationMetric(long allocationTime = default(long), Dictionary<string, int> classExhausted = default(Dictionary<string, int>), Dictionary<string, int> classFiltered = default(Dictionary<string, int>), int coalescedFailures = default(int), Dictionary<string, int> constraintFiltered = default(Dictionary<string, int>), Dictionary<string, int> dimensionExhausted = default(Dictionary<string, int>), Dictionary<string, int> nodesAvailable = default(Dictionary<string, int>), int nodesEvaluated = default(int), int nodesExhausted = default(int), int nodesFiltered = default(int), List<string> quotaExhausted = default(List<string>), Dictionary<string, Resources> resourcesExhausted = default(Dictionary<string, Resources>), List<NodeScoreMeta> scoreMetaData = default(List<NodeScoreMeta>), Dictionary<string, double> scores = default(Dictionary<string, double>))
        {
            this.AllocationTime = allocationTime;
            this.ClassExhausted = classExhausted;
            this.ClassFiltered = classFiltered;
            this.CoalescedFailures = coalescedFailures;
            this.ConstraintFiltered = constraintFiltered;
            this.DimensionExhausted = dimensionExhausted;
            this.NodesAvailable = nodesAvailable;
            this.NodesEvaluated = nodesEvaluated;
            this.NodesExhausted = nodesExhausted;
            this.NodesFiltered = nodesFiltered;
            this.QuotaExhausted = quotaExhausted;
            this.ResourcesExhausted = resourcesExhausted;
            this.ScoreMetaData = scoreMetaData;
            this.Scores = scores;
        }

        /// <summary>
        /// Gets or Sets AllocationTime
        /// </summary>
        [DataMember(Name = "AllocationTime", EmitDefaultValue = false)]
        public long AllocationTime { get; set; }

        /// <summary>
        /// Gets or Sets ClassExhausted
        /// </summary>
        [DataMember(Name = "ClassExhausted", EmitDefaultValue = false)]
        public Dictionary<string, int> ClassExhausted { get; set; }

        /// <summary>
        /// Gets or Sets ClassFiltered
        /// </summary>
        [DataMember(Name = "ClassFiltered", EmitDefaultValue = false)]
        public Dictionary<string, int> ClassFiltered { get; set; }

        /// <summary>
        /// Gets or Sets CoalescedFailures
        /// </summary>
        [DataMember(Name = "CoalescedFailures", EmitDefaultValue = false)]
        public int CoalescedFailures { get; set; }

        /// <summary>
        /// Gets or Sets ConstraintFiltered
        /// </summary>
        [DataMember(Name = "ConstraintFiltered", EmitDefaultValue = false)]
        public Dictionary<string, int> ConstraintFiltered { get; set; }

        /// <summary>
        /// Gets or Sets DimensionExhausted
        /// </summary>
        [DataMember(Name = "DimensionExhausted", EmitDefaultValue = false)]
        public Dictionary<string, int> DimensionExhausted { get; set; }

        /// <summary>
        /// Gets or Sets NodesAvailable
        /// </summary>
        [DataMember(Name = "NodesAvailable", EmitDefaultValue = false)]
        public Dictionary<string, int> NodesAvailable { get; set; }

        /// <summary>
        /// Gets or Sets NodesEvaluated
        /// </summary>
        [DataMember(Name = "NodesEvaluated", EmitDefaultValue = false)]
        public int NodesEvaluated { get; set; }

        /// <summary>
        /// Gets or Sets NodesExhausted
        /// </summary>
        [DataMember(Name = "NodesExhausted", EmitDefaultValue = false)]
        public int NodesExhausted { get; set; }

        /// <summary>
        /// Gets or Sets NodesFiltered
        /// </summary>
        [DataMember(Name = "NodesFiltered", EmitDefaultValue = false)]
        public int NodesFiltered { get; set; }

        /// <summary>
        /// Gets or Sets QuotaExhausted
        /// </summary>
        [DataMember(Name = "QuotaExhausted", EmitDefaultValue = false)]
        public List<string> QuotaExhausted { get; set; }

        /// <summary>
        /// Gets or Sets ResourcesExhausted
        /// </summary>
        [DataMember(Name = "ResourcesExhausted", EmitDefaultValue = false)]
        public Dictionary<string, Resources> ResourcesExhausted { get; set; }

        /// <summary>
        /// Gets or Sets ScoreMetaData
        /// </summary>
        [DataMember(Name = "ScoreMetaData", EmitDefaultValue = false)]
        public List<NodeScoreMeta> ScoreMetaData { get; set; }

        /// <summary>
        /// Gets or Sets Scores
        /// </summary>
        [DataMember(Name = "Scores", EmitDefaultValue = false)]
        public Dictionary<string, double> Scores { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class AllocationMetric {\n");
            sb.Append("  AllocationTime: ").Append(AllocationTime).Append("\n");
            sb.Append("  ClassExhausted: ").Append(ClassExhausted).Append("\n");
            sb.Append("  ClassFiltered: ").Append(ClassFiltered).Append("\n");
            sb.Append("  CoalescedFailures: ").Append(CoalescedFailures).Append("\n");
            sb.Append("  ConstraintFiltered: ").Append(ConstraintFiltered).Append("\n");
            sb.Append("  DimensionExhausted: ").Append(DimensionExhausted).Append("\n");
            sb.Append("  NodesAvailable: ").Append(NodesAvailable).Append("\n");
            sb.Append("  NodesEvaluated: ").Append(NodesEvaluated).Append("\n");
            sb.Append("  NodesExhausted: ").Append(NodesExhausted).Append("\n");
            sb.Append("  NodesFiltered: ").Append(NodesFiltered).Append("\n");
            sb.Append("  QuotaExhausted: ").Append(QuotaExhausted).Append("\n");
            sb.Append("  ResourcesExhausted: ").Append(ResourcesExhausted).Append("\n");
            sb.Append("  ScoreMetaData: ").Append(ScoreMetaData).Append("\n");
            sb.Append("  Scores: ").Append(Scores).Append("\n");
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
            return this.Equals(input as AllocationMetric);
        }

        /// <summary>
        /// Returns true if AllocationMetric instances are equal
        /// </summary>
        /// <param name="input">Instance of AllocationMetric to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(AllocationMetric input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.AllocationTime == input.AllocationTime ||
                    this.AllocationTime.Equals(input.AllocationTime)
                ) && 
                (
                    this.ClassExhausted == input.ClassExhausted ||
                    this.ClassExhausted != null &&
                    input.ClassExhausted != null &&
                    this.ClassExhausted.SequenceEqual(input.ClassExhausted)
                ) && 
                (
                    this.ClassFiltered == input.ClassFiltered ||
                    this.ClassFiltered != null &&
                    input.ClassFiltered != null &&
                    this.ClassFiltered.SequenceEqual(input.ClassFiltered)
                ) && 
                (
                    this.CoalescedFailures == input.CoalescedFailures ||
                    this.CoalescedFailures.Equals(input.CoalescedFailures)
                ) && 
                (
                    this.ConstraintFiltered == input.ConstraintFiltered ||
                    this.ConstraintFiltered != null &&
                    input.ConstraintFiltered != null &&
                    this.ConstraintFiltered.SequenceEqual(input.ConstraintFiltered)
                ) && 
                (
                    this.DimensionExhausted == input.DimensionExhausted ||
                    this.DimensionExhausted != null &&
                    input.DimensionExhausted != null &&
                    this.DimensionExhausted.SequenceEqual(input.DimensionExhausted)
                ) && 
                (
                    this.NodesAvailable == input.NodesAvailable ||
                    this.NodesAvailable != null &&
                    input.NodesAvailable != null &&
                    this.NodesAvailable.SequenceEqual(input.NodesAvailable)
                ) && 
                (
                    this.NodesEvaluated == input.NodesEvaluated ||
                    this.NodesEvaluated.Equals(input.NodesEvaluated)
                ) && 
                (
                    this.NodesExhausted == input.NodesExhausted ||
                    this.NodesExhausted.Equals(input.NodesExhausted)
                ) && 
                (
                    this.NodesFiltered == input.NodesFiltered ||
                    this.NodesFiltered.Equals(input.NodesFiltered)
                ) && 
                (
                    this.QuotaExhausted == input.QuotaExhausted ||
                    this.QuotaExhausted != null &&
                    input.QuotaExhausted != null &&
                    this.QuotaExhausted.SequenceEqual(input.QuotaExhausted)
                ) && 
                (
                    this.ResourcesExhausted == input.ResourcesExhausted ||
                    this.ResourcesExhausted != null &&
                    input.ResourcesExhausted != null &&
                    this.ResourcesExhausted.SequenceEqual(input.ResourcesExhausted)
                ) && 
                (
                    this.ScoreMetaData == input.ScoreMetaData ||
                    this.ScoreMetaData != null &&
                    input.ScoreMetaData != null &&
                    this.ScoreMetaData.SequenceEqual(input.ScoreMetaData)
                ) && 
                (
                    this.Scores == input.Scores ||
                    this.Scores != null &&
                    input.Scores != null &&
                    this.Scores.SequenceEqual(input.Scores)
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
                hashCode = hashCode * 59 + this.AllocationTime.GetHashCode();
                if (this.ClassExhausted != null)
                    hashCode = hashCode * 59 + this.ClassExhausted.GetHashCode();
                if (this.ClassFiltered != null)
                    hashCode = hashCode * 59 + this.ClassFiltered.GetHashCode();
                hashCode = hashCode * 59 + this.CoalescedFailures.GetHashCode();
                if (this.ConstraintFiltered != null)
                    hashCode = hashCode * 59 + this.ConstraintFiltered.GetHashCode();
                if (this.DimensionExhausted != null)
                    hashCode = hashCode * 59 + this.DimensionExhausted.GetHashCode();
                if (this.NodesAvailable != null)
                    hashCode = hashCode * 59 + this.NodesAvailable.GetHashCode();
                hashCode = hashCode * 59 + this.NodesEvaluated.GetHashCode();
                hashCode = hashCode * 59 + this.NodesExhausted.GetHashCode();
                hashCode = hashCode * 59 + this.NodesFiltered.GetHashCode();
                if (this.QuotaExhausted != null)
                    hashCode = hashCode * 59 + this.QuotaExhausted.GetHashCode();
                if (this.ResourcesExhausted != null)
                    hashCode = hashCode * 59 + this.ResourcesExhausted.GetHashCode();
                if (this.ScoreMetaData != null)
                    hashCode = hashCode * 59 + this.ScoreMetaData.GetHashCode();
                if (this.Scores != null)
                    hashCode = hashCode * 59 + this.Scores.GetHashCode();
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
