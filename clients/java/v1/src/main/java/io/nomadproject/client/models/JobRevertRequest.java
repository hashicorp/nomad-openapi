/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

/*
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.nomadproject.client.models;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * JobRevertRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class JobRevertRequest {
  public static final String SERIALIZED_NAME_CONSUL_TOKEN = "ConsulToken";
  @SerializedName(SERIALIZED_NAME_CONSUL_TOKEN)
  private String consulToken;

  public static final String SERIALIZED_NAME_ENFORCE_PRIOR_VERSION = "EnforcePriorVersion";
  @SerializedName(SERIALIZED_NAME_ENFORCE_PRIOR_VERSION)
  private Integer enforcePriorVersion;

  public static final String SERIALIZED_NAME_JOB_I_D = "JobID";
  @SerializedName(SERIALIZED_NAME_JOB_I_D)
  private String jobID;

  public static final String SERIALIZED_NAME_JOB_VERSION = "JobVersion";
  @SerializedName(SERIALIZED_NAME_JOB_VERSION)
  private Integer jobVersion;

  public static final String SERIALIZED_NAME_NAMESPACE = "Namespace";
  @SerializedName(SERIALIZED_NAME_NAMESPACE)
  private String namespace;

  public static final String SERIALIZED_NAME_REGION = "Region";
  @SerializedName(SERIALIZED_NAME_REGION)
  private String region;

  public static final String SERIALIZED_NAME_SECRET_I_D = "SecretID";
  @SerializedName(SERIALIZED_NAME_SECRET_I_D)
  private String secretID;

  public static final String SERIALIZED_NAME_VAULT_TOKEN = "VaultToken";
  @SerializedName(SERIALIZED_NAME_VAULT_TOKEN)
  private String vaultToken;


  public JobRevertRequest consulToken(String consulToken) {
    
    this.consulToken = consulToken;
    return this;
  }

   /**
   * Get consulToken
   * @return consulToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getConsulToken() {
    return consulToken;
  }


  public void setConsulToken(String consulToken) {
    this.consulToken = consulToken;
  }


  public JobRevertRequest enforcePriorVersion(Integer enforcePriorVersion) {
    
    this.enforcePriorVersion = enforcePriorVersion;
    return this;
  }

   /**
   * Get enforcePriorVersion
   * minimum: 0
   * maximum: 384
   * @return enforcePriorVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getEnforcePriorVersion() {
    return enforcePriorVersion;
  }


  public void setEnforcePriorVersion(Integer enforcePriorVersion) {
    this.enforcePriorVersion = enforcePriorVersion;
  }


  public JobRevertRequest jobID(String jobID) {
    
    this.jobID = jobID;
    return this;
  }

   /**
   * Get jobID
   * @return jobID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getJobID() {
    return jobID;
  }


  public void setJobID(String jobID) {
    this.jobID = jobID;
  }


  public JobRevertRequest jobVersion(Integer jobVersion) {
    
    this.jobVersion = jobVersion;
    return this;
  }

   /**
   * Get jobVersion
   * minimum: 0
   * maximum: 384
   * @return jobVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getJobVersion() {
    return jobVersion;
  }


  public void setJobVersion(Integer jobVersion) {
    this.jobVersion = jobVersion;
  }


  public JobRevertRequest namespace(String namespace) {
    
    this.namespace = namespace;
    return this;
  }

   /**
   * Get namespace
   * @return namespace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNamespace() {
    return namespace;
  }


  public void setNamespace(String namespace) {
    this.namespace = namespace;
  }


  public JobRevertRequest region(String region) {
    
    this.region = region;
    return this;
  }

   /**
   * Get region
   * @return region
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRegion() {
    return region;
  }


  public void setRegion(String region) {
    this.region = region;
  }


  public JobRevertRequest secretID(String secretID) {
    
    this.secretID = secretID;
    return this;
  }

   /**
   * Get secretID
   * @return secretID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSecretID() {
    return secretID;
  }


  public void setSecretID(String secretID) {
    this.secretID = secretID;
  }


  public JobRevertRequest vaultToken(String vaultToken) {
    
    this.vaultToken = vaultToken;
    return this;
  }

   /**
   * Get vaultToken
   * @return vaultToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVaultToken() {
    return vaultToken;
  }


  public void setVaultToken(String vaultToken) {
    this.vaultToken = vaultToken;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    JobRevertRequest jobRevertRequest = (JobRevertRequest) o;
    return Objects.equals(this.consulToken, jobRevertRequest.consulToken) &&
        Objects.equals(this.enforcePriorVersion, jobRevertRequest.enforcePriorVersion) &&
        Objects.equals(this.jobID, jobRevertRequest.jobID) &&
        Objects.equals(this.jobVersion, jobRevertRequest.jobVersion) &&
        Objects.equals(this.namespace, jobRevertRequest.namespace) &&
        Objects.equals(this.region, jobRevertRequest.region) &&
        Objects.equals(this.secretID, jobRevertRequest.secretID) &&
        Objects.equals(this.vaultToken, jobRevertRequest.vaultToken);
  }

  @Override
  public int hashCode() {
    return Objects.hash(consulToken, enforcePriorVersion, jobID, jobVersion, namespace, region, secretID, vaultToken);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class JobRevertRequest {\n");
    sb.append("    consulToken: ").append(toIndentedString(consulToken)).append("\n");
    sb.append("    enforcePriorVersion: ").append(toIndentedString(enforcePriorVersion)).append("\n");
    sb.append("    jobID: ").append(toIndentedString(jobID)).append("\n");
    sb.append("    jobVersion: ").append(toIndentedString(jobVersion)).append("\n");
    sb.append("    namespace: ").append(toIndentedString(namespace)).append("\n");
    sb.append("    region: ").append(toIndentedString(region)).append("\n");
    sb.append("    secretID: ").append(toIndentedString(secretID)).append("\n");
    sb.append("    vaultToken: ").append(toIndentedString(vaultToken)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

