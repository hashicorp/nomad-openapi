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
import io.nomadproject.client.models.ConsulGatewayBindAddress;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * ConsulGatewayProxy
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ConsulGatewayProxy {
  public static final String SERIALIZED_NAME_CONFIG = "Config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Map<String, Object> config = null;

  public static final String SERIALIZED_NAME_CONNECT_TIMEOUT = "ConnectTimeout";
  @SerializedName(SERIALIZED_NAME_CONNECT_TIMEOUT)
  private Long connectTimeout;

  public static final String SERIALIZED_NAME_ENVOY_D_N_S_DISCOVERY_TYPE = "EnvoyDNSDiscoveryType";
  @SerializedName(SERIALIZED_NAME_ENVOY_D_N_S_DISCOVERY_TYPE)
  private String envoyDNSDiscoveryType;

  public static final String SERIALIZED_NAME_ENVOY_GATEWAY_BIND_ADDRESSES = "EnvoyGatewayBindAddresses";
  @SerializedName(SERIALIZED_NAME_ENVOY_GATEWAY_BIND_ADDRESSES)
  private Map<String, ConsulGatewayBindAddress> envoyGatewayBindAddresses = null;

  public static final String SERIALIZED_NAME_ENVOY_GATEWAY_BIND_TAGGED_ADDRESSES = "EnvoyGatewayBindTaggedAddresses";
  @SerializedName(SERIALIZED_NAME_ENVOY_GATEWAY_BIND_TAGGED_ADDRESSES)
  private Boolean envoyGatewayBindTaggedAddresses;

  public static final String SERIALIZED_NAME_ENVOY_GATEWAY_NO_DEFAULT_BIND = "EnvoyGatewayNoDefaultBind";
  @SerializedName(SERIALIZED_NAME_ENVOY_GATEWAY_NO_DEFAULT_BIND)
  private Boolean envoyGatewayNoDefaultBind;


  public ConsulGatewayProxy config(Map<String, Object> config) {
    
    this.config = config;
    return this;
  }

  public ConsulGatewayProxy putConfigItem(String key, Object configItem) {
    if (this.config == null) {
      this.config = new HashMap<String, Object>();
    }
    this.config.put(key, configItem);
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getConfig() {
    return config;
  }


  public void setConfig(Map<String, Object> config) {
    this.config = config;
  }


  public ConsulGatewayProxy connectTimeout(Long connectTimeout) {
    
    this.connectTimeout = connectTimeout;
    return this;
  }

   /**
   * Get connectTimeout
   * @return connectTimeout
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getConnectTimeout() {
    return connectTimeout;
  }


  public void setConnectTimeout(Long connectTimeout) {
    this.connectTimeout = connectTimeout;
  }


  public ConsulGatewayProxy envoyDNSDiscoveryType(String envoyDNSDiscoveryType) {
    
    this.envoyDNSDiscoveryType = envoyDNSDiscoveryType;
    return this;
  }

   /**
   * Get envoyDNSDiscoveryType
   * @return envoyDNSDiscoveryType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEnvoyDNSDiscoveryType() {
    return envoyDNSDiscoveryType;
  }


  public void setEnvoyDNSDiscoveryType(String envoyDNSDiscoveryType) {
    this.envoyDNSDiscoveryType = envoyDNSDiscoveryType;
  }


  public ConsulGatewayProxy envoyGatewayBindAddresses(Map<String, ConsulGatewayBindAddress> envoyGatewayBindAddresses) {
    
    this.envoyGatewayBindAddresses = envoyGatewayBindAddresses;
    return this;
  }

  public ConsulGatewayProxy putEnvoyGatewayBindAddressesItem(String key, ConsulGatewayBindAddress envoyGatewayBindAddressesItem) {
    if (this.envoyGatewayBindAddresses == null) {
      this.envoyGatewayBindAddresses = new HashMap<String, ConsulGatewayBindAddress>();
    }
    this.envoyGatewayBindAddresses.put(key, envoyGatewayBindAddressesItem);
    return this;
  }

   /**
   * Get envoyGatewayBindAddresses
   * @return envoyGatewayBindAddresses
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, ConsulGatewayBindAddress> getEnvoyGatewayBindAddresses() {
    return envoyGatewayBindAddresses;
  }


  public void setEnvoyGatewayBindAddresses(Map<String, ConsulGatewayBindAddress> envoyGatewayBindAddresses) {
    this.envoyGatewayBindAddresses = envoyGatewayBindAddresses;
  }


  public ConsulGatewayProxy envoyGatewayBindTaggedAddresses(Boolean envoyGatewayBindTaggedAddresses) {
    
    this.envoyGatewayBindTaggedAddresses = envoyGatewayBindTaggedAddresses;
    return this;
  }

   /**
   * Get envoyGatewayBindTaggedAddresses
   * @return envoyGatewayBindTaggedAddresses
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnvoyGatewayBindTaggedAddresses() {
    return envoyGatewayBindTaggedAddresses;
  }


  public void setEnvoyGatewayBindTaggedAddresses(Boolean envoyGatewayBindTaggedAddresses) {
    this.envoyGatewayBindTaggedAddresses = envoyGatewayBindTaggedAddresses;
  }


  public ConsulGatewayProxy envoyGatewayNoDefaultBind(Boolean envoyGatewayNoDefaultBind) {
    
    this.envoyGatewayNoDefaultBind = envoyGatewayNoDefaultBind;
    return this;
  }

   /**
   * Get envoyGatewayNoDefaultBind
   * @return envoyGatewayNoDefaultBind
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnvoyGatewayNoDefaultBind() {
    return envoyGatewayNoDefaultBind;
  }


  public void setEnvoyGatewayNoDefaultBind(Boolean envoyGatewayNoDefaultBind) {
    this.envoyGatewayNoDefaultBind = envoyGatewayNoDefaultBind;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ConsulGatewayProxy consulGatewayProxy = (ConsulGatewayProxy) o;
    return Objects.equals(this.config, consulGatewayProxy.config) &&
        Objects.equals(this.connectTimeout, consulGatewayProxy.connectTimeout) &&
        Objects.equals(this.envoyDNSDiscoveryType, consulGatewayProxy.envoyDNSDiscoveryType) &&
        Objects.equals(this.envoyGatewayBindAddresses, consulGatewayProxy.envoyGatewayBindAddresses) &&
        Objects.equals(this.envoyGatewayBindTaggedAddresses, consulGatewayProxy.envoyGatewayBindTaggedAddresses) &&
        Objects.equals(this.envoyGatewayNoDefaultBind, consulGatewayProxy.envoyGatewayNoDefaultBind);
  }

  @Override
  public int hashCode() {
    return Objects.hash(config, connectTimeout, envoyDNSDiscoveryType, envoyGatewayBindAddresses, envoyGatewayBindTaggedAddresses, envoyGatewayNoDefaultBind);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ConsulGatewayProxy {\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    connectTimeout: ").append(toIndentedString(connectTimeout)).append("\n");
    sb.append("    envoyDNSDiscoveryType: ").append(toIndentedString(envoyDNSDiscoveryType)).append("\n");
    sb.append("    envoyGatewayBindAddresses: ").append(toIndentedString(envoyGatewayBindAddresses)).append("\n");
    sb.append("    envoyGatewayBindTaggedAddresses: ").append(toIndentedString(envoyGatewayBindTaggedAddresses)).append("\n");
    sb.append("    envoyGatewayNoDefaultBind: ").append(toIndentedString(envoyGatewayNoDefaultBind)).append("\n");
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

