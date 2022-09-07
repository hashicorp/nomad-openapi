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
import io.nomadproject.client.models.AllocDeploymentStatus;
import io.nomadproject.client.models.AllocatedResources;
import io.nomadproject.client.models.RescheduleTracker;
import io.nomadproject.client.models.TaskState;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import org.openapitools.jackson.nullable.JsonNullable;

/**
 * AllocationListStub
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class AllocationListStub {
  public static final String SERIALIZED_NAME_ALLOCATED_RESOURCES = "AllocatedResources";
  @SerializedName(SERIALIZED_NAME_ALLOCATED_RESOURCES)
  private AllocatedResources allocatedResources;

  public static final String SERIALIZED_NAME_CLIENT_DESCRIPTION = "ClientDescription";
  @SerializedName(SERIALIZED_NAME_CLIENT_DESCRIPTION)
  private String clientDescription;

  public static final String SERIALIZED_NAME_CLIENT_STATUS = "ClientStatus";
  @SerializedName(SERIALIZED_NAME_CLIENT_STATUS)
  private String clientStatus;

  public static final String SERIALIZED_NAME_CREATE_INDEX = "CreateIndex";
  @SerializedName(SERIALIZED_NAME_CREATE_INDEX)
  private Integer createIndex;

  public static final String SERIALIZED_NAME_CREATE_TIME = "CreateTime";
  @SerializedName(SERIALIZED_NAME_CREATE_TIME)
  private Long createTime;

  public static final String SERIALIZED_NAME_DEPLOYMENT_STATUS = "DeploymentStatus";
  @SerializedName(SERIALIZED_NAME_DEPLOYMENT_STATUS)
  private AllocDeploymentStatus deploymentStatus;

  public static final String SERIALIZED_NAME_DESIRED_DESCRIPTION = "DesiredDescription";
  @SerializedName(SERIALIZED_NAME_DESIRED_DESCRIPTION)
  private String desiredDescription;

  public static final String SERIALIZED_NAME_DESIRED_STATUS = "DesiredStatus";
  @SerializedName(SERIALIZED_NAME_DESIRED_STATUS)
  private String desiredStatus;

  public static final String SERIALIZED_NAME_EVAL_I_D = "EvalID";
  @SerializedName(SERIALIZED_NAME_EVAL_I_D)
  private String evalID;

  public static final String SERIALIZED_NAME_FOLLOWUP_EVAL_I_D = "FollowupEvalID";
  @SerializedName(SERIALIZED_NAME_FOLLOWUP_EVAL_I_D)
  private String followupEvalID;

  public static final String SERIALIZED_NAME_I_D = "ID";
  @SerializedName(SERIALIZED_NAME_I_D)
  private String ID;

  public static final String SERIALIZED_NAME_JOB_I_D = "JobID";
  @SerializedName(SERIALIZED_NAME_JOB_I_D)
  private String jobID;

  public static final String SERIALIZED_NAME_JOB_TYPE = "JobType";
  @SerializedName(SERIALIZED_NAME_JOB_TYPE)
  private String jobType;

  public static final String SERIALIZED_NAME_JOB_VERSION = "JobVersion";
  @SerializedName(SERIALIZED_NAME_JOB_VERSION)
  private Integer jobVersion;

  public static final String SERIALIZED_NAME_MODIFY_INDEX = "ModifyIndex";
  @SerializedName(SERIALIZED_NAME_MODIFY_INDEX)
  private Integer modifyIndex;

  public static final String SERIALIZED_NAME_MODIFY_TIME = "ModifyTime";
  @SerializedName(SERIALIZED_NAME_MODIFY_TIME)
  private Long modifyTime;

  public static final String SERIALIZED_NAME_NAME = "Name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_NAMESPACE = "Namespace";
  @SerializedName(SERIALIZED_NAME_NAMESPACE)
  private String namespace;

  public static final String SERIALIZED_NAME_NODE_I_D = "NodeID";
  @SerializedName(SERIALIZED_NAME_NODE_I_D)
  private String nodeID;

  public static final String SERIALIZED_NAME_NODE_NAME = "NodeName";
  @SerializedName(SERIALIZED_NAME_NODE_NAME)
  private String nodeName;

  public static final String SERIALIZED_NAME_PREEMPTED_ALLOCATIONS = "PreemptedAllocations";
  @SerializedName(SERIALIZED_NAME_PREEMPTED_ALLOCATIONS)
  private List<String> preemptedAllocations = null;

  public static final String SERIALIZED_NAME_PREEMPTED_BY_ALLOCATION = "PreemptedByAllocation";
  @SerializedName(SERIALIZED_NAME_PREEMPTED_BY_ALLOCATION)
  private String preemptedByAllocation;

  public static final String SERIALIZED_NAME_RESCHEDULE_TRACKER = "RescheduleTracker";
  @SerializedName(SERIALIZED_NAME_RESCHEDULE_TRACKER)
  private RescheduleTracker rescheduleTracker;

  public static final String SERIALIZED_NAME_TASK_GROUP = "TaskGroup";
  @SerializedName(SERIALIZED_NAME_TASK_GROUP)
  private String taskGroup;

  public static final String SERIALIZED_NAME_TASK_STATES = "TaskStates";
  @SerializedName(SERIALIZED_NAME_TASK_STATES)
  private Map<String, TaskState> taskStates = null;

  public AllocationListStub() { 
  }

  public AllocationListStub allocatedResources(AllocatedResources allocatedResources) {
    
    this.allocatedResources = allocatedResources;
    return this;
  }

   /**
   * Get allocatedResources
   * @return allocatedResources
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public AllocatedResources getAllocatedResources() {
    return allocatedResources;
  }


  public void setAllocatedResources(AllocatedResources allocatedResources) {
    this.allocatedResources = allocatedResources;
  }


  public AllocationListStub clientDescription(String clientDescription) {
    
    this.clientDescription = clientDescription;
    return this;
  }

   /**
   * Get clientDescription
   * @return clientDescription
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientDescription() {
    return clientDescription;
  }


  public void setClientDescription(String clientDescription) {
    this.clientDescription = clientDescription;
  }


  public AllocationListStub clientStatus(String clientStatus) {
    
    this.clientStatus = clientStatus;
    return this;
  }

   /**
   * Get clientStatus
   * @return clientStatus
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientStatus() {
    return clientStatus;
  }


  public void setClientStatus(String clientStatus) {
    this.clientStatus = clientStatus;
  }


  public AllocationListStub createIndex(Integer createIndex) {
    
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


  public AllocationListStub createTime(Long createTime) {
    
    this.createTime = createTime;
    return this;
  }

   /**
   * Get createTime
   * @return createTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCreateTime() {
    return createTime;
  }


  public void setCreateTime(Long createTime) {
    this.createTime = createTime;
  }


  public AllocationListStub deploymentStatus(AllocDeploymentStatus deploymentStatus) {
    
    this.deploymentStatus = deploymentStatus;
    return this;
  }

   /**
   * Get deploymentStatus
   * @return deploymentStatus
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public AllocDeploymentStatus getDeploymentStatus() {
    return deploymentStatus;
  }


  public void setDeploymentStatus(AllocDeploymentStatus deploymentStatus) {
    this.deploymentStatus = deploymentStatus;
  }


  public AllocationListStub desiredDescription(String desiredDescription) {
    
    this.desiredDescription = desiredDescription;
    return this;
  }

   /**
   * Get desiredDescription
   * @return desiredDescription
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDesiredDescription() {
    return desiredDescription;
  }


  public void setDesiredDescription(String desiredDescription) {
    this.desiredDescription = desiredDescription;
  }


  public AllocationListStub desiredStatus(String desiredStatus) {
    
    this.desiredStatus = desiredStatus;
    return this;
  }

   /**
   * Get desiredStatus
   * @return desiredStatus
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDesiredStatus() {
    return desiredStatus;
  }


  public void setDesiredStatus(String desiredStatus) {
    this.desiredStatus = desiredStatus;
  }


  public AllocationListStub evalID(String evalID) {
    
    this.evalID = evalID;
    return this;
  }

   /**
   * Get evalID
   * @return evalID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEvalID() {
    return evalID;
  }


  public void setEvalID(String evalID) {
    this.evalID = evalID;
  }


  public AllocationListStub followupEvalID(String followupEvalID) {
    
    this.followupEvalID = followupEvalID;
    return this;
  }

   /**
   * Get followupEvalID
   * @return followupEvalID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFollowupEvalID() {
    return followupEvalID;
  }


  public void setFollowupEvalID(String followupEvalID) {
    this.followupEvalID = followupEvalID;
  }


  public AllocationListStub ID(String ID) {
    
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


  public AllocationListStub jobID(String jobID) {
    
    this.jobID = jobID;
    return this;
  }

   /**
   * Get jobID
   * @return jobID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getJobID() {
    return jobID;
  }


  public void setJobID(String jobID) {
    this.jobID = jobID;
  }


  public AllocationListStub jobType(String jobType) {
    
    this.jobType = jobType;
    return this;
  }

   /**
   * Get jobType
   * @return jobType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getJobType() {
    return jobType;
  }


  public void setJobType(String jobType) {
    this.jobType = jobType;
  }


  public AllocationListStub jobVersion(Integer jobVersion) {
    
    this.jobVersion = jobVersion;
    return this;
  }

   /**
   * Get jobVersion
   * minimum: 0
   * maximum: 384
   * @return jobVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getJobVersion() {
    return jobVersion;
  }


  public void setJobVersion(Integer jobVersion) {
    this.jobVersion = jobVersion;
  }


  public AllocationListStub modifyIndex(Integer modifyIndex) {
    
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


  public AllocationListStub modifyTime(Long modifyTime) {
    
    this.modifyTime = modifyTime;
    return this;
  }

   /**
   * Get modifyTime
   * @return modifyTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getModifyTime() {
    return modifyTime;
  }


  public void setModifyTime(Long modifyTime) {
    this.modifyTime = modifyTime;
  }


  public AllocationListStub name(String name) {
    
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


  public AllocationListStub namespace(String namespace) {
    
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


  public AllocationListStub nodeID(String nodeID) {
    
    this.nodeID = nodeID;
    return this;
  }

   /**
   * Get nodeID
   * @return nodeID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNodeID() {
    return nodeID;
  }


  public void setNodeID(String nodeID) {
    this.nodeID = nodeID;
  }


  public AllocationListStub nodeName(String nodeName) {
    
    this.nodeName = nodeName;
    return this;
  }

   /**
   * Get nodeName
   * @return nodeName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNodeName() {
    return nodeName;
  }


  public void setNodeName(String nodeName) {
    this.nodeName = nodeName;
  }


  public AllocationListStub preemptedAllocations(List<String> preemptedAllocations) {
    
    this.preemptedAllocations = preemptedAllocations;
    return this;
  }

  public AllocationListStub addPreemptedAllocationsItem(String preemptedAllocationsItem) {
    if (this.preemptedAllocations == null) {
      this.preemptedAllocations = new ArrayList<String>();
    }
    this.preemptedAllocations.add(preemptedAllocationsItem);
    return this;
  }

   /**
   * Get preemptedAllocations
   * @return preemptedAllocations
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getPreemptedAllocations() {
    return preemptedAllocations;
  }


  public void setPreemptedAllocations(List<String> preemptedAllocations) {
    this.preemptedAllocations = preemptedAllocations;
  }


  public AllocationListStub preemptedByAllocation(String preemptedByAllocation) {
    
    this.preemptedByAllocation = preemptedByAllocation;
    return this;
  }

   /**
   * Get preemptedByAllocation
   * @return preemptedByAllocation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPreemptedByAllocation() {
    return preemptedByAllocation;
  }


  public void setPreemptedByAllocation(String preemptedByAllocation) {
    this.preemptedByAllocation = preemptedByAllocation;
  }


  public AllocationListStub rescheduleTracker(RescheduleTracker rescheduleTracker) {
    
    this.rescheduleTracker = rescheduleTracker;
    return this;
  }

   /**
   * Get rescheduleTracker
   * @return rescheduleTracker
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public RescheduleTracker getRescheduleTracker() {
    return rescheduleTracker;
  }


  public void setRescheduleTracker(RescheduleTracker rescheduleTracker) {
    this.rescheduleTracker = rescheduleTracker;
  }


  public AllocationListStub taskGroup(String taskGroup) {
    
    this.taskGroup = taskGroup;
    return this;
  }

   /**
   * Get taskGroup
   * @return taskGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTaskGroup() {
    return taskGroup;
  }


  public void setTaskGroup(String taskGroup) {
    this.taskGroup = taskGroup;
  }


  public AllocationListStub taskStates(Map<String, TaskState> taskStates) {
    
    this.taskStates = taskStates;
    return this;
  }

  public AllocationListStub putTaskStatesItem(String key, TaskState taskStatesItem) {
    if (this.taskStates == null) {
      this.taskStates = new HashMap<String, TaskState>();
    }
    this.taskStates.put(key, taskStatesItem);
    return this;
  }

   /**
   * Get taskStates
   * @return taskStates
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, TaskState> getTaskStates() {
    return taskStates;
  }


  public void setTaskStates(Map<String, TaskState> taskStates) {
    this.taskStates = taskStates;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AllocationListStub allocationListStub = (AllocationListStub) o;
    return Objects.equals(this.allocatedResources, allocationListStub.allocatedResources) &&
        Objects.equals(this.clientDescription, allocationListStub.clientDescription) &&
        Objects.equals(this.clientStatus, allocationListStub.clientStatus) &&
        Objects.equals(this.createIndex, allocationListStub.createIndex) &&
        Objects.equals(this.createTime, allocationListStub.createTime) &&
        Objects.equals(this.deploymentStatus, allocationListStub.deploymentStatus) &&
        Objects.equals(this.desiredDescription, allocationListStub.desiredDescription) &&
        Objects.equals(this.desiredStatus, allocationListStub.desiredStatus) &&
        Objects.equals(this.evalID, allocationListStub.evalID) &&
        Objects.equals(this.followupEvalID, allocationListStub.followupEvalID) &&
        Objects.equals(this.ID, allocationListStub.ID) &&
        Objects.equals(this.jobID, allocationListStub.jobID) &&
        Objects.equals(this.jobType, allocationListStub.jobType) &&
        Objects.equals(this.jobVersion, allocationListStub.jobVersion) &&
        Objects.equals(this.modifyIndex, allocationListStub.modifyIndex) &&
        Objects.equals(this.modifyTime, allocationListStub.modifyTime) &&
        Objects.equals(this.name, allocationListStub.name) &&
        Objects.equals(this.namespace, allocationListStub.namespace) &&
        Objects.equals(this.nodeID, allocationListStub.nodeID) &&
        Objects.equals(this.nodeName, allocationListStub.nodeName) &&
        Objects.equals(this.preemptedAllocations, allocationListStub.preemptedAllocations) &&
        Objects.equals(this.preemptedByAllocation, allocationListStub.preemptedByAllocation) &&
        Objects.equals(this.rescheduleTracker, allocationListStub.rescheduleTracker) &&
        Objects.equals(this.taskGroup, allocationListStub.taskGroup) &&
        Objects.equals(this.taskStates, allocationListStub.taskStates);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(allocatedResources, clientDescription, clientStatus, createIndex, createTime, deploymentStatus, desiredDescription, desiredStatus, evalID, followupEvalID, ID, jobID, jobType, jobVersion, modifyIndex, modifyTime, name, namespace, nodeID, nodeName, preemptedAllocations, preemptedByAllocation, rescheduleTracker, taskGroup, taskStates);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AllocationListStub {\n");
    sb.append("    allocatedResources: ").append(toIndentedString(allocatedResources)).append("\n");
    sb.append("    clientDescription: ").append(toIndentedString(clientDescription)).append("\n");
    sb.append("    clientStatus: ").append(toIndentedString(clientStatus)).append("\n");
    sb.append("    createIndex: ").append(toIndentedString(createIndex)).append("\n");
    sb.append("    createTime: ").append(toIndentedString(createTime)).append("\n");
    sb.append("    deploymentStatus: ").append(toIndentedString(deploymentStatus)).append("\n");
    sb.append("    desiredDescription: ").append(toIndentedString(desiredDescription)).append("\n");
    sb.append("    desiredStatus: ").append(toIndentedString(desiredStatus)).append("\n");
    sb.append("    evalID: ").append(toIndentedString(evalID)).append("\n");
    sb.append("    followupEvalID: ").append(toIndentedString(followupEvalID)).append("\n");
    sb.append("    ID: ").append(toIndentedString(ID)).append("\n");
    sb.append("    jobID: ").append(toIndentedString(jobID)).append("\n");
    sb.append("    jobType: ").append(toIndentedString(jobType)).append("\n");
    sb.append("    jobVersion: ").append(toIndentedString(jobVersion)).append("\n");
    sb.append("    modifyIndex: ").append(toIndentedString(modifyIndex)).append("\n");
    sb.append("    modifyTime: ").append(toIndentedString(modifyTime)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    namespace: ").append(toIndentedString(namespace)).append("\n");
    sb.append("    nodeID: ").append(toIndentedString(nodeID)).append("\n");
    sb.append("    nodeName: ").append(toIndentedString(nodeName)).append("\n");
    sb.append("    preemptedAllocations: ").append(toIndentedString(preemptedAllocations)).append("\n");
    sb.append("    preemptedByAllocation: ").append(toIndentedString(preemptedByAllocation)).append("\n");
    sb.append("    rescheduleTracker: ").append(toIndentedString(rescheduleTracker)).append("\n");
    sb.append("    taskGroup: ").append(toIndentedString(taskGroup)).append("\n");
    sb.append("    taskStates: ").append(toIndentedString(taskStates)).append("\n");
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

