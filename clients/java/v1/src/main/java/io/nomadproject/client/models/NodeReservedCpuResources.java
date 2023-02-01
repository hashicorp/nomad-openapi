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
 * NodeReservedCpuResources
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NodeReservedCpuResources {
  public static final String SERIALIZED_NAME_CPU_SHARES = "CpuShares";
  @SerializedName(SERIALIZED_NAME_CPU_SHARES)
  private Integer cpuShares;


  public NodeReservedCpuResources cpuShares(Integer cpuShares) {
    
    this.cpuShares = cpuShares;
    return this;
  }

   /**
   * Get cpuShares
   * minimum: 0
   * maximum: 384
   * @return cpuShares
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getCpuShares() {
    return cpuShares;
  }


  public void setCpuShares(Integer cpuShares) {
    this.cpuShares = cpuShares;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NodeReservedCpuResources nodeReservedCpuResources = (NodeReservedCpuResources) o;
    return Objects.equals(this.cpuShares, nodeReservedCpuResources.cpuShares);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cpuShares);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NodeReservedCpuResources {\n");
    sb.append("    cpuShares: ").append(toIndentedString(cpuShares)).append("\n");
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

