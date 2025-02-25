/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Kratos.Client.Client.OpenAPIDateConverter;

namespace Ory.Kratos.Client.Model
{
    /// <summary>
    /// VolumeUsageData Usage details about the volume. This information is used by the &#x60;GET /system/df&#x60; endpoint, and omitted in other endpoints.
    /// </summary>
    [DataContract(Name = "VolumeUsageData")]
    public partial class KratosVolumeUsageData : IEquatable<KratosVolumeUsageData>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosVolumeUsageData" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosVolumeUsageData()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosVolumeUsageData" /> class.
        /// </summary>
        /// <param name="refCount">The number of containers referencing this volume. This field is set to &#x60;-1&#x60; if the reference-count is not available. (required).</param>
        /// <param name="size">Amount of disk space used by the volume (in bytes). This information is only available for volumes created with the &#x60;\&quot;local\&quot;&#x60; volume driver. For volumes created with other volume drivers, this field is set to &#x60;-1&#x60; (\&quot;not available\&quot;) (required).</param>
        public KratosVolumeUsageData(long refCount = default(long), long size = default(long))
        {
            this.RefCount = refCount;
            this.Size = size;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The number of containers referencing this volume. This field is set to &#x60;-1&#x60; if the reference-count is not available.
        /// </summary>
        /// <value>The number of containers referencing this volume. This field is set to &#x60;-1&#x60; if the reference-count is not available.</value>
        [DataMember(Name = "RefCount", IsRequired = true, EmitDefaultValue = false)]
        public long RefCount { get; set; }

        /// <summary>
        /// Amount of disk space used by the volume (in bytes). This information is only available for volumes created with the &#x60;\&quot;local\&quot;&#x60; volume driver. For volumes created with other volume drivers, this field is set to &#x60;-1&#x60; (\&quot;not available\&quot;)
        /// </summary>
        /// <value>Amount of disk space used by the volume (in bytes). This information is only available for volumes created with the &#x60;\&quot;local\&quot;&#x60; volume driver. For volumes created with other volume drivers, this field is set to &#x60;-1&#x60; (\&quot;not available\&quot;)</value>
        [DataMember(Name = "Size", IsRequired = true, EmitDefaultValue = false)]
        public long Size { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class KratosVolumeUsageData {\n");
            sb.Append("  RefCount: ").Append(RefCount).Append("\n");
            sb.Append("  Size: ").Append(Size).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as KratosVolumeUsageData);
        }

        /// <summary>
        /// Returns true if KratosVolumeUsageData instances are equal
        /// </summary>
        /// <param name="input">Instance of KratosVolumeUsageData to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(KratosVolumeUsageData input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.RefCount == input.RefCount ||
                    this.RefCount.Equals(input.RefCount)
                ) && 
                (
                    this.Size == input.Size ||
                    this.Size.Equals(input.Size)
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                hashCode = hashCode * 59 + this.RefCount.GetHashCode();
                hashCode = hashCode * 59 + this.Size.GetHashCode();
                if (this.AdditionalProperties != null)
                    hashCode = hashCode * 59 + this.AdditionalProperties.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
