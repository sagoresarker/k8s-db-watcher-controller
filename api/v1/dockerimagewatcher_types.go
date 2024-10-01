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

// DockerImageWatcherSpec defines the desired state of DockerImageWatcher
type DockerImageWatcherSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DockerImageWatcher. Edit dockerimagewatcher_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DockerImageWatcherStatus defines the observed state of DockerImageWatcher
type DockerImageWatcherStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DockerImageWatcher is the Schema for the dockerimagewatchers API
type DockerImageWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DockerImageWatcherSpec   `json:"spec,omitempty"`
	Status DockerImageWatcherStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DockerImageWatcherList contains a list of DockerImageWatcher
type DockerImageWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DockerImageWatcher `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DockerImageWatcher{}, &DockerImageWatcherList{})
}
