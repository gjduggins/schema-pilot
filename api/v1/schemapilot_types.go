/*
Copyright 2025.

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

// SchemaPilotSpec defines the desired state of SchemaPilot.
type SchemaPilotSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Enum=Oracle;SQLServer
	DatabaseType     string          `json:"databaseType,omitempty"` // Oracle, SQLServer
	ConnectionString SecretReference `json:"connectionString"`
	Repository       RepositorySpec  `json:"repository"`
	SchemaFiles      []SchemaFile    `json:"schemaFiles,omitempty"`
}

// RepositorySpec defines the repository configuration.
type RepositorySpec struct {
	Url    string          `json:"url,omitempty"`
	Branch string          `json:"branch,omitempty"`
	Owner  string          `json:"owner,omitempty"`
	Repo   string          `json:"repo,omitempty"`
	Token  SecretReference `json:"token"`
}

type SecretReference struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type SchemaFile struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	// +kubebuilder:validation:Enum=Table;View;Procedure;Function;Trigger;Index;Grant;Synonym;PackageHeader;PackageBody;Other
	ObjectType string `json:"objectType,omitempty"` // Table, View, Procedure, etc.
}

// SchemaPilotStatus defines the observed state of SchemaPilot.
type SchemaPilotStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SchemaPilot is the Schema for the schemapilots API.
type SchemaPilot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SchemaPilotSpec   `json:"spec,omitempty"`
	Status SchemaPilotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaPilotList contains a list of SchemaPilot.
type SchemaPilotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchemaPilot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SchemaPilot{}, &SchemaPilotList{})
}
