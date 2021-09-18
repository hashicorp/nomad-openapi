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
 * ParameterizedJobConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ParameterizedJobConfig {
  public static final String SERIALIZED_NAME_META_OPTIONAL = "MetaOptional";
  @SerializedName(SERIALIZED_NAME_META_OPTIONAL)
  private List<String> metaOptional = null;

  public static final String SERIALIZED_NAME_META_REQUIRED = "MetaRequired";
  @SerializedName(SERIALIZED_NAME_META_REQUIRED)
  private List<String> metaRequired = null;

  public static final String SERIALIZED_NAME_PAYLOAD = "Payload";
  @SerializedName(SERIALIZED_NAME_PAYLOAD)
  private String payload;


  public ParameterizedJobConfig metaOptional(List<String> metaOptional) {
    
    this.metaOptional = metaOptional;
    return this;
  }

  public ParameterizedJobConfig addMetaOptionalItem(String metaOptionalItem) {
    if (this.metaOptional == null) {
      this.metaOptional = new ArrayList<String>();
    }
    this.metaOptional.add(metaOptionalItem);
    return this;
  }

   /**
   * Get metaOptional
   * @return metaOptional
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getMetaOptional() {
    return metaOptional;
  }


  public void setMetaOptional(List<String> metaOptional) {
    this.metaOptional = metaOptional;
  }


  public ParameterizedJobConfig metaRequired(List<String> metaRequired) {
    
    this.metaRequired = metaRequired;
    return this;
  }

  public ParameterizedJobConfig addMetaRequiredItem(String metaRequiredItem) {
    if (this.metaRequired == null) {
      this.metaRequired = new ArrayList<String>();
    }
    this.metaRequired.add(metaRequiredItem);
    return this;
  }

   /**
   * Get metaRequired
   * @return metaRequired
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getMetaRequired() {
    return metaRequired;
  }


  public void setMetaRequired(List<String> metaRequired) {
    this.metaRequired = metaRequired;
  }


  public ParameterizedJobConfig payload(String payload) {
    
    this.payload = payload;
    return this;
  }

   /**
   * Get payload
   * @return payload
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPayload() {
    return payload;
  }


  public void setPayload(String payload) {
    this.payload = payload;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ParameterizedJobConfig parameterizedJobConfig = (ParameterizedJobConfig) o;
    return Objects.equals(this.metaOptional, parameterizedJobConfig.metaOptional) &&
        Objects.equals(this.metaRequired, parameterizedJobConfig.metaRequired) &&
        Objects.equals(this.payload, parameterizedJobConfig.payload);
  }

  @Override
  public int hashCode() {
    return Objects.hash(metaOptional, metaRequired, payload);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ParameterizedJobConfig {\n");
    sb.append("    metaOptional: ").append(toIndentedString(metaOptional)).append("\n");
    sb.append("    metaRequired: ").append(toIndentedString(metaRequired)).append("\n");
    sb.append("    payload: ").append(toIndentedString(payload)).append("\n");
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

