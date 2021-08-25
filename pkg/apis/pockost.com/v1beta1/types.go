package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SshPipe is a top-level type
type SshPipe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SshPipeSpec `json:"spec"`
}

// SshPipeSpec is the spec for a SshPipe resource
type SshPipeSpec struct {
	Users  []string   `json:"users"`
	Target TargetSpec `json:"target"`
	SecretName string `json:"secretName"`
}
type TargetSpec struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	User string `json:"user"`
	Namespace string `json:"namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SshPipeList is a list of SshPipe resources
type SshPipeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `son:"metadata"`

	Items []SshPipe `json:"items"`
}
