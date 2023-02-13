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
import java.util.ArrayList;
import java.util.List;

/**
 * NodeCpuResources
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NodeCpuResources {
  public static final String SERIALIZED_NAME_CPU_SHARES = "CpuShares";
  @SerializedName(SERIALIZED_NAME_CPU_SHARES)
  private Long cpuShares;

  public static final String SERIALIZED_NAME_RESERVABLE_CPU_CORES = "ReservableCpuCores";
  @SerializedName(SERIALIZED_NAME_RESERVABLE_CPU_CORES)
  private List<Integer> reservableCpuCores = null;

  public static final String SERIALIZED_NAME_TOTAL_CPU_CORES = "TotalCpuCores";
  @SerializedName(SERIALIZED_NAME_TOTAL_CPU_CORES)
  private Integer totalCpuCores;


  public NodeCpuResources cpuShares(Long cpuShares) {
    
    this.cpuShares = cpuShares;
    return this;
  }

   /**
   * Get cpuShares
   * @return cpuShares
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCpuShares() {
    return cpuShares;
  }


  public void setCpuShares(Long cpuShares) {
    this.cpuShares = cpuShares;
  }


  public NodeCpuResources reservableCpuCores(List<Integer> reservableCpuCores) {
    
    this.reservableCpuCores = reservableCpuCores;
    return this;
  }

  public NodeCpuResources addReservableCpuCoresItem(Integer reservableCpuCoresItem) {
    if (this.reservableCpuCores == null) {
      this.reservableCpuCores = new ArrayList<Integer>();
    }
    this.reservableCpuCores.add(reservableCpuCoresItem);
    return this;
  }

   /**
   * Get reservableCpuCores
   * @return reservableCpuCores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Integer> getReservableCpuCores() {
    return reservableCpuCores;
  }


  public void setReservableCpuCores(List<Integer> reservableCpuCores) {
    this.reservableCpuCores = reservableCpuCores;
  }


  public NodeCpuResources totalCpuCores(Integer totalCpuCores) {
    
    this.totalCpuCores = totalCpuCores;
    return this;
  }

   /**
   * Get totalCpuCores
   * minimum: 0
   * maximum: 65535
   * @return totalCpuCores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getTotalCpuCores() {
    return totalCpuCores;
  }


  public void setTotalCpuCores(Integer totalCpuCores) {
    this.totalCpuCores = totalCpuCores;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NodeCpuResources nodeCpuResources = (NodeCpuResources) o;
    return Objects.equals(this.cpuShares, nodeCpuResources.cpuShares) &&
        Objects.equals(this.reservableCpuCores, nodeCpuResources.reservableCpuCores) &&
        Objects.equals(this.totalCpuCores, nodeCpuResources.totalCpuCores);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cpuShares, reservableCpuCores, totalCpuCores);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NodeCpuResources {\n");
    sb.append("    cpuShares: ").append(toIndentedString(cpuShares)).append("\n");
    sb.append("    reservableCpuCores: ").append(toIndentedString(reservableCpuCores)).append("\n");
    sb.append("    totalCpuCores: ").append(toIndentedString(totalCpuCores)).append("\n");
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

