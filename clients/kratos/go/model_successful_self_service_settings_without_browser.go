/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SuccessfulSelfServiceSettingsWithoutBrowser The Response for Settings Flows via API
type SuccessfulSelfServiceSettingsWithoutBrowser struct {
	Flow SelfServiceSettingsFlow `json:"flow"`
	Identity Identity `json:"identity"`
}

// NewSuccessfulSelfServiceSettingsWithoutBrowser instantiates a new SuccessfulSelfServiceSettingsWithoutBrowser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulSelfServiceSettingsWithoutBrowser(flow SelfServiceSettingsFlow, identity Identity) *SuccessfulSelfServiceSettingsWithoutBrowser {
	this := SuccessfulSelfServiceSettingsWithoutBrowser{}
	this.Flow = flow
	this.Identity = identity
	return &this
}

// NewSuccessfulSelfServiceSettingsWithoutBrowserWithDefaults instantiates a new SuccessfulSelfServiceSettingsWithoutBrowser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulSelfServiceSettingsWithoutBrowserWithDefaults() *SuccessfulSelfServiceSettingsWithoutBrowser {
	this := SuccessfulSelfServiceSettingsWithoutBrowser{}
	return &this
}

// GetFlow returns the Flow field value
func (o *SuccessfulSelfServiceSettingsWithoutBrowser) GetFlow() SelfServiceSettingsFlow {
	if o == nil {
		var ret SelfServiceSettingsFlow
		return ret
	}

	return o.Flow
}

// GetFlowOk returns a tuple with the Flow field value
// and a boolean to check if the value has been set.
func (o *SuccessfulSelfServiceSettingsWithoutBrowser) GetFlowOk() (*SelfServiceSettingsFlow, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Flow, true
}

// SetFlow sets field value
func (o *SuccessfulSelfServiceSettingsWithoutBrowser) SetFlow(v SelfServiceSettingsFlow) {
	o.Flow = v
}

// GetIdentity returns the Identity field value
func (o *SuccessfulSelfServiceSettingsWithoutBrowser) GetIdentity() Identity {
	if o == nil {
		var ret Identity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *SuccessfulSelfServiceSettingsWithoutBrowser) GetIdentityOk() (*Identity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *SuccessfulSelfServiceSettingsWithoutBrowser) SetIdentity(v Identity) {
	o.Identity = v
}

func (o SuccessfulSelfServiceSettingsWithoutBrowser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["flow"] = o.Flow
	}
	if true {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessfulSelfServiceSettingsWithoutBrowser struct {
	value *SuccessfulSelfServiceSettingsWithoutBrowser
	isSet bool
}

func (v NullableSuccessfulSelfServiceSettingsWithoutBrowser) Get() *SuccessfulSelfServiceSettingsWithoutBrowser {
	return v.value
}

func (v *NullableSuccessfulSelfServiceSettingsWithoutBrowser) Set(val *SuccessfulSelfServiceSettingsWithoutBrowser) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessfulSelfServiceSettingsWithoutBrowser) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessfulSelfServiceSettingsWithoutBrowser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessfulSelfServiceSettingsWithoutBrowser(val *SuccessfulSelfServiceSettingsWithoutBrowser) *NullableSuccessfulSelfServiceSettingsWithoutBrowser {
	return &NullableSuccessfulSelfServiceSettingsWithoutBrowser{value: val, isSet: true}
}

func (v NullableSuccessfulSelfServiceSettingsWithoutBrowser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessfulSelfServiceSettingsWithoutBrowser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


