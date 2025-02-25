/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.15
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateRecoveryLink struct for CreateRecoveryLink
type CreateRecoveryLink struct {
	// Link Expires In  The recovery link will expire at that point in time. Defaults to the configuration value of `selfservice.flows.recovery.request_lifespan`.
	ExpiresIn *string `json:"expires_in,omitempty"`
	IdentityId string `json:"identity_id"`
}

// NewCreateRecoveryLink instantiates a new CreateRecoveryLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecoveryLink(identityId string) *CreateRecoveryLink {
	this := CreateRecoveryLink{}
	this.IdentityId = identityId
	return &this
}

// NewCreateRecoveryLinkWithDefaults instantiates a new CreateRecoveryLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecoveryLinkWithDefaults() *CreateRecoveryLink {
	this := CreateRecoveryLink{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *CreateRecoveryLink) GetExpiresIn() string {
	if o == nil || o.ExpiresIn == nil {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecoveryLink) GetExpiresInOk() (*string, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *CreateRecoveryLink) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *CreateRecoveryLink) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetIdentityId returns the IdentityId field value
func (o *CreateRecoveryLink) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *CreateRecoveryLink) GetIdentityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *CreateRecoveryLink) SetIdentityId(v string) {
	o.IdentityId = v
}

func (o CreateRecoveryLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRecoveryLink struct {
	value *CreateRecoveryLink
	isSet bool
}

func (v NullableCreateRecoveryLink) Get() *CreateRecoveryLink {
	return v.value
}

func (v *NullableCreateRecoveryLink) Set(val *CreateRecoveryLink) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecoveryLink) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecoveryLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecoveryLink(val *CreateRecoveryLink) *NullableCreateRecoveryLink {
	return &NullableCreateRecoveryLink{value: val, isSet: true}
}

func (v NullableCreateRecoveryLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecoveryLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


