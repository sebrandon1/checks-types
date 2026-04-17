package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BestPracticeScannerSpec struct {
	// +optional
	TargetNamespace string `json:"targetNamespace,omitempty"`
	// +optional
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
	// +optional
	// +kubebuilder:validation:Pattern=`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`
	ScanInterval string `json:"scanInterval,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Checks []string `json:"checks,omitempty"`
	// +optional
	Suspend bool `json:"suspend,omitempty"`
}

type ScannerPhase string

const (
	PhaseIdle      ScannerPhase = "Idle"
	PhaseScanning  ScannerPhase = "Scanning"
	PhaseCompleted ScannerPhase = "Completed"
	PhaseError     ScannerPhase = "Error"
)

type ScanSummary struct {
	Total        int `json:"total"`
	Compliant    int `json:"compliant"`
	NonCompliant int `json:"nonCompliant"`
	Error        int `json:"error"`
	Skipped      int `json:"skipped"`
}

const (
	ConditionScanComplete   = "ScanComplete"
	ConditionProbeAvailable = "ProbeAvailable"
)

const (
	ReasonScanStarted      = "ScanStarted"
	ReasonScanSucceeded    = "ScanSucceeded"
	ReasonScanFailed       = "ScanFailed"
	ReasonScanCompleted    = "ScanCompleted"
	ReasonProbePodsReady   = "ProbePodsReady"
	ReasonProbePodsFailed  = "ProbePodsFailed"
	ReasonProbeUnavailable = "ProbeUnavailable"
)

type BestPracticeScannerStatus struct {
	// +optional
	Phase ScannerPhase `json:"phase,omitempty"`
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// +optional
	LastScanTime *metav1.Time `json:"lastScanTime,omitempty"`
	// +optional
	NextScanTime *metav1.Time `json:"nextScanTime,omitempty"`
	// +optional
	Summary *ScanSummary `json:"summary,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=bps
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Interval",type=string,JSONPath=`.spec.scanInterval`,priority=0
// +kubebuilder:printcolumn:name="Compliant",type=integer,JSONPath=`.status.summary.compliant`
// +kubebuilder:printcolumn:name="NonCompliant",type=integer,JSONPath=`.status.summary.nonCompliant`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type BestPracticeScanner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BestPracticeScannerSpec   `json:"spec,omitempty"`
	Status BestPracticeScannerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type BestPracticeScannerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BestPracticeScanner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BestPracticeScanner{}, &BestPracticeScannerList{})
}
