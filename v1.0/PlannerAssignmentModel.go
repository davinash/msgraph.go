// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PlannerAssignment undocumented
type PlannerAssignment struct {
	// AssignedBy undocumented
	AssignedBy *IdentitySet `json:"assignedBy,omitempty"`
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
}