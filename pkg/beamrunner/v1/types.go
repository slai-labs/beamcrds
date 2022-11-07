package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BeamRunner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              BeamRunnerSpec `json:"spec"`
}

type Trigger struct {
	Inputs      map[string]map[string]string `json:"inputs"`
	Outputs     map[string]map[string]string `json:"outputs"`
	TriggerType string                       `json:"trigger_type"`
	Handler     string                       `json:"handler"`
	Loader      string                       `json:"loader"`
	Expression  string                       `json:"expression"` // Cron expression
}

type BeamRunnerSpec struct {
	IdentityId      string  `json:"identityId"`
	Image           string  `json:"image"`
	AppId           string  `json:"appId"`
	AppVersion      int     `json:"appVersion"`
	AppSessionId    string  `json:"appSessionId"`
	AppDeploymentId string  `json:"appDeploymentId"`
	Trigger         Trigger `json:"trigger"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BeamRunnerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []BeamRunner `json:"items"`
}
