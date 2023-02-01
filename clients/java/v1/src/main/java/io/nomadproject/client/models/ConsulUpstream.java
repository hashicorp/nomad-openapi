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
import io.nomadproject.client.models.ConsulMeshGateway;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * ConsulUpstream
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ConsulUpstream {
  public static final String SERIALIZED_NAME_DATACENTER = "Datacenter";
  @SerializedName(SERIALIZED_NAME_DATACENTER)
  private String datacenter;

  public static final String SERIALIZED_NAME_DESTINATION_NAME = "DestinationName";
  @SerializedName(SERIALIZED_NAME_DESTINATION_NAME)
  private String destinationName;

  public static final String SERIALIZED_NAME_DESTINATION_NAMESPACE = "DestinationNamespace";
  @SerializedName(SERIALIZED_NAME_DESTINATION_NAMESPACE)
  private String destinationNamespace;

  public static final String SERIALIZED_NAME_LOCAL_BIND_ADDRESS = "LocalBindAddress";
  @SerializedName(SERIALIZED_NAME_LOCAL_BIND_ADDRESS)
  private String localBindAddress;

  public static final String SERIALIZED_NAME_LOCAL_BIND_PORT = "LocalBindPort";
  @SerializedName(SERIALIZED_NAME_LOCAL_BIND_PORT)
  private Integer localBindPort;

  public static final String SERIALIZED_NAME_MESH_GATEWAY = "MeshGateway";
  @SerializedName(SERIALIZED_NAME_MESH_GATEWAY)
  private ConsulMeshGateway meshGateway;


  public ConsulUpstream datacenter(String datacenter) {
    
    this.datacenter = datacenter;
    return this;
  }

   /**
   * Get datacenter
   * @return datacenter
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDatacenter() {
    return datacenter;
  }


  public void setDatacenter(String datacenter) {
    this.datacenter = datacenter;
  }


  public ConsulUpstream destinationName(String destinationName) {
    
    this.destinationName = destinationName;
    return this;
  }

   /**
   * Get destinationName
   * @return destinationName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDestinationName() {
    return destinationName;
  }


  public void setDestinationName(String destinationName) {
    this.destinationName = destinationName;
  }


  public ConsulUpstream destinationNamespace(String destinationNamespace) {
    
    this.destinationNamespace = destinationNamespace;
    return this;
  }

   /**
   * Get destinationNamespace
   * @return destinationNamespace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDestinationNamespace() {
    return destinationNamespace;
  }


  public void setDestinationNamespace(String destinationNamespace) {
    this.destinationNamespace = destinationNamespace;
  }


  public ConsulUpstream localBindAddress(String localBindAddress) {
    
    this.localBindAddress = localBindAddress;
    return this;
  }

   /**
   * Get localBindAddress
   * @return localBindAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLocalBindAddress() {
    return localBindAddress;
  }


  public void setLocalBindAddress(String localBindAddress) {
    this.localBindAddress = localBindAddress;
  }


  public ConsulUpstream localBindPort(Integer localBindPort) {
    
    this.localBindPort = localBindPort;
    return this;
  }

   /**
   * Get localBindPort
   * @return localBindPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getLocalBindPort() {
    return localBindPort;
  }


  public void setLocalBindPort(Integer localBindPort) {
    this.localBindPort = localBindPort;
  }


  public ConsulUpstream meshGateway(ConsulMeshGateway meshGateway) {
    
    this.meshGateway = meshGateway;
    return this;
  }

   /**
   * Get meshGateway
   * @return meshGateway
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ConsulMeshGateway getMeshGateway() {
    return meshGateway;
  }


  public void setMeshGateway(ConsulMeshGateway meshGateway) {
    this.meshGateway = meshGateway;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ConsulUpstream consulUpstream = (ConsulUpstream) o;
    return Objects.equals(this.datacenter, consulUpstream.datacenter) &&
        Objects.equals(this.destinationName, consulUpstream.destinationName) &&
        Objects.equals(this.destinationNamespace, consulUpstream.destinationNamespace) &&
        Objects.equals(this.localBindAddress, consulUpstream.localBindAddress) &&
        Objects.equals(this.localBindPort, consulUpstream.localBindPort) &&
        Objects.equals(this.meshGateway, consulUpstream.meshGateway);
  }

  @Override
  public int hashCode() {
    return Objects.hash(datacenter, destinationName, destinationNamespace, localBindAddress, localBindPort, meshGateway);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ConsulUpstream {\n");
    sb.append("    datacenter: ").append(toIndentedString(datacenter)).append("\n");
    sb.append("    destinationName: ").append(toIndentedString(destinationName)).append("\n");
    sb.append("    destinationNamespace: ").append(toIndentedString(destinationNamespace)).append("\n");
    sb.append("    localBindAddress: ").append(toIndentedString(localBindAddress)).append("\n");
    sb.append("    localBindPort: ").append(toIndentedString(localBindPort)).append("\n");
    sb.append("    meshGateway: ").append(toIndentedString(meshGateway)).append("\n");
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

