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
import io.nomadproject.client.models.TaskEvent;
import io.nomadproject.client.models.TaskHandle;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import org.threeten.bp.OffsetDateTime;

/**
 * TaskState
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TaskState {
  public static final String SERIALIZED_NAME_EVENTS = "Events";
  @SerializedName(SERIALIZED_NAME_EVENTS)
  private List<TaskEvent> events = null;

  public static final String SERIALIZED_NAME_FAILED = "Failed";
  @SerializedName(SERIALIZED_NAME_FAILED)
  private Boolean failed;

  public static final String SERIALIZED_NAME_FINISHED_AT = "FinishedAt";
  @SerializedName(SERIALIZED_NAME_FINISHED_AT)
  private OffsetDateTime finishedAt;

  public static final String SERIALIZED_NAME_LAST_RESTART = "LastRestart";
  @SerializedName(SERIALIZED_NAME_LAST_RESTART)
  private OffsetDateTime lastRestart;

  public static final String SERIALIZED_NAME_RESTARTS = "Restarts";
  @SerializedName(SERIALIZED_NAME_RESTARTS)
  private Integer restarts;

  public static final String SERIALIZED_NAME_STARTED_AT = "StartedAt";
  @SerializedName(SERIALIZED_NAME_STARTED_AT)
  private OffsetDateTime startedAt;

  public static final String SERIALIZED_NAME_STATE = "State";
  @SerializedName(SERIALIZED_NAME_STATE)
  private String state;

  public static final String SERIALIZED_NAME_TASK_HANDLE = "TaskHandle";
  @SerializedName(SERIALIZED_NAME_TASK_HANDLE)
  private TaskHandle taskHandle;


  public TaskState events(List<TaskEvent> events) {
    
    this.events = events;
    return this;
  }

  public TaskState addEventsItem(TaskEvent eventsItem) {
    if (this.events == null) {
      this.events = new ArrayList<TaskEvent>();
    }
    this.events.add(eventsItem);
    return this;
  }

   /**
   * Get events
   * @return events
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<TaskEvent> getEvents() {
    return events;
  }


  public void setEvents(List<TaskEvent> events) {
    this.events = events;
  }


  public TaskState failed(Boolean failed) {
    
    this.failed = failed;
    return this;
  }

   /**
   * Get failed
   * @return failed
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getFailed() {
    return failed;
  }


  public void setFailed(Boolean failed) {
    this.failed = failed;
  }


  public TaskState finishedAt(OffsetDateTime finishedAt) {
    
    this.finishedAt = finishedAt;
    return this;
  }

   /**
   * Get finishedAt
   * @return finishedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getFinishedAt() {
    return finishedAt;
  }


  public void setFinishedAt(OffsetDateTime finishedAt) {
    this.finishedAt = finishedAt;
  }


  public TaskState lastRestart(OffsetDateTime lastRestart) {
    
    this.lastRestart = lastRestart;
    return this;
  }

   /**
   * Get lastRestart
   * @return lastRestart
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastRestart() {
    return lastRestart;
  }


  public void setLastRestart(OffsetDateTime lastRestart) {
    this.lastRestart = lastRestart;
  }


  public TaskState restarts(Integer restarts) {
    
    this.restarts = restarts;
    return this;
  }

   /**
   * Get restarts
   * minimum: 0
   * maximum: 384
   * @return restarts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getRestarts() {
    return restarts;
  }


  public void setRestarts(Integer restarts) {
    this.restarts = restarts;
  }


  public TaskState startedAt(OffsetDateTime startedAt) {
    
    this.startedAt = startedAt;
    return this;
  }

   /**
   * Get startedAt
   * @return startedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getStartedAt() {
    return startedAt;
  }


  public void setStartedAt(OffsetDateTime startedAt) {
    this.startedAt = startedAt;
  }


  public TaskState state(String state) {
    
    this.state = state;
    return this;
  }

   /**
   * Get state
   * @return state
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getState() {
    return state;
  }


  public void setState(String state) {
    this.state = state;
  }


  public TaskState taskHandle(TaskHandle taskHandle) {
    
    this.taskHandle = taskHandle;
    return this;
  }

   /**
   * Get taskHandle
   * @return taskHandle
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public TaskHandle getTaskHandle() {
    return taskHandle;
  }


  public void setTaskHandle(TaskHandle taskHandle) {
    this.taskHandle = taskHandle;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskState taskState = (TaskState) o;
    return Objects.equals(this.events, taskState.events) &&
        Objects.equals(this.failed, taskState.failed) &&
        Objects.equals(this.finishedAt, taskState.finishedAt) &&
        Objects.equals(this.lastRestart, taskState.lastRestart) &&
        Objects.equals(this.restarts, taskState.restarts) &&
        Objects.equals(this.startedAt, taskState.startedAt) &&
        Objects.equals(this.state, taskState.state) &&
        Objects.equals(this.taskHandle, taskState.taskHandle);
  }

  @Override
  public int hashCode() {
    return Objects.hash(events, failed, finishedAt, lastRestart, restarts, startedAt, state, taskHandle);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskState {\n");
    sb.append("    events: ").append(toIndentedString(events)).append("\n");
    sb.append("    failed: ").append(toIndentedString(failed)).append("\n");
    sb.append("    finishedAt: ").append(toIndentedString(finishedAt)).append("\n");
    sb.append("    lastRestart: ").append(toIndentedString(lastRestart)).append("\n");
    sb.append("    restarts: ").append(toIndentedString(restarts)).append("\n");
    sb.append("    startedAt: ").append(toIndentedString(startedAt)).append("\n");
    sb.append("    state: ").append(toIndentedString(state)).append("\n");
    sb.append("    taskHandle: ").append(toIndentedString(taskHandle)).append("\n");
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

