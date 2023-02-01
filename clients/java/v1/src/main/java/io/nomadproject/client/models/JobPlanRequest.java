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
import io.nomadproject.client.models.Job;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * JobPlanRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class JobPlanRequest {
  public static final String SERIALIZED_NAME_DIFF = "Diff";
  @SerializedName(SERIALIZED_NAME_DIFF)
  private Boolean diff;

  public static final String SERIALIZED_NAME_JOB = "Job";
  @SerializedName(SERIALIZED_NAME_JOB)
  private Job job;

  public static final String SERIALIZED_NAME_NAMESPACE = "Namespace";
  @SerializedName(SERIALIZED_NAME_NAMESPACE)
  private String namespace;

  public static final String SERIALIZED_NAME_POLICY_OVERRIDE = "PolicyOverride";
  @SerializedName(SERIALIZED_NAME_POLICY_OVERRIDE)
  private Boolean policyOverride;

  public static final String SERIALIZED_NAME_REGION = "Region";
  @SerializedName(SERIALIZED_NAME_REGION)
  private String region;

  public static final String SERIALIZED_NAME_SECRET_I_D = "SecretID";
  @SerializedName(SERIALIZED_NAME_SECRET_I_D)
  private String secretID;


  public JobPlanRequest diff(Boolean diff) {
    
    this.diff = diff;
    return this;
  }

   /**
   * Get diff
   * @return diff
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDiff() {
    return diff;
  }


  public void setDiff(Boolean diff) {
    this.diff = diff;
  }


  public JobPlanRequest job(Job job) {
    
    this.job = job;
    return this;
  }

   /**
   * Get job
   * @return job
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Job getJob() {
    return job;
  }


  public void setJob(Job job) {
    this.job = job;
  }


  public JobPlanRequest namespace(String namespace) {
    
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


  public JobPlanRequest policyOverride(Boolean policyOverride) {
    
    this.policyOverride = policyOverride;
    return this;
  }

   /**
   * Get policyOverride
   * @return policyOverride
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getPolicyOverride() {
    return policyOverride;
  }


  public void setPolicyOverride(Boolean policyOverride) {
    this.policyOverride = policyOverride;
  }


  public JobPlanRequest region(String region) {
    
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


  public JobPlanRequest secretID(String secretID) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    JobPlanRequest jobPlanRequest = (JobPlanRequest) o;
    return Objects.equals(this.diff, jobPlanRequest.diff) &&
        Objects.equals(this.job, jobPlanRequest.job) &&
        Objects.equals(this.namespace, jobPlanRequest.namespace) &&
        Objects.equals(this.policyOverride, jobPlanRequest.policyOverride) &&
        Objects.equals(this.region, jobPlanRequest.region) &&
        Objects.equals(this.secretID, jobPlanRequest.secretID);
  }

  @Override
  public int hashCode() {
    return Objects.hash(diff, job, namespace, policyOverride, region, secretID);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class JobPlanRequest {\n");
    sb.append("    diff: ").append(toIndentedString(diff)).append("\n");
    sb.append("    job: ").append(toIndentedString(job)).append("\n");
    sb.append("    namespace: ").append(toIndentedString(namespace)).append("\n");
    sb.append("    policyOverride: ").append(toIndentedString(policyOverride)).append("\n");
    sb.append("    region: ").append(toIndentedString(region)).append("\n");
    sb.append("    secretID: ").append(toIndentedString(secretID)).append("\n");
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

