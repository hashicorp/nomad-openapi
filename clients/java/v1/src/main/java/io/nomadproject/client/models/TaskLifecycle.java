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
 * TaskLifecycle
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TaskLifecycle {
  public static final String SERIALIZED_NAME_HOOK = "Hook";
  @SerializedName(SERIALIZED_NAME_HOOK)
  private String hook;

  public static final String SERIALIZED_NAME_SIDECAR = "Sidecar";
  @SerializedName(SERIALIZED_NAME_SIDECAR)
  private Boolean sidecar;


  public TaskLifecycle hook(String hook) {
    
    this.hook = hook;
    return this;
  }

   /**
   * Get hook
   * @return hook
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHook() {
    return hook;
  }


  public void setHook(String hook) {
    this.hook = hook;
  }


  public TaskLifecycle sidecar(Boolean sidecar) {
    
    this.sidecar = sidecar;
    return this;
  }

   /**
   * Get sidecar
   * @return sidecar
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSidecar() {
    return sidecar;
  }


  public void setSidecar(Boolean sidecar) {
    this.sidecar = sidecar;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskLifecycle taskLifecycle = (TaskLifecycle) o;
    return Objects.equals(this.hook, taskLifecycle.hook) &&
        Objects.equals(this.sidecar, taskLifecycle.sidecar);
  }

  @Override
  public int hashCode() {
    return Objects.hash(hook, sidecar);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskLifecycle {\n");
    sb.append("    hook: ").append(toIndentedString(hook)).append("\n");
    sb.append("    sidecar: ").append(toIndentedString(sidecar)).append("\n");
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

