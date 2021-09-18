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
 * DesiredUpdates
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class DesiredUpdates {
  public static final String SERIALIZED_NAME_CANARY = "Canary";
  @SerializedName(SERIALIZED_NAME_CANARY)
  private Integer canary;

  public static final String SERIALIZED_NAME_DESTRUCTIVE_UPDATE = "DestructiveUpdate";
  @SerializedName(SERIALIZED_NAME_DESTRUCTIVE_UPDATE)
  private Integer destructiveUpdate;

  public static final String SERIALIZED_NAME_IGNORE = "Ignore";
  @SerializedName(SERIALIZED_NAME_IGNORE)
  private Integer ignore;

  public static final String SERIALIZED_NAME_IN_PLACE_UPDATE = "InPlaceUpdate";
  @SerializedName(SERIALIZED_NAME_IN_PLACE_UPDATE)
  private Integer inPlaceUpdate;

  public static final String SERIALIZED_NAME_MIGRATE = "Migrate";
  @SerializedName(SERIALIZED_NAME_MIGRATE)
  private Integer migrate;

  public static final String SERIALIZED_NAME_PLACE = "Place";
  @SerializedName(SERIALIZED_NAME_PLACE)
  private Integer place;

  public static final String SERIALIZED_NAME_PREEMPTIONS = "Preemptions";
  @SerializedName(SERIALIZED_NAME_PREEMPTIONS)
  private Integer preemptions;

  public static final String SERIALIZED_NAME_STOP = "Stop";
  @SerializedName(SERIALIZED_NAME_STOP)
  private Integer stop;


  public DesiredUpdates canary(Integer canary) {
    
    this.canary = canary;
    return this;
  }

   /**
   * Get canary
   * minimum: 0
   * maximum: 384
   * @return canary
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getCanary() {
    return canary;
  }


  public void setCanary(Integer canary) {
    this.canary = canary;
  }


  public DesiredUpdates destructiveUpdate(Integer destructiveUpdate) {
    
    this.destructiveUpdate = destructiveUpdate;
    return this;
  }

   /**
   * Get destructiveUpdate
   * minimum: 0
   * maximum: 384
   * @return destructiveUpdate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getDestructiveUpdate() {
    return destructiveUpdate;
  }


  public void setDestructiveUpdate(Integer destructiveUpdate) {
    this.destructiveUpdate = destructiveUpdate;
  }


  public DesiredUpdates ignore(Integer ignore) {
    
    this.ignore = ignore;
    return this;
  }

   /**
   * Get ignore
   * minimum: 0
   * maximum: 384
   * @return ignore
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getIgnore() {
    return ignore;
  }


  public void setIgnore(Integer ignore) {
    this.ignore = ignore;
  }


  public DesiredUpdates inPlaceUpdate(Integer inPlaceUpdate) {
    
    this.inPlaceUpdate = inPlaceUpdate;
    return this;
  }

   /**
   * Get inPlaceUpdate
   * minimum: 0
   * maximum: 384
   * @return inPlaceUpdate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getInPlaceUpdate() {
    return inPlaceUpdate;
  }


  public void setInPlaceUpdate(Integer inPlaceUpdate) {
    this.inPlaceUpdate = inPlaceUpdate;
  }


  public DesiredUpdates migrate(Integer migrate) {
    
    this.migrate = migrate;
    return this;
  }

   /**
   * Get migrate
   * minimum: 0
   * maximum: 384
   * @return migrate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getMigrate() {
    return migrate;
  }


  public void setMigrate(Integer migrate) {
    this.migrate = migrate;
  }


  public DesiredUpdates place(Integer place) {
    
    this.place = place;
    return this;
  }

   /**
   * Get place
   * minimum: 0
   * maximum: 384
   * @return place
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getPlace() {
    return place;
  }


  public void setPlace(Integer place) {
    this.place = place;
  }


  public DesiredUpdates preemptions(Integer preemptions) {
    
    this.preemptions = preemptions;
    return this;
  }

   /**
   * Get preemptions
   * minimum: 0
   * maximum: 384
   * @return preemptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getPreemptions() {
    return preemptions;
  }


  public void setPreemptions(Integer preemptions) {
    this.preemptions = preemptions;
  }


  public DesiredUpdates stop(Integer stop) {
    
    this.stop = stop;
    return this;
  }

   /**
   * Get stop
   * minimum: 0
   * maximum: 384
   * @return stop
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getStop() {
    return stop;
  }


  public void setStop(Integer stop) {
    this.stop = stop;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DesiredUpdates desiredUpdates = (DesiredUpdates) o;
    return Objects.equals(this.canary, desiredUpdates.canary) &&
        Objects.equals(this.destructiveUpdate, desiredUpdates.destructiveUpdate) &&
        Objects.equals(this.ignore, desiredUpdates.ignore) &&
        Objects.equals(this.inPlaceUpdate, desiredUpdates.inPlaceUpdate) &&
        Objects.equals(this.migrate, desiredUpdates.migrate) &&
        Objects.equals(this.place, desiredUpdates.place) &&
        Objects.equals(this.preemptions, desiredUpdates.preemptions) &&
        Objects.equals(this.stop, desiredUpdates.stop);
  }

  @Override
  public int hashCode() {
    return Objects.hash(canary, destructiveUpdate, ignore, inPlaceUpdate, migrate, place, preemptions, stop);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DesiredUpdates {\n");
    sb.append("    canary: ").append(toIndentedString(canary)).append("\n");
    sb.append("    destructiveUpdate: ").append(toIndentedString(destructiveUpdate)).append("\n");
    sb.append("    ignore: ").append(toIndentedString(ignore)).append("\n");
    sb.append("    inPlaceUpdate: ").append(toIndentedString(inPlaceUpdate)).append("\n");
    sb.append("    migrate: ").append(toIndentedString(migrate)).append("\n");
    sb.append("    place: ").append(toIndentedString(place)).append("\n");
    sb.append("    preemptions: ").append(toIndentedString(preemptions)).append("\n");
    sb.append("    stop: ").append(toIndentedString(stop)).append("\n");
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

