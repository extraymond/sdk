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

// PluginConfig struct for PluginConfig
type PluginConfig struct {
	Args PluginConfigArgs `json:"Args"`
	// description
	Description string `json:"Description"`
	// Docker Version used to create the plugin
	DockerVersion *string `json:"DockerVersion,omitempty"`
	// documentation
	Documentation string `json:"Documentation"`
	// entrypoint
	Entrypoint []string `json:"Entrypoint"`
	// env
	Env []PluginEnv `json:"Env"`
	Interface PluginConfigInterface `json:"Interface"`
	// ipc host
	IpcHost bool `json:"IpcHost"`
	Linux PluginConfigLinux `json:"Linux"`
	// mounts
	Mounts []PluginMount `json:"Mounts"`
	Network PluginConfigNetwork `json:"Network"`
	// pid host
	PidHost bool `json:"PidHost"`
	// propagated mount
	PropagatedMount string `json:"PropagatedMount"`
	User *PluginConfigUser `json:"User,omitempty"`
	// work dir
	WorkDir string `json:"WorkDir"`
	Rootfs *PluginConfigRootfs `json:"rootfs,omitempty"`
}

// NewPluginConfig instantiates a new PluginConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfig(args PluginConfigArgs, description string, documentation string, entrypoint []string, env []PluginEnv, interface_ PluginConfigInterface, ipcHost bool, linux PluginConfigLinux, mounts []PluginMount, network PluginConfigNetwork, pidHost bool, propagatedMount string, workDir string) *PluginConfig {
	this := PluginConfig{}
	this.Args = args
	this.Description = description
	this.Documentation = documentation
	this.Entrypoint = entrypoint
	this.Env = env
	this.Interface = interface_
	this.IpcHost = ipcHost
	this.Linux = linux
	this.Mounts = mounts
	this.Network = network
	this.PidHost = pidHost
	this.PropagatedMount = propagatedMount
	this.WorkDir = workDir
	return &this
}

// NewPluginConfigWithDefaults instantiates a new PluginConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigWithDefaults() *PluginConfig {
	this := PluginConfig{}
	return &this
}

// GetArgs returns the Args field value
func (o *PluginConfig) GetArgs() PluginConfigArgs {
	if o == nil {
		var ret PluginConfigArgs
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetArgsOk() (*PluginConfigArgs, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *PluginConfig) SetArgs(v PluginConfigArgs) {
	o.Args = v
}

// GetDescription returns the Description field value
func (o *PluginConfig) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PluginConfig) SetDescription(v string) {
	o.Description = v
}

// GetDockerVersion returns the DockerVersion field value if set, zero value otherwise.
func (o *PluginConfig) GetDockerVersion() string {
	if o == nil || o.DockerVersion == nil {
		var ret string
		return ret
	}
	return *o.DockerVersion
}

// GetDockerVersionOk returns a tuple with the DockerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetDockerVersionOk() (*string, bool) {
	if o == nil || o.DockerVersion == nil {
		return nil, false
	}
	return o.DockerVersion, true
}

// HasDockerVersion returns a boolean if a field has been set.
func (o *PluginConfig) HasDockerVersion() bool {
	if o != nil && o.DockerVersion != nil {
		return true
	}

	return false
}

// SetDockerVersion gets a reference to the given string and assigns it to the DockerVersion field.
func (o *PluginConfig) SetDockerVersion(v string) {
	o.DockerVersion = &v
}

// GetDocumentation returns the Documentation field value
func (o *PluginConfig) GetDocumentation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetDocumentationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Documentation, true
}

// SetDocumentation sets field value
func (o *PluginConfig) SetDocumentation(v string) {
	o.Documentation = v
}

// GetEntrypoint returns the Entrypoint field value
func (o *PluginConfig) GetEntrypoint() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetEntrypointOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Entrypoint, true
}

// SetEntrypoint sets field value
func (o *PluginConfig) SetEntrypoint(v []string) {
	o.Entrypoint = v
}

// GetEnv returns the Env field value
func (o *PluginConfig) GetEnv() []PluginEnv {
	if o == nil {
		var ret []PluginEnv
		return ret
	}

	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetEnvOk() ([]PluginEnv, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Env, true
}

// SetEnv sets field value
func (o *PluginConfig) SetEnv(v []PluginEnv) {
	o.Env = v
}

// GetInterface returns the Interface field value
func (o *PluginConfig) GetInterface() PluginConfigInterface {
	if o == nil {
		var ret PluginConfigInterface
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetInterfaceOk() (*PluginConfigInterface, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *PluginConfig) SetInterface(v PluginConfigInterface) {
	o.Interface = v
}

// GetIpcHost returns the IpcHost field value
func (o *PluginConfig) GetIpcHost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IpcHost
}

// GetIpcHostOk returns a tuple with the IpcHost field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetIpcHostOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpcHost, true
}

// SetIpcHost sets field value
func (o *PluginConfig) SetIpcHost(v bool) {
	o.IpcHost = v
}

// GetLinux returns the Linux field value
func (o *PluginConfig) GetLinux() PluginConfigLinux {
	if o == nil {
		var ret PluginConfigLinux
		return ret
	}

	return o.Linux
}

// GetLinuxOk returns a tuple with the Linux field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetLinuxOk() (*PluginConfigLinux, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Linux, true
}

// SetLinux sets field value
func (o *PluginConfig) SetLinux(v PluginConfigLinux) {
	o.Linux = v
}

// GetMounts returns the Mounts field value
func (o *PluginConfig) GetMounts() []PluginMount {
	if o == nil {
		var ret []PluginMount
		return ret
	}

	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetMountsOk() ([]PluginMount, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mounts, true
}

// SetMounts sets field value
func (o *PluginConfig) SetMounts(v []PluginMount) {
	o.Mounts = v
}

// GetNetwork returns the Network field value
func (o *PluginConfig) GetNetwork() PluginConfigNetwork {
	if o == nil {
		var ret PluginConfigNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetNetworkOk() (*PluginConfigNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *PluginConfig) SetNetwork(v PluginConfigNetwork) {
	o.Network = v
}

// GetPidHost returns the PidHost field value
func (o *PluginConfig) GetPidHost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PidHost
}

// GetPidHostOk returns a tuple with the PidHost field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetPidHostOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PidHost, true
}

// SetPidHost sets field value
func (o *PluginConfig) SetPidHost(v bool) {
	o.PidHost = v
}

// GetPropagatedMount returns the PropagatedMount field value
func (o *PluginConfig) GetPropagatedMount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropagatedMount
}

// GetPropagatedMountOk returns a tuple with the PropagatedMount field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetPropagatedMountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropagatedMount, true
}

// SetPropagatedMount sets field value
func (o *PluginConfig) SetPropagatedMount(v string) {
	o.PropagatedMount = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PluginConfig) GetUser() PluginConfigUser {
	if o == nil || o.User == nil {
		var ret PluginConfigUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetUserOk() (*PluginConfigUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PluginConfig) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given PluginConfigUser and assigns it to the User field.
func (o *PluginConfig) SetUser(v PluginConfigUser) {
	o.User = &v
}

// GetWorkDir returns the WorkDir field value
func (o *PluginConfig) GetWorkDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkDir
}

// GetWorkDirOk returns a tuple with the WorkDir field value
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetWorkDirOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkDir, true
}

// SetWorkDir sets field value
func (o *PluginConfig) SetWorkDir(v string) {
	o.WorkDir = v
}

// GetRootfs returns the Rootfs field value if set, zero value otherwise.
func (o *PluginConfig) GetRootfs() PluginConfigRootfs {
	if o == nil || o.Rootfs == nil {
		var ret PluginConfigRootfs
		return ret
	}
	return *o.Rootfs
}

// GetRootfsOk returns a tuple with the Rootfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfig) GetRootfsOk() (*PluginConfigRootfs, bool) {
	if o == nil || o.Rootfs == nil {
		return nil, false
	}
	return o.Rootfs, true
}

// HasRootfs returns a boolean if a field has been set.
func (o *PluginConfig) HasRootfs() bool {
	if o != nil && o.Rootfs != nil {
		return true
	}

	return false
}

// SetRootfs gets a reference to the given PluginConfigRootfs and assigns it to the Rootfs field.
func (o *PluginConfig) SetRootfs(v PluginConfigRootfs) {
	o.Rootfs = &v
}

func (o PluginConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Args"] = o.Args
	}
	if true {
		toSerialize["Description"] = o.Description
	}
	if o.DockerVersion != nil {
		toSerialize["DockerVersion"] = o.DockerVersion
	}
	if true {
		toSerialize["Documentation"] = o.Documentation
	}
	if true {
		toSerialize["Entrypoint"] = o.Entrypoint
	}
	if true {
		toSerialize["Env"] = o.Env
	}
	if true {
		toSerialize["Interface"] = o.Interface
	}
	if true {
		toSerialize["IpcHost"] = o.IpcHost
	}
	if true {
		toSerialize["Linux"] = o.Linux
	}
	if true {
		toSerialize["Mounts"] = o.Mounts
	}
	if true {
		toSerialize["Network"] = o.Network
	}
	if true {
		toSerialize["PidHost"] = o.PidHost
	}
	if true {
		toSerialize["PropagatedMount"] = o.PropagatedMount
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}
	if true {
		toSerialize["WorkDir"] = o.WorkDir
	}
	if o.Rootfs != nil {
		toSerialize["rootfs"] = o.Rootfs
	}
	return json.Marshal(toSerialize)
}

type NullablePluginConfig struct {
	value *PluginConfig
	isSet bool
}

func (v NullablePluginConfig) Get() *PluginConfig {
	return v.value
}

func (v *NullablePluginConfig) Set(val *PluginConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfig(val *PluginConfig) *NullablePluginConfig {
	return &NullablePluginConfig{value: val, isSet: true}
}

func (v NullablePluginConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


