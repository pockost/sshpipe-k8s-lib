package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SshPipe is a top-level type
type SshPipe struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	
	// +optional
	Status SshPipeStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec SshPipeSpec `json:"spec,omitempty"`
}

// custom spec 
type SshPipeSpec struct {
  Users []string `json:"users"`
  Target TargetSpec `json:"target"`
}
type TargetSpec struct {
  Name string `json:"name"`
  Port int    `json:"port:omitempty"`
}

// custom status
type SshPipeStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type SshPipeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []SshPipe `json:"items"`
}
