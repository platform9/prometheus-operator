// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
)

// SafeTLSConfigApplyConfiguration represents a declarative configuration of the SafeTLSConfig type for use
// with apply.
type SafeTLSConfigApplyConfiguration struct {
	CA                 *SecretOrConfigMapApplyConfiguration `json:"ca,omitempty"`
	Cert               *SecretOrConfigMapApplyConfiguration `json:"cert,omitempty"`
	KeySecret          *corev1.SecretKeySelector            `json:"keySecret,omitempty"`
	ServerName         *string                              `json:"serverName,omitempty"`
	InsecureSkipVerify *bool                                `json:"insecureSkipVerify,omitempty"`
	MinVersion         *monitoringv1.TLSVersion             `json:"minVersion,omitempty"`
	MaxVersion         *monitoringv1.TLSVersion             `json:"maxVersion,omitempty"`
}

// SafeTLSConfigApplyConfiguration constructs a declarative configuration of the SafeTLSConfig type for use with
// apply.
func SafeTLSConfig() *SafeTLSConfigApplyConfiguration {
	return &SafeTLSConfigApplyConfiguration{}
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *SafeTLSConfigApplyConfiguration) WithCA(value *SecretOrConfigMapApplyConfiguration) *SafeTLSConfigApplyConfiguration {
	b.CA = value
	return b
}

// WithCert sets the Cert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cert field is set to the value of the last call.
func (b *SafeTLSConfigApplyConfiguration) WithCert(value *SecretOrConfigMapApplyConfiguration) *SafeTLSConfigApplyConfiguration {
	b.Cert = value
	return b
}

// WithKeySecret sets the KeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeySecret field is set to the value of the last call.
func (b *SafeTLSConfigApplyConfiguration) WithKeySecret(value corev1.SecretKeySelector) *SafeTLSConfigApplyConfiguration {
	b.KeySecret = &value
	return b
}

// WithServerName sets the ServerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerName field is set to the value of the last call.
func (b *SafeTLSConfigApplyConfiguration) WithServerName(value string) *SafeTLSConfigApplyConfiguration {
	b.ServerName = &value
	return b
}

// WithInsecureSkipVerify sets the InsecureSkipVerify field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InsecureSkipVerify field is set to the value of the last call.
func (b *SafeTLSConfigApplyConfiguration) WithInsecureSkipVerify(value bool) *SafeTLSConfigApplyConfiguration {
	b.InsecureSkipVerify = &value
	return b
}

// WithMinVersion sets the MinVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinVersion field is set to the value of the last call.
func (b *SafeTLSConfigApplyConfiguration) WithMinVersion(value monitoringv1.TLSVersion) *SafeTLSConfigApplyConfiguration {
	b.MinVersion = &value
	return b
}

// WithMaxVersion sets the MaxVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxVersion field is set to the value of the last call.
func (b *SafeTLSConfigApplyConfiguration) WithMaxVersion(value monitoringv1.TLSVersion) *SafeTLSConfigApplyConfiguration {
	b.MaxVersion = &value
	return b
}
