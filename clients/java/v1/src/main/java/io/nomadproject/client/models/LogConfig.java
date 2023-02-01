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
 * LogConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class LogConfig {
  public static final String SERIALIZED_NAME_MAX_FILE_SIZE_M_B = "MaxFileSizeMB";
  @SerializedName(SERIALIZED_NAME_MAX_FILE_SIZE_M_B)
  private Integer maxFileSizeMB;

  public static final String SERIALIZED_NAME_MAX_FILES = "MaxFiles";
  @SerializedName(SERIALIZED_NAME_MAX_FILES)
  private Integer maxFiles;


  public LogConfig maxFileSizeMB(Integer maxFileSizeMB) {
    
    this.maxFileSizeMB = maxFileSizeMB;
    return this;
  }

   /**
   * Get maxFileSizeMB
   * @return maxFileSizeMB
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMaxFileSizeMB() {
    return maxFileSizeMB;
  }


  public void setMaxFileSizeMB(Integer maxFileSizeMB) {
    this.maxFileSizeMB = maxFileSizeMB;
  }


  public LogConfig maxFiles(Integer maxFiles) {
    
    this.maxFiles = maxFiles;
    return this;
  }

   /**
   * Get maxFiles
   * @return maxFiles
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMaxFiles() {
    return maxFiles;
  }


  public void setMaxFiles(Integer maxFiles) {
    this.maxFiles = maxFiles;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LogConfig logConfig = (LogConfig) o;
    return Objects.equals(this.maxFileSizeMB, logConfig.maxFileSizeMB) &&
        Objects.equals(this.maxFiles, logConfig.maxFiles);
  }

  @Override
  public int hashCode() {
    return Objects.hash(maxFileSizeMB, maxFiles);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LogConfig {\n");
    sb.append("    maxFileSizeMB: ").append(toIndentedString(maxFileSizeMB)).append("\n");
    sb.append("    maxFiles: ").append(toIndentedString(maxFiles)).append("\n");
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

