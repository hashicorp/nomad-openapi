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
 * NamespaceCapabilities
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NamespaceCapabilities {
  public static final String SERIALIZED_NAME_DISABLED_TASK_DRIVERS = "DisabledTaskDrivers";
  @SerializedName(SERIALIZED_NAME_DISABLED_TASK_DRIVERS)
  private List<String> disabledTaskDrivers = null;

  public static final String SERIALIZED_NAME_ENABLED_TASK_DRIVERS = "EnabledTaskDrivers";
  @SerializedName(SERIALIZED_NAME_ENABLED_TASK_DRIVERS)
  private List<String> enabledTaskDrivers = null;

  public NamespaceCapabilities() { 
  }

  public NamespaceCapabilities disabledTaskDrivers(List<String> disabledTaskDrivers) {
    
    this.disabledTaskDrivers = disabledTaskDrivers;
    return this;
  }

  public NamespaceCapabilities addDisabledTaskDriversItem(String disabledTaskDriversItem) {
    if (this.disabledTaskDrivers == null) {
      this.disabledTaskDrivers = new ArrayList<String>();
    }
    this.disabledTaskDrivers.add(disabledTaskDriversItem);
    return this;
  }

   /**
   * Get disabledTaskDrivers
   * @return disabledTaskDrivers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getDisabledTaskDrivers() {
    return disabledTaskDrivers;
  }


  public void setDisabledTaskDrivers(List<String> disabledTaskDrivers) {
    this.disabledTaskDrivers = disabledTaskDrivers;
  }


  public NamespaceCapabilities enabledTaskDrivers(List<String> enabledTaskDrivers) {
    
    this.enabledTaskDrivers = enabledTaskDrivers;
    return this;
  }

  public NamespaceCapabilities addEnabledTaskDriversItem(String enabledTaskDriversItem) {
    if (this.enabledTaskDrivers == null) {
      this.enabledTaskDrivers = new ArrayList<String>();
    }
    this.enabledTaskDrivers.add(enabledTaskDriversItem);
    return this;
  }

   /**
   * Get enabledTaskDrivers
   * @return enabledTaskDrivers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getEnabledTaskDrivers() {
    return enabledTaskDrivers;
  }


  public void setEnabledTaskDrivers(List<String> enabledTaskDrivers) {
    this.enabledTaskDrivers = enabledTaskDrivers;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NamespaceCapabilities namespaceCapabilities = (NamespaceCapabilities) o;
    return Objects.equals(this.disabledTaskDrivers, namespaceCapabilities.disabledTaskDrivers) &&
        Objects.equals(this.enabledTaskDrivers, namespaceCapabilities.enabledTaskDrivers);
  }

  @Override
  public int hashCode() {
    return Objects.hash(disabledTaskDrivers, enabledTaskDrivers);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NamespaceCapabilities {\n");
    sb.append("    disabledTaskDrivers: ").append(toIndentedString(disabledTaskDrivers)).append("\n");
    sb.append("    enabledTaskDrivers: ").append(toIndentedString(enabledTaskDrivers)).append("\n");
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

