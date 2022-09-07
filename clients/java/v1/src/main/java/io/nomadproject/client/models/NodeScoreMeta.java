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
 * NodeScoreMeta
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NodeScoreMeta {
  public static final String SERIALIZED_NAME_NODE_I_D = "NodeID";
  @SerializedName(SERIALIZED_NAME_NODE_I_D)
  private String nodeID;

  public static final String SERIALIZED_NAME_NORM_SCORE = "NormScore";
  @SerializedName(SERIALIZED_NAME_NORM_SCORE)
  private Double normScore;

  public static final String SERIALIZED_NAME_SCORES = "Scores";
  @SerializedName(SERIALIZED_NAME_SCORES)
  private Map<String, Double> scores = null;

  public NodeScoreMeta() { 
  }

  public NodeScoreMeta nodeID(String nodeID) {
    
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


  public NodeScoreMeta normScore(Double normScore) {
    
    this.normScore = normScore;
    return this;
  }

   /**
   * Get normScore
   * @return normScore
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Double getNormScore() {
    return normScore;
  }


  public void setNormScore(Double normScore) {
    this.normScore = normScore;
  }


  public NodeScoreMeta scores(Map<String, Double> scores) {
    
    this.scores = scores;
    return this;
  }

  public NodeScoreMeta putScoresItem(String key, Double scoresItem) {
    if (this.scores == null) {
      this.scores = new HashMap<String, Double>();
    }
    this.scores.put(key, scoresItem);
    return this;
  }

   /**
   * Get scores
   * @return scores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Double> getScores() {
    return scores;
  }


  public void setScores(Map<String, Double> scores) {
    this.scores = scores;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NodeScoreMeta nodeScoreMeta = (NodeScoreMeta) o;
    return Objects.equals(this.nodeID, nodeScoreMeta.nodeID) &&
        Objects.equals(this.normScore, nodeScoreMeta.normScore) &&
        Objects.equals(this.scores, nodeScoreMeta.scores);
  }

  @Override
  public int hashCode() {
    return Objects.hash(nodeID, normScore, scores);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NodeScoreMeta {\n");
    sb.append("    nodeID: ").append(toIndentedString(nodeID)).append("\n");
    sb.append("    normScore: ").append(toIndentedString(normScore)).append("\n");
    sb.append("    scores: ").append(toIndentedString(scores)).append("\n");
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

