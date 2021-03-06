/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// Listener Information about the listener.
type Listener struct {
	// The port on which the back-end VM is listening (between `1` and `65535`, both included).
	BackendPort *int64 `json:"BackendPort,omitempty"`
	// The protocol for routing traffic to back-end VMs (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	BackendProtocol *string `json:"BackendProtocol,omitempty"`
	// The port on which the load balancer is listening (between 1 and `65535`, both included).
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty"`
	// The routing protocol (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	LoadBalancerProtocol *string `json:"LoadBalancerProtocol,omitempty"`
	// The names of the policies. If there are no policies enabled, the list is empty.
	PolicyNames *[]string `json:"PolicyNames,omitempty"`
	// The ID of the server certificate.
	ServerCertificateId *string `json:"ServerCertificateId,omitempty"`
}

// GetBackendPort returns the BackendPort field value if set, zero value otherwise.
func (o *Listener) GetBackendPort() int64 {
	if o == nil || o.BackendPort == nil {
		var ret int64
		return ret
	}
	return *o.BackendPort
}

// GetBackendPortOk returns a tuple with the BackendPort field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetBackendPortOk() (int64, bool) {
	if o == nil || o.BackendPort == nil {
		var ret int64
		return ret, false
	}
	return *o.BackendPort, true
}

// HasBackendPort returns a boolean if a field has been set.
func (o *Listener) HasBackendPort() bool {
	if o != nil && o.BackendPort != nil {
		return true
	}

	return false
}

// SetBackendPort gets a reference to the given int64 and assigns it to the BackendPort field.
func (o *Listener) SetBackendPort(v int64) {
	o.BackendPort = &v
}

// GetBackendProtocol returns the BackendProtocol field value if set, zero value otherwise.
func (o *Listener) GetBackendProtocol() string {
	if o == nil || o.BackendProtocol == nil {
		var ret string
		return ret
	}
	return *o.BackendProtocol
}

// GetBackendProtocolOk returns a tuple with the BackendProtocol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetBackendProtocolOk() (string, bool) {
	if o == nil || o.BackendProtocol == nil {
		var ret string
		return ret, false
	}
	return *o.BackendProtocol, true
}

// HasBackendProtocol returns a boolean if a field has been set.
func (o *Listener) HasBackendProtocol() bool {
	if o != nil && o.BackendProtocol != nil {
		return true
	}

	return false
}

// SetBackendProtocol gets a reference to the given string and assigns it to the BackendProtocol field.
func (o *Listener) SetBackendProtocol(v string) {
	o.BackendProtocol = &v
}

// GetLoadBalancerPort returns the LoadBalancerPort field value if set, zero value otherwise.
func (o *Listener) GetLoadBalancerPort() int64 {
	if o == nil || o.LoadBalancerPort == nil {
		var ret int64
		return ret
	}
	return *o.LoadBalancerPort
}

// GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetLoadBalancerPortOk() (int64, bool) {
	if o == nil || o.LoadBalancerPort == nil {
		var ret int64
		return ret, false
	}
	return *o.LoadBalancerPort, true
}

// HasLoadBalancerPort returns a boolean if a field has been set.
func (o *Listener) HasLoadBalancerPort() bool {
	if o != nil && o.LoadBalancerPort != nil {
		return true
	}

	return false
}

// SetLoadBalancerPort gets a reference to the given int64 and assigns it to the LoadBalancerPort field.
func (o *Listener) SetLoadBalancerPort(v int64) {
	o.LoadBalancerPort = &v
}

// GetLoadBalancerProtocol returns the LoadBalancerProtocol field value if set, zero value otherwise.
func (o *Listener) GetLoadBalancerProtocol() string {
	if o == nil || o.LoadBalancerProtocol == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerProtocol
}

// GetLoadBalancerProtocolOk returns a tuple with the LoadBalancerProtocol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetLoadBalancerProtocolOk() (string, bool) {
	if o == nil || o.LoadBalancerProtocol == nil {
		var ret string
		return ret, false
	}
	return *o.LoadBalancerProtocol, true
}

// HasLoadBalancerProtocol returns a boolean if a field has been set.
func (o *Listener) HasLoadBalancerProtocol() bool {
	if o != nil && o.LoadBalancerProtocol != nil {
		return true
	}

	return false
}

// SetLoadBalancerProtocol gets a reference to the given string and assigns it to the LoadBalancerProtocol field.
func (o *Listener) SetLoadBalancerProtocol(v string) {
	o.LoadBalancerProtocol = &v
}

// GetPolicyNames returns the PolicyNames field value if set, zero value otherwise.
func (o *Listener) GetPolicyNames() []string {
	if o == nil || o.PolicyNames == nil {
		var ret []string
		return ret
	}
	return *o.PolicyNames
}

// GetPolicyNamesOk returns a tuple with the PolicyNames field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetPolicyNamesOk() ([]string, bool) {
	if o == nil || o.PolicyNames == nil {
		var ret []string
		return ret, false
	}
	return *o.PolicyNames, true
}

// HasPolicyNames returns a boolean if a field has been set.
func (o *Listener) HasPolicyNames() bool {
	if o != nil && o.PolicyNames != nil {
		return true
	}

	return false
}

// SetPolicyNames gets a reference to the given []string and assigns it to the PolicyNames field.
func (o *Listener) SetPolicyNames(v []string) {
	o.PolicyNames = &v
}

// GetServerCertificateId returns the ServerCertificateId field value if set, zero value otherwise.
func (o *Listener) GetServerCertificateId() string {
	if o == nil || o.ServerCertificateId == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificateId
}

// GetServerCertificateIdOk returns a tuple with the ServerCertificateId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetServerCertificateIdOk() (string, bool) {
	if o == nil || o.ServerCertificateId == nil {
		var ret string
		return ret, false
	}
	return *o.ServerCertificateId, true
}

// HasServerCertificateId returns a boolean if a field has been set.
func (o *Listener) HasServerCertificateId() bool {
	if o != nil && o.ServerCertificateId != nil {
		return true
	}

	return false
}

// SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.
func (o *Listener) SetServerCertificateId(v string) {
	o.ServerCertificateId = &v
}

type NullableListener struct {
	Value        Listener
	ExplicitNull bool
}

func (v NullableListener) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableListener) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
