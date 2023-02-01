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
 * NodeReservedDiskResources
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NodeReservedDiskResources {
  public static final String SERIALIZED_NAME_DISK_M_B = "DiskMB";
  @SerializedName(SERIALIZED_NAME_DISK_M_B)
  private Integer diskMB;


  public NodeReservedDiskResources diskMB(Integer diskMB) {
    
    this.diskMB = diskMB;
    return this;
  }

   /**
   * Get diskMB
   * minimum: 0
   * maximum: 384
   * @return diskMB
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getDiskMB() {
    return diskMB;
  }


  public void setDiskMB(Integer diskMB) {
    this.diskMB = diskMB;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NodeReservedDiskResources nodeReservedDiskResources = (NodeReservedDiskResources) o;
    return Objects.equals(this.diskMB, nodeReservedDiskResources.diskMB);
  }

  @Override
  public int hashCode() {
    return Objects.hash(diskMB);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NodeReservedDiskResources {\n");
    sb.append("    diskMB: ").append(toIndentedString(diskMB)).append("\n");
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

