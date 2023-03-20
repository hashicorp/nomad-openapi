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
import io.nomadproject.client.models.ConsulProxy;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * ConsulSidecarService
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ConsulSidecarService {
  public static final String SERIALIZED_NAME_DISABLE_DEFAULT_T_C_P_CHECK = "DisableDefaultTCPCheck";
  @SerializedName(SERIALIZED_NAME_DISABLE_DEFAULT_T_C_P_CHECK)
  private Boolean disableDefaultTCPCheck;

  public static final String SERIALIZED_NAME_PORT = "Port";
  @SerializedName(SERIALIZED_NAME_PORT)
  private String port;

  public static final String SERIALIZED_NAME_PROXY = "Proxy";
  @SerializedName(SERIALIZED_NAME_PROXY)
  private ConsulProxy proxy;

  public static final String SERIALIZED_NAME_TAGS = "Tags";
  @SerializedName(SERIALIZED_NAME_TAGS)
  private List<String> tags = null;


  public ConsulSidecarService disableDefaultTCPCheck(Boolean disableDefaultTCPCheck) {
    
    this.disableDefaultTCPCheck = disableDefaultTCPCheck;
    return this;
  }

   /**
   * Get disableDefaultTCPCheck
   * @return disableDefaultTCPCheck
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDisableDefaultTCPCheck() {
    return disableDefaultTCPCheck;
  }


  public void setDisableDefaultTCPCheck(Boolean disableDefaultTCPCheck) {
    this.disableDefaultTCPCheck = disableDefaultTCPCheck;
  }


  public ConsulSidecarService port(String port) {
    
    this.port = port;
    return this;
  }

   /**
   * Get port
   * @return port
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPort() {
    return port;
  }


  public void setPort(String port) {
    this.port = port;
  }


  public ConsulSidecarService proxy(ConsulProxy proxy) {
    
    this.proxy = proxy;
    return this;
  }

   /**
   * Get proxy
   * @return proxy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ConsulProxy getProxy() {
    return proxy;
  }


  public void setProxy(ConsulProxy proxy) {
    this.proxy = proxy;
  }


  public ConsulSidecarService tags(List<String> tags) {
    
    this.tags = tags;
    return this;
  }

  public ConsulSidecarService addTagsItem(String tagsItem) {
    if (this.tags == null) {
      this.tags = new ArrayList<String>();
    }
    this.tags.add(tagsItem);
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getTags() {
    return tags;
  }


  public void setTags(List<String> tags) {
    this.tags = tags;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ConsulSidecarService consulSidecarService = (ConsulSidecarService) o;
    return Objects.equals(this.disableDefaultTCPCheck, consulSidecarService.disableDefaultTCPCheck) &&
        Objects.equals(this.port, consulSidecarService.port) &&
        Objects.equals(this.proxy, consulSidecarService.proxy) &&
        Objects.equals(this.tags, consulSidecarService.tags);
  }

  @Override
  public int hashCode() {
    return Objects.hash(disableDefaultTCPCheck, port, proxy, tags);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ConsulSidecarService {\n");
    sb.append("    disableDefaultTCPCheck: ").append(toIndentedString(disableDefaultTCPCheck)).append("\n");
    sb.append("    port: ").append(toIndentedString(port)).append("\n");
    sb.append("    proxy: ").append(toIndentedString(proxy)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
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

