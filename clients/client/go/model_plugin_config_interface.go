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

// PluginConfigInterface PluginConfigInterface The interface between Docker and the plugin
type PluginConfigInterface struct {
	// socket
	Socket string `json:"Socket"`
	// types
	Types []PluginInterfaceType `json:"Types"`
}

// NewPluginConfigInterface instantiates a new PluginConfigInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfigInterface(socket string, types []PluginInterfaceType) *PluginConfigInterface {
	this := PluginConfigInterface{}
	this.Socket = socket
	this.Types = types
	return &this
}

// NewPluginConfigInterfaceWithDefaults instantiates a new PluginConfigInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigInterfaceWithDefaults() *PluginConfigInterface {
	this := PluginConfigInterface{}
	return &this
}

// GetSocket returns the Socket field value
func (o *PluginConfigInterface) GetSocket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Socket
}

// GetSocketOk returns a tuple with the Socket field value
// and a boolean to check if the value has been set.
func (o *PluginConfigInterface) GetSocketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Socket, true
}

// SetSocket sets field value
func (o *PluginConfigInterface) SetSocket(v string) {
	o.Socket = v
}

// GetTypes returns the Types field value
func (o *PluginConfigInterface) GetTypes() []PluginInterfaceType {
	if o == nil {
		var ret []PluginInterfaceType
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *PluginConfigInterface) GetTypesOk() ([]PluginInterfaceType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *PluginConfigInterface) SetTypes(v []PluginInterfaceType) {
	o.Types = v
}

func (o PluginConfigInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Socket"] = o.Socket
	}
	if true {
		toSerialize["Types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullablePluginConfigInterface struct {
	value *PluginConfigInterface
	isSet bool
}

func (v NullablePluginConfigInterface) Get() *PluginConfigInterface {
	return v.value
}

func (v *NullablePluginConfigInterface) Set(val *PluginConfigInterface) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfigInterface) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfigInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfigInterface(val *PluginConfigInterface) *NullablePluginConfigInterface {
	return &NullablePluginConfigInterface{value: val, isSet: true}
}

func (v NullablePluginConfigInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfigInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


