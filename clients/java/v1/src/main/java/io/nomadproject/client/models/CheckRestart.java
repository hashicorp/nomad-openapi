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
 * CheckRestart
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class CheckRestart {
  public static final String SERIALIZED_NAME_GRACE = "Grace";
  @SerializedName(SERIALIZED_NAME_GRACE)
  private Long grace;

  public static final String SERIALIZED_NAME_IGNORE_WARNINGS = "IgnoreWarnings";
  @SerializedName(SERIALIZED_NAME_IGNORE_WARNINGS)
  private Boolean ignoreWarnings;

  public static final String SERIALIZED_NAME_LIMIT = "Limit";
  @SerializedName(SERIALIZED_NAME_LIMIT)
  private Integer limit;


  public CheckRestart grace(Long grace) {
    
    this.grace = grace;
    return this;
  }

   /**
   * Get grace
   * @return grace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getGrace() {
    return grace;
  }


  public void setGrace(Long grace) {
    this.grace = grace;
  }


  public CheckRestart ignoreWarnings(Boolean ignoreWarnings) {
    
    this.ignoreWarnings = ignoreWarnings;
    return this;
  }

   /**
   * Get ignoreWarnings
   * @return ignoreWarnings
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIgnoreWarnings() {
    return ignoreWarnings;
  }


  public void setIgnoreWarnings(Boolean ignoreWarnings) {
    this.ignoreWarnings = ignoreWarnings;
  }


  public CheckRestart limit(Integer limit) {
    
    this.limit = limit;
    return this;
  }

   /**
   * Get limit
   * @return limit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getLimit() {
    return limit;
  }


  public void setLimit(Integer limit) {
    this.limit = limit;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CheckRestart checkRestart = (CheckRestart) o;
    return Objects.equals(this.grace, checkRestart.grace) &&
        Objects.equals(this.ignoreWarnings, checkRestart.ignoreWarnings) &&
        Objects.equals(this.limit, checkRestart.limit);
  }

  @Override
  public int hashCode() {
    return Objects.hash(grace, ignoreWarnings, limit);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CheckRestart {\n");
    sb.append("    grace: ").append(toIndentedString(grace)).append("\n");
    sb.append("    ignoreWarnings: ").append(toIndentedString(ignoreWarnings)).append("\n");
    sb.append("    limit: ").append(toIndentedString(limit)).append("\n");
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

