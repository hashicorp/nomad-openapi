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
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * TaskArtifact
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TaskArtifact {
  public static final String SERIALIZED_NAME_GETTER_HEADERS = "GetterHeaders";
  @SerializedName(SERIALIZED_NAME_GETTER_HEADERS)
  private Map<String, String> getterHeaders = null;

  public static final String SERIALIZED_NAME_GETTER_MODE = "GetterMode";
  @SerializedName(SERIALIZED_NAME_GETTER_MODE)
  private String getterMode;

  public static final String SERIALIZED_NAME_GETTER_OPTIONS = "GetterOptions";
  @SerializedName(SERIALIZED_NAME_GETTER_OPTIONS)
  private Map<String, String> getterOptions = null;

  public static final String SERIALIZED_NAME_GETTER_SOURCE = "GetterSource";
  @SerializedName(SERIALIZED_NAME_GETTER_SOURCE)
  private String getterSource;

  public static final String SERIALIZED_NAME_RELATIVE_DEST = "RelativeDest";
  @SerializedName(SERIALIZED_NAME_RELATIVE_DEST)
  private String relativeDest;


  public TaskArtifact getterHeaders(Map<String, String> getterHeaders) {
    
    this.getterHeaders = getterHeaders;
    return this;
  }

  public TaskArtifact putGetterHeadersItem(String key, String getterHeadersItem) {
    if (this.getterHeaders == null) {
      this.getterHeaders = new HashMap<String, String>();
    }
    this.getterHeaders.put(key, getterHeadersItem);
    return this;
  }

   /**
   * Get getterHeaders
   * @return getterHeaders
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, String> getGetterHeaders() {
    return getterHeaders;
  }


  public void setGetterHeaders(Map<String, String> getterHeaders) {
    this.getterHeaders = getterHeaders;
  }


  public TaskArtifact getterMode(String getterMode) {
    
    this.getterMode = getterMode;
    return this;
  }

   /**
   * Get getterMode
   * @return getterMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGetterMode() {
    return getterMode;
  }


  public void setGetterMode(String getterMode) {
    this.getterMode = getterMode;
  }


  public TaskArtifact getterOptions(Map<String, String> getterOptions) {
    
    this.getterOptions = getterOptions;
    return this;
  }

  public TaskArtifact putGetterOptionsItem(String key, String getterOptionsItem) {
    if (this.getterOptions == null) {
      this.getterOptions = new HashMap<String, String>();
    }
    this.getterOptions.put(key, getterOptionsItem);
    return this;
  }

   /**
   * Get getterOptions
   * @return getterOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, String> getGetterOptions() {
    return getterOptions;
  }


  public void setGetterOptions(Map<String, String> getterOptions) {
    this.getterOptions = getterOptions;
  }


  public TaskArtifact getterSource(String getterSource) {
    
    this.getterSource = getterSource;
    return this;
  }

   /**
   * Get getterSource
   * @return getterSource
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGetterSource() {
    return getterSource;
  }


  public void setGetterSource(String getterSource) {
    this.getterSource = getterSource;
  }


  public TaskArtifact relativeDest(String relativeDest) {
    
    this.relativeDest = relativeDest;
    return this;
  }

   /**
   * Get relativeDest
   * @return relativeDest
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRelativeDest() {
    return relativeDest;
  }


  public void setRelativeDest(String relativeDest) {
    this.relativeDest = relativeDest;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskArtifact taskArtifact = (TaskArtifact) o;
    return Objects.equals(this.getterHeaders, taskArtifact.getterHeaders) &&
        Objects.equals(this.getterMode, taskArtifact.getterMode) &&
        Objects.equals(this.getterOptions, taskArtifact.getterOptions) &&
        Objects.equals(this.getterSource, taskArtifact.getterSource) &&
        Objects.equals(this.relativeDest, taskArtifact.relativeDest);
  }

  @Override
  public int hashCode() {
    return Objects.hash(getterHeaders, getterMode, getterOptions, getterSource, relativeDest);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskArtifact {\n");
    sb.append("    getterHeaders: ").append(toIndentedString(getterHeaders)).append("\n");
    sb.append("    getterMode: ").append(toIndentedString(getterMode)).append("\n");
    sb.append("    getterOptions: ").append(toIndentedString(getterOptions)).append("\n");
    sb.append("    getterSource: ").append(toIndentedString(getterSource)).append("\n");
    sb.append("    relativeDest: ").append(toIndentedString(relativeDest)).append("\n");
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

