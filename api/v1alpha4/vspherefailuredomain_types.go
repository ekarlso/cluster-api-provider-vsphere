/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FailureDomainType string

const (
	HostGroupFailureDomain      FailureDomainType = "HostGroup"
	ComputeClusterFailureDomain FailureDomainType = "ComputeCluster"
	DatacenterFailureDomain     FailureDomainType = "Datacenter"
)

// VSphereFailureDomainSpec defines the desired state of VSphereFailureDomain
type VSphereFailureDomainSpec struct {

	// Region defines the name and type of a region
	Region FailureDomain `json:"region"`

	// Zone defines the name and type of a zone
	Zone FailureDomain `json:"zone"`

	// Topology is the what describes a given failure domain using vSphere constructs
	Topology Topology `json:"topology"`
}

type FailureDomain struct {
	// Name is the name of the tag that represents this failure domain
	Name string `json:"name"`

	// Type is the type of failure domain, the current values are "Datacenter", "ComputeCluster" and "HostGroup"
	// +kubebuilder:validation:Enum=Datacenter;ComputeCluster;HostGroup
	Type FailureDomainType `json:"type"`

	// TagCategory is the category used for the tag
	TagCategory string `json:"tagCategory"`

	// AutoConfigure tags the Type which is specified in the Topology
	AutoConfigure *bool `json:"autoConfigure,omitempty"`
}

type Topology struct {
	// The underlying infrastructure for this failure domain
	// Datacenter as the failure domain
	Datacenter string `json:"datacenter"`

	// ComputeCluster as the failure domain
	// +optional
	ComputeCluster *string `json:"computeCluster,omitempty"`

	// HostGroup as the failure domain
	// +optional
	HostGroup *FailureDomainHostGroup `json:"hostGroup,omitempty"`
}

// FailureDomainHostGroup as the failure domain
type FailureDomainHostGroup struct {
	// name of the host group
	Name string `json:"name"`

	// AutoConfigure creates the given hostGroup based on the supplied zone tagging
	// +optional
	AutoConfigure *bool `json:"autoConfigure,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:path=vspherefailuredomains,scope=Cluster,categories=cluster-api

// VSphereFailureDomain is the Schema for the vspherefailuredomains API
type VSphereFailureDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VSphereFailureDomainSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereFailureDomainList contains a list of VSphereFailureDomain
type VSphereFailureDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereFailureDomain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VSphereFailureDomain{}, &VSphereFailureDomainList{})
}
