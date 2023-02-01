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
import io.nomadproject.client.models.FieldDiff;
import io.nomadproject.client.models.ObjectDiff;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * TaskDiff
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TaskDiff {
  public static final String SERIALIZED_NAME_ANNOTATIONS = "Annotations";
  @SerializedName(SERIALIZED_NAME_ANNOTATIONS)
  private List<String> annotations = null;

  public static final String SERIALIZED_NAME_FIELDS = "Fields";
  @SerializedName(SERIALIZED_NAME_FIELDS)
  private List<FieldDiff> fields = null;

  public static final String SERIALIZED_NAME_NAME = "Name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_OBJECTS = "Objects";
  @SerializedName(SERIALIZED_NAME_OBJECTS)
  private List<ObjectDiff> objects = null;

  public static final String SERIALIZED_NAME_TYPE = "Type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;


  public TaskDiff annotations(List<String> annotations) {
    
    this.annotations = annotations;
    return this;
  }

  public TaskDiff addAnnotationsItem(String annotationsItem) {
    if (this.annotations == null) {
      this.annotations = new ArrayList<String>();
    }
    this.annotations.add(annotationsItem);
    return this;
  }

   /**
   * Get annotations
   * @return annotations
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getAnnotations() {
    return annotations;
  }


  public void setAnnotations(List<String> annotations) {
    this.annotations = annotations;
  }


  public TaskDiff fields(List<FieldDiff> fields) {
    
    this.fields = fields;
    return this;
  }

  public TaskDiff addFieldsItem(FieldDiff fieldsItem) {
    if (this.fields == null) {
      this.fields = new ArrayList<FieldDiff>();
    }
    this.fields.add(fieldsItem);
    return this;
  }

   /**
   * Get fields
   * @return fields
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<FieldDiff> getFields() {
    return fields;
  }


  public void setFields(List<FieldDiff> fields) {
    this.fields = fields;
  }


  public TaskDiff name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public TaskDiff objects(List<ObjectDiff> objects) {
    
    this.objects = objects;
    return this;
  }

  public TaskDiff addObjectsItem(ObjectDiff objectsItem) {
    if (this.objects == null) {
      this.objects = new ArrayList<ObjectDiff>();
    }
    this.objects.add(objectsItem);
    return this;
  }

   /**
   * Get objects
   * @return objects
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ObjectDiff> getObjects() {
    return objects;
  }


  public void setObjects(List<ObjectDiff> objects) {
    this.objects = objects;
  }


  public TaskDiff type(String type) {
    
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
    TaskDiff taskDiff = (TaskDiff) o;
    return Objects.equals(this.annotations, taskDiff.annotations) &&
        Objects.equals(this.fields, taskDiff.fields) &&
        Objects.equals(this.name, taskDiff.name) &&
        Objects.equals(this.objects, taskDiff.objects) &&
        Objects.equals(this.type, taskDiff.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(annotations, fields, name, objects, type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskDiff {\n");
    sb.append("    annotations: ").append(toIndentedString(annotations)).append("\n");
    sb.append("    fields: ").append(toIndentedString(fields)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    objects: ").append(toIndentedString(objects)).append("\n");
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

