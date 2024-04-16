/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeCleanUpSpec defines the desired state of NodeCleanUp
type NodeCleanUpSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of NodeCleanUp. Edit nodecleanup_types.go to remove/update
	Name    string `json:"name,omitempty"`
	Site    string `json:"site,omitempty"`
	Address string `json:"address,omitempty"`
	Zone    string `json:"zone,omitempty"`
	URL     string `json:"url,omitempty"`
}

// NodeCleanUpStatus defines the observed state of NodeCleanUp
type NodeCleanUpStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Cleanup bool `json:"cleanup"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NodeCleanUp is the Schema for the nodecleanups API
type NodeCleanUp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeCleanUpSpec   `json:"spec,omitempty"`
	Status NodeCleanUpStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NodeCleanUpList contains a list of NodeCleanUp
type NodeCleanUpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeCleanUp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeCleanUp{}, &NodeCleanUpList{})
}
