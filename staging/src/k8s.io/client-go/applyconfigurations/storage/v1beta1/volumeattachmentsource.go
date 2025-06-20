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

package v1beta1

import (
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// VolumeAttachmentSourceApplyConfiguration represents a declarative configuration of the VolumeAttachmentSource type for use
// with apply.
type VolumeAttachmentSourceApplyConfiguration struct {
	PersistentVolumeName *string                                    `json:"persistentVolumeName,omitempty"`
	InlineVolumeSpec     *v1.PersistentVolumeSpecApplyConfiguration `json:"inlineVolumeSpec,omitempty"`
}

// VolumeAttachmentSourceApplyConfiguration constructs a declarative configuration of the VolumeAttachmentSource type for use with
// apply.
func VolumeAttachmentSource() *VolumeAttachmentSourceApplyConfiguration {
	return &VolumeAttachmentSourceApplyConfiguration{}
}
func (b VolumeAttachmentSourceApplyConfiguration) IsApplyConfiguration() {}

// WithPersistentVolumeName sets the PersistentVolumeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PersistentVolumeName field is set to the value of the last call.
func (b *VolumeAttachmentSourceApplyConfiguration) WithPersistentVolumeName(value string) *VolumeAttachmentSourceApplyConfiguration {
	b.PersistentVolumeName = &value
	return b
}

// WithInlineVolumeSpec sets the InlineVolumeSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InlineVolumeSpec field is set to the value of the last call.
func (b *VolumeAttachmentSourceApplyConfiguration) WithInlineVolumeSpec(value *v1.PersistentVolumeSpecApplyConfiguration) *VolumeAttachmentSourceApplyConfiguration {
	b.InlineVolumeSpec = value
	return b
}
