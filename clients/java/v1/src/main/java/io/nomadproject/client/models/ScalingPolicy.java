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
 * ScalingPolicy
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ScalingPolicy {
  public static final String SERIALIZED_NAME_CREATE_INDEX = "CreateIndex";
  @SerializedName(SERIALIZED_NAME_CREATE_INDEX)
  private Integer createIndex;

  public static final String SERIALIZED_NAME_ENABLED = "Enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_I_D = "ID";
  @SerializedName(SERIALIZED_NAME_I_D)
  private String ID;

  public static final String SERIALIZED_NAME_MAX = "Max";
  @SerializedName(SERIALIZED_NAME_MAX)
  private Long max;

  public static final String SERIALIZED_NAME_MIN = "Min";
  @SerializedName(SERIALIZED_NAME_MIN)
  private Long min;

  public static final String SERIALIZED_NAME_MODIFY_INDEX = "ModifyIndex";
  @SerializedName(SERIALIZED_NAME_MODIFY_INDEX)
  private Integer modifyIndex;

  public static final String SERIALIZED_NAME_NAMESPACE = "Namespace";
  @SerializedName(SERIALIZED_NAME_NAMESPACE)
  private String namespace;

  public static final String SERIALIZED_NAME_POLICY = "Policy";
  @SerializedName(SERIALIZED_NAME_POLICY)
  private Map<String, Object> policy = null;

  public static final String SERIALIZED_NAME_TARGET = "Target";
  @SerializedName(SERIALIZED_NAME_TARGET)
  private Map<String, String> target = null;

  public static final String SERIALIZED_NAME_TYPE = "Type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;


  public ScalingPolicy createIndex(Integer createIndex) {
    
    this.createIndex = createIndex;
    return this;
  }

   /**
   * Get createIndex
   * minimum: 0
   * maximum: 384
   * @return createIndex
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getCreateIndex() {
    return createIndex;
  }


  public void setCreateIndex(Integer createIndex) {
    this.createIndex = createIndex;
  }


  public ScalingPolicy enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Get enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public ScalingPolicy ID(String ID) {
    
    this.ID = ID;
    return this;
  }

   /**
   * Get ID
   * @return ID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getID() {
    return ID;
  }


  public void setID(String ID) {
    this.ID = ID;
  }


  public ScalingPolicy max(Long max) {
    
    this.max = max;
    return this;
  }

   /**
   * Get max
   * @return max
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMax() {
    return max;
  }


  public void setMax(Long max) {
    this.max = max;
  }


  public ScalingPolicy min(Long min) {
    
    this.min = min;
    return this;
  }

   /**
   * Get min
   * @return min
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMin() {
    return min;
  }


  public void setMin(Long min) {
    this.min = min;
  }


  public ScalingPolicy modifyIndex(Integer modifyIndex) {
    
    this.modifyIndex = modifyIndex;
    return this;
  }

   /**
   * Get modifyIndex
   * minimum: 0
   * maximum: 384
   * @return modifyIndex
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getModifyIndex() {
    return modifyIndex;
  }


  public void setModifyIndex(Integer modifyIndex) {
    this.modifyIndex = modifyIndex;
  }


  public ScalingPolicy namespace(String namespace) {
    
    this.namespace = namespace;
    return this;
  }

   /**
   * Get namespace
   * @return namespace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNamespace() {
    return namespace;
  }


  public void setNamespace(String namespace) {
    this.namespace = namespace;
  }


  public ScalingPolicy policy(Map<String, Object> policy) {
    
    this.policy = policy;
    return this;
  }

  public ScalingPolicy putPolicyItem(String key, Object policyItem) {
    if (this.policy == null) {
      this.policy = new HashMap<String, Object>();
    }
    this.policy.put(key, policyItem);
    return this;
  }

   /**
   * Get policy
   * @return policy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getPolicy() {
    return policy;
  }


  public void setPolicy(Map<String, Object> policy) {
    this.policy = policy;
  }


  public ScalingPolicy target(Map<String, String> target) {
    
    this.target = target;
    return this;
  }

  public ScalingPolicy putTargetItem(String key, String targetItem) {
    if (this.target == null) {
      this.target = new HashMap<String, String>();
    }
    this.target.put(key, targetItem);
    return this;
  }

   /**
   * Get target
   * @return target
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, String> getTarget() {
    return target;
  }


  public void setTarget(Map<String, String> target) {
    this.target = target;
  }


  public ScalingPolicy type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ScalingPolicy scalingPolicy = (ScalingPolicy) o;
    return Objects.equals(this.createIndex, scalingPolicy.createIndex) &&
        Objects.equals(this.enabled, scalingPolicy.enabled) &&
        Objects.equals(this.ID, scalingPolicy.ID) &&
        Objects.equals(this.max, scalingPolicy.max) &&
        Objects.equals(this.min, scalingPolicy.min) &&
        Objects.equals(this.modifyIndex, scalingPolicy.modifyIndex) &&
        Objects.equals(this.namespace, scalingPolicy.namespace) &&
        Objects.equals(this.policy, scalingPolicy.policy) &&
        Objects.equals(this.target, scalingPolicy.target) &&
        Objects.equals(this.type, scalingPolicy.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(createIndex, enabled, ID, max, min, modifyIndex, namespace, policy, target, type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ScalingPolicy {\n");
    sb.append("    createIndex: ").append(toIndentedString(createIndex)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    ID: ").append(toIndentedString(ID)).append("\n");
    sb.append("    max: ").append(toIndentedString(max)).append("\n");
    sb.append("    min: ").append(toIndentedString(min)).append("\n");
    sb.append("    modifyIndex: ").append(toIndentedString(modifyIndex)).append("\n");
    sb.append("    namespace: ").append(toIndentedString(namespace)).append("\n");
    sb.append("    policy: ").append(toIndentedString(policy)).append("\n");
    sb.append("    target: ").append(toIndentedString(target)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
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

