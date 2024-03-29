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
 * EvalOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class EvalOptions {
  public static final String SERIALIZED_NAME_FORCE_RESCHEDULE = "ForceReschedule";
  @SerializedName(SERIALIZED_NAME_FORCE_RESCHEDULE)
  private Boolean forceReschedule;


  public EvalOptions forceReschedule(Boolean forceReschedule) {
    
    this.forceReschedule = forceReschedule;
    return this;
  }

   /**
   * Get forceReschedule
   * @return forceReschedule
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getForceReschedule() {
    return forceReschedule;
  }


  public void setForceReschedule(Boolean forceReschedule) {
    this.forceReschedule = forceReschedule;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    EvalOptions evalOptions = (EvalOptions) o;
    return Objects.equals(this.forceReschedule, evalOptions.forceReschedule);
  }

  @Override
  public int hashCode() {
    return Objects.hash(forceReschedule);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class EvalOptions {\n");
    sb.append("    forceReschedule: ").append(toIndentedString(forceReschedule)).append("\n");
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

