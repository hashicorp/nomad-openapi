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
    /// CSIVolumeListStub
    /// </summary>
    [DataContract(Name = "CSIVolumeListStub")]
    public partial class CSIVolumeListStub : IEquatable<CSIVolumeListStub>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="CSIVolumeListStub" /> class.
        /// </summary>
        /// <param name="accessMode">accessMode.</param>
        /// <param name="attachmentMode">attachmentMode.</param>
        /// <param name="controllerRequired">controllerRequired.</param>
        /// <param name="controllersExpected">controllersExpected.</param>
        /// <param name="controllersHealthy">controllersHealthy.</param>
        /// <param name="createIndex">createIndex.</param>
        /// <param name="externalID">externalID.</param>
        /// <param name="iD">iD.</param>
        /// <param name="modifyIndex">modifyIndex.</param>
        /// <param name="name">name.</param>
        /// <param name="_namespace">_namespace.</param>
        /// <param name="nodesExpected">nodesExpected.</param>
        /// <param name="nodesHealthy">nodesHealthy.</param>
        /// <param name="pluginID">pluginID.</param>
        /// <param name="provider">provider.</param>
        /// <param name="resourceExhausted">resourceExhausted.</param>
        /// <param name="schedulable">schedulable.</param>
        /// <param name="topologies">topologies.</param>
        public CSIVolumeListStub(string accessMode = default(string), string attachmentMode = default(string), bool controllerRequired = default(bool), int controllersExpected = default(int), int controllersHealthy = default(int), int createIndex = default(int), string externalID = default(string), string iD = default(string), int modifyIndex = default(int), string name = default(string), string _namespace = default(string), int nodesExpected = default(int), int nodesHealthy = default(int), string pluginID = default(string), string provider = default(string), DateTime resourceExhausted = default(DateTime), bool schedulable = default(bool), List<CSITopology> topologies = default(List<CSITopology>))
        {
            this.AccessMode = accessMode;
            this.AttachmentMode = attachmentMode;
            this.ControllerRequired = controllerRequired;
            this.ControllersExpected = controllersExpected;
            this.ControllersHealthy = controllersHealthy;
            this.CreateIndex = createIndex;
            this.ExternalID = externalID;
            this.ID = iD;
            this.ModifyIndex = modifyIndex;
            this.Name = name;
            this.Namespace = _namespace;
            this.NodesExpected = nodesExpected;
            this.NodesHealthy = nodesHealthy;
            this.PluginID = pluginID;
            this.Provider = provider;
            this.ResourceExhausted = resourceExhausted;
            this.Schedulable = schedulable;
            this.Topologies = topologies;
        }

        /// <summary>
        /// Gets or Sets AccessMode
        /// </summary>
        [DataMember(Name = "AccessMode", EmitDefaultValue = false)]
        public string AccessMode { get; set; }

        /// <summary>
        /// Gets or Sets AttachmentMode
        /// </summary>
        [DataMember(Name = "AttachmentMode", EmitDefaultValue = false)]
        public string AttachmentMode { get; set; }

        /// <summary>
        /// Gets or Sets ControllerRequired
        /// </summary>
        [DataMember(Name = "ControllerRequired", EmitDefaultValue = true)]
        public bool ControllerRequired { get; set; }

        /// <summary>
        /// Gets or Sets ControllersExpected
        /// </summary>
        [DataMember(Name = "ControllersExpected", EmitDefaultValue = false)]
        public int ControllersExpected { get; set; }

        /// <summary>
        /// Gets or Sets ControllersHealthy
        /// </summary>
        [DataMember(Name = "ControllersHealthy", EmitDefaultValue = false)]
        public int ControllersHealthy { get; set; }

        /// <summary>
        /// Gets or Sets CreateIndex
        /// </summary>
        [DataMember(Name = "CreateIndex", EmitDefaultValue = false)]
        public int CreateIndex { get; set; }

        /// <summary>
        /// Gets or Sets ExternalID
        /// </summary>
        [DataMember(Name = "ExternalID", EmitDefaultValue = false)]
        public string ExternalID { get; set; }

        /// <summary>
        /// Gets or Sets ID
        /// </summary>
        [DataMember(Name = "ID", EmitDefaultValue = false)]
        public string ID { get; set; }

        /// <summary>
        /// Gets or Sets ModifyIndex
        /// </summary>
        [DataMember(Name = "ModifyIndex", EmitDefaultValue = false)]
        public int ModifyIndex { get; set; }

        /// <summary>
        /// Gets or Sets Name
        /// </summary>
        [DataMember(Name = "Name", EmitDefaultValue = false)]
        public string Name { get; set; }

        /// <summary>
        /// Gets or Sets Namespace
        /// </summary>
        [DataMember(Name = "Namespace", EmitDefaultValue = false)]
        public string Namespace { get; set; }

        /// <summary>
        /// Gets or Sets NodesExpected
        /// </summary>
        [DataMember(Name = "NodesExpected", EmitDefaultValue = false)]
        public int NodesExpected { get; set; }

        /// <summary>
        /// Gets or Sets NodesHealthy
        /// </summary>
        [DataMember(Name = "NodesHealthy", EmitDefaultValue = false)]
        public int NodesHealthy { get; set; }

        /// <summary>
        /// Gets or Sets PluginID
        /// </summary>
        [DataMember(Name = "PluginID", EmitDefaultValue = false)]
        public string PluginID { get; set; }

        /// <summary>
        /// Gets or Sets Provider
        /// </summary>
        [DataMember(Name = "Provider", EmitDefaultValue = false)]
        public string Provider { get; set; }

        /// <summary>
        /// Gets or Sets ResourceExhausted
        /// </summary>
        [DataMember(Name = "ResourceExhausted", EmitDefaultValue = false)]
        public DateTime ResourceExhausted { get; set; }

        /// <summary>
        /// Gets or Sets Schedulable
        /// </summary>
        [DataMember(Name = "Schedulable", EmitDefaultValue = true)]
        public bool Schedulable { get; set; }

        /// <summary>
        /// Gets or Sets Topologies
        /// </summary>
        [DataMember(Name = "Topologies", EmitDefaultValue = false)]
        public List<CSITopology> Topologies { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class CSIVolumeListStub {\n");
            sb.Append("  AccessMode: ").Append(AccessMode).Append("\n");
            sb.Append("  AttachmentMode: ").Append(AttachmentMode).Append("\n");
            sb.Append("  ControllerRequired: ").Append(ControllerRequired).Append("\n");
            sb.Append("  ControllersExpected: ").Append(ControllersExpected).Append("\n");
            sb.Append("  ControllersHealthy: ").Append(ControllersHealthy).Append("\n");
            sb.Append("  CreateIndex: ").Append(CreateIndex).Append("\n");
            sb.Append("  ExternalID: ").Append(ExternalID).Append("\n");
            sb.Append("  ID: ").Append(ID).Append("\n");
            sb.Append("  ModifyIndex: ").Append(ModifyIndex).Append("\n");
            sb.Append("  Name: ").Append(Name).Append("\n");
            sb.Append("  Namespace: ").Append(Namespace).Append("\n");
            sb.Append("  NodesExpected: ").Append(NodesExpected).Append("\n");
            sb.Append("  NodesHealthy: ").Append(NodesHealthy).Append("\n");
            sb.Append("  PluginID: ").Append(PluginID).Append("\n");
            sb.Append("  Provider: ").Append(Provider).Append("\n");
            sb.Append("  ResourceExhausted: ").Append(ResourceExhausted).Append("\n");
            sb.Append("  Schedulable: ").Append(Schedulable).Append("\n");
            sb.Append("  Topologies: ").Append(Topologies).Append("\n");
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
            return this.Equals(input as CSIVolumeListStub);
        }

        /// <summary>
        /// Returns true if CSIVolumeListStub instances are equal
        /// </summary>
        /// <param name="input">Instance of CSIVolumeListStub to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(CSIVolumeListStub input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.AccessMode == input.AccessMode ||
                    (this.AccessMode != null &&
                    this.AccessMode.Equals(input.AccessMode))
                ) && 
                (
                    this.AttachmentMode == input.AttachmentMode ||
                    (this.AttachmentMode != null &&
                    this.AttachmentMode.Equals(input.AttachmentMode))
                ) && 
                (
                    this.ControllerRequired == input.ControllerRequired ||
                    this.ControllerRequired.Equals(input.ControllerRequired)
                ) && 
                (
                    this.ControllersExpected == input.ControllersExpected ||
                    this.ControllersExpected.Equals(input.ControllersExpected)
                ) && 
                (
                    this.ControllersHealthy == input.ControllersHealthy ||
                    this.ControllersHealthy.Equals(input.ControllersHealthy)
                ) && 
                (
                    this.CreateIndex == input.CreateIndex ||
                    this.CreateIndex.Equals(input.CreateIndex)
                ) && 
                (
                    this.ExternalID == input.ExternalID ||
                    (this.ExternalID != null &&
                    this.ExternalID.Equals(input.ExternalID))
                ) && 
                (
                    this.ID == input.ID ||
                    (this.ID != null &&
                    this.ID.Equals(input.ID))
                ) && 
                (
                    this.ModifyIndex == input.ModifyIndex ||
                    this.ModifyIndex.Equals(input.ModifyIndex)
                ) && 
                (
                    this.Name == input.Name ||
                    (this.Name != null &&
                    this.Name.Equals(input.Name))
                ) && 
                (
                    this.Namespace == input.Namespace ||
                    (this.Namespace != null &&
                    this.Namespace.Equals(input.Namespace))
                ) && 
                (
                    this.NodesExpected == input.NodesExpected ||
                    this.NodesExpected.Equals(input.NodesExpected)
                ) && 
                (
                    this.NodesHealthy == input.NodesHealthy ||
                    this.NodesHealthy.Equals(input.NodesHealthy)
                ) && 
                (
                    this.PluginID == input.PluginID ||
                    (this.PluginID != null &&
                    this.PluginID.Equals(input.PluginID))
                ) && 
                (
                    this.Provider == input.Provider ||
                    (this.Provider != null &&
                    this.Provider.Equals(input.Provider))
                ) && 
                (
                    this.ResourceExhausted == input.ResourceExhausted ||
                    (this.ResourceExhausted != null &&
                    this.ResourceExhausted.Equals(input.ResourceExhausted))
                ) && 
                (
                    this.Schedulable == input.Schedulable ||
                    this.Schedulable.Equals(input.Schedulable)
                ) && 
                (
                    this.Topologies == input.Topologies ||
                    this.Topologies != null &&
                    input.Topologies != null &&
                    this.Topologies.SequenceEqual(input.Topologies)
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
                if (this.AccessMode != null)
                    hashCode = hashCode * 59 + this.AccessMode.GetHashCode();
                if (this.AttachmentMode != null)
                    hashCode = hashCode * 59 + this.AttachmentMode.GetHashCode();
                hashCode = hashCode * 59 + this.ControllerRequired.GetHashCode();
                hashCode = hashCode * 59 + this.ControllersExpected.GetHashCode();
                hashCode = hashCode * 59 + this.ControllersHealthy.GetHashCode();
                hashCode = hashCode * 59 + this.CreateIndex.GetHashCode();
                if (this.ExternalID != null)
                    hashCode = hashCode * 59 + this.ExternalID.GetHashCode();
                if (this.ID != null)
                    hashCode = hashCode * 59 + this.ID.GetHashCode();
                hashCode = hashCode * 59 + this.ModifyIndex.GetHashCode();
                if (this.Name != null)
                    hashCode = hashCode * 59 + this.Name.GetHashCode();
                if (this.Namespace != null)
                    hashCode = hashCode * 59 + this.Namespace.GetHashCode();
                hashCode = hashCode * 59 + this.NodesExpected.GetHashCode();
                hashCode = hashCode * 59 + this.NodesHealthy.GetHashCode();
                if (this.PluginID != null)
                    hashCode = hashCode * 59 + this.PluginID.GetHashCode();
                if (this.Provider != null)
                    hashCode = hashCode * 59 + this.Provider.GetHashCode();
                if (this.ResourceExhausted != null)
                    hashCode = hashCode * 59 + this.ResourceExhausted.GetHashCode();
                hashCode = hashCode * 59 + this.Schedulable.GetHashCode();
                if (this.Topologies != null)
                    hashCode = hashCode * 59 + this.Topologies.GetHashCode();
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
