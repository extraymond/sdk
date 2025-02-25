/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.15
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

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
 * ImageDeleteResponseItem image delete response item
 */
@ApiModel(description = "ImageDeleteResponseItem image delete response item")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-07-07T10:29:07.243350551Z[Etc/UTC]")
public class ImageDeleteResponseItem {
  public static final String SERIALIZED_NAME_DELETED = "Deleted";
  @SerializedName(SERIALIZED_NAME_DELETED)
  private String deleted;

  public static final String SERIALIZED_NAME_UNTAGGED = "Untagged";
  @SerializedName(SERIALIZED_NAME_UNTAGGED)
  private String untagged;


  public ImageDeleteResponseItem deleted(String deleted) {
    
    this.deleted = deleted;
    return this;
  }

   /**
   * The image ID of an image that was deleted
   * @return deleted
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The image ID of an image that was deleted")

  public String getDeleted() {
    return deleted;
  }


  public void setDeleted(String deleted) {
    this.deleted = deleted;
  }


  public ImageDeleteResponseItem untagged(String untagged) {
    
    this.untagged = untagged;
    return this;
  }

   /**
   * The image ID of an image that was untagged
   * @return untagged
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The image ID of an image that was untagged")

  public String getUntagged() {
    return untagged;
  }


  public void setUntagged(String untagged) {
    this.untagged = untagged;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ImageDeleteResponseItem imageDeleteResponseItem = (ImageDeleteResponseItem) o;
    return Objects.equals(this.deleted, imageDeleteResponseItem.deleted) &&
        Objects.equals(this.untagged, imageDeleteResponseItem.untagged);
  }

  @Override
  public int hashCode() {
    return Objects.hash(deleted, untagged);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ImageDeleteResponseItem {\n");
    sb.append("    deleted: ").append(toIndentedString(deleted)).append("\n");
    sb.append("    untagged: ").append(toIndentedString(untagged)).append("\n");
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

