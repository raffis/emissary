/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	types "k8s.io/apimachinery/pkg/types"
)

// StatusDetailsApplyConfiguration represents an declarative configuration of the StatusDetails type for use
// with apply.
type StatusDetailsApplyConfiguration struct {
	Name              *string                         `json:"name,omitempty"`
	Group             *string                         `json:"group,omitempty"`
	Kind              *string                         `json:"kind,omitempty"`
	UID               *types.UID                      `json:"uid,omitempty"`
	Causes            []StatusCauseApplyConfiguration `json:"causes,omitempty"`
	RetryAfterSeconds *int32                          `json:"retryAfterSeconds,omitempty"`
}

// StatusDetailsApplyConfiguration constructs an declarative configuration of the StatusDetails type for use with
// apply.
func StatusDetails() *StatusDetailsApplyConfiguration {
	return &StatusDetailsApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *StatusDetailsApplyConfiguration) WithName(value string) *StatusDetailsApplyConfiguration {
	b.Name = &value
	return b
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *StatusDetailsApplyConfiguration) WithGroup(value string) *StatusDetailsApplyConfiguration {
	b.Group = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *StatusDetailsApplyConfiguration) WithKind(value string) *StatusDetailsApplyConfiguration {
	b.Kind = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *StatusDetailsApplyConfiguration) WithUID(value types.UID) *StatusDetailsApplyConfiguration {
	b.UID = &value
	return b
}

// WithCauses adds the given value to the Causes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Causes field.
func (b *StatusDetailsApplyConfiguration) WithCauses(values ...*StatusCauseApplyConfiguration) *StatusDetailsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCauses")
		}
		b.Causes = append(b.Causes, *values[i])
	}
	return b
}

// WithRetryAfterSeconds sets the RetryAfterSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RetryAfterSeconds field is set to the value of the last call.
func (b *StatusDetailsApplyConfiguration) WithRetryAfterSeconds(value int32) *StatusDetailsApplyConfiguration {
	b.RetryAfterSeconds = &value
	return b
}
