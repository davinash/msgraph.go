// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CredentialUserRegistrationCount undocumented
type CredentialUserRegistrationCount struct {
	Entity
	// TotalUserCount undocumented
	TotalUserCount *int `json:"totalUserCount,omitempty"`
	// UserRegistrationCounts undocumented
	UserRegistrationCounts []UserRegistrationCount `json:"userRegistrationCounts,omitempty"`
}