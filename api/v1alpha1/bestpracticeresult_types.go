package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComplianceStatus string

const (
	StatusCompliant    ComplianceStatus = "Compliant"
	StatusNonCompliant ComplianceStatus = "NonCompliant"
	StatusError        ComplianceStatus = "Error"
	StatusSkipped      ComplianceStatus = "Skipped"
)

type ResourceDetail struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Compliant bool   `json:"compliant"`
	Message   string `json:"message"`
}

type BestPracticeResultSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	ScannerRef string `json:"scannerRef"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	CheckName        string           `json:"checkName"`
	Category         string           `json:"category"`
	Description      string           `json:"description"`
	ComplianceStatus ComplianceStatus `json:"complianceStatus"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	Remediation string `json:"remediation,omitempty"`
	// +optional
	CatalogURL string `json:"catalogURL,omitempty"`
	// +optional
	Details   []ResourceDetail `json:"details,omitempty"`
	Timestamp metav1.Time      `json:"timestamp"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=bpr
// +kubebuilder:printcolumn:name="Check",type=string,JSONPath=`.spec.checkName`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.spec.complianceStatus`
// +kubebuilder:printcolumn:name="Scanner",type=string,JSONPath=`.spec.scannerRef`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type BestPracticeResult struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BestPracticeResultSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
type BestPracticeResultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BestPracticeResult `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BestPracticeResult{}, &BestPracticeResultList{})
}
