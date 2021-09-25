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
 * OneTimeTokenExchangeRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class OneTimeTokenExchangeRequest {
  public static final String SERIALIZED_NAME_ONE_TIME_SECRET_I_D = "OneTimeSecretID";
  @SerializedName(SERIALIZED_NAME_ONE_TIME_SECRET_I_D)
  private String oneTimeSecretID;


  public OneTimeTokenExchangeRequest oneTimeSecretID(String oneTimeSecretID) {
    
    this.oneTimeSecretID = oneTimeSecretID;
    return this;
  }

   /**
   * Get oneTimeSecretID
   * @return oneTimeSecretID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOneTimeSecretID() {
    return oneTimeSecretID;
  }


  public void setOneTimeSecretID(String oneTimeSecretID) {
    this.oneTimeSecretID = oneTimeSecretID;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    OneTimeTokenExchangeRequest oneTimeTokenExchangeRequest = (OneTimeTokenExchangeRequest) o;
    return Objects.equals(this.oneTimeSecretID, oneTimeTokenExchangeRequest.oneTimeSecretID);
  }

  @Override
  public int hashCode() {
    return Objects.hash(oneTimeSecretID);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OneTimeTokenExchangeRequest {\n");
    sb.append("    oneTimeSecretID: ").append(toIndentedString(oneTimeSecretID)).append("\n");
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

