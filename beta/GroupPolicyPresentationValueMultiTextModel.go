// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyPresentationValueMultiText The entity represents a string value of a multi-line text box presentation on a policy definition.
type GroupPolicyPresentationValueMultiText struct {
	GroupPolicyPresentationValue
	// Values A collection of non-empty strings for the associated presentation.
	Values []string `json:"values,omitempty"`
}