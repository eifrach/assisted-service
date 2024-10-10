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

package v2

import (
	v2 "k8s.io/api/autoscaling/v2"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// MetricTargetApplyConfiguration represents a declarative configuration of the MetricTarget type for use
// with apply.
type MetricTargetApplyConfiguration struct {
	Type               *v2.MetricTargetType `json:"type,omitempty"`
	Value              *resource.Quantity   `json:"value,omitempty"`
	AverageValue       *resource.Quantity   `json:"averageValue,omitempty"`
	AverageUtilization *int32               `json:"averageUtilization,omitempty"`
}

// MetricTargetApplyConfiguration constructs a declarative configuration of the MetricTarget type for use with
// apply.
func MetricTarget() *MetricTargetApplyConfiguration {
	return &MetricTargetApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *MetricTargetApplyConfiguration) WithType(value v2.MetricTargetType) *MetricTargetApplyConfiguration {
	b.Type = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *MetricTargetApplyConfiguration) WithValue(value resource.Quantity) *MetricTargetApplyConfiguration {
	b.Value = &value
	return b
}

// WithAverageValue sets the AverageValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AverageValue field is set to the value of the last call.
func (b *MetricTargetApplyConfiguration) WithAverageValue(value resource.Quantity) *MetricTargetApplyConfiguration {
	b.AverageValue = &value
	return b
}

// WithAverageUtilization sets the AverageUtilization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AverageUtilization field is set to the value of the last call.
func (b *MetricTargetApplyConfiguration) WithAverageUtilization(value int32) *MetricTargetApplyConfiguration {
	b.AverageUtilization = &value
	return b
}
