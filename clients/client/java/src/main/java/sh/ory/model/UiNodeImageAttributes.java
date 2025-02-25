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
 * UiNodeImageAttributes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-07-07T10:29:07.243350551Z[Etc/UTC]")
public class UiNodeImageAttributes {
  public static final String SERIALIZED_NAME_SRC = "src";
  @SerializedName(SERIALIZED_NAME_SRC)
  private String src;


  public UiNodeImageAttributes src(String src) {
    
    this.src = src;
    return this;
  }

   /**
   * The image&#39;s source URL.  format: uri
   * @return src
  **/
  @ApiModelProperty(required = true, value = "The image's source URL.  format: uri")

  public String getSrc() {
    return src;
  }


  public void setSrc(String src) {
    this.src = src;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UiNodeImageAttributes uiNodeImageAttributes = (UiNodeImageAttributes) o;
    return Objects.equals(this.src, uiNodeImageAttributes.src);
  }

  @Override
  public int hashCode() {
    return Objects.hash(src);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UiNodeImageAttributes {\n");
    sb.append("    src: ").append(toIndentedString(src)).append("\n");
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

